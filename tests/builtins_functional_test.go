package tests

import (
"testing"
)

func useFunctional(code string) string {
return "use " + `"` + "functional" + `"` + "; use " + `"` + "math" + `"` + "; " + code
}

func TestFunctionalCall(t *testing.T) {
src := useFunctional(`print(functional.call(math.abs, -42));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "42" { t.Fatalf("expected '42', got %q", out) }
}

func TestFunctionalApply(t *testing.T) {
src := useFunctional(`let args = [-9]; print(functional.apply(math.abs, args));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "9" { t.Fatalf("expected '9', got %q", out) }
}

func TestFunctionalPartial(t *testing.T) {
src := useFunctional(`let pow2 = functional.partial(math.pow, 2); print(pow2(10));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "1024" { t.Fatalf("expected '1024', got %q", out) }
}

func TestFunctionalOnce(t *testing.T) {
src := useFunctional(`let f = functional.once(math.abs); print(f(-5)); print(f(-999));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "5\n5" { t.Fatalf("expected '5\\n5', got %q", out) }
}

func TestFunctionalMemo(t *testing.T) {
src := useFunctional(`let f = functional.memo(math.sqrt); print(f(16)); print(f(16)); print(f(25));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "4\n4\n5" { t.Fatalf("expected '4\\n4\\n5', got %q", out) }
}

func TestFunctionalCallWrongType(t *testing.T) {
src := useFunctional(`functional.call("not_fn", 1);`)
_, err := runFig(t, src)
if err == nil { t.Fatalf("expected error for non-function argument") }
}

func TestFunctionalPartialMultiArgs(t *testing.T) {
src := useFunctional(`let clamp5 = functional.partial(math.clamp, 5); print(clamp5(0, 10));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "5" { t.Fatalf("expected '5', got %q", out) }
}
