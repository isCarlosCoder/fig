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
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// ffiLogLevel caches the parsed FFI_LOG_LEVEL for the client side.
// 0=error, 1=warn, 2=info, 3=debug. Default: warn (1).
var ffiLogLevel = func() int {
	switch strings.ToLower(os.Getenv("FFI_LOG_LEVEL")) {
	case "error":
		return 0
	case "info":
		return 2
	case "debug":
		return 3
	default:
		return 1 // warn
	}
}()

func ffiDebug(format string, args ...interface{}) {
	if ffiLogLevel >= 3 {
		fmt.Fprintf(os.Stderr, "ffi client: [DEBUG] "+format+"\n", args...)
	}
}

func ffiWarn(format string, args ...interface{}) {
	if ffiLogLevel >= 1 {
		fmt.Fprintf(os.Stderr, "ffi client: [WARN] "+format+"\n", args...)
	}
}

// FFIError represents a structured error from the FFI helper.
type FFIError struct {
	Code    string
	Message string
}

func (e *FFIError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// ffiParseError extracts a structured error from a helper response.
// It handles both the new format {"code":"ERR_...","message":"..."} and
// falls back to a plain string for backwards compatibility.
func ffiParseError(resp map[string]interface{}, prefix string) error {
	raw := resp["error"]
	if raw == nil {
		return fmt.Errorf("%s: unknown error", prefix)
	}
	if m, ok := raw.(map[string]interface{}); ok {
		code, _ := m["code"].(string)
		msg, _ := m["message"].(string)
		if code != "" {
			return &FFIError{Code: code, Message: msg}
		}
	}
	// fallback: plain string or Sprintf
	return fmt.Errorf("%s: %v", prefix, raw)
}

var helpersMu sync.Mutex
var helpers = map[string]*helperClient{}

// FFIProtocolVersion is the protocol version the client expects from the helper.
const FFIProtocolVersion = "1.0"

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
			ffiDebug("received invoke_callback id=%s cb=%s args=%v", id, cbid, args)
			// run callback
			res, err := runCallback(cbid, args)
			if err != nil {
				ffiWarn("callback error: %v", err)
				_ = h.sendRaw(map[string]interface{}{"ok": false, "error": err.Error(), "id": id})
				continue
			}
			ffiDebug("callback result for id=%s res=%v", id, res)
			_ = h.sendRaw(map[string]interface{}{"ok": true, "resp": res, "id": id})
			continue
		}
		// unknown incoming message; ignore
		ffiWarn("ignoring incoming message: %v", resp)
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
		if err := newHc.performHandshake(); err != nil {
			_ = newHc.Stop()
			helpersMu.Lock()
			delete(helpers, projectRoot)
			helpersMu.Unlock()
			return nil, err
		}
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
	if err := newHc.performHandshake(); err != nil {
		_ = newHc.Stop()
		helpersMu.Lock()
		delete(helpers, projectRoot)
		helpersMu.Unlock()
		return nil, err
	}
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

// performHandshake sends a handshake command to verify protocol version compatibility.
// Returns nil on success or if the helper doesn't support handshake (legacy).
func (h *helperClient) performHandshake() error {
	resp, err := h.call(map[string]interface{}{
		"cmd":     "handshake",
		"version": FFIProtocolVersion,
	})
	if err != nil {
		// If the call itself fails (e.g. timeout, connection error), propagate
		ffiWarn("handshake failed: %v", err)
		return nil // treat as legacy helper
	}
	if ok, _ := resp["ok"].(bool); !ok {
		// helper doesn't recognize "handshake" → legacy helper, warn and continue
		ffiWarn("helper does not support handshake (legacy mode)")
		return nil
	}
	// Extract version from result
	if result, ok := resp["result"].(map[string]interface{}); ok {
		helperVer, _ := result["version"].(string)
		if helperVer != "" && helperVer != FFIProtocolVersion {
			// Major version mismatch
			major := func(v string) string {
				if i := len(v); i > 0 {
					for j := 0; j < len(v); j++ {
						if v[j] == '.' {
							return v[:j]
						}
					}
				}
				return v
			}
			if major(helperVer) != major(FFIProtocolVersion) {
				return fmt.Errorf("FFI protocol version mismatch: client=%s, helper=%s", FFIProtocolVersion, helperVer)
			}
			ffiWarn("FFI protocol minor version difference: client=%s, helper=%s", FFIProtocolVersion, helperVer)
		}
		ffiDebug("handshake ok: helper version=%s", helperVer)
	}
	return nil
}

// Load asks the helper to load a library at path and returns a handle string.
func (h *helperClient) Load(path string) (string, error) {
	resp, err := h.call(map[string]interface{}{"cmd": "load", "path": path})
	if err != nil {
		return "", err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return "", ffiParseError(resp, "load")
	}
	// new envelope: result.handle; old: handle
	if result, ok := resp["result"].(map[string]interface{}); ok {
		if h, ok := result["handle"].(string); ok {
			return h, nil
		}
		if h, ok := result["handle"]; ok {
			return fmt.Sprintf("%v", h), nil
		}
	}
	if handle := resp["handle"]; handle != nil {
		return fmt.Sprintf("%v", handle), nil
	}
	return "", fmt.Errorf("load: missing handle in response")
}

