package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/iscarloscoder/fig/builtins"
)

// waitForServer waits until a TCP server is accepting connections.
func waitForServer(addr string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", addr, 100*time.Millisecond)
		if err == nil {
			conn.Close()
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return fmt.Errorf("server at %s did not start within %v", addr, timeout)
}

func TestHTTPRoute(t *testing.T) {
	src := `use "http"
http.route("GET", "/", fn(req, res) { res.send("ok") })
print("routed")`

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "routed" {
		t.Errorf("expected routed, got %q", out)
	}
}

func TestHTTPRouteArgError(t *testing.T) {
	src := `use "http"
http.route("GET")`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for route() with wrong args")
	}
}

func TestHTTPRouteHandlerNotFunction(t *testing.T) {
	src := `use "http"
http.route("GET", "/", "not-a-fn")`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for route() with non-function handler")
	}
}

func TestHTTPListenArgError(t *testing.T) {
	src := `use "http"
http.listen("bad")`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for listen() with string arg")
	}
}

func TestHTTPRender(t *testing.T) {
	tmpDir := t.TempDir()
	tmplPath := tmpDir + "/test.html"
	os.WriteFile(tmplPath, []byte("<h1>Hello {{name}}! You are {{age}} years old.</h1>"), 0644)

	src := fmt.Sprintf(`use "http"
let html = http.render("%s", {"name": "Ana", "age": 25})
print(html)`, tmplPath)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "<h1>Hello Ana! You are 25 years old.</h1>" {
		t.Errorf("unexpected render output: %q", out)
	}
}

func TestHTTPRenderNoData(t *testing.T) {
	tmpDir := t.TempDir()
	tmplPath := tmpDir + "/test.html"
	os.WriteFile(tmplPath, []byte("<p>static</p>"), 0644)

	src := fmt.Sprintf(`use "http"
let html = http.render("%s")
print(html)`, tmplPath)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "<p>static</p>" {
		t.Errorf("expected <p>static</p>, got %q", out)
	}
}

func TestHTTPRenderFileNotFound(t *testing.T) {
	src := `use "http"
http.render("/nonexistent/file.html")`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for missing template")
	}
}

func TestHTTPRenderArgError(t *testing.T) {
	src := `use "http"
http.render(123)`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for render with number arg")
	}
}

// ── Server integration tests ──

