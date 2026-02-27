package main

/*
#cgo linux LDFLAGS: -ldl -rdynamic
#cgo darwin LDFLAGS: -ldl
#cgo windows LDFLAGS: -lkernel32
#include "dl_portable.h"
#include <stdlib.h>
#include <string.h>

// helper-exposed function that C libraries can call to invoke a registered Fig callback
// (call_cb_fn is implemented as a thin wrapper that forwards to the Go export call_cb_from_go)
char* call_cb_fn(const char* cbid, const char* arg);

// prototypes for wrapper functions provided in wrappers.c
// int wrappers
int ffi_call_int_fn0(void* fn);
int ffi_call_int_fn1(void* fn, int a);
int ffi_call_int_fn(void* fn, int a, int b);
int ffi_call_int_fn3(void* fn, int a, int b, int c);
// mixed-type int wrappers
int ffi_call_int_fn1_str(void* fn, char* a);
int ffi_call_int_fn2_str_str(void* fn, char* a, char* b);
int ffi_call_int_fn2_str_int(void* fn, char* a, int b);
int ffi_call_int_fn2_int_str(void* fn, int a, char* b);
int ffi_call_int_fn4_iisi(void* fn, int a, int b, char* c, int d);
// double wrappers
double ffi_call_double_fn0(void* fn);
double ffi_call_double_fn1(void* fn, double a);
double ffi_call_double_fn(void* fn, double a, double b);
// void wrappers
void ffi_call_void_fn0(void* fn);
void ffi_call_void_fn1_int(void* fn, int a);
void ffi_call_void_fn1_str(void* fn, char* a);
void ffi_call_void_fn2_str(void* fn, char* a, char* b);
// string wrappers
char* ffi_call_str_fn0(void* fn);
char* ffi_call_str_fn1(void* fn, char* a);
char* ffi_call_str_fn2(void* fn, char* a, char* b);
char* ffi_call_str_fn2_intint(void* fn, int a, int b);
char* ffi_call_str_fn3(void* fn, char* a, int b, double c);
char* ffi_call_str_fn3_intint(void* fn, char* a, int b, int c);
char* ffi_call_str_fn3_strs(void* fn, char* a, char* b, char* c);
char* ffi_call_str_fn4_sisi(void* fn, char* a, int b, char* c, int d);
char* ffi_call_str_fn4_ssss(void* fn, char* a, char* b, char* c, char* d);
char* ffi_call_str_fn4_siid(void* fn, char* a, int b, int c, double d);

// pointer-return wrappers (struct pointers are opaque to Fig)
void* ffi_call_ptr_fn0(void* fn);
void* ffi_call_ptr_fn1_int(void* fn, int a);
void* ffi_call_ptr_fn1_str(void* fn, char* a);
void* ffi_call_ptr_fn2_double_str(void* fn, double a, char* b);
// pointer argument helpers
int ffi_call_int_fn1_ptr(void* fn, void* p);
double ffi_call_double_fn1_ptr(void* fn, void* p);
char* ffi_call_str_fn1_ptr(void* fn, void* p);

*/
import "C"

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// LogLevel represents the severity of a log message.
type LogLevel int

const (
	LogError LogLevel = iota
	LogWarn
	LogInfo
	LogDebug
)

// Logger provides leveled logging to stderr.
type Logger struct {
	level  LogLevel
	prefix string
}

// NewLogger creates a Logger from a level string (error|warn|info|debug).
// Defaults to warn if unrecognized.
func NewLogger(level string) *Logger {
	l := &Logger{prefix: "ffi-helper"}
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "error":
		l.level = LogError
	case "warn", "warning":
		l.level = LogWarn
	case "info":
		l.level = LogInfo
	case "debug":
		l.level = LogDebug
	default:
		l.level = LogWarn
	}
	return l
}

func (l *Logger) Error(args ...interface{}) {
	if l.level >= LogError {
		fmt.Fprintln(os.Stderr, append([]interface{}{l.prefix + ": [ERROR]"}, args...)...)
	}
}

func (l *Logger) Warn(args ...interface{}) {
	if l.level >= LogWarn {
		fmt.Fprintln(os.Stderr, append([]interface{}{l.prefix + ": [WARN]"}, args...)...)
	}
}

func (l *Logger) Info(args ...interface{}) {
	if l.level >= LogInfo {
		fmt.Fprintln(os.Stderr, append([]interface{}{l.prefix + ": [INFO]"}, args...)...)
	}
}

func (l *Logger) Debug(args ...interface{}) {
	if l.level >= LogDebug {
		fmt.Fprintln(os.Stderr, append([]interface{}{l.prefix + ": [DEBUG]"}, args...)...)
	}
}

// log is the global logger, initialized in main().
var log = NewLogger("warn")

// FFI protocol version for handshake compatibility
const FFIProtocolVersion = "1.0"

// Error code constants for structured error responses
const (
	ErrInvalidJSON     = "ERR_INVALID_JSON"
	ErrUnknownCmd      = "ERR_UNKNOWN_CMD"
	ErrMissingParam    = "ERR_MISSING_PARAM"
	ErrInvalidHandle   = "ERR_INVALID_HANDLE"
	ErrInvalidSymbol   = "ERR_INVALID_SYMBOL"
	ErrDlopenFailed    = "ERR_DLOPEN_FAILED"
	ErrDlsymFailed     = "ERR_DLSYM_FAILED"
	ErrCallFailed      = "ERR_CALL_FAILED"
	ErrTypeError       = "ERR_TYPE_ERROR"
	ErrUnsupportedArgs = "ERR_UNSUPPORTED_ARGS"
	ErrMallocFailed    = "ERR_MALLOC_FAILED"
	ErrInvalidMemID    = "ERR_INVALID_MEM_ID"
	ErrOutOfBounds     = "ERR_OUT_OF_BOUNDS"
	ErrInvalidBase64   = "ERR_INVALID_BASE64"
	ErrVersionMismatch = "ERR_VERSION_MISMATCH"
)

