package builtins

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/iscarloscoder/fig/environment"
)

// httpRoute stores a registered route handler.
type httpRoute struct {
	method  string
	pattern string
	handler environment.Value // FunctionType value
}

// figResponse wraps an http.ResponseWriter with status tracking.
type figResponse struct {
	w          http.ResponseWriter
	statusCode int
	sent       bool
}

// httpState holds global config for the http module.
var httpState struct {
	timeout        time.Duration
	defaultHeaders map[string]string
	routes         []httpRoute
	mu             sync.Mutex
}

// FigCaller is set by the interpreter to call Fig functions from HTTP handlers.
// Returns any runtime error that occurred during the function call.
var FigCaller func(fn environment.Value, args []environment.Value) error

// httpServer holds the current server instance for shutdown.
var httpServer *http.Server

// ServerAddr holds the actual address the server is listening on (useful when port 0 is used).
var ServerAddr string

// ShutdownServer stops the running HTTP server.
func ShutdownServer() {
	if httpServer != nil {
		httpServer.Close()
		httpServer = nil
	}
}

// ResetHTTPState resets all HTTP module state (for testing).
func ResetHTTPState() {
	httpState.mu.Lock()
	defer httpState.mu.Unlock()
	httpState.timeout = 30 * time.Second
	httpState.defaultHeaders = map[string]string{}
	httpState.routes = nil
	ServerAddr = ""
	ShutdownServer()
}

