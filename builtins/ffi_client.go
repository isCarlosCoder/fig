package builtins

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var helpersMu sync.Mutex
var helpers = map[string]*helperClient{}

// startHelperOnce starts the helper at path, runs it with --server and returns a client
// The helper is short-lived for now (we start/stop per call).
func startHelperOnce(path string) (*helperClient, error) {
	// start helper process (no context here to avoid killing it when function exits)
	cmd := exec.Command(path, "--server")
	in, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("cannot get stdin pipe: %v", err)
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("cannot get stdout pipe: %v", err)
	}
	errp, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("cannot get stderr pipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("cannot start helper: %v", err)
	}
	// capture stderr asynchronously
	var stderrBuf []byte
	go func() {
		sc := bufio.NewScanner(errp)
		for sc.Scan() {
			stderrBuf = append(stderrBuf, sc.Bytes()...)
			stderrBuf = append(stderrBuf, '\n')
		}
	}()
	// small delay to let helper initialize
	time.Sleep(20 * time.Millisecond)
	hc := &helperClient{cmd: cmd, stdin: in, stdout: bufio.NewReader(out), stderr: &stderrBuf, pending: map[string]chan map[string]interface{}{}}
	go hc.readLoop()
	return hc, nil
}

func (h *helperClient) readLoop() {
	sc := bufio.NewScanner(h.stdout)
	for sc.Scan() {
		var resp map[string]interface{}
		if err := json.Unmarshal(sc.Bytes(), &resp); err != nil {
			continue
		}
		idv, _ := resp["id"]
		var id string
		if idv != nil {
			switch t := idv.(type) {
			case string:
				id = t
			case float64:
				id = fmt.Sprintf("%d", int64(t))
			default:
				id = fmt.Sprintf("%v", t)
			}
		}
		if id == "" {
			// no id — nothing to route to pending; ignore
			continue
		}
		h.pmu.Lock()
		ch, ok := h.pending[id]
		h.pmu.Unlock()
		if ok {
			select {
			case ch <- resp:
			default:
			}
			// pending entry will be cleaned by caller
			continue
		}
		// no pending entry -> treat as incoming request from helper
		if cmd, _ := resp["cmd"].(string); cmd == "invoke_callback" {
			// extract cb id and args
			cbid, _ := resp["cb"].(string)
			args, _ := resp["args"].([]interface{})
			fmt.Fprintf(os.Stderr, "ffi client: received invoke_callback id=%s cb=%s args=%v\n", id, cbid, args)
			// run callback
			res, err := runCallback(cbid, args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ffi client: callback error: %v\n", err)
				_ = h.sendRaw(map[string]interface{}{"ok": false, "error": err.Error(), "id": id})
				continue
			}
			fmt.Fprintf(os.Stderr, "ffi client: callback result for id=%s res=%v\n", id, res)
			_ = h.sendRaw(map[string]interface{}{"ok": true, "resp": res, "id": id})
			continue
		}
		// unknown incoming message; ignore
		fmt.Fprintf(os.Stderr, "ffi client: ignoring incoming message: %v\n", resp)
		continue
	}
	// if scanner stopped, notify pending channels of error
	if err := sc.Err(); err != nil {
		h.pmu.Lock()
		for id, ch := range h.pending {
			select {
			case ch <- map[string]interface{}{"ok": false, "error": "helper disconnected"}:
			default:
			}
			delete(h.pending, id)
		}
		h.pmu.Unlock()
	}
}

