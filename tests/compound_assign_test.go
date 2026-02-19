package tests

import "testing"

func TestVarCompoundAssignPlus(t *testing.T) {
	out, err := runFig(t, "let x = 10; x += 5; print(x);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "15" {
		t.Fatalf("expected '15', got %q", out)
	}
}

func TestVarCompoundAssignMinus(t *testing.T) {
	out, err := runFig(t, "let x = 10; x -= 3; print(x);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "7" {
		t.Fatalf("expected '7', got %q", out)
	}
}

func TestVarCompoundAssignMultiplyDivideModulo(t *testing.T) {
	out, err := runFig(t, "let a = 3; a *= 4; print(a); let b = 20; b /= 4; print(b); let c = 10; c %= 3; print(c);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "12\n5\n1" {
		t.Fatalf("expected '12\\n5\\n1', got %q", out)
	}
}

func TestStringPlusAssign(t *testing.T) {
	out, err := runFig(t, "let s = \"a\"; s += \"b\"; print(s);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "ab" {
		t.Fatalf("expected 'ab', got %q", out)
	}
}

func TestMemberCompoundAssign(t *testing.T) {
	out, err := runFig(t, "let o = {x: 1}; o.x += 2; print(o.x);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "3" {
		t.Fatalf("expected '3', got %q", out)
	}

	out, err = runFig(t, "let a = [1,2,3]; a[1] *= 3; print(a);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[1, 6, 3]" {
		t.Fatalf("expected '[1, 6, 3]', got %q", out)
	}
}

func TestCompoundAssignUndefinedVarError(t *testing.T) {
	_, err := runFig(t, "x += 1")
	if err == nil {
		t.Fatal("expected runtime error for undefined variable")
	}
}

func TestCompoundAssignDestructuringError(t *testing.T) {
	_, err := runFig(t, "let [a, b] = [1,2]; [a, b] += [1,1];")
	if err == nil {
		t.Fatal("expected runtime error for compound assignment on destructuring")
	}
}
