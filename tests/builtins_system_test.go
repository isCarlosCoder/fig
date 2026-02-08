package tests

import (
"strconv"
"strings"
"testing"
)

func useSystem(code string) string {
return "use " + `"` + "system" + `"` + "; " + code
}

func TestSystemNow(t *testing.T) {
src := useSystem(`print(system.now());`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
n, e := strconv.ParseFloat(out, 64)
if e != nil { t.Fatalf("expected number, got %q", out) }
if n < 1e12 { t.Fatalf("expected timestamp in ms, got %v", n) }
}

func TestSystemClock(t *testing.T) {
src := useSystem(`print(system.clock());`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
n, e := strconv.ParseFloat(out, 64)
if e != nil { t.Fatalf("expected number, got %q", out) }
if n < 1e9 { t.Fatalf("expected clock in seconds, got %v", n) }
}

func TestSystemVersion(t *testing.T) {
src := useSystem(`print(system.version());`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "0.1.0" { t.Fatalf("expected '0.1.0', got %q", out) }
}

func TestSystemPlatform(t *testing.T) {
src := useSystem(`print(system.platform());`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "linux" && out != "darwin" && out != "windows" { t.Fatalf("unexpected platform %q", out) }
}

func TestSystemEnv(t *testing.T) {
src := useSystem(`print(system.env("HOME"));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out == "" || out == "null" { t.Fatalf("expected HOME path, got %q", out) }
}

func TestSystemEnvMissing(t *testing.T) {
src := useSystem(`print(system.env("FIGTEST_NONEXISTENT_ABC"));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "null" { t.Fatalf("expected 'null', got %q", out) }
}

func TestSystemArgs(t *testing.T) {
src := useSystem(`let a = system.args(); print(a);`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if !strings.HasPrefix(out, "[") { t.Fatalf("expected array, got %q", out) }
}

func TestSystemSleep(t *testing.T) {
src := useSystem(`system.sleep(1); print("ok");`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "ok" { t.Fatalf("expected 'ok', got %q", out) }
}
