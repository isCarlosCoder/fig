package tests

import (
	"strings"
	"testing"
)

func TestStructBasicFields(t *testing.T) {
	src := "struct Point { x; y }\nlet p = Point()\nprint(p.x)\nprint(p.y)"
	out, err := runFig(t, src)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 || lines[0] != "null" || lines[1] != "null" {
		t.Errorf("expected null/null, got %q", out)
	}
}

func TestStructFieldAssign(t *testing.T) {
	src := "struct Point { x; y }\nlet p = Point()\np.x = 10\np.y = 20\nprint(p.x)\nprint(p.y)"
	out, err := runFig(t, src)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 || lines[0] != "10" || lines[1] != "20" {
		t.Errorf("expected 10/20, got %q", out)
	}
}

func TestStructDefaultValues(t *testing.T) {
	src := "struct User { name = \"anon\"; age = 0 }\nlet u = User()\nprint(u.name)\nprint(u.age)"
	out, err := runFig(t, src)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 || lines[0] != "anon" || lines[1] != "0" {
		t.Errorf("expected anon/0, got %q", out)
	}
}

func TestStructInit(t *testing.T) {
	src := "struct Point { x; y\nfn init(x, y) { this.x = x; this.y = y } }\nlet p = Point(10, 20)\nprint(p.x)\nprint(p.y)"
	out, err := runFig(t, src)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 || lines[0] != "10" || lines[1] != "20" {
		t.Errorf("expected 10/20, got %q", out)
	}
}

func TestStructInitArityError(t *testing.T) {
	src := "struct Point { x; y\nfn init(x, y) { this.x = x; this.y = y } }\nlet p = Point(10)"
	_, err := runFig(t, src)
	if err == nil { t.Fatal("expected error for wrong init arity") }
}

func TestStructNoInitArgsError(t *testing.T) {
	src := "struct Empty { x }\nlet e = Empty(1, 2, 3)"
	_, err := runFig(t, src)
	if err == nil { t.Fatal("expected error for args without init") }
}

func TestStructMethod(t *testing.T) {
	src := "struct Point { x; y\nfn init(x, y) { this.x = x; this.y = y }\nfn move(dx, dy) { this.x = this.x + dx; this.y = this.y + dy } }\nlet p = Point(10, 20)\np.move(5, 10)\nprint(p.x)\nprint(p.y)"
	out, err := runFig(t, src)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 || lines[0] != "15" || lines[1] != "30" {
		t.Errorf("expected 15/30, got %q", out)
	}
}

func TestStructMethodReturn(t *testing.T) {
	src := "struct Counter { count = 0\nfn increment() { this.count = this.count + 1 }\nfn get() { return this.count } }\nlet c = Counter()\nc.increment()\nc.increment()\nc.increment()\nprint(c.get())"
	out, err := runFig(t, src)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	if strings.TrimSpace(out) != "3" { t.Errorf("expected 3, got %q", out) }
}

func TestStructMultipleInstances(t *testing.T) {
	src := "struct Point { x; y\nfn init(x, y) { this.x = x; this.y = y } }\nlet a = Point(1, 2)\nlet b = Point(3, 4)\nprint(a.x)\nprint(b.x)\na.x = 100\nprint(a.x)\nprint(b.x)"
	out, err := runFig(t, src)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 4 || lines[0] != "1" || lines[1] != "3" || lines[2] != "100" || lines[3] != "3" {
		t.Errorf("expected 1/3/100/3, got %q", out)
	}
}

func TestStructThisOutsideMethod(t *testing.T) {
	_, err := runFig(t, "print(this)")
	if err == nil { t.Fatal("expected error for this outside method") }
}

func TestStructPrint(t *testing.T) {
	src := "struct Point { x; y\nfn init(x, y) { this.x = x; this.y = y } }\nlet p = Point(5, 10)\nprint(p)"
	out, err := runFig(t, src)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	if !strings.HasPrefix(strings.TrimSpace(out), "Point{") { t.Errorf("expected Point{...}, got %q", out) }
}

func TestStructWithLoop(t *testing.T) {
	src := "struct Item { name\nfn init(n) { this.name = n }\nfn getName() { return this.name } }\nlet items = [Item(\"a\"), Item(\"b\"), Item(\"c\")]\nlet r = \"\"\nfor item in items {\nr = r + item.getName()\n}\nprint(r)"
	out, err := runFig(t, src)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	if strings.TrimSpace(out) != "abc" { t.Errorf("expected abc, got %q", out) }
}

func TestStructMethodArityError(t *testing.T) {
	src := "struct Foo { fn bar(x) { print(x) } }\nlet f = Foo()\nf.bar()"
	_, err := runFig(t, src)
	if err == nil { t.Fatal("expected error for wrong method arity") }
}

func TestStructNested(t *testing.T) {
	src := "struct Address { city\nfn init(c) { this.city = c } }\nstruct Person { name; addr\nfn init(n, c) { this.name = n; this.addr = Address(c) } }\nlet p = Person(\"Carlos\", \"SP\")\nprint(p.name)\nprint(p.addr.city)"
	out, err := runFig(t, src)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 || lines[0] != "Carlos" || lines[1] != "SP" {
		t.Errorf("expected Carlos/SP, got %q", out)
	}
}

func TestStructMethodCalc(t *testing.T) {
	src := "struct Calc { result = 0\nfn add(n) { this.result = this.result + n }\nfn getResult() { return this.result } }\nlet c = Calc()\nc.add(10)\nc.add(20)\nc.add(30)\nprint(c.getResult())"
	out, err := runFig(t, src)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	if strings.TrimSpace(out) != "60" { t.Errorf("expected 60, got %q", out) }
}