func init() {
	httpState.timeout = 30 * time.Second
	httpState.defaultHeaders = map[string]string{}

	register(newModule("http",
		// ── core ──

		// request(method, url, body, headers) → {status, headers, body}
		fn("request", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 2 || len(args) > 4 {
				return environment.NewNil(), fmt.Errorf("request() expects 2-4 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("request() method must be a string")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("request() url must be a string")
			}

			method := strings.ToUpper(args[0].Str)
			url := args[1].Str

			// body (optional, 3rd arg)
			var bodyReader io.Reader
			if len(args) >= 3 && args[2].Type == environment.StringType {
				bodyReader = strings.NewReader(args[2].Str)
			}

			req, err := http.NewRequest(method, url, bodyReader)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("request() error creating request: %v", err)
			}

			// apply default headers first
			for k, v := range httpState.defaultHeaders {
				req.Header.Set(k, v)
			}

			// apply per-request headers (4th arg, object)
			if len(args) >= 4 && args[3].Type == environment.ObjectType {
				obj := args[3].Obj
				for _, k := range obj.Keys {
					v := obj.Entries[k]
					if v.Type == environment.StringType {
						req.Header.Set(k, v.Str)
					}
				}
			}

			client := &http.Client{Timeout: httpState.timeout}
			resp, err := client.Do(req)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("request() error: %v", err)
			}
			defer resp.Body.Close()

			return buildResponse(resp)
		}),

		// ── shortcuts ──

		// get(url) → {status, headers, body}
		fn("get", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("get() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("get() url must be a string")
			}

			req, err := http.NewRequest("GET", args[0].Str, nil)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("get() error: %v", err)
			}
			for k, v := range httpState.defaultHeaders {
				req.Header.Set(k, v)
			}

			client := &http.Client{Timeout: httpState.timeout}
			resp, err := client.Do(req)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("get() error: %v", err)
			}
			defer resp.Body.Close()

			return buildResponse(resp)
		}),

		// post(url, body) → {status, headers, body}
		fn("post", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 || len(args) > 2 {
				return environment.NewNil(), fmt.Errorf("post() expects 1-2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("post() url must be a string")
			}

			var bodyReader io.Reader
			if len(args) == 2 && args[1].Type == environment.StringType {
				bodyReader = strings.NewReader(args[1].Str)
			}

			req, err := http.NewRequest("POST", args[0].Str, bodyReader)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("post() error: %v", err)
			}
			for k, v := range httpState.defaultHeaders {
				req.Header.Set(k, v)
			}

			client := &http.Client{Timeout: httpState.timeout}
			resp, err := client.Do(req)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("post() error: %v", err)
			}
			defer resp.Body.Close()

			return buildResponse(resp)
		}),

		// ── utilities ──

		// download(url, path) → nil
		fn("download", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("download() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("download() url must be a string")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("download() path must be a string")
			}

			req, err := http.NewRequest("GET", args[0].Str, nil)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("download() error: %v", err)
			}
			for k, v := range httpState.defaultHeaders {
				req.Header.Set(k, v)
			}

			client := &http.Client{Timeout: httpState.timeout}
			resp, err := client.Do(req)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("download() error: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode >= 400 {
				return environment.NewNil(), fmt.Errorf("download() server returned status %d", resp.StatusCode)
			}

			file, err := os.Create(args[1].Str)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("download() cannot create file: %v", err)
			}
			defer file.Close()

			_, err = io.Copy(file, resp.Body)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("download() error writing file: %v", err)
			}

			return environment.NewNil(), nil
		}),

		// setTimeout(ms) → nil
		fn("setTimeout", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("setTimeout() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("setTimeout() argument must be a number")
			}
			httpState.timeout = time.Duration(args[0].Num) * time.Millisecond
			return environment.NewNil(), nil
		}),

		// setHeader(key, value) → nil
		fn("setHeader", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("setHeader() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("setHeader() key must be a string")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("setHeader() value must be a string")
			}
			httpState.defaultHeaders[args[0].Str] = args[1].Str
			return environment.NewNil(), nil
		}),

		// clearHeaders() → nil
		fn("clearHeaders", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("clearHeaders() expects 0 arguments, got %d", len(args))
			}
			httpState.defaultHeaders = map[string]string{}
			return environment.NewNil(), nil
		}),

		// ── response helpers ──

		// isOk(res) → bool (status 200-299)
		fn("isOk", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isOk() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ObjectType {
				return environment.NewNil(), fmt.Errorf("isOk() argument must be a response object")
			}
			statusVal, ok := args[0].Obj.Entries["status"]
			if !ok || statusVal.Type != environment.NumberType {
				return environment.NewBool(false), nil
			}
			code := int(statusVal.Num)
			return environment.NewBool(code >= 200 && code <= 299), nil
		}),

		// raiseForStatus(res) → nil or error
		fn("raiseForStatus", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("raiseForStatus() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ObjectType {
				return environment.NewNil(), fmt.Errorf("raiseForStatus() argument must be a response object")
			}
			statusVal, ok := args[0].Obj.Entries["status"]
			if !ok || statusVal.Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("raiseForStatus() response has no valid status")
			}
			code := int(statusVal.Num)
			if code >= 400 {
				return environment.NewNil(), fmt.Errorf("HTTP error: status %d", code)
			}
			return environment.NewNil(), nil
		}),

		// ── server ──

		// route(method, path, handler) → nil
		fn("route", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("route() expects 3 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("route() method must be a string")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("route() path must be a string")
			}
			if args[2].Type != environment.FunctionType {
				return environment.NewNil(), fmt.Errorf("route() handler must be a function")
			}
			httpState.routes = append(httpState.routes, httpRoute{
				method:  strings.ToUpper(args[0].Str),
				pattern: args[1].Str,
				handler: args[2],
			})
			return environment.NewNil(), nil
		}),

		// listen(port [, onStart]) → blocks serving HTTP. Optional `onStart` is a fn() called during initialization.
		fn("listen", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 || len(args) > 2 {
				return environment.NewNil(), fmt.Errorf("listen() expects 1 or 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("listen() port must be a number")
			}
			if FigCaller == nil {
				return environment.NewNil(), fmt.Errorf("listen() requires the http module to be loaded via 'use'")
			}

			port := int(args[0].Num)
			addr := fmt.Sprintf(":%d", port)

			// optional onStart callback: must be a function
			if len(args) == 2 {
				if args[1].Type != environment.FunctionType {
					return environment.NewNil(), fmt.Errorf("listen() onStart must be a function")
				}
				// call the callback so it can perform initialization (e.g., register routes)
				if err := FigCaller(args[1], nil); err != nil {
					return environment.NewNil(), fmt.Errorf("listen() onStart callback error: %v", err)
				}
			}

			mux := buildServerMux()

			ln, err := net.Listen("tcp", addr)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("listen() error: %v", err)
			}
			ServerAddr = ln.Addr().String()

			httpServer = &http.Server{Handler: mux}
			err = httpServer.Serve(ln)
			if err != nil && err != http.ErrServerClosed {
				return environment.NewNil(), fmt.Errorf("listen() error: %v", err)
			}
			return environment.NewNil(), nil
		}),

		// render(path, data) → string (HTML template rendering)
		fn("render", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 || len(args) > 2 {
				return environment.NewNil(), fmt.Errorf("render() expects 1-2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("render() path must be a string")
			}
			data, err := os.ReadFile(args[0].Str)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("render() cannot read template: %v", err)
			}
			html := string(data)
			if len(args) == 2 && args[1].Type == environment.ObjectType {
				for _, k := range args[1].Obj.Keys {
					re := regexp.MustCompile(`\{\{\s*` + regexp.QuoteMeta(k) + `\s*\}\}`)
					val := args[1].Obj.Entries[k]
					var replacement string
					if val.Type == environment.StringType {
						replacement = val.Str
					} else {
						replacement = val.String()
					}
					html = re.ReplaceAllString(html, replacement)
				}
			}
			return environment.NewString(html), nil
		}),
	))
}

