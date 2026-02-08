package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// newTestServer creates a simple HTTP test server for testing the http module.
func newTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name":"fig","version":1}`)
	})

	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "text/plain")
		w.Write(body)
	})

	mux.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		hdrs := map[string]string{}
		for k := range r.Header {
			hdrs[strings.ToLower(k)] = r.Header.Get(k)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hdrs)
	})

	mux.HandleFunc("/status/", func(w http.ResponseWriter, r *http.Request) {
		code := 200
		fmt.Sscanf(r.URL.Path, "/status/%d", &code)
		w.WriteHeader(code)
		fmt.Fprintf(w, "status %d", code)
	})

	mux.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "file-content-here")
	})

	return httptest.NewServer(mux)
}

func TestHTTPGet(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.get("%s/hello")
print(res.status)
print(res.body)`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) != 2 || lines[0] != "200" || lines[1] != "hello world" {
		t.Errorf("expected 200/hello world, got %q", out)
	}
}

func TestHTTPRequest(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.request("GET", "%s/hello")
print(res.status)
print(res.body)`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) != 2 || lines[0] != "200" || lines[1] != "hello world" {
		t.Errorf("expected 200/hello world, got %q", out)
	}
}

func TestHTTPPost(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.post("%s/echo", "hello from fig")
print(res.status)
print(res.body)`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) != 2 || lines[0] != "200" || lines[1] != "hello from fig" {
		t.Errorf("expected 200/hello from fig, got %q", out)
	}
}

func TestHTTPRequestWithBody(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.request("POST", "%s/echo", "post-body")
print(res.body)`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "post-body" {
		t.Errorf("expected post-body, got %q", out)
	}
}

func TestHTTPRequestWithHeaders(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
use "json"
let res = http.request("GET", "%s/headers", null, {"X-Custom": "fig-value"})
let hdrs = json.parse(res.body)
print(hdrs["x-custom"])`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "fig-value" {
		t.Errorf("expected fig-value, got %q", out)
	}
}

func TestHTTPSetHeader(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
use "json"
http.setHeader("X-Global", "global-val")
let res = http.get("%s/headers")
let hdrs = json.parse(res.body)
print(hdrs["x-global"])
http.clearHeaders()`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "global-val" {
		t.Errorf("expected global-val, got %q", out)
	}
}

func TestHTTPClearHeaders(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
use "json"
http.setHeader("X-Temp", "temp-val")
http.clearHeaders()
let res = http.get("%s/headers")
let hdrs = json.parse(res.body)
let has = hdrs["x-temp"]
print(has)`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "null" {
		t.Errorf("expected null (header cleared), got %q", out)
	}
}

func TestHTTPSetTimeout(t *testing.T) {
	src := `use "http"
http.setTimeout(5000)
print("ok")`

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "ok" {
		t.Errorf("expected ok, got %q", out)
	}
}

func TestHTTPIsOkTrue(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.get("%s/hello")
print(http.isOk(res))`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "true" {
		t.Errorf("expected true, got %q", out)
	}
}

func TestHTTPIsOkFalse(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.get("%s/status/404")
print(http.isOk(res))`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "false" {
		t.Errorf("expected false, got %q", out)
	}
}

func TestHTTPRaiseForStatusOk(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.get("%s/hello")
http.raiseForStatus(res)
print("ok")`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "ok" {
		t.Errorf("expected ok, got %q", out)
	}
}

func TestHTTPRaiseForStatusError(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.get("%s/status/500")
http.raiseForStatus(res)`, srv.URL)

	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for status 500")
	}
}

func TestHTTPDownload(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	tmpFile := t.TempDir() + "/downloaded.txt"

	src := fmt.Sprintf(`use "http"
http.download("%s/file", "%s")
print("done")`, srv.URL, tmpFile)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "done" {
		t.Errorf("expected done, got %q", out)
	}

	data, readErr := os.ReadFile(tmpFile)
	if readErr != nil {
		t.Fatalf("cannot read downloaded file: %v", readErr)
	}
	if string(data) != "file-content-here" {
		t.Errorf("expected file-content-here, got %q", string(data))
	}
}

func TestHTTPResponseHeaders(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.get("%s/json")
print(res.headers["content-type"])`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "application/json") {
		t.Errorf("expected application/json in content-type, got %q", out)
	}
}

func TestHTTPPostNoBody(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.post("%s/echo")
print(res.status)`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "200" {
		t.Errorf("expected 200, got %q", out)
	}
}

func TestHTTPRequestPUT(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.request("PUT", "%s/echo", "put-data")
print(res.body)`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "put-data" {
		t.Errorf("expected put-data, got %q", out)
	}
}

func TestHTTPRequestDELETE(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	src := fmt.Sprintf(`use "http"
let res = http.request("DELETE", "%s/hello")
print(res.status)`, srv.URL)

	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "200" {
		t.Errorf("expected 200, got %q", out)
	}
}

func TestHTTPGetArgError(t *testing.T) {
	_, err := runFig(t, `use "http"
http.get()`)
	if err == nil {
		t.Fatal("expected error for get() with no args")
	}
}

func TestHTTPPostArgError(t *testing.T) {
	_, err := runFig(t, `use "http"
http.post()`)
	if err == nil {
		t.Fatal("expected error for post() with no args")
	}
}

func TestHTTPRequestArgError(t *testing.T) {
	_, err := runFig(t, `use "http"
http.request()`)
	if err == nil {
		t.Fatal("expected error for request() with no args")
	}
}

func TestHTTPSetTimeoutArgError(t *testing.T) {
	_, err := runFig(t, `use "http"
http.setTimeout("bad")`)
	if err == nil {
		t.Fatal("expected error for setTimeout with string")
	}
}

func TestHTTPIsOkArgError(t *testing.T) {
	_, err := runFig(t, `use "http"
http.isOk("not-an-object")`)
	if err == nil {
		t.Fatal("expected error for isOk with string")
	}
}

func TestHTTPDownloadArgError(t *testing.T) {
	_, err := runFig(t, `use "http"
http.download(123)`)
	if err == nil {
		t.Fatal("expected error for download with wrong args")
	}
}