// cbRequest used to send callback invocation requests from C into Fig via the helper
type cbRequest struct {
	cbid string
	args []interface{}
	resp chan map[string]interface{}
}

var cbReqCh = make(chan cbRequest, 16)
var cbPending = map[string]chan map[string]interface{}{}
var cbPendingMu sync.Mutex
var cbNext uint64

// memory managed by helper (for ownership semantics)
var memByID = map[string]unsafe.Pointer{}
var memSize = map[string]int{}
var memMu sync.Mutex
var memNext uint64

// safeFloat64 extracts a float64 from an interface{} value safely.
func safeFloat64(v interface{}) (float64, error) {
	switch n := v.(type) {
	case float64:
		return n, nil
	case int:
		return float64(n), nil
	case int64:
		return float64(n), nil
	case nil:
		return 0, fmt.Errorf("expected number, got nil")
	default:
		return 0, fmt.Errorf("expected number, got %T", v)
	}
}

// safeInt extracts an int from an interface{} value via float64 truncation.
func safeInt(v interface{}) (int, error) {
	f, err := safeFloat64(v)
	if err != nil {
		return 0, err
	}
	return int(f), nil
}

// safeString extracts a string from an interface{} value safely.
func safeString(v interface{}) (string, error) {
	switch s := v.(type) {
	case string:
		return s, nil
	case nil:
		return "", fmt.Errorf("expected string, got nil")
	default:
		return "", fmt.Errorf("expected string, got %T", v)
	}
}

func serve(r io.Reader, w io.Writer) error {
	// notify on stderr that server started
	log.Info("server started")
	defer func() {
		if r := recover(); r != nil {
			log.Error("panic:", r)
		}
	}()
	sc := bufio.NewScanner(r)
	enc := json.NewEncoder(w)
	var encMu sync.Mutex

	// request channel so we can read from the socket while a C call is blocking
	reqCh := make(chan map[string]interface{}, 16)

	// start a goroutine to read incoming messages and route responses to cbPending
	go func() {
		for sc.Scan() {
			var req map[string]interface{}
			if err := json.Unmarshal(sc.Bytes(), &req); err != nil {
				// log invalid JSON for debugging
				log.Warn("invalid json:", string(sc.Bytes()), "err:", err)
				resp := map[string]interface{}{"ok": false, "error": "invalid json"}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					log.Error("encode error:", err)
					encMu.Unlock()
					return
				}
				encMu.Unlock()
				continue
			}
			// if message is a response to helper-initiated request (has 'ok' field), route to pending
			if _, isResp := req["ok"]; isResp {
				if idRaw, ok := req["id"]; ok {
					id := fmt.Sprintf("%v", idRaw)
					cbPendingMu.Lock()
					ch, has := cbPending[id]
					cbPendingMu.Unlock()
					if has {
						log.Debug("got cb response id=", id, "resp=", req)
						select {
						case ch <- req:
						default:
						}
					}
				}
				continue
			}
			// otherwise, forward to the request processing goroutine
			reqCh <- req
		}
		close(reqCh)
	}()

	// callback request processing uses package-level channel and pending map
	go func() {
		for req := range cbReqCh {
			id := fmt.Sprintf("%d", atomic.AddUint64(&cbNext, 1))
			m := map[string]interface{}{"cmd": "invoke_callback", "cb": req.cbid, "args": req.args, "id": id}
			ch := make(chan map[string]interface{}, 1)
			cbPendingMu.Lock()
			cbPending[id] = ch
			cbPendingMu.Unlock()
			encMu.Lock()
			if err := enc.Encode(m); err != nil {
				log.Error("failed to send invoke_callback to fig:", err)
				ch <- map[string]interface{}{"ok": false, "error": err.Error()}
				encMu.Unlock()
			} else {
				log.Debug("req invoke_callback id=", id)
				encMu.Unlock()
			}
			select {
			case r := <-ch:
				log.Debug("cb goroutine received response id=", id, "resp=", r)
				req.resp <- r
			case <-time.After(3 * time.Second):
				req.resp <- map[string]interface{}{"ok": false, "error": "timeout"}
			}
			cbPendingMu.Lock()
			delete(cbPending, id)
			cbPendingMu.Unlock()
		}
	}()

	// in-memory state for load/sym/call
	var nextHandle uint64 = 0

	symbols := map[string]map[string]string{}
	var nextSymbol uint64 = 0
	ptrs := map[string]unsafe.Pointer{}
	symbolByID := map[string]unsafe.Pointer{}
	symbolTypeByID := map[string]string{}

	// storage for struct pointers returned by calls.  each id maps to a
	// C pointer which may later be passed back as an argument.  we protect
	// the map with a mutex since calls may happen concurrently.
	structPtrs := map[string]unsafe.Pointer{}
	var nextPtr uint64 = 0
	var structPtrsMu sync.Mutex