// buildResponse converts an *http.Response into a Fig object {status, headers, body}.
func buildResponse(resp *http.Response) (environment.Value, error) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return environment.NewNil(), fmt.Errorf("error reading response body: %v", err)
	}

	// Build headers as a Fig object (only first value per header key)
	hdrEntries := make(map[string]environment.Value)
	hdrKeys := make([]string, 0, len(resp.Header))
	for k := range resp.Header {
		lk := strings.ToLower(k)
		hdrEntries[lk] = environment.NewString(resp.Header.Get(k))
		hdrKeys = append(hdrKeys, lk)
	}

	// Build the response object
	entries := map[string]environment.Value{
		"status":  environment.NewNumber(float64(resp.StatusCode)),
		"headers": environment.NewObject(hdrEntries, hdrKeys),
		"body":    environment.NewString(string(bodyBytes)),
	}
	keys := []string{"status", "headers", "body"}

	return environment.NewObject(entries, keys), nil
}

// ── server helpers ──

// buildServerMux creates an http.ServeMux from registered routes.
func buildServerMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Group routes by path pattern
	type routeGroup struct {
		routes []httpRoute
		order  int
	}
	groups := map[string]*routeGroup{}
	idx := 0
	for _, r := range httpState.routes {
		g, ok := groups[r.pattern]
		if !ok {
			g = &routeGroup{order: idx}
			groups[r.pattern] = g
			idx++
		}
		g.routes = append(g.routes, r)
	}

	for path, group := range groups {
		rs := group.routes // capture for closure
		mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			for _, r := range rs {
				if r.method == req.Method || r.method == "*" {
					handleFigRequest(w, req, r.handler)
					return
				}
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method Not Allowed")
		})
	}

	return mux
}

// handleFigRequest processes an incoming HTTP request using a Fig handler function.
func handleFigRequest(w http.ResponseWriter, r *http.Request, handler environment.Value) {
	if FigCaller == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error: Fig caller not set")
		return
	}

	reqObj := buildReqObject(r)
	fr := &figResponse{w: w, statusCode: 200, sent: false}
	resObj := buildResObject(fr)

	httpState.mu.Lock()
	err := FigCaller(handler, []environment.Value{reqObj, resObj})
	httpState.mu.Unlock()

	if !fr.sent {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal Server Error")
		} else {
			w.WriteHeader(fr.statusCode)
		}
	}
}

// buildReqObject creates a Fig object representing the HTTP request.
func buildReqObject(r *http.Request) environment.Value {
	bodyBytes, _ := io.ReadAll(r.Body)

	// Request headers
	hdrEntries := make(map[string]environment.Value)
	hdrKeys := make([]string, 0, len(r.Header))
	for k := range r.Header {
		lk := strings.ToLower(k)
		hdrEntries[lk] = environment.NewString(r.Header.Get(k))
		hdrKeys = append(hdrKeys, lk)
	}

	// Query parameters
	qEntries := make(map[string]environment.Value)
	qKeys := make([]string, 0, len(r.URL.Query()))
	for k, v := range r.URL.Query() {
		qEntries[k] = environment.NewString(v[0])
		qKeys = append(qKeys, k)
	}

	entries := map[string]environment.Value{
		"path":    environment.NewString(r.URL.Path),
		"method":  environment.NewString(r.Method),
		"body":    environment.NewString(string(bodyBytes)),
		"headers": environment.NewObject(hdrEntries, hdrKeys),
		"query":   environment.NewObject(qEntries, qKeys),
	}
	keys := []string{"path", "method", "body", "headers", "query"}
	return environment.NewObject(entries, keys)
}

// buildResObject creates a Fig object with send(), json(), status() methods.
func buildResObject(fr *figResponse) environment.Value {
	entries := map[string]environment.Value{
		// send(body) — writes text/html response
		"send": environment.NewBuiltinFn("send", func(args []environment.Value) (environment.Value, error) {
			if fr.sent {
				return environment.NewNil(), nil
			}
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("send() expects 1 argument, got %d", len(args))
			}
			fr.sent = true
			fr.w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fr.w.WriteHeader(fr.statusCode)
			if args[0].Type == environment.StringType {
				fmt.Fprint(fr.w, args[0].Str)
			} else {
				fmt.Fprint(fr.w, args[0].String())
			}
			return environment.NewNil(), nil
		}),

		// json(obj) — writes JSON response
		"json": environment.NewBuiltinFn("json", func(args []environment.Value) (environment.Value, error) {
			if fr.sent {
				return environment.NewNil(), nil
			}
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("json() expects 1 argument, got %d", len(args))
			}
			fr.sent = true
			goVal := figToGo(args[0])
			data, err := json.Marshal(goVal)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("json() error: %v", err)
			}
			fr.w.Header().Set("Content-Type", "application/json")
			fr.w.WriteHeader(fr.statusCode)
			fr.w.Write(data)
			return environment.NewNil(), nil
		}),

		// status(code) — sets the HTTP status code
		"status": environment.NewBuiltinFn("status", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("status() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("status() argument must be a number")
			}
			fr.statusCode = int(args[0].Num)
			return environment.NewNil(), nil
		}),
	}
	keys := []string{"send", "json", "status"}
	return environment.NewObject(entries, keys)
}