func TestHTTPServerSendText(t *testing.T) {
	builtins.ResetHTTPState()

	src := `use "http"
http.route("GET", "/", fn(req, res) {
	res.send("hello from fig")
})
http.listen(0, fn() {
	http.route("GET", "/cb", fn(req, res) { res.send("cb-ok") })
})`

	// start server and verify route added before and by onStart callback
	go runFig(t, src)

	// Wait for server to start and get its address
	time.Sleep(200 * time.Millisecond)
	addr := builtins.ServerAddr
	if addr == "" {
		t.Fatal("server did not start")
	}
	if err := waitForServer(addr, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer builtins.ShutdownServer()

	// existing route
	resp, err := http.Get("http://" + addr + "/")
	if err != nil {
		t.Fatalf("GET / error: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if string(body) != "hello from fig" {
		t.Errorf("expected 'hello from fig', got %q", string(body))
	}
	if ct := resp.Header.Get("Content-Type"); !strings.Contains(ct, "text/html") {
		t.Errorf("expected text/html content-type, got %q", ct)
	}

	// route added by callback
	resp2, err := http.Get("http://" + addr + "/cb")
	if err != nil {
		t.Fatalf("GET /cb error: %v", err)
	}
	defer resp2.Body.Close()
	body2, _ := io.ReadAll(resp2.Body)
	if resp2.StatusCode != 200 || string(body2) != "cb-ok" {
		t.Fatalf("expected cb-ok from callback-added route, got status %d body %q", resp2.StatusCode, string(body2))
	}
}

func TestHTTPServerJSON(t *testing.T) {
	builtins.ResetHTTPState()

	src := `use "http"
http.route("GET", "/data", fn(req, res) {
	res.json({"name": "fig", "version": 1})
})
http.listen(0)`

	go runFig(t, src)
	time.Sleep(200 * time.Millisecond)
	addr := builtins.ServerAddr
	if addr == "" {
		t.Fatal("server did not start")
	}
	if err := waitForServer(addr, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer builtins.ShutdownServer()

	resp, err := http.Get("http://" + addr + "/data")
	if err != nil {
		t.Fatalf("GET /data error: %v", err)
	}
	defer resp.Body.Close()

	if ct := resp.Header.Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Errorf("expected application/json, got %q", ct)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	if result["name"] != "fig" {
		t.Errorf("expected name=fig, got %v", result["name"])
	}
	if result["version"] != float64(1) {
		t.Errorf("expected version=1, got %v", result["version"])
	}
}

func TestHTTPServerStatus(t *testing.T) {
	builtins.ResetHTTPState()

	src := `use "http"
http.route("GET", "/created", fn(req, res) {
	res.status(201)
	res.send("created")
})
http.listen(0)`

	go runFig(t, src)
	time.Sleep(200 * time.Millisecond)
	addr := builtins.ServerAddr
	if addr == "" {
		t.Fatal("server did not start")
	}
	if err := waitForServer(addr, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer builtins.ShutdownServer()

	resp, err := http.Get("http://" + addr + "/created")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		t.Errorf("expected 201, got %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if string(body) != "created" {
		t.Errorf("expected 'created', got %q", string(body))
	}
}

func TestHTTPServerReqFields(t *testing.T) {
	builtins.ResetHTTPState()

	src := `use "http"
http.route("POST", "/echo", fn(req, res) {
	res.send(req.method + " " + req.path + " " + req.body)
})
http.listen(0)`

	go runFig(t, src)
	time.Sleep(200 * time.Millisecond)
	addr := builtins.ServerAddr
	if addr == "" {
		t.Fatal("server did not start")
	}
	if err := waitForServer(addr, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer builtins.ShutdownServer()

	resp, err := http.Post("http://"+addr+"/echo", "text/plain", strings.NewReader("test-body"))
	if err != nil {
		t.Fatalf("POST /echo error: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "POST /echo test-body" {
		t.Errorf("unexpected body: %q", string(body))
	}
}

func TestHTTPServerQueryParams(t *testing.T) {
	builtins.ResetHTTPState()

	src := `use "http"
http.route("GET", "/search", fn(req, res) {
	res.send(req.query["q"])
})
http.listen(0)`

	go runFig(t, src)
	time.Sleep(200 * time.Millisecond)
	addr := builtins.ServerAddr
	if addr == "" {
		t.Fatal("server did not start")
	}
	if err := waitForServer(addr, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer builtins.ShutdownServer()

	resp, err := http.Get("http://" + addr + "/search?q=fig-lang")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "fig-lang" {
		t.Errorf("expected fig-lang, got %q", string(body))
	}
}

func TestHTTPServerMethodNotAllowed(t *testing.T) {
	builtins.ResetHTTPState()

	src := `use "http"
http.route("POST", "/only-post", fn(req, res) {
	res.send("ok")
})
http.listen(0)`

	go runFig(t, src)
	time.Sleep(200 * time.Millisecond)
	addr := builtins.ServerAddr
	if addr == "" {
		t.Fatal("server did not start")
	}
	if err := waitForServer(addr, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer builtins.ShutdownServer()

	resp, err := http.Get("http://" + addr + "/only-post")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 405 {
		t.Errorf("expected 405, got %d", resp.StatusCode)
	}
}

func TestHTTPServerMultipleRoutes(t *testing.T) {
	builtins.ResetHTTPState()

	src := `use "http"
http.route("GET", "/a", fn(req, res) { res.send("route-a") })
http.route("GET", "/b", fn(req, res) { res.send("route-b") })
http.listen(0)`

	go runFig(t, src)
	time.Sleep(200 * time.Millisecond)
	addr := builtins.ServerAddr
	if addr == "" {
		t.Fatal("server did not start")
	}
	if err := waitForServer(addr, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer builtins.ShutdownServer()

	resp1, _ := http.Get("http://" + addr + "/a")
	body1, _ := io.ReadAll(resp1.Body)
	resp1.Body.Close()

	resp2, _ := http.Get("http://" + addr + "/b")
	body2, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()

	if string(body1) != "route-a" {
		t.Errorf("expected route-a, got %q", string(body1))
	}
	if string(body2) != "route-b" {
		t.Errorf("expected route-b, got %q", string(body2))
	}
}

func TestHTTPServerReqHeaders(t *testing.T) {
	builtins.ResetHTTPState()

	src := `use "http"
http.route("GET", "/hdrs", fn(req, res) {
	res.send(req.headers["x-test"])
})
http.listen(0)`

	go runFig(t, src)
	time.Sleep(200 * time.Millisecond)
	addr := builtins.ServerAddr
	if addr == "" {
		t.Fatal("server did not start")
	}
	if err := waitForServer(addr, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer builtins.ShutdownServer()

	req, _ := http.NewRequest("GET", "http://"+addr+"/hdrs", nil)
	req.Header.Set("X-Test", "header-value")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "header-value" {
		t.Errorf("expected header-value, got %q", string(body))
	}
}