// getHelperForProject returns a persistent helper client for the given projectRoot.
// It will restart the helper if it appears dead, up to the sandbox max_restarts limit.
func getHelperForProject(projectRoot string, helperPath string) (*helperClient, error) {
	helpersMu.Lock()
	hc, ok := helpers[projectRoot]
	helpersMu.Unlock()
	if ok {
		// quick health check: try a small ping with timeout
		ch := make(chan error, 1)
		go func() {
			_, err := hc.call(map[string]interface{}{"cmd": "ping"})
			ch <- err
		}()
		select {
		case err := <-ch:
			if err == nil {
				return hc, nil
			}
			// failed: check restart limit before restarting
			sbCfg, _, _ := readSandboxConfig()
			count := IncrementRestartCount(projectRoot)
			if sbCfg.MaxRestarts > 0 && count > int64(sbCfg.MaxRestarts) {
				return nil, fmt.Errorf("sandbox: helper restart limit exceeded (%d/%d)", count, sbCfg.MaxRestarts)
			}
			_ = hc.Stop()
			helpersMu.Lock()
			delete(helpers, projectRoot)
			helpersMu.Unlock()
		case <-time.After(300 * time.Millisecond):
			// timeout: assume helper dead; check restart limit
			sbCfg, _, _ := readSandboxConfig()
			count := IncrementRestartCount(projectRoot)
			if sbCfg.MaxRestarts > 0 && count > int64(sbCfg.MaxRestarts) {
				return nil, fmt.Errorf("sandbox: helper restart limit exceeded (%d/%d)", count, sbCfg.MaxRestarts)
			}
			_ = hc.Stop()
			helpersMu.Lock()
			delete(helpers, projectRoot)
			helpersMu.Unlock()
		}
	}
	// start new helper — on Windows use stdio mode, on POSIX use unix socket
	if runtime.GOOS == "windows" {
		// Windows: use stdin/stdout pipe mode (no unix sockets)
		newHc, err := startHelperOnce(helperPath)
		if err != nil {
			helpersMu.Lock()
			delete(helpers, projectRoot+":starting")
			helpersMu.Unlock()
			return nil, err
		}
		helpersMu.Lock()
		delete(helpers, projectRoot+":starting")
		helpers[projectRoot] = newHc
		helpersMu.Unlock()
		return newHc, nil
	}
	// POSIX: use unix socket mode
	sockPath := filepath.Join(projectRoot, ".fig", "ffi", "ffi.sock")
	// ensure parent dir exists
	_ = os.MkdirAll(filepath.Dir(sockPath), 0755)
	// serialize helper startup to avoid races where multiple goroutines try to start
	for {
		helpersMu.Lock()
		if hc, ok := helpers[projectRoot]; ok {
			helpersMu.Unlock()
			return hc, nil
		}
		// use a sentinel in the helpers map to indicate startup in progress
		if _, starting := helpers[projectRoot+":starting"]; starting {
			helpersMu.Unlock()
			// wait and retry
			time.Sleep(10 * time.Millisecond)
			continue
		}
		// mark as starting
		helpers[projectRoot+":starting"] = &helperClient{}
		helpersMu.Unlock()
		break
	}

	newHc, err := startHelperDaemon(helperPath, sockPath)
	if err != nil {
		// clear starting marker
		helpersMu.Lock()
		delete(helpers, projectRoot+":starting")
		helpersMu.Unlock()
		return nil, err
	}

	helpersMu.Lock()
	delete(helpers, projectRoot+":starting")
	helpers[projectRoot] = newHc
	helpersMu.Unlock()
	return newHc, nil
}

func startHelperDaemon(path string, sockPath string) (*helperClient, error) {
	// remove stale socket
	_ = os.Remove(sockPath)
	cmd := exec.Command(path, "--socket", sockPath)
	errp, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("cannot get stderr pipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("cannot start helper daemon: %v", err)
	}
	// capture stderr asynchronously
	var stderrBuf []byte
	go func() {
		sc := bufio.NewScanner(errp)
		for sc.Scan() {
			stderrBuf = append(stderrBuf, sc.Bytes()...)
			stderrBuf = append(stderrBuf, '\n')
		}
	}()
	// wait for socket to become connectable (retry loop)
	deadline := time.Now().Add(2 * time.Second)
	var conn net.Conn
	for {
		if time.Now().After(deadline) {
			_ = cmd.Process.Kill()
			return nil, fmt.Errorf("helper daemon did not create/connect socket in time; stderr: %s", string(stderrBuf))
		}
		c, err := net.Dial("unix", sockPath)
		if err == nil {
			conn = c
			break
		}
		// not ready yet
		time.Sleep(20 * time.Millisecond)
	}
	hc := &helperClient{cmd: cmd, stdin: nil, stdout: bufio.NewReader(conn), stderr: &stderrBuf, conn: conn, pending: map[string]chan map[string]interface{}{}}
	go hc.readLoop()
	return hc, nil
}

