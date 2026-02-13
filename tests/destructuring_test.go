package tests

import "testing"

func TestArrayDestructuringDeclaration(t *testing.T) {
	src := `let [a, b] = [10, 20]; print(a); print(b);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "10\n20" {
		t.Fatalf("expected '10\n20', got %q", out)
	}
}

func TestArrayDestructuringIgnoreAndMissing(t *testing.T) {
	src := `let [_, a] = [10, 20]; print(a);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "20" {
		t.Fatalf("expected '20', got %q", out)
	}

	// missing element becomes null
	src2 := `let [x, y] = [1]; print(x); print(y);`
	out2, err := runFigSource(t, src2)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out2 != "1\nnull" {
		t.Fatalf("expected '1\nnull', got %q", out2)
	}
}

func TestObjectDestructuringDeclaration(t *testing.T) {
	src := `let {name, age} = { name: "Carlos", age: 23 }; print(name); print(age);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "Carlos\n23" {
		t.Fatalf("expected 'Carlos\n23', got %q", out)
	}
}

func TestDestructuringAssignmentToExistingVars(t *testing.T) {
	src := `let arr = [7,8]; let a = 0; let b = 0; [a, b] = arr; print(a); print(b);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "7\n8" {
		t.Fatalf("expected '7\n8', got %q", out)
	}
}

func TestNestedArrayDestructuring(t *testing.T) {
	src := `let [a, [b, c]] = [1, [2, 3]]; print(a); print(b); print(c);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "1\n2\n3" {
		t.Fatalf("expected '1\n2\n3', got %q", out)
	}
}

func TestNestedDestructuringAssignment(t *testing.T) {
	src := `let arr = [1, [2,3]]; let a = 0; let b = 0; let c = 0; [a, [b, c]] = arr; print(a); print(b); print(c);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "1\n2\n3" {
		t.Fatalf("expected '1\n2\n3', got %q", out)
	}
}

func TestArrayContainingObjectPattern(t *testing.T) {
	src := `let [ {name, age}, x ] = [ {name: "C", age: 23}, 5 ]; print(name); print(age); print(x);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "C\n23\n5" {
		t.Fatalf("expected 'C\n23\n5', got %q", out)
	}
}
