package tests

import (
	"strings"
	"testing"
)

func TestEvalStepLimitEnforced(t *testing.T) {
	// This loop should hit the eval step limit and produce an error
	src := `use "system"
let i = 0
while (i < 100000) { i = i + 1 }`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for exceeding eval steps, got none")
	}
	if !strings.Contains(err.Error(), "maximum evaluation steps exceeded") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDisableEnableStepLimit(t *testing.T) {
	// disable the limit, run a heavy loop, then re-enable
	src := `use "system"
system.disableStepLimit()
print("disabled=" + system.isStepLimitDisabled())
let i = 0
while (i < 30000) { i = i + 1 }
print("doneLoop")
system.enableStepLimit()
print("afterEnabled=" + system.isStepLimitDisabled())
print(i)`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error while disabled: %v", err)
	}
	if !strings.Contains(out, "disabled=true") {
		t.Fatalf("expected disabled=true in output, got %q", out)
	}
	if !strings.Contains(out, "doneLoop") {
		t.Fatalf("loop did not complete, output: %q", out)
	}
	if !strings.Contains(out, "afterEnabled=false") {
		t.Fatalf("expected afterEnabled=false in output, got %q", out)
	}
	if !strings.Contains(out, "30000") {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestWithoutStepLimitRequiresTask(t *testing.T) {
	// without loading task, calling withoutStepLimit should error
	src := `use "system"
fn heavy() { let i = 0; while (i < 30000) { i = i + 1 } return i }
system.withoutStepLimit(heavy)`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error when task module not loaded")
	}
}

func TestWithoutStepLimitWithTask(t *testing.T) {
	// with task loaded it should execute
	src := `use "task"
use "system"
fn heavy() { let i = 0; while (i < 30000) { i = i + 1 } return i }
let r = system.withoutStepLimit(heavy)
print(r)`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error while withoutStepLimit: %v", err)
	}
	if strings.TrimSpace(out) != "30000" {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestDisableStepLimitReferenceCount(t *testing.T) {
	// If the caller disables the step limit, a nested withoutStepLimit must not
	// re-enable it on completion. This verifies the reference-counted behavior.
	src := `use "task"
use "system"
print("start=" + system.isStepLimitDisabled())
system.disableStepLimit()
print("afterDisable=" + system.isStepLimitDisabled())
fn heavy() { let i = 0; while (i < 30000) { i = i + 1 } return i }
let r = system.withoutStepLimit(heavy)
print("afterWithout=" + system.isStepLimitDisabled())
print(r)
system.enableStepLimit()
print("afterEnabled=" + system.isStepLimitDisabled())`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error in nested disable test: %v", err)
	}
	// Be tolerant of extra parser/debug output; ensure expected pieces are present
	if !strings.Contains(out, "afterDisable=true") {
		t.Fatalf("expected afterDisable=true in output, got %q", out)
	}
	if !strings.Contains(out, "afterWithout=true") {
		t.Fatalf("expected afterWithout=true in output, got %q", out)
	}
	if !strings.Contains(out, "30000") {
		t.Fatalf("expected heavy return 30000 in output, got %q", out)
	}
	if !strings.Contains(out, "afterEnabled=false") {
		t.Fatalf("expected afterEnabled=false in output, got %q", out)
	}
}