// StopAllHelpers stops and clears all running helper processes and resets FFI state.
func StopAllHelpers() {
	helpersMu.Lock()
	defer helpersMu.Unlock()
	for k, h := range helpers {
		_ = h.Stop()
		delete(helpers, k)
	}
	ResetFfiState()
	ResetSandboxCounters()
}

// HelperStderrFor returns the stderr buffer for the helper associated with projectRoot
// (useful for debugging in tests).
func HelperStderrFor(projectRoot string) string {
	helpersMu.Lock()
	defer helpersMu.Unlock()
	h, ok := helpers[projectRoot]
	if !ok || h == nil {
		return ""
	}
	return h.stderrString()
}

type helperClient struct {
	cmd         *exec.Cmd
	stdin       io.WriteCloser
	stdout      *bufio.Reader
	stderr      *[]byte
	conn        io.ReadWriteCloser
	pending     map[string]chan map[string]interface{}
	pmu         sync.Mutex
	writeMu     sync.Mutex
	callTimeout time.Duration // configurable call timeout, defaults to 3s
}

// sendRaw sends a raw JSON map to the helper (thread-safe)
func (h *helperClient) sendRaw(m map[string]interface{}) error {
	b, _ := json.Marshal(m)
	b = append(b, '\n')
	h.writeMu.Lock()
	defer h.writeMu.Unlock()
	if h.conn != nil {
		_, err := h.conn.Write(b)
		return err
	}
	if h.stdin != nil {
		_, err := h.stdin.Write(b)
		return err
	}
	return fmt.Errorf("no helper connection")
}

var nextReqID uint64

func (h *helperClient) stderrString() string {
	if h.stderr == nil {
		return ""
	}
	return string(*h.stderr)
}

// getCallTimeout returns the configured timeout or the 3s default
func (h *helperClient) getCallTimeout() time.Duration {
	if h.callTimeout > 0 {
		return h.callTimeout
	}
	return 3 * time.Second
}

func (h *helperClient) Stop() error {
	// best-effort kill
	if h.conn != nil {
		_ = h.conn.Close()
	}
	if h.cmd != nil && h.cmd.Process != nil {
		_ = h.cmd.Process.Kill()
		_ = h.cmd.Wait()
	}
	if h.stdin != nil {
		_ = h.stdin.Close()
	}
	return nil
}

func (h *helperClient) call(req map[string]interface{}) (map[string]interface{}, error) {
	id := strconv.FormatUint(atomic.AddUint64(&nextReqID, 1), 10)
	req["id"] = id
	ch := make(chan map[string]interface{}, 1)
	h.pmu.Lock()
	h.pending[id] = ch
	h.pmu.Unlock()
	defer func() {
		h.pmu.Lock()
		delete(h.pending, id)
		h.pmu.Unlock()
	}()

	b, _ := json.Marshal(req)
	b = append(b, '\n')
	if h.conn != nil {
		if _, err := h.conn.Write(b); err != nil {
			return nil, fmt.Errorf("write failed: %v; helper stderr: %s", err, h.stderrString())
		}
	} else {
		if h.stdin == nil {
			return nil, fmt.Errorf("helper stdin not available")
		}
		if _, err := h.stdin.Write(b); err != nil {
			return nil, fmt.Errorf("write failed: %v; helper stderr: %s", err, h.stderrString())
		}
	}

	select {
	case resp := <-ch:
		return resp, nil
	case <-time.After(h.getCallTimeout()):
		return nil, fmt.Errorf("timeout waiting for helper response; stderr: %s", h.stderrString())
	}
}

// Load asks the helper to load a library at path and returns a handle string.
func (h *helperClient) Load(path string) (string, error) {
	resp, err := h.call(map[string]interface{}{"cmd": "load", "path": path})
	if err != nil {
		return "", err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return "", fmt.Errorf("load failed: %v", resp["error"])
	}
	// handle is numeric; stringify it
	handle := resp["handle"]
	return fmt.Sprintf("%v", handle), nil
}

