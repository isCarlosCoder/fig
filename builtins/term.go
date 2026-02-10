package builtins

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/iscarloscoder/fig/environment"
	"golang.org/x/sys/unix"
	"golang.org/x/term"
)

// Minimal term builtin: startRaw / stopRaw / isRaw
// Note: This implementation intentionally provides only the requested primitives.

var termMu sync.Mutex
var termOldState *term.State
var termFd int = -1
var termSigCh chan os.Signal

// resize callback management
var resizeMu sync.Mutex
var resizeCb environment.Value
var resizeSigCh chan os.Signal

// onKey management
var onKeyMu sync.Mutex
var onKeyCb environment.Value
var onKeyStopCh chan struct{}
var onKeyMadeRaw bool

// screen buffer for drawing; use bufferMu to protect
var bufferMu sync.Mutex
var screenBuf = bytes.Buffer{}

func startRawMode() error {
	termMu.Lock()
	defer termMu.Unlock()
	if termOldState != nil {
		// already in raw
		return nil
	}
	f := os.Stdin
	if !term.IsTerminal(int(f.Fd())) {
		// not a TTY — nothing to do
		return nil
	}
	fd := int(f.Fd())
	old, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	termOldState = old
	termFd = fd
	termSigCh = make(chan os.Signal, 1)
	signal.Notify(termSigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for sig := range termSigCh {
			termMu.Lock()
			if termOldState != nil {
				_ = term.Restore(termFd, termOldState)
				termOldState = nil
			}
			termMu.Unlock()
			// re-raise the signal so default behavior applies
			if s, ok := sig.(syscall.Signal); ok {
				_ = syscall.Kill(syscall.Getpid(), s)
			} else {
				_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
			}
		}
	}()
	return nil
}

func stopRawMode() error {
	termMu.Lock()
	defer termMu.Unlock()
	if termOldState == nil {
		return nil
	}
	err := term.Restore(termFd, termOldState)
	termOldState = nil
	termFd = -1
	if termSigCh != nil {
		signal.Stop(termSigCh)
		close(termSigCh)
		termSigCh = nil
	}
	return err
}

func isRawMode() bool {
	termMu.Lock()
	defer termMu.Unlock()
	return termOldState != nil
}

// read raw bytes from stdin with optional timeout. returns nil, nil when timeout and no data.
// Implementation uses non-blocking reads when timeout > 0 to poll for availability.
func readRawBytes(timeout time.Duration) ([]byte, error) {
	f := os.Stdin
	fd := int(f.Fd())
	// blocking read
	if timeout == 0 {
		b := make([]byte, 1)
		n, err := f.Read(b)
		if n > 0 {
			buf := append([]byte(nil), b[:n]...)
			// try to drain additional immediate bytes (non-blocking)
			_ = unix.SetNonblock(fd, true)
			tmp := make([]byte, 128)
			for {
				n2, _ := f.Read(tmp)
				if n2 <= 0 {
					break
				}
				buf = append(buf, tmp[:n2]...)
				if len(buf) >= 64 {
					break
				}
			}
			_ = unix.SetNonblock(fd, false)
			return buf, nil
		}
		return nil, err
	}

	// timeout > 0: poll using non-blocking reads until timeout
	_ = unix.SetNonblock(fd, true)
	defer unix.SetNonblock(fd, false)
	deadline := time.Now().Add(timeout)
	buf := make([]byte, 0, 64)
	tmp := make([]byte, 128)
	for {
		n, err := f.Read(tmp)
		if n > 0 {
			buf = append(buf, tmp[:n]...)
			return buf, nil
		}
		if err != nil {
			// treat EAGAIN/EWOULDBLOCK as no-data-yet using errors.Is to unwrap
			if errors.Is(err, syscall.EAGAIN) || errors.Is(err, syscall.EWOULDBLOCK) {
				// continue polling until timeout
			} else if errors.Is(err, io.EOF) {
				return nil, nil
			} else {
				// unknown error
				return nil, err
			}
		}
		if time.Now().After(deadline) {
			return nil, nil
		}
		// small sleep to avoid busy-loop
		time.Sleep(1 * time.Millisecond)
	}
}

