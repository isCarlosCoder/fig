package tests

import (
"strings"
"testing"
)

func useRuntime(code string) string {
return `use "runtime"; ` + code
}

func TestRuntimeGC(t *testing.T) {
out, err := runFig(t, useRuntime(`runtime.gc(); print("ok")`))
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if strings.TrimSpace(out) != "ok" {
t.Errorf("expected ok, got %q", out)
}
}

func TestRuntimeMemUsage(t *testing.T) {
out, err := runFig(t, useRuntime(`let m = runtime.memUsage(); print(m.alloc > 0)`))
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if strings.TrimSpace(out) != "true" {
t.Errorf("expected true, got %q", out)
}
}

func TestRuntimeMemUsageTotalAlloc(t *testing.T) {
out, err := runFig(t, useRuntime(`let m = runtime.memUsage(); print(m.totalAlloc > 0)`))
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if strings.TrimSpace(out) != "true" {
t.Errorf("expected true, got %q", out)
}
}

func TestRuntimeMemUsageSys(t *testing.T) {
out, err := runFig(t, useRuntime(`let m = runtime.memUsage(); print(m.sys > 0)`))
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if strings.TrimSpace(out) != "true" {
t.Errorf("expected true, got %q", out)
}
}

func TestRuntimeMemUsageNumGC(t *testing.T) {
out, err := runFig(t, useRuntime(`let m = runtime.memUsage(); print(m.numGC >= 0)`))
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if strings.TrimSpace(out) != "true" {
t.Errorf("expected true, got %q", out)
}
}

func TestRuntimeVersion(t *testing.T) {
out, err := runFig(t, useRuntime(`let v = runtime.version(); print(v)`))
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if !strings.HasPrefix(strings.TrimSpace(out), "go") {
t.Errorf("expected version starting with 'go', got %q", out)
}
}

func TestRuntimePlatform(t *testing.T) {
out, err := runFig(t, useRuntime(`let p = runtime.platform(); print(p)`))
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
trimmed := strings.TrimSpace(out)
if !strings.Contains(trimmed, "/") {
t.Errorf("expected platform like 'os/arch', got %q", trimmed)
}
}

func TestRuntimeNumCPU(t *testing.T) {
out, err := runFig(t, useRuntime(`let n = runtime.numCPU(); print(n > 0)`))
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if strings.TrimSpace(out) != "true" {
t.Errorf("expected true, got %q", out)
}
}

func TestRuntimeGCThenMemUsage(t *testing.T) {
out, err := runFig(t, useRuntime(`runtime.gc(); let m = runtime.memUsage(); print(m.alloc > 0)`))
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if strings.TrimSpace(out) != "true" {
t.Errorf("expected true, got %q", out)
}
}

func TestRuntimeGCErrorArgs(t *testing.T) {
_, err := runFig(t, useRuntime(`runtime.gc(1)`))
if err == nil {
t.Fatal("expected error for gc() with args")
}
}