// Call asks the helper to call a symbol-less invocation (fallback) with args and returns the raw response.
func (h *helperClient) Call(args []interface{}) (interface{}, error) {
	resp, err := h.call(map[string]interface{}{"cmd": "call", "args": args})
	if err != nil {
		return nil, err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return nil, fmt.Errorf("call failed: %v", resp["error"])
	}
	return resp["resp"], nil
}

// Sym resolves a symbol name on a loaded handle and returns a symbol id string.
func (h *helperClient) Sym(handle string, name string, rtype string) (string, error) {
	hid, err := strconv.ParseInt(handle, 10, 64)
	if err != nil {
		return "", fmt.Errorf("invalid handle: %v", err)
	}
	resp, err := h.call(map[string]interface{}{"cmd": "sym", "handle": float64(hid), "name": name, "rtype": rtype})
	if err != nil {
		return "", err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return "", fmt.Errorf("sym failed: %v", resp["error"])
	}
	sym := resp["symbol"]
	return fmt.Sprintf("%v", sym), nil
}

// CallSymbol asks the helper to call a previously resolved symbol id with args.
func (h *helperClient) CallSymbol(symbol string, args []interface{}, argTypes []string) (interface{}, error) {
	sid, err := strconv.ParseInt(symbol, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid symbol id: %v", err)
	}
	req := map[string]interface{}{"cmd": "call", "symbol": float64(sid), "args": args}
	if len(argTypes) > 0 {
		req["arg_types"] = argTypes
	}
	resp, err := h.call(req)
	if err != nil {
		return nil, err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return nil, fmt.Errorf("call failed: %v", resp["error"])
	}
	return resp["resp"], nil
}

// Alloc allocates memory in the helper and returns a mem id
func (h *helperClient) Alloc(size int) (string, error) {
	resp, err := h.call(map[string]interface{}{"cmd": "alloc", "size": float64(size)})
	if err != nil {
		return "", err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return "", fmt.Errorf("alloc failed: %v", resp["error"])
	}
	id := fmt.Sprintf("%v", resp["mem_id"])
	return id, nil
}

// Strdup duplicates a Go string as a C-allocated string in the helper and returns a mem id
func (h *helperClient) Strdup(s string) (string, error) {
	resp, err := h.call(map[string]interface{}{"cmd": "strdup", "data": s})
	if err != nil {
		return "", err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return "", fmt.Errorf("strdup failed: %v", resp["error"])
	}
	id := fmt.Sprintf("%v", resp["mem_id"])
	return id, nil
}

// Free releases a previously allocated mem id
func (h *helperClient) Free(id string) error {
	resp, err := h.call(map[string]interface{}{"cmd": "free", "mem_id": id})
	if err != nil {
		return err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return fmt.Errorf("free failed: %v", resp["error"])
	}
	return nil
}

// MemWrite writes bytes (base64 encoded) to a helper memory id at offset
func (h *helperClient) MemWrite(id string, b64 string, offset int) error {
	req := map[string]interface{}{"cmd": "mem_write", "mem_id": id, "data": map[string]interface{}{"__bytes__": b64}, "offset": float64(offset)}
	resp, err := h.call(req)
	if err != nil {
		return err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return fmt.Errorf("mem_write failed: %v", resp["error"])
	}
	return nil
}

// MemRead reads bytes from helper memory id at offset with length bytes and returns base64 string
func (h *helperClient) MemRead(id string, offset int, length int) (string, error) {
	resp, err := h.call(map[string]interface{}{"cmd": "mem_read", "mem_id": id, "offset": float64(offset), "len": float64(length)})
	if err != nil {
		return "", err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return "", fmt.Errorf("mem_read failed: %v", resp["error"])
	}
	if m, ok := resp["resp"].(map[string]interface{}); ok {
		if b64, _ := m["__bytes__"].(string); b64 != "" {
			return b64, nil
		}
	}
	return "", fmt.Errorf("invalid mem_read response")
}