// Call asks the helper to call a symbol-less invocation (fallback) with args and returns the raw response.
func (h *helperClient) Call(args []interface{}) (interface{}, error) {
	resp, err := h.call(map[string]interface{}{"cmd": "call", "args": args})
	if err != nil {
		return nil, err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return nil, ffiParseError(resp, "call")
	}
	// new envelope: result; old: resp
	if result, ok := resp["result"]; ok {
		return result, nil
	}
	return resp["resp"], nil
}

// Sym resolves a symbol name on a loaded handle and returns a symbol id string.
func (h *helperClient) Sym(handle string, name string, rtype string) (string, error) {
	resp, err := h.call(map[string]interface{}{"cmd": "sym", "handle": handle, "name": name, "rtype": rtype})
	if err != nil {
		return "", err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return "", ffiParseError(resp, "sym")
	}
	// new envelope: result.symbol; old: symbol
	if result, ok := resp["result"].(map[string]interface{}); ok {
		if s, ok := result["symbol"].(string); ok {
			return s, nil
		}
		if s, ok := result["symbol"]; ok {
			return fmt.Sprintf("%v", s), nil
		}
	}
	if sym := resp["symbol"]; sym != nil {
		return fmt.Sprintf("%v", sym), nil
	}
	return "", fmt.Errorf("sym: missing symbol in response")
}

// CallSymbol asks the helper to call a previously resolved symbol id with args.
func (h *helperClient) CallSymbol(symbol string, args []interface{}, argTypes []string) (interface{}, error) {
	req := map[string]interface{}{"cmd": "call", "symbol": symbol, "args": args}
	if len(argTypes) > 0 {
		req["arg_types"] = argTypes
	}
	resp, err := h.call(req)
	if err != nil {
		return nil, err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return nil, ffiParseError(resp, "call")
	}
	// new envelope: result; old: resp
	if result, ok := resp["result"]; ok {
		return result, nil
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
		return "", ffiParseError(resp, "alloc")
	}
	// new envelope: result.mem_id; old: mem_id
	if result, ok := resp["result"].(map[string]interface{}); ok {
		if mid, ok := result["mem_id"]; ok {
			return fmt.Sprintf("%v", mid), nil
		}
	}
	if mid := resp["mem_id"]; mid != nil {
		return fmt.Sprintf("%v", mid), nil
	}
	return "", fmt.Errorf("alloc: missing mem_id in response")
}

// Strdup duplicates a Go string as a C-allocated string in the helper and returns a mem id
func (h *helperClient) Strdup(s string) (string, error) {
	resp, err := h.call(map[string]interface{}{"cmd": "strdup", "data": s})
	if err != nil {
		return "", err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return "", ffiParseError(resp, "strdup")
	}
	// new envelope: result.mem_id; old: mem_id
	if result, ok := resp["result"].(map[string]interface{}); ok {
		if mid, ok := result["mem_id"]; ok {
			return fmt.Sprintf("%v", mid), nil
		}
	}
	if mid := resp["mem_id"]; mid != nil {
		return fmt.Sprintf("%v", mid), nil
	}
	return "", fmt.Errorf("strdup: missing mem_id in response")
}

// Free releases a previously allocated mem id
func (h *helperClient) Free(id string) error {
	resp, err := h.call(map[string]interface{}{"cmd": "free", "mem_id": id})
	if err != nil {
		return err
	}
	if ok, _ := resp["ok"].(bool); !ok {
		return ffiParseError(resp, "free")
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
		return ffiParseError(resp, "mem_write")
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
		return "", ffiParseError(resp, "mem_read")
	}
	// new envelope: result.__bytes__; old: resp.__bytes__
	if result, ok := resp["result"].(map[string]interface{}); ok {
		if b64, _ := result["__bytes__"].(string); b64 != "" {
			return b64, nil
		}
	}
	if m, ok := resp["resp"].(map[string]interface{}); ok {
		if b64, _ := m["__bytes__"].(string); b64 != "" {
			return b64, nil
		}
	}
	return "", fmt.Errorf("invalid mem_read response")
}

// HelperForTest is an exported wrapper around helperClient for use in tests.
type HelperForTest struct {
	hc *helperClient
}

func (h *HelperForTest) Load(path string) (string, error) { return h.hc.Load(path) }
func (h *HelperForTest) Free(id string) error             { return h.hc.Free(id) }
func (h *HelperForTest) Sym(handle, name, rtype string) (string, error) {
	return h.hc.Sym(handle, name, rtype)
}
func (h *HelperForTest) CallSymbol(sym string, args []interface{}, argTypes []string) (interface{}, error) {
	return h.hc.CallSymbol(sym, args, argTypes)
}

// GetHelperForTest returns a HelperForTest client for the given project root
// using the specified helper binary path. For unit tests only.
func GetHelperForTest(projectRoot, helperBin string) (*HelperForTest, error) {
	hc, err := getHelperForProject(projectRoot, helperBin)
	if err != nil {
		return nil, err
	}
	return &HelperForTest{hc: hc}, nil
}