requestLoop:
	for req := range reqCh {
		cmd, _ := req["cmd"].(string)
		id := req["id"]
		log.Debug("req", cmd, "id=", id)

		// sendErr sends a structured error response for the current request
		sendErr := func(code string, msg string) {
			resp := map[string]interface{}{
				"ok": false,
				"error": map[string]interface{}{
					"code":    code,
					"message": msg,
				},
			}
			if id != nil {
				resp["id"] = id
			}
			encMu.Lock()
			enc.Encode(resp)
			encMu.Unlock()
		}

		// sendOK sends a success response with the given result payload
		sendOK := func(result interface{}) error {
			resp := map[string]interface{}{"ok": true, "result": result}
			if id != nil {
				resp["id"] = id
			}
			encMu.Lock()
			defer encMu.Unlock()
			if err := enc.Encode(resp); err != nil {
				log.Error("encode error:", err)
				return err
			}
			return nil
		}
		_ = sendErr // avoid unused warning in branches that don't use it

		switch cmd {
		case "ping":
			if err := sendOK("pong"); err != nil {
				return err
			}
		case "handshake":
			clientVer, _ := req["version"].(string)
			log.Info("handshake from client version=", clientVer)
			supportedOps := []string{
				"ping", "handshake", "load", "sym", "call", "alloc", "free",
				"strdup", "mem_write", "mem_read", "sleep", "crash",
			}
			result := map[string]interface{}{
				"version":       FFIProtocolVersion,
				"supported_ops": supportedOps,
			}
			if err := sendOK(result); err != nil {
				return err
			}
		case "sleep":
			msf, _ := req["ms"].(float64)
			ms := int(msf)
			log.Debug("sleeping", ms, "ms")
			time.Sleep(time.Duration(ms) * time.Millisecond)
			if err := sendOK("slept"); err != nil {
				return err
			}
		case "crash":
			log.Warn("crashing as requested")
			os.Exit(1)
		case "load":
			// load via dlopen
			p, _ := req["path"].(string)
			if p == "" {
				sendErr(ErrMissingParam, "missing path")
				continue
			}
			cp := C.CString(p)
			hdl := C.dl_open(cp)
			C.free(unsafe.Pointer(cp))
			if hdl == nil {
				errstr := C.GoString(C.dl_error())
				sendErr(ErrDlopenFailed, fmt.Sprintf("dlopen failed: %s", errstr))
				continue
			}
			nextHandle++
			handle := fmt.Sprintf("lib-%d", nextHandle)
			if symbols[handle] == nil {
				symbols[handle] = map[string]string{}
			}
			ptrs[handle] = unsafe.Pointer(hdl)
			if err := sendOK(map[string]interface{}{"handle": handle}); err != nil {
				return err
			}
		case "sym":
			hid, _ := req["handle"].(string)
			name, _ := req["name"].(string)
			rtype, _ := req["rtype"].(string)
			if rtype == "" {
				rtype = "int"
			}
			if _, ok := symbols[hid]; !ok {
				sendErr(ErrInvalidHandle, "invalid handle")
				continue
			}
			cname := C.CString(name)
			sym := C.dl_sym(C.dl_handle(ptrs[hid]), cname)
			C.free(unsafe.Pointer(cname))
			if sym == nil {
				errstr := C.GoString(C.dl_error())
				sendErr(ErrDlsymFailed, fmt.Sprintf("dlsym failed: %s", errstr))
				continue
			}
			nextSymbol++
			sid := fmt.Sprintf("sym-%d", nextSymbol)
			symbols[hid][name] = sid
			symbolByID[sid] = unsafe.Pointer(sym)
			symbolTypeByID[sid] = rtype
			if err := sendOK(map[string]interface{}{"symbol": sid}); err != nil {
				return err
			}
		case "call":
			// if a symbol id is provided, route to C function
			if symRaw, ok := req["symbol"]; ok {
				sid, sidOk := symRaw.(string)
				if !sidOk || sid == "" {
					sendErr(ErrTypeError, "type error: symbol must be a string")
					continue
				}
				ptr := symbolByID[sid]
				if ptr == nil {
					sendErr(ErrInvalidSymbol, "invalid symbol")
					continue
				}
				rtype := symbolTypeByID[sid]
				args, _ := req["args"].([]interface{})
				// struct returns: we allocate a handle for the returned C pointer and
				// hand back a lightweight marker to the Fig runtime.
				if strings.HasPrefix(rtype, "struct:") {
					name := strings.TrimPrefix(rtype, "struct:")
					var resPtr unsafe.Pointer
					switch len(args) {
					case 0:
						resPtr = C.ffi_call_ptr_fn0(ptr)
					case 1:
						// unify handling of numeric/string/pointer arguments
						if n, ok := args[0].(float64); ok {
							resPtr = C.ffi_call_ptr_fn1_int(ptr, C.int(int(n)))
						} else if s, ok := args[0].(string); ok {
							cs := C.CString(s)
							resPtr = C.ffi_call_ptr_fn1_str(ptr, cs)
							C.free(unsafe.Pointer(cs))
						} else if m, ok := args[0].(map[string]interface{}); ok {
							if pid, ok := m["__ptrid__"].(string); ok {
								structPtrsMu.Lock()
								p, has := structPtrs[pid]
								structPtrsMu.Unlock()
								if has {
									resPtr = p
								} else {
									// unknown pointer id, pass NULL
									resPtr = nil
								}
							}
						}
					case 2:
						// expect (double, string-like)
						a0, aErr := safeFloat64(args[0])
						if aErr != nil {
							resPtr = nil
							break
						}
						a := C.double(a0)
						// convert second arg using extract logic similar to string case
						var arg1 *C.char
						if m, ok := args[1].(map[string]interface{}); ok {
							if pid, ok := m["__ptrid__"].(string); ok {
								structPtrsMu.Lock()
								p, has := structPtrs[pid]
								structPtrsMu.Unlock()
								if has {
									arg1 = (*C.char)(p)
								}
							}
						}
						if arg1 == nil {
							// fallback to generic string conversion
							if s, ok := args[1].(string); ok {
								arg1 = C.CString(s)
								defer C.free(unsafe.Pointer(arg1))
							} else if n, ok := args[1].(float64); ok {
								s2 := fmt.Sprintf("%v", n)
								arg1 = C.CString(s2)
								defer C.free(unsafe.Pointer(arg1))
							} else {
								arg1 = C.CString("")
								defer C.free(unsafe.Pointer(arg1))
							}
						}
						resPtr = C.ffi_call_ptr_fn2_double_str(ptr, a, arg1)
					default:
						// for now we don't support >2 args for struct return
					}
					// assign id and remember pointer
					var idstr string
					structPtrsMu.Lock()
					nextPtr++
					idstr = fmt.Sprintf("p-%d", nextPtr)
					structPtrs[idstr] = resPtr
					structPtrsMu.Unlock()
					if err := sendOK(map[string]interface{}{"__struct__": name, "__ptrid__": idstr, "__rtype__": rtype}); err != nil {
						return err
					}
					continue requestLoop
				}
				switch rtype {
				case "int":
					// read arg_types metadata for mixed-type dispatch
					var intArgTypes []string
					if atRaw, ok := req["arg_types"].([]interface{}); ok {
						for _, v := range atRaw {
							if s, ok := v.(string); ok {
								intArgTypes = append(intArgTypes, s)
							}
						}
					}
					// helper: check if type is string-like
					isStrType := func(t string) bool {
						return t == "string" || t == "str"
					}
					isIntType := func(t string) bool {
						return t == "int" || t == "integer"
					}

					if len(args) == 0 {
						res := C.ffi_call_int_fn0(ptr)
						if err := sendOK(float64(res)); err != nil {
							return err
						}
					} else if len(args) == 1 {
						// pointer argument? use dedicated wrapper
						if m, ok := args[0].(map[string]interface{}); ok {
							if pid, ok := m["__ptrid__"].(string); ok {
								structPtrsMu.Lock()
								p, has := structPtrs[pid]
								structPtrsMu.Unlock()
								if has {
									res := C.ffi_call_int_fn1_ptr(ptr, p)
									if err := sendOK(float64(res)); err != nil {
										return err
									}
									continue
								}
							}
						}
						// check arg_types for mixed dispatch
						if len(intArgTypes) >= 1 && isStrType(intArgTypes[0]) {
							s0, ok := args[0].(string)
							if !ok {
								sendErr(ErrTypeError, fmt.Sprintf("type error: arg 0: expected string, got %T", args[0]))
								continue
							}
							cs := C.CString(s0)
							res := C.ffi_call_int_fn1_str(ptr, cs)
							C.free(unsafe.Pointer(cs))
							if err := sendOK(float64(res)); err != nil {
								return err
							}
						} else {
							a0, aErr := safeInt(args[0])
							if aErr != nil {
								sendErr(ErrTypeError, fmt.Sprintf("type error: arg 0: %v", aErr))
								continue
							}
							a := C.int(a0)
							res := C.ffi_call_int_fn1(ptr, a)
							if err := sendOK(float64(res)); err != nil {
								return err
							}
						}
					} else if len(args) == 2 {
						// build type signature from arg_types
						sig := ""
						if len(intArgTypes) >= 2 {
							for _, t := range intArgTypes {
								if isStrType(t) {
									sig += "s"
								} else if isIntType(t) {
									sig += "i"
								} else {
									sig += "i" // default to int
								}
							}
						} else {
							sig = "ii" // default: both int
						}
						switch sig {
						case "ss":
							s0, ok0 := args[0].(string)
							s1, ok1 := args[1].(string)
							if !ok0 || !ok1 {
								sendErr(ErrTypeError, "type error: expected string args for int(str,str) call")
								continue
							}
							cs0 := C.CString(s0)
							cs1 := C.CString(s1)
							res := C.ffi_call_int_fn2_str_str(ptr, cs0, cs1)
							C.free(unsafe.Pointer(cs0))
							C.free(unsafe.Pointer(cs1))
							if err := sendOK(float64(res)); err != nil {
								return err
							}
						case "si":
							s0, ok0 := args[0].(string)
							if !ok0 {
								sendErr(ErrTypeError, "type error: expected string for arg 0")
								continue
							}
							i1, iErr := safeInt(args[1])
							if iErr != nil {
								sendErr(ErrTypeError, fmt.Sprintf("type error: arg 1: %v", iErr))
								continue
							}
							cs0 := C.CString(s0)
							res := C.ffi_call_int_fn2_str_int(ptr, cs0, C.int(i1))
							C.free(unsafe.Pointer(cs0))
							if err := sendOK(float64(res)); err != nil {
								return err
							}
						case "is":
							i0, iErr := safeInt(args[0])
							if iErr != nil {
								sendErr(ErrTypeError, fmt.Sprintf("type error: arg 0: %v", iErr))
								continue
							}
							s1, ok1 := args[1].(string)
							if !ok1 {
								sendErr(ErrTypeError, "type error: expected string for arg 1")
								continue
							}
							cs1 := C.CString(s1)
							res := C.ffi_call_int_fn2_int_str(ptr, C.int(i0), cs1)
							C.free(unsafe.Pointer(cs1))
							if err := sendOK(float64(res)); err != nil {
								return err
							}
						default: // "ii"
							a0, aErr := safeInt(args[0])
							if aErr != nil {
								sendErr(ErrTypeError, fmt.Sprintf("type error: arg 0: %v", aErr))
								continue
							}
							b0, bErr := safeInt(args[1])
							if bErr != nil {
								sendErr(ErrTypeError, fmt.Sprintf("type error: arg 1: %v", bErr))
								continue
							}
							a := C.int(a0)
							b := C.int(b0)
							res := C.ffi_call_int_fn(ptr, a, b)
							if err := sendOK(float64(res)); err != nil {
								return err
							}
						}
					} else if len(args) == 3 {
						a0, aErr := safeInt(args[0])
						if aErr != nil {
							sendErr(ErrTypeError, fmt.Sprintf("type error: arg 0: %v", aErr))
							continue
						}
						b0, bErr := safeInt(args[1])
						if bErr != nil {
							sendErr(ErrTypeError, fmt.Sprintf("type error: arg 1: %v", bErr))
							continue
						}
						c0, cErr := safeInt(args[2])
						if cErr != nil {
							sendErr(ErrTypeError, fmt.Sprintf("type error: arg 2: %v", cErr))
							continue
						}
						a := C.int(a0)
						b := C.int(b0)
						c := C.int(c0)
						res := C.ffi_call_int_fn3(ptr, a, b, c)
						if err := sendOK(float64(res)); err != nil {
							return err
						}
					} else if len(args) == 4 {
						// build type signature from arg_types (expect iisi for bind_text)
						sig := ""
						if len(intArgTypes) >= 4 {
							for i := 0; i < 4; i++ {
								if isStrType(intArgTypes[i]) {
									sig += "s"
								} else if isIntType(intArgTypes[i]) {
									sig += "i"
								} else {
									sig += "i"
								}
							}
						}
						switch sig {
						case "iisi":
							i0, i0Err := safeInt(args[0])
							if i0Err != nil {
								sendErr(ErrTypeError, fmt.Sprintf("type error: arg 0: %v", i0Err))
								continue
							}
							i1, i1Err := safeInt(args[1])
							if i1Err != nil {
								sendErr(ErrTypeError, fmt.Sprintf("type error: arg 1: %v", i1Err))
								continue
							}
							s2, ok2 := args[2].(string)
							if !ok2 {
								sendErr(ErrTypeError, "type error: expected string for arg 2")
								continue
							}
							i3, i3Err := safeInt(args[3])
							if i3Err != nil {
								sendErr(ErrTypeError, fmt.Sprintf("type error: arg 3: %v", i3Err))
								continue
							}
							cs2 := C.CString(s2)
							res := C.ffi_call_int_fn4_iisi(ptr, C.int(i0), C.int(i1), cs2, C.int(i3))
							C.free(unsafe.Pointer(cs2))
							if err := sendOK(float64(res)); err != nil {
								return err
							}
						default:
							sendErr(ErrUnsupportedArgs, fmt.Sprintf("unsupported 4-arg int signature: %s", sig))
							continue
						}
					} else {
						sendErr(ErrUnsupportedArgs, fmt.Sprintf("int call with %d args not supported", len(args)))
					}
				case "double":
					if len(args) == 0 {
						res := C.ffi_call_double_fn0(ptr)
						if err := sendOK(float64(res)); err != nil {
							return err
						}
					} else if len(args) == 1 {
						// pointer argument support
						if m, ok := args[0].(map[string]interface{}); ok {
							if pid, ok := m["__ptrid__"].(string); ok {
								structPtrsMu.Lock()
								p, has := structPtrs[pid]
								structPtrsMu.Unlock()
								if has {
									res := C.ffi_call_double_fn1_ptr(ptr, p)
									if err := sendOK(float64(res)); err != nil {
										return err
									}
									continue
								}
							}
						}
						a0, aErr := safeFloat64(args[0])
						if aErr != nil {
							sendErr(ErrTypeError, fmt.Sprintf("type error: arg 0: %v", aErr))
							continue
						}
						a := C.double(a0)
						res := C.ffi_call_double_fn1(ptr, a)
						if err := sendOK(float64(res)); err != nil {
							return err
						}
					} else if len(args) == 2 {
						a0, aErr := safeFloat64(args[0])
						if aErr != nil {
							sendErr(ErrTypeError, fmt.Sprintf("type error: arg 0: %v", aErr))
							continue
						}
						b0, bErr := safeFloat64(args[1])
						if bErr != nil {
							sendErr(ErrTypeError, fmt.Sprintf("type error: arg 1: %v", bErr))
							continue
						}
						a := C.double(a0)
						b := C.double(b0)
						res := C.ffi_call_double_fn(ptr, a, b)
						if err := sendOK(float64(res)); err != nil {
							return err
						}
					} else {
						sendErr(ErrUnsupportedArgs, fmt.Sprintf("double call with %d args not supported", len(args)))
					}
				case "void":
					// void return type — call function and return nil
					if len(args) == 0 {
						C.ffi_call_void_fn0(ptr)
					} else if len(args) == 1 {
						// check arg type
						if n, ok := args[0].(float64); ok {
							C.ffi_call_void_fn1_int(ptr, C.int(int(n)))
						} else if s, ok := args[0].(string); ok {
							cs := C.CString(s)
							C.ffi_call_void_fn1_str(ptr, cs)
							C.free(unsafe.Pointer(cs))
						} else {
							sendErr(ErrUnsupportedArgs, "void: unsupported arg type")
							continue
						}
					} else if len(args) == 2 {
						s0, ok0 := args[0].(string)
						s1, ok1 := args[1].(string)
						if ok0 && ok1 {
							cs0 := C.CString(s0)
							cs1 := C.CString(s1)
							C.ffi_call_void_fn2_str(ptr, cs0, cs1)
							C.free(unsafe.Pointer(cs0))
							C.free(unsafe.Pointer(cs1))
						} else {
							sendErr(ErrUnsupportedArgs, "void: unsupported 2-arg combination")
							continue
						}
					} else {
						sendErr(ErrUnsupportedArgs, fmt.Sprintf("void call with %d args not supported", len(args)))
						continue
					}
					if err := sendOK(nil); err != nil {
						return err
					}
				case "string", "str":
					if err := func() error {
						var toFree []*C.char
						defer func() {
							for _, p := range toFree {
								C.free(unsafe.Pointer(p))
							}
						}()
						// helper to extract a C string from a value (string, callback, mem ptr, map→JSON,
						// or struct pointer marker)
						extract := func(x interface{}) (*C.char, error) {
							if s, ok := x.(string); ok {
								cs := C.CString(s)
								toFree = append(toFree, cs)
								return cs, nil
							}
							// accept numbers by converting to string (makes string-return
							// calls more forgiving and matches earlier semantics)
							if n, ok := x.(float64); ok {
								s := fmt.Sprintf("%v", n)
								cs := C.CString(s)
								toFree = append(toFree, cs)
								return cs, nil
							}
							if m, ok := x.(map[string]interface{}); ok {
								// pointer marker?
								if pid, ok := m["__ptrid__"].(string); ok {
									structPtrsMu.Lock()
									p, has := structPtrs[pid]
									structPtrsMu.Unlock()
									if has {
										return (*C.char)(p), nil
									}
									return nil, fmt.Errorf("unknown pointer id: %s", pid)
								}
								if cbid, ok := m["__cb__"].(string); ok {
									cs := C.CString(cbid)
									toFree = append(toFree, cs)
									return cs, nil
								}
								if mid, ok := m["__mem__"].(string); ok {
									memMu.Lock()
									p, has := memByID[mid]
									memMu.Unlock()
									if has && p != nil {
										return (*C.char)(p), nil
									}
									return nil, fmt.Errorf("invalid memory id: %s", mid)
								}
								// generic map → JSON string
								js, err := json.Marshal(m)
								if err != nil {
									return nil, fmt.Errorf("cannot marshal object to json: %v", err)
								}
								cs := C.CString(string(js))
								toFree = append(toFree, cs)
								return cs, nil
							}
							return nil, fmt.Errorf("expected string or callback object")
						}

						// read arg_types metadata if provided
						var argTypesList []string
						if atRaw, ok := req["arg_types"].([]interface{}); ok {
							for _, v := range atRaw {
								if s, ok := v.(string); ok {
									argTypesList = append(argTypesList, s)
								}
							}
						}

						var cres *C.char
						// dispatch based on arg count and types
						if len(args) == 0 {
							cres = C.ffi_call_str_fn0(ptr)
						} else {
							if len(args) == 1 {
								// extract arg0 (string-like)
								arg0Str, err := extract(args[0])
								if err != nil {
									sendErr(ErrTypeError, err.Error())
									return nil
								}
								cres = C.ffi_call_str_fn1(ptr, arg0Str)
							} else if len(args) == 2 {
								if len(argTypesList) == 2 {
									isInt0 := argTypesList[0] == "int" || argTypesList[0] == "integer"
									isInt1 := argTypesList[1] == "int" || argTypesList[1] == "integer"
									if isInt0 && isInt1 {
										n0, n0Err := safeInt(args[0])
										if n0Err != nil {
											sendErr(ErrTypeError, fmt.Sprintf("type error: arg 0: %v", n0Err))
											return nil
										}
										n1, n1Err := safeInt(args[1])
										if n1Err != nil {
											sendErr(ErrTypeError, fmt.Sprintf("type error: arg 1: %v", n1Err))
											return nil
										}
										cres = C.ffi_call_str_fn2_intint(ptr, C.int(n0), C.int(n1))
									} else {
										arg0Str, err := extract(args[0])
										if err != nil {
											sendErr(ErrTypeError, err.Error())
											return nil
										}
										a1Str, err := extract(args[1])
										if err != nil {
											sendErr(ErrTypeError, err.Error())
											return nil
										}
										cres = C.ffi_call_str_fn2(ptr, arg0Str, a1Str)
									}
								} else {
									arg0Str, err := extract(args[0])
									if err != nil {
										sendErr(ErrTypeError, err.Error())
										return nil
									}
									a1Str, err := extract(args[1])
									if err != nil {
										sendErr(ErrTypeError, err.Error())
										return nil
									}
									cres = C.ffi_call_str_fn2(ptr, arg0Str, a1Str)
								}
							} else if len(args) == 3 {
								// extract arg0 (string-like)
								arg0Str, err := extract(args[0])
								if err != nil {
									sendErr(ErrTypeError, err.Error())
									return nil
								}
								if len(argTypesList) == 3 {
									t1 := argTypesList[1]
									t2 := argTypesList[2]
									isInt1 := t1 == "int" || t1 == "integer"
									isInt2 := t2 == "int" || t2 == "integer"
									isNum2 := t2 == "number" || t2 == "double" || t2 == "float"
									if isInt1 && isInt2 {
										n1, n1Err := safeInt(args[1])
										if n1Err != nil {
											sendErr(ErrTypeError, fmt.Sprintf("type error: arg 1: %v", n1Err))
											return nil
										}
										n2, n2Err := safeInt(args[2])
										if n2Err != nil {
											sendErr(ErrTypeError, fmt.Sprintf("type error: arg 2: %v", n2Err))
											return nil
										}
										cres = C.ffi_call_str_fn3_intint(ptr, arg0Str, C.int(n1), C.int(n2))
									} else if isInt1 && isNum2 {
										n1, n1Err := safeInt(args[1])
										if n1Err != nil {
											sendErr(ErrTypeError, fmt.Sprintf("type error: arg 1: %v", n1Err))
											return nil
										}
										n2, n2Err := safeFloat64(args[2])
										if n2Err != nil {
											sendErr(ErrTypeError, fmt.Sprintf("type error: arg 2: %v", n2Err))
											return nil
										}
										cres = C.ffi_call_str_fn3(ptr, arg0Str, C.int(n1), C.double(n2))
									} else {
										a1Str, err := extract(args[1])
										if err != nil {
											sendErr(ErrTypeError, err.Error())
											return nil
										}
										a2Str, err := extract(args[2])
										if err != nil {
											sendErr(ErrTypeError, err.Error())
											return nil
										}
										cres = C.ffi_call_str_fn3_strs(ptr, arg0Str, a1Str, a2Str)
									}
								} else {
									_, isStr1 := args[1].(string)
									_, isMap1 := args[1].(map[string]interface{})
									_, isStr2 := args[2].(string)
									_, isMap2 := args[2].(map[string]interface{})
									n1, isNum1 := args[1].(float64)
									n2, isNum2 := args[2].(float64)
									if (isStr1 || isMap1) && (isStr2 || isMap2) {
										a1Str, err := extract(args[1])
										if err != nil {
											sendErr(ErrTypeError, err.Error())
											return nil
										}
										a2Str, err := extract(args[2])
										if err != nil {
											sendErr(ErrTypeError, err.Error())
											return nil
										}
										cres = C.ffi_call_str_fn3_strs(ptr, arg0Str, a1Str, a2Str)
									} else if isNum1 && isNum2 {
										cres = C.ffi_call_str_fn3_intint(ptr, arg0Str, C.int(int(n1)), C.int(int(n2)))
									} else {
										sendErr(ErrUnsupportedArgs, fmt.Sprintf("unsupported 3-arg string call: arg types %T, %T, %T", args[0], args[1], args[2]))
										return nil
									}
								}
							} else if len(args) == 4 {
								// extract arg0 (string-like)
								arg0Str, err := extract(args[0])
								if err != nil {
									sendErr(ErrTypeError, err.Error())
									return nil
								}
								// 4-arg string call: dispatch based on arg_types
								if len(argTypesList) == 4 {
									sig := ""
									for _, t := range argTypesList {
										switch {
										case t == "string" || t == "str":
											sig += "s"
										case t == "int" || t == "integer":
											sig += "i"
										case t == "double" || t == "float" || t == "number":
											sig += "d"
										default:
											sig += "?"
										}
									}
									switch sig {
									case "sisi":
										n1, n1Err := safeInt(args[1])
										if n1Err != nil {
											sendErr(ErrTypeError, fmt.Sprintf("type error: arg 1: %v", n1Err))
											return nil
										}
										a2Str, err := extract(args[2])
										if err != nil {
											sendErr(ErrTypeError, err.Error())
											return nil
										}
										n3, n3Err := safeInt(args[3])
										if n3Err != nil {
											sendErr(ErrTypeError, fmt.Sprintf("type error: arg 3: %v", n3Err))
											return nil
										}
										cres = C.ffi_call_str_fn4_sisi(ptr, arg0Str, C.int(n1), a2Str, C.int(n3))
									case "ssss":
										a1Str, err := extract(args[1])
										if err != nil {
											sendErr(ErrTypeError, err.Error())
											return nil
										}
										a2Str, err := extract(args[2])
										if err != nil {
											sendErr(ErrTypeError, err.Error())
											return nil
										}
										a3Str, err := extract(args[3])
										if err != nil {
											sendErr(ErrTypeError, err.Error())
											return nil
										}
										cres = C.ffi_call_str_fn4_ssss(ptr, arg0Str, a1Str, a2Str, a3Str)
									case "siid":
										n1, n1Err := safeInt(args[1])
										if n1Err != nil {
											sendErr(ErrTypeError, fmt.Sprintf("type error: arg 1: %v", n1Err))
											return nil
										}
										n2, n2Err := safeInt(args[2])
										if n2Err != nil {
											sendErr(ErrTypeError, fmt.Sprintf("type error: arg 2: %v", n2Err))
											return nil
										}
										n3, n3Err := safeFloat64(args[3])
										if n3Err != nil {
											sendErr(ErrTypeError, fmt.Sprintf("type error: arg 3: %v", n3Err))
											return nil
										}
										cres = C.ffi_call_str_fn4_siid(ptr, arg0Str, C.int(n1), C.int(n2), C.double(n3))
									default:
										sendErr(ErrUnsupportedArgs, fmt.Sprintf("unsupported 4-arg string signature: %s", sig))
										return nil
									}
								} else {
									sendErr(ErrUnsupportedArgs, "4-arg string call requires arg_types metadata")
									return nil
								}
							} else {
								sendErr(ErrUnsupportedArgs, fmt.Sprintf("string call with %d args not supported", len(args)))
								return nil
							}
						} // end of 0-arg vs 1+ arg dispatch
						if cres == nil {
							sendErr(ErrCallFailed, "call returned NULL")
							return nil
						}
						goStr := C.GoString(cres)
						C.free(unsafe.Pointer(cres))
						if err := sendOK(goStr); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				default:
					sendErr(ErrUnsupportedArgs, "unsupported symbol type")
				}
			} else {
				// fallback: sum/echo as before
				args, _ := req["args"].([]interface{})
				sum := float64(0)
				num := true
				for _, a := range args {
					switch v := a.(type) {
					case float64:
						sum += v
					default:
						num = false
					}
				}
				if num {
					if err := sendOK(sum); err != nil {
						return err
					}
				} else {
					if err := sendOK(args); err != nil {
						return err
					}
				}
			}
		case "alloc":
			szf, _ := req["size"].(float64)
			sz := int(szf)
			if sz <= 0 {
				sendErr(ErrMissingParam, "invalid alloc size")
				continue
			}
			ptr := C.malloc(C.size_t(sz))
			if ptr != nil {
				C.memset(ptr, 0, C.size_t(sz))
			}
			if ptr == nil {
				sendErr(ErrMallocFailed, "malloc failed")
				continue
			}
			mid := fmt.Sprintf("m-%d", atomic.AddUint64(&memNext, 1))
			memMu.Lock()
			memByID[mid] = ptr
			memSize[mid] = sz
			memMu.Unlock()
			if err := sendOK(map[string]interface{}{"mem_id": mid, "size": sz}); err != nil {
				return err
			}
		case "strdup":
			// strdup(data) -> allocates C string copy, returns mem_id
			s, _ := req["data"].(string)
			if s == "" {
				sendErr(ErrMissingParam, "strdup requires non-empty data string")
				continue
			}
			cs := C.CString(s)
			sz := len(s) + 1 // includes NUL terminator
			mid := fmt.Sprintf("m-%d", atomic.AddUint64(&memNext, 1))
			memMu.Lock()
			memByID[mid] = unsafe.Pointer(cs)
			memSize[mid] = sz
			memMu.Unlock()
			if err := sendOK(map[string]interface{}{"mem_id": mid, "size": sz}); err != nil {
				return err
			}
		case "free":
			mid, _ := req["mem_id"].(string)
			if mid == "" {
				// fallback: try "id" for backwards compatibility
				mid, _ = req["id"].(string)
			}
			if mid == "" {
				sendErr(ErrMissingParam, "free requires id")
				continue
			}
			memMu.Lock()
			ptr, ok := memByID[mid]
			if ok {
				C.free(ptr)
				delete(memByID, mid)
				delete(memSize, mid)
			}
			memMu.Unlock()
			if !ok {
				sendErr(ErrInvalidMemID, "invalid mem id")
				continue
			}
			if err := sendOK(nil); err != nil {
				return err
			}
		case "mem_write":
			mid, _ := req["mem_id"].(string)
			log.Debug("mem_write requested mem_id=", mid)
			data, _ := req["data"].(map[string]interface{})
			b64, _ := data["__bytes__"].(string)
			offf, _ := req["offset"].(float64)
			off := int(offf)
			if mid == "" || b64 == "" {
				sendErr(ErrMissingParam, "mem_write requires id and data")
				continue
			}
			raw, err := base64.StdEncoding.DecodeString(b64)
			if err != nil {
				sendErr(ErrInvalidBase64, "invalid base64")
				continue
			}
			memMu.Lock()
			ptr, has := memByID[mid]
			sz, _ := memSize[mid]
			memMu.Unlock()
			if !has || ptr == nil {
				sendErr(ErrInvalidMemID, "invalid mem id")
				continue
			}
			if off < 0 || off+len(raw) > sz {
				sendErr(ErrOutOfBounds, "write out of bounds")
				continue
			}
			C.memcpy(unsafe.Pointer(uintptr(ptr)+uintptr(off)), unsafe.Pointer(&raw[0]), C.size_t(len(raw)))
			if err := sendOK(nil); err != nil {
				return err
			}
		case "mem_read":
			mid, _ := req["mem_id"].(string)
			offf, _ := req["offset"].(float64)
			lnf, _ := req["len"].(float64)
			off := int(offf)
			ln := int(lnf)
			memMu.Lock()
			ptr, has := memByID[mid]
			sz, _ := memSize[mid]
			memMu.Unlock()
			if !has || ptr == nil {
				sendErr(ErrInvalidMemID, "invalid mem id")
				continue
			}
			if off < 0 || off+ln > sz {
				sendErr(ErrOutOfBounds, "read out of bounds")
				continue
			}
			buf := C.malloc(C.size_t(ln))
			C.memcpy(buf, unsafe.Pointer(uintptr(ptr)+uintptr(off)), C.size_t(ln))
			b := C.GoBytes(buf, C.int(ln))
			C.free(buf)
			encB64 := base64.StdEncoding.EncodeToString(b)
			if err := sendOK(map[string]interface{}{"__bytes__": encB64}); err != nil {
				return err
			}
		default:
			sendErr(ErrUnknownCmd, "unknown command")
		}
	}
	return sc.Err()
}