func init() {
	register(newModule("term",
		fn("startRaw", func(args []environment.Value) (environment.Value, error) {
			if err := startRawMode(); err != nil {
				return environment.NewNil(), fmt.Errorf("startRaw() failed: %v", err)
			}
			return environment.NewNil(), nil
		}),

		fn("stopRaw", func(args []environment.Value) (environment.Value, error) {
			if err := stopRawMode(); err != nil {
				return environment.NewNil(), fmt.Errorf("stopRaw() failed: %v", err)
			}
			return environment.NewNil(), nil
		}),

		fn("isRaw", func(args []environment.Value) (environment.Value, error) {
			return environment.NewBool(isRawMode()), nil
		}),

		// readKey(blocking=true, timeout_ms=0) -> string | nil (raw bytes as string)
		fn("readKey", func(args []environment.Value) (environment.Value, error) {
			blocking := true
			timeoutMs := float64(0)
			if len(args) >= 1 {
				if b, err := args[0].AsBool(); err == nil {
					blocking = b
				}
			}
			if len(args) >= 2 {
				if n, err := args[1].AsNumber(); err == nil {
					timeoutMs = n
				}
			}
			var to time.Duration
			if !blocking && timeoutMs == 0 {
				to = 1 * time.Millisecond
			} else if timeoutMs > 0 {
				to = time.Duration(int(timeoutMs)) * time.Millisecond
			} else {
				to = 0
			}
			data, err := readRawBytes(to)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("readKey() failed: %v", err)
			}
			if data == nil || len(data) == 0 {
				return environment.NewNil(), nil
			}
			return environment.NewString(string(data)), nil
		}),

		// readKeyBlocking() — like readKey but guarantees blocking until a key is available
		fn("readKeyBlocking", func(args []environment.Value) (environment.Value, error) {
			// ignore args, always block
			data, err := readRawBytes(0)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("readKeyBlocking() failed: %v", err)
			}
			if data == nil || len(data) == 0 {
				// in theory should not happen, but return nil to be safe
				return environment.NewNil(), nil
			}
			return environment.NewString(string(data)), nil
		}),

		// pollKey() -> string | nil : non-blocking, returns immediately
		fn("pollKey", func(args []environment.Value) (environment.Value, error) {
			data, err := readRawBytes(1 * time.Millisecond)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("pollKey() failed: %v", err)
			}
			if data == nil || len(data) == 0 {
				return environment.NewNil(), nil
			}
			return environment.NewString(string(data)), nil
		}),

		// keyPressed() -> bool : check if there is data available on stdin without consuming it
		fn("keyPressed", func(args []environment.Value) (environment.Value, error) {
			f := os.Stdin
			fd := int32(f.Fd())
			pfd := []unix.PollFd{{Fd: fd, Events: unix.POLLIN}}
			_, err := unix.Poll(pfd, 0)
			if err != nil {
				return environment.NewBool(false), nil
			}
			return environment.NewBool((pfd[0].Revents & unix.POLLIN) != 0), nil
		}),

		// clear() -> clears the screen and moves cursor to origin
		fn("clear", func(args []environment.Value) (environment.Value, error) {
			seq := "\x1b[2J" + "\x1b[H"
			// write directly to stdout
			_, _ = os.Stdout.Write([]byte(seq))
			return environment.NewNil(), nil
		}),

		// write(s) -> append to screen buffer (does not flush)
		fn("write", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("write() expects 1 argument")
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("write() argument must be a string")
			}
			bufferMu.Lock()
			screenBuf.WriteString(s)
			bufferMu.Unlock()
			return environment.NewNil(), nil
		}),

		// writeln(s) -> append s + '\n' to buffer
		fn("writeln", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("writeln() expects 1 argument")
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("writeln() argument must be a string")
			}
			bufferMu.Lock()
			screenBuf.WriteString(s + "\n")
			bufferMu.Unlock()
			return environment.NewNil(), nil
		}),

		// refresh() -> flush screen buffer to terminal and clear it
		fn("refresh", func(args []environment.Value) (environment.Value, error) {
			bufferMu.Lock()
			if screenBuf.Len() > 0 {
				_, _ = os.Stdout.Write(screenBuf.Bytes())
				screenBuf.Reset()
			}
			bufferMu.Unlock()
			return environment.NewNil(), nil
		}),

		// flush() -> force immediate write of buffer to stdout (clears buffer)
		fn("flush", func(args []environment.Value) (environment.Value, error) {
			bufferMu.Lock()
			if screenBuf.Len() > 0 {
				_, _ = os.Stdout.Write(screenBuf.Bytes())
				screenBuf.Reset()
			}
			bufferMu.Unlock()
			return environment.NewNil(), nil
		}),

		// move(row, col) -> move cursor to absolute position using ANSI escape
		fn("move", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("move() expects 2 numeric arguments: row, col")
			}
			rn, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("move() first argument must be a number")
			}
			cn, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("move() second argument must be a number")
			}
			row := int(rn)
			col := int(cn)
			if row < 1 {
				row = 1
			}
			if col < 1 {
				col = 1
			}
			seq := fmt.Sprintf("\x1b[%d;%dH", row, col)
			_, _ = os.Stdout.Write([]byte(seq))
			return environment.NewNil(), nil
		}),

		// writeAt(row, col, s) -> move cursor to absolute position and write string immediately
		fn("writeAt", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("writeAt() expects 3 arguments: row, col, str")
			}
			rn, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("writeAt() first argument must be a number")
			}
			cn, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("writeAt() second argument must be a number")
			}
			s, err := args[2].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("writeAt() third argument must be a string")
			}
			row := int(rn)
			col := int(cn)
			if row < 1 {
				row = 1
			}
			if col < 1 {
				col = 1
			}
			seq := fmt.Sprintf("\x1b[%d;%dH%s", row, col, s)
			_, _ = os.Stdout.Write([]byte(seq))
			return environment.NewNil(), nil
		}),

		// createBuffer(w, h) -> returns a buffer object with methods: set, clear, fill, copy
		fn("createBuffer", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("createBuffer() expects 2 numeric arguments: w, h")
			}
			wn, err1 := args[0].AsNumber()
			hn, err2 := args[1].AsNumber()
			if err1 != nil || err2 != nil {
				return environment.NewNil(), fmt.Errorf("createBuffer() expects numeric w and h")
			}
			w := int(wn)
			h := int(hn)
			if w <= 0 || h <= 0 {
				return environment.NewNil(), fmt.Errorf("createBuffer() width and height must be > 0")
			}
			rows := make([]environment.Value, h)
			blank := strings.Repeat(" ", w)
			for i := 0; i < h; i++ {
				rows[i] = environment.NewString(blank)
			}

			entries := map[string]environment.Value{"w": environment.NewNumber(float64(w)), "h": environment.NewNumber(float64(h)), "rows": environment.Value{Type: environment.ArrayType, Arr: &rows}}
			// Methods bound to this buffer (closures capture rows, w, h)
			entries["set"] = environment.NewBuiltinFn("buffer.set", func(a []environment.Value) (environment.Value, error) {
				if len(a) != 3 {
					return environment.NewNil(), fmt.Errorf("buffer.set(row, col, char) expects 3 arguments")
				}
				rn, err := a[0].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("buffer.set: first arg must be number (row)")
				}
				cn, err := a[1].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("buffer.set: second arg must be number (col)")
				}
				s, err := a[2].AsString()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("buffer.set: third arg must be a string char")
				}
				row := int(rn) - 1
				col := int(cn) - 1
				if row < 0 || row >= h || col < 0 || col >= w {
					return environment.NewNil(), fmt.Errorf("buffer.set: coordinates out of bounds")
				}
				// get current row as rune slice
				cur := []rune("")
				if rows[row].Type == environment.StringType {
					cur = []rune(rows[row].Str)
				}
				if len(cur) < w {
					pad := make([]rune, w-len(cur))
					for i := range pad {
						pad[i] = ' '
					}
					cur = append(cur, pad...)
				} else if len(cur) > w {
					cur = cur[:w]
				}
				chRunes := []rune(s)
				if len(chRunes) == 0 {
					return environment.NewNil(), fmt.Errorf("buffer.set: empty char")
				}
				cur[col] = chRunes[0]
				rows[row] = environment.NewString(string(cur))
				return environment.NewNil(), nil
			})

			entries["clear"] = environment.NewBuiltinFn("buffer.clear", func(a []environment.Value) (environment.Value, error) {
				blank := strings.Repeat(" ", w)
				for i := 0; i < h; i++ {
					rows[i] = environment.NewString(blank)
				}
				return environment.NewNil(), nil
			})

			entries["fill"] = environment.NewBuiltinFn("buffer.fill", func(a []environment.Value) (environment.Value, error) {
				if len(a) != 1 {
					return environment.NewNil(), fmt.Errorf("buffer.fill(char) expects 1 string argument")
				}
				s, err := a[0].AsString()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("buffer.fill: arg must be a string char")
				}
				runes := []rune(s)
				if len(runes) == 0 {
					return environment.NewNil(), fmt.Errorf("buffer.fill: empty char")
				}
				ch := runes[0]
				line := string([]rune{ch})
				line = strings.Repeat(line, w)
				for i := 0; i < h; i++ {
					rows[i] = environment.NewString(line)
				}
				return environment.NewNil(), nil
			})

			entries["copy"] = environment.NewBuiltinFn("buffer.copy", func(a []environment.Value) (environment.Value, error) {
				if len(a) != 1 {
					return environment.NewNil(), fmt.Errorf("buffer.copy(otherBuf) expects 1 argument")
				}
				src := a[0]
				if src.Type != environment.ObjectType || src.Obj == nil {
					return environment.NewNil(), fmt.Errorf("buffer.copy: argument must be a buffer object")
				}
				wVal, ok := src.Obj.Entries["w"]
				hVal, ok2 := src.Obj.Entries["h"]
				rowsVal, ok3 := src.Obj.Entries["rows"]
				if !ok || !ok2 || !ok3 || rowsVal.Type != environment.ArrayType {
					return environment.NewNil(), fmt.Errorf("buffer.copy: invalid source buffer")
				}
				sw := int(wVal.Num)
				sh := int(hVal.Num)
				srows := *rowsVal.Arr
				// copy clipped to destination size
				for i := 0; i < h && i < sh; i++ {
					destRunes := []rune("")
					if rows[i].Type == environment.StringType {
						destRunes = []rune(rows[i].Str)
					}
					if len(destRunes) < w {
						pad := make([]rune, w-len(destRunes))
						for k := range pad {
							pad[k] = ' '
						}
						destRunes = append(destRunes, pad...)
					} else if len(destRunes) > w {
						destRunes = destRunes[:w]
					}
					srcRunes := []rune("")
					if srows[i].Type == environment.StringType {
						srcRunes = []rune(srows[i].Str)
					}
					for j := 0; j < w && j < sw && j < len(srcRunes); j++ {
						destRunes[j] = srcRunes[j]
					}
					rows[i] = environment.NewString(string(destRunes))
				}
				return environment.NewNil(), nil
			})

			keys := []string{"w", "h", "rows", "set", "clear", "fill", "copy"}
			return environment.NewObject(entries, keys), nil
		}),

		// drawBuffer(buf) -> draws the buffer object to the terminal at absolute position (1,1)
		fn("drawBuffer", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("drawBuffer() expects 1 argument: buffer object")
			}
			buf := args[0]
			if buf.Type != environment.ObjectType || buf.Obj == nil {
				return environment.NewNil(), fmt.Errorf("drawBuffer() expects an object created by createBuffer()")
			}
			wVal, ok := buf.Obj.Entries["w"]
			hVal, ok2 := buf.Obj.Entries["h"]
			rowsVal, ok3 := buf.Obj.Entries["rows"]
			if !ok || !ok2 || !ok3 || rowsVal.Type != environment.ArrayType {
				return environment.NewNil(), fmt.Errorf("drawBuffer() invalid buffer object")
			}
			w := int(wVal.Num)
			h := int(hVal.Num)
			rows := *rowsVal.Arr
			// Move to origin
			_, _ = os.Stdout.Write([]byte("\x1b[H"))
			for i := 0; i < h && i < len(rows); i++ {
				rowStr := ""
				if rows[i].Type == environment.StringType {
					rowStr = rows[i].Str
				}
				// pad or trim to width
				if len(rowStr) < w {
					rowStr = rowStr + strings.Repeat(" ", w-len(rowStr))
				} else if len(rowStr) > w {
					rowStr = rowStr[:w]
				}
				seq := fmt.Sprintf("\x1b[%d;1H%s", i+1, rowStr)
				_, _ = os.Stdout.Write([]byte(seq))
			}
			return environment.NewNil(), nil
		}),

		// moveUp(n) -> move cursor up by n rows
		fn("moveUp", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("moveUp() expects 1 numeric argument: n")
			}
			nn, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("moveUp() argument must be a number")
			}
			n := int(nn)
			if n < 1 {
				n = 1
			}
			seq := fmt.Sprintf("\x1b[%dA", n)
			_, _ = os.Stdout.Write([]byte(seq))
			return environment.NewNil(), nil
		}),

		// moveDown(n) -> move cursor down by n rows
		fn("moveDown", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("moveDown() expects 1 numeric argument: n")
			}
			nn, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("moveDown() argument must be a number")
			}
			n := int(nn)
			if n < 1 {
				n = 1
			}
			seq := fmt.Sprintf("\x1b[%dB", n)
			_, _ = os.Stdout.Write([]byte(seq))
			return environment.NewNil(), nil
		}),

		// moveRight(n) -> move cursor right by n columns
		fn("moveRight", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("moveRight() expects 1 numeric argument: n")
			}
			nn, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("moveRight() argument must be a number")
			}
			n := int(nn)
			if n < 1 {
				n = 1
			}
			seq := fmt.Sprintf("\x1b[%dC", n)
			_, _ = os.Stdout.Write([]byte(seq))
			return environment.NewNil(), nil
		}),

		// moveLeft(n) -> move cursor left by n columns
		fn("moveLeft", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("moveLeft() expects 1 numeric argument: n")
			}
			nn, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("moveLeft() argument must be a number")
			}
			n := int(nn)
			if n < 1 {
				n = 1
			}
			seq := fmt.Sprintf("\x1b[%dD", n)
			_, _ = os.Stdout.Write([]byte(seq))
			return environment.NewNil(), nil
		}),

		// hideCursor() -> hide the terminal cursor
		fn("hideCursor", func(args []environment.Value) (environment.Value, error) {
			// ANSI: CSI ? 25 l
			_, _ = os.Stdout.Write([]byte("\x1b[?25l"))
			return environment.NewNil(), nil
		}),

		// showCursor() -> show the terminal cursor
		fn("showCursor", func(args []environment.Value) (environment.Value, error) {
			// ANSI: CSI ? 25 h
			_, _ = os.Stdout.Write([]byte("\x1b[?25h"))
			return environment.NewNil(), nil
		}),

		// fg(color) -> set foreground color. Accepts hex string "#rrggbb" or three numeric args (r,g,b)
		fn("fg", func(args []environment.Value) (environment.Value, error) {
			var r, g, b int
			if len(args) == 1 {
				// try hex string
				if s, err := args[0].AsString(); err == nil {
					if len(s) == 7 && s[0] == '#' {
						var rv, gv, bv uint64
						var e error
						rv, e = strconv.ParseUint(s[1:3], 16, 8)
						if e != nil {
							return environment.NewNil(), fmt.Errorf("fg() invalid hex color")
						}
						gv, e = strconv.ParseUint(s[3:5], 16, 8)
						if e != nil {
							return environment.NewNil(), fmt.Errorf("fg() invalid hex color")
						}
						bv, e = strconv.ParseUint(s[5:7], 16, 8)
						if e != nil {
							return environment.NewNil(), fmt.Errorf("fg() invalid hex color")
						}
						r, g, b = int(rv), int(gv), int(bv)
					} else {
						return environment.NewNil(), fmt.Errorf("fg() expects '#rrggbb' when given a single string")
					}
				} else {
					return environment.NewNil(), fmt.Errorf("fg() expects a hex string like '#rrggbb' or three numeric args")
				}
			} else if len(args) == 3 {
				// numeric r,g,b
				rn, err1 := args[0].AsNumber()
				gn, err2 := args[1].AsNumber()
				bn, err3 := args[2].AsNumber()
				if err1 != nil || err2 != nil || err3 != nil {
					return environment.NewNil(), fmt.Errorf("fg() expects numeric r,g,b values")
				}
				r, g, b = int(rn), int(gn), int(bn)
			} else {
				return environment.NewNil(), fmt.Errorf("fg() expects either '#rrggbb' or three numbers (r,g,b)")
			}
			seq := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
			_, _ = os.Stdout.Write([]byte(seq))
			return environment.NewNil(), nil
		}),

		// bg(color) -> set background color. Same ABI as fg()
		fn("bg", func(args []environment.Value) (environment.Value, error) {
			var r, g, b int
			if len(args) == 1 {
				if s, err := args[0].AsString(); err == nil {
					if len(s) == 7 && s[0] == '#' {
						var rv, gv, bv uint64
						var e error
						rv, e = strconv.ParseUint(s[1:3], 16, 8)
						if e != nil {
							return environment.NewNil(), fmt.Errorf("bg() invalid hex color")
						}
						gv, e = strconv.ParseUint(s[3:5], 16, 8)
						if e != nil {
							return environment.NewNil(), fmt.Errorf("bg() invalid hex color")
						}
						bv, e = strconv.ParseUint(s[5:7], 16, 8)
						if e != nil {
							return environment.NewNil(), fmt.Errorf("bg() invalid hex color")
						}
						r, g, b = int(rv), int(gv), int(bv)
					} else {
						return environment.NewNil(), fmt.Errorf("bg() expects '#rrggbb' when given a single string")
					}
				} else {
					return environment.NewNil(), fmt.Errorf("bg() expects a hex string like '#rrggbb' or three numeric args")
				}
			} else if len(args) == 3 {
				rn, err1 := args[0].AsNumber()
				gn, err2 := args[1].AsNumber()
				bn, err3 := args[2].AsNumber()
				if err1 != nil || err2 != nil || err3 != nil {
					return environment.NewNil(), fmt.Errorf("bg() expects numeric r,g,b values")
				}
				r, g, b = int(rn), int(gn), int(bn)
			} else {
				return environment.NewNil(), fmt.Errorf("bg() expects either '#rrggbb' or three numbers (r,g,b)")
			}
			seq := fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b)
			_, _ = os.Stdout.Write([]byte(seq))
			return environment.NewNil(), nil
		}),

		// resetStyle() -> reset colors and attributes
		fn("resetStyle", func(args []environment.Value) (environment.Value, error) {
			_, _ = os.Stdout.Write([]byte("\x1b[0m"))
			return environment.NewNil(), nil
		}),

		// sleep(ms) -> pause for ms milliseconds
		fn("sleep", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("sleep() expects 1 numeric argument: ms")
			}
			ms, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("sleep() expects a number (ms)")
			}
			time.Sleep(time.Duration(int(ms)) * time.Millisecond)
			return environment.NewNil(), nil
		}),

		// now() -> current timestamp in milliseconds
		fn("now", func(args []environment.Value) (environment.Value, error) {
			return environment.NewNumber(float64(time.Now().UnixNano() / int64(time.Millisecond))), nil
		}),

		// frameLimit(fps) -> returns a function to call each frame to enforce FPS
		fn("frameLimit", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("frameLimit() expects 1 numeric argument: fps")
			}
			fps, err := args[0].AsNumber()
			if err != nil || fps <= 0 {
				return environment.NewNil(), fmt.Errorf("frameLimit() expects a positive number for fps")
			}
			// support fractional fps
			frameNs := int64(float64(time.Second) / fps)
			dur := time.Duration(frameNs) * time.Nanosecond
			last := time.Now()
			fnName := fmt.Sprintf("frameWait@%d", time.Now().UnixNano())
			frameFn := func(a []environment.Value) (environment.Value, error) {
				elapsed := time.Since(last)
				if elapsed < dur {
					time.Sleep(dur - elapsed)
				}
				last = time.Now()
				return environment.NewNil(), nil
			}
			return environment.NewBuiltinFn(fnName, frameFn), nil
		}),

		// onResize(fn) -> register a callback function invoked when terminal resizes (SIGWINCH)
		// Pass `null` to unregister.
		fn("onResize", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("onResize() expects 1 argument: function or null")
			}
			if args[0].Type == environment.NilType {
				// unregister
				resizeMu.Lock()
				resizeCb = environment.NewNil()
				if resizeSigCh != nil {
					signal.Stop(resizeSigCh)
					close(resizeSigCh)
					resizeSigCh = nil
				}
				resizeMu.Unlock()
				return environment.NewNil(), nil
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("onResize() argument must be a function or null")
			}
			resizeMu.Lock()
			resizeCb = args[0]
			// if we don't already have a channel/goroutine listening, start it
			if resizeSigCh == nil {
				resizeSigCh = make(chan os.Signal, 1)
				signal.Notify(resizeSigCh, syscall.SIGWINCH)
				go func() {
					for range resizeSigCh {
						resizeMu.Lock()
						cb := resizeCb
						resizeMu.Unlock()
						if cb.Type == environment.NilType {
							continue
						}
						// Use TaskSpawner (recommended) to run user function in interpreter goroutine
						if TaskSpawner != nil {
							resultCh := make(chan TaskResult, 1)
							TaskSpawner(cb, resultCh)
							// we won't wait for the result; drop it when it arrives
							go func() { <-resultCh }()
						} else {
							// TaskSpawner not available; notify on stdout
							fmt.Fprintln(os.Stdout, "term.onResize: task module not loaded; callback not invoked")
						}
					}
				}()
			}
			resizeMu.Unlock()
			return environment.NewNil(), nil
		}),

		// onKey(fn) -> register a callback when any key is pressed; pass `null` to unregister
		fn("onKey", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("onKey() expects 1 argument: function or null")
			}
			if args[0].Type == environment.NilType {
				// unregister
				onKeyMu.Lock()
				if onKeyStopCh != nil {
					close(onKeyStopCh)
					onKeyStopCh = nil
				}
				onKeyCb = environment.NewNil()
				// restore raw mode if we enabled it
				if onKeyMadeRaw {
					_ = stopRawMode()
					onKeyMadeRaw = false
				}
				onKeyMu.Unlock()
				return environment.NewNil(), nil
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("onKey() argument must be a function or null")
			}
			onKeyMu.Lock()
			onKeyCb = args[0]
			if onKeyStopCh == nil {
				onKeyStopCh = make(chan struct{})
				// ensure raw mode for immediate key events
				if !isRawMode() {
					_ = startRawMode()
					onKeyMadeRaw = true
				}
				stopCh := onKeyStopCh
				go func() {
					fd := int(os.Stdin.Fd())
					pfd := []unix.PollFd{{Fd: int32(fd), Events: unix.POLLIN}}
					for {
						// poll with timeout so we can check stop channel periodically
						n, err := unix.Poll(pfd, 100)
						if err != nil {
							time.Sleep(50 * time.Millisecond)
							continue
						}
						if n > 0 && (pfd[0].Revents&unix.POLLIN) != 0 {
							// invoke callback via TaskSpawner
							onKeyMu.Lock()
							cb := onKeyCb
							onKeyMu.Unlock()
							if cb.Type == environment.NilType {
								continue
							}
							if TaskSpawner != nil {
								resultCh := make(chan TaskResult, 1)
								TaskSpawner(cb, resultCh)
								go func() { <-resultCh }()
							} else {
								fmt.Fprintln(os.Stdout, "term.onKey: task module not loaded; callback not invoked")
							}
						}
						select {
						case <-stopCh:
							return
						default:
						}
					}
				}()
			}
			onKeyMu.Unlock()
			return environment.NewNil(), nil
		}),

		// bold(enabled) -> enable/disable bold (intensity)
		fn("bold", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("bold() expects 1 boolean argument")
			}
			b, err := args[0].AsBool()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("bold() argument must be boolean")
			}
			if b {
				_, _ = os.Stdout.Write([]byte("\x1b[1m"))
			} else {
				_, _ = os.Stdout.Write([]byte("\x1b[22m"))
			}
			return environment.NewNil(), nil
		}),

		// underline(enabled) -> enable/disable underline
		fn("underline", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("underline() expects 1 boolean argument")
			}
			b, err := args[0].AsBool()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("underline() argument must be boolean")
			}
			if b {
				_, _ = os.Stdout.Write([]byte("\x1b[4m"))
			} else {
				_, _ = os.Stdout.Write([]byte("\x1b[24m"))
			}
			return environment.NewNil(), nil
		}),

		// invert(enabled) -> enable/disable reverse video (swap fg/bg)
		fn("invert", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("invert() expects 1 boolean argument")
			}
			b, err := args[0].AsBool()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("invert() argument must be boolean")
			}
			if b {
				_, _ = os.Stdout.Write([]byte("\x1b[7m"))
			} else {
				_, _ = os.Stdout.Write([]byte("\x1b[27m"))
			}
			return environment.NewNil(), nil
		}),

		// size() -> {rows, cols}
		fn("size", func(args []environment.Value) (environment.Value, error) {
			ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
			rows := 0
			cols := 0
			if err == nil && ws != nil {
				cols = int(ws.Col)
				rows = int(ws.Row)
			}
			entries := map[string]environment.Value{"rows": environment.NewNumber(float64(rows)), "cols": environment.NewNumber(float64(cols))}
			keys := []string{"rows", "cols"}
			return environment.NewObject(entries, keys), nil
		}),
	))
}