func serveSocket(sockPath string) error {
	// ensure parent exists
	_ = os.Remove(sockPath)
	ln, err := net.Listen("unix", sockPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = ln.Close()
		_ = os.Remove(sockPath)
	}()
	log.Info("socket server listening", sockPath)
	conn, err := ln.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()
	return serve(conn, conn)
}

//export call_cb_from_go
func call_cb_from_go(cbid *C.char, carg *C.char) *C.char {
	id := C.GoString(cbid)
	arg := ""
	if carg != nil {
		arg = C.GoString(carg)
	}
	rch := make(chan map[string]interface{}, 1)
	log.Debug("call_cb_from_go received cb=", id, "arg=", arg)
	cbReqCh <- cbRequest{cbid: id, args: []interface{}{arg}, resp: rch}
	log.Debug("call_cb_from_go queued request for cb=", id)

	var resp map[string]interface{}
	select {
	case resp = <-rch:
		log.Debug("call_cb_from_go got response for cb=", id, "resp=", resp)
	case <-time.After(5 * time.Second):
		log.Warn("call_cb_from_go timeout for cb=", id)
		return nil
	}

	if ok, _ := resp["ok"].(bool); !ok {
		return nil
	}
	if s, ok := resp["resp"].(string); ok {
		return C.CString(s)
	}
	b, _ := json.Marshal(resp["resp"])
	return C.CString(string(b))
}

func main() {
	server := flag.Bool("server", false, "run as server reading JSON lines from stdin and responding")
	sock := flag.String("socket", "", "path to unix socket to listen on")
	logLevel := flag.String("log-level", "", "log level: error|warn|info|debug (default: warn, or FFI_LOG_LEVEL env)")
	flag.Parse()

	// Resolve log level: flag > env > default (warn)
	lvl := *logLevel
	if lvl == "" {
		lvl = os.Getenv("FFI_LOG_LEVEL")
	}
	if lvl == "" {
		lvl = "warn"
	}
	log = NewLogger(lvl)
	if *server {
		if err := serve(os.Stdin, os.Stdout); err != nil {
			log.Error("server error:", err)
			os.Exit(1)
		}
		return
	}
	if *sock != "" {
		if err := serveSocket(*sock); err != nil {
			log.Error("socket server error:", err)
			os.Exit(1)
		}
		return
	}
	fmt.Println("ffi-helper: stub helper (run with --server or --socket to act as RPC helper)")
}
