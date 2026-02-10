package tests

import (
	"strings"
	"testing"
)

// --- 1. Basic try/onerror: division by zero fallback ---

func TestTryDivisionByZeroFallback(t *testing.T) {
	out, err := runFig(t, `
		let x = try 10 / 0 onerror() {
			return 0
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "0" {
		t.Fatalf("expected '0', got %q", out)
	}
}

// --- 2. Capture error message ---

func TestTryCaptureErrorMessage(t *testing.T) {
	out, err := runFig(t, `
		let x = try 10 / 0 onerror(e) {
			print(e)
			return -1
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) < 2 {
		t.Fatalf("expected 2 lines, got: %q", out)
	}
	if !strings.Contains(lines[0], "division by zero") {
		t.Fatalf("expected error message about division by zero, got: %q", lines[0])
	}
	if lines[1] != "-1" {
		t.Fatalf("expected '-1', got %q", lines[1])
	}
}

// --- 3. No error — returns original value ---

func TestTryNoError(t *testing.T) {
	out, err := runFig(t, `
		let x = try 10 / 2 onerror() {
			return 0
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "5" {
		t.Fatalf("expected '5', got %q", out)
	}
}

// --- 4. Try with function that errors ---

func TestTryWithFunctionError(t *testing.T) {
	out, err := runFig(t, `
		fn safeDiv(a, b) {
			return try a / b onerror(e) {
				return 0
			}
		}
		print(safeDiv(10, 0))
		print(safeDiv(10, 2))
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "0\n5" {
		t.Fatalf("expected '0\\n5', got %q", out)
	}
}

// --- 5. Try in array literal ---

func TestTryInArrayLiteral(t *testing.T) {
	out, err := runFig(t, `
		let arr = [1, 2, try 10 / 0 onerror { return 0 }, 4]
		print(arr)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "[1, 2, 0, 4]" {
		t.Fatalf("expected '[1, 2, 0, 4]', got %q", out)
	}
}

// --- 6. Try in object literal ---

func TestTryInObjectLiteral(t *testing.T) {
	out, err := runFig(t, `
		let obj = {
			"val": try 10 / 0 onerror { return 99 }
		}
		print(obj.val)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "99" {
		t.Fatalf("expected '99', got %q", out)
	}
}

// --- 7. Nested try fallback cascade ---

func TestTryNestedCascade(t *testing.T) {
	out, err := runFig(t, `
		let val = try 10 / 0 onerror {
			return try 20 / 0 onerror {
				return 42
			}
		}
		print(val)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "42" {
		t.Fatalf("expected '42', got %q", out)
	}
}

// --- 8. Try without error variable ---

func TestTryWithoutErrorVar(t *testing.T) {
	out, err := runFig(t, `
		let x = try 10 / 0 onerror {
			return -1
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "-1" {
		t.Fatalf("expected '-1', got %q", out)
	}
}

// --- 9. Try in if condition ---

func TestTryInIfCondition(t *testing.T) {
	out, err := runFig(t, `
		if (try 10 / 0 onerror { return false }) {
			print("yes")
		} else {
			print("no")
		}
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "no" {
		t.Fatalf("expected 'no', got %q", out)
	}
}

// --- 10. Try in function argument ---

func TestTryInFunctionArg(t *testing.T) {
	out, err := runFig(t, `
		fn double(n) { return n * 2 }
		let result = double(try 10 / 0 onerror { return 5 })
		print(result)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "10" {
		t.Fatalf("expected '10', got %q", out)
	}
}

// --- 11. Try with unary minus error ---

func TestTryUnaryMinusOnString(t *testing.T) {
	out, err := runFig(t, `
		let x = try -"hello" onerror(e) {
			return 0
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "0" {
		t.Fatalf("expected '0', got %q", out)
	}
}

// --- 12. Try with dot access on null ---

func TestTryDotAccessOnNull(t *testing.T) {
	out, err := runFig(t, `
		let x = null
		let y = try x.foo onerror {
			return "default"
		}
		print(y)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "default" {
		t.Fatalf("expected 'default', got %q", out)
	}
}

// --- 13. Try with comparison error ---

func TestTryComparisonError(t *testing.T) {
	out, err := runFig(t, `
		let x = try ("hello" > 5) onerror {
			return false
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "false" {
		t.Fatalf("expected 'false', got %q", out)
	}
}

// --- 14. Multiple try in same expression ---

func TestMultipleTrySameExpression(t *testing.T) {
	out, err := runFig(t, `
		let a = try 10 / 0 onerror { return 1 }
		let b = try 20 / 0 onerror { return 2 }
		print(a + b)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "3" {
		t.Fatalf("expected '3', got %q", out)
	}
}

// --- 15. Try onerror with continue in loop ---

func TestTryOnerrorContinue(t *testing.T) {
	out, err := runFig(t, `
		use "arrays"
		let results = []
		let inputs = [2, 0, 5, 0, 3]
		for x in inputs {
			let n = try 10 / x onerror {
				continue
			}
			arrays.push(results, n)
		}
		print(results)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// inputs: 2→5, 0→skip, 5→2, 0→skip, 3→3.33...
	if !strings.HasPrefix(out, "[5, 2, 3.") {
		t.Fatalf("expected '[5, 2, 3.33...]', got %q", out)
	}
}

// --- 16. Try onerror with break in loop ---

func TestTryOnerrorBreak(t *testing.T) {
	out, err := runFig(t, `
		use "arrays"
		let results = []
		let inputs = [2, 4, 0, 8]
		for x in inputs {
			let n = try 10 / x onerror {
				break
			}
			arrays.push(results, n)
		}
		print(results)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "[5, 2.5]" {
		t.Fatalf("expected '[5, 2.5]', got %q", out)
	}
}

// --- 17. Try with anonymous function ---

func TestTryWithAnonymousFunction(t *testing.T) {
	out, err := runFig(t, `
		fn safe(f) {
			return fn(x) {
				return try f(x) onerror { return null }
			}
		}
		fn risky(x) {
			return 10 / x
		}
		let safeFn = safe(risky)
		print(safeFn(2))
		print(safeFn(0))
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "5\nnull" {
		t.Fatalf("expected '5\\nnull', got %q", out)
	}
}

// --- 18. Try preserves onerror scope ---

func TestTryOnerrorScope(t *testing.T) {
	out, err := runFig(t, `
		let e = "outer"
		let x = try 10 / 0 onerror(e) {
			return e
		}
		print(x)
		print(e)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) < 2 {
		t.Fatalf("expected 2 lines, got: %q", out)
	}
	if !strings.Contains(lines[0], "division by zero") {
		t.Fatalf("expected error message in x, got: %q", lines[0])
	}
	if lines[1] != "outer" {
		t.Fatalf("expected outer e unchanged, got: %q", lines[1])
	}
}

// --- 19. Try onerror block without return gives null ---

func TestTryOnerrorNoReturn(t *testing.T) {
	out, err := runFig(t, `
		let x = try 10 / 0 onerror {
			print("caught")
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "caught\nnull" {
		t.Fatalf("expected 'caught\\nnull', got %q", out)
	}
}

// --- 20. Try as expression statement (discard value) ---

func TestTryAsExpressionStatement(t *testing.T) {
	out, err := runFig(t, `
		try 10 / 0 onerror(e) {
			print("error: " + e)
		}
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "error: division by zero" {
		t.Fatalf("expected 'error: division by zero', got %q", out)
	}
}

// --- 21. Try with subtraction on strings ---

func TestTrySubtractStrings(t *testing.T) {
	out, err := runFig(t, `
		let x = try "a" - "b" onerror {
			return "cannot subtract"
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "cannot subtract" {
		t.Fatalf("expected 'cannot subtract', got %q", out)
	}
}

// --- 22. Try with multiply on strings ---

func TestTryMultiplyStrings(t *testing.T) {
	out, err := runFig(t, `
		let x = try "a" * "b" onerror {
			return 0
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "0" {
		t.Fatalf("expected '0', got %q", out)
	}
}

// --- 23. Try with index out of range ---

func TestTryIndexOutOfRange(t *testing.T) {
	out, err := runFig(t, `
		let arr = [1, 2, 3]
		let x = try arr[10] onerror {
			return -1
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "-1" {
		t.Fatalf("expected '-1', got %q", out)
	}
}

// --- 24. Try in while loop ---

func TestTryInWhileLoop(t *testing.T) {
	out, err := runFig(t, `
		use "arrays"
		let i = 3
		let results = []
		while (i >= 0) {
			let n = try 10 / i onerror {
				i = i - 1
				continue
			}
			arrays.push(results, n)
			i = i - 1
		}
		print(results)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// i=3 → 10/3=3.33, i=2 → 10/2=5, i=1 → 10/1=10, i=0 → error→continue
	if !strings.HasPrefix(out, "[3.") {
		t.Fatalf("expected array starting with 3.33..., got %q", out)
	}
	if !strings.Contains(out, "5") || !strings.Contains(out, "10") {
		t.Fatalf("expected 5 and 10 in output, got %q", out)
	}
}

// --- 25. Struct with try in init ---

func TestTryInStructInit(t *testing.T) {
	out, err := runFig(t, `
		struct User {
			name
			age

			fn init(data) {
				this.name = try data.name onerror { return "anon" }
				this.age  = try data.age  onerror { return 0 }
			}
		}
		let u = User({})
		print(u.name)
		print(u.age)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Accessing .name on an empty object returns null (not an error), so no onerror triggered
	// But let's test with a number where dot access would error
	if out != "null\nnull" {
		t.Fatalf("expected 'null\\nnull' (object key not found returns null), got %q", out)
	}
}

// --- 26. Struct with try catching actual error ---

func TestTryInStructInitWithError(t *testing.T) {
	out, err := runFig(t, `
		struct Config {
			val

			fn init(raw) {
				this.val = try raw / 2 onerror { return 0 }
			}
		}
		let c1 = Config(10)
		let c2 = Config("bad")
		print(c1.val)
		print(c2.val)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "5\n0" {
		t.Fatalf("expected '5\\n0', got %q", out)
	}
}

// --- 27. Try with guarded block that returns ---

func TestTryGuardedBlockReturn(t *testing.T) {
	out, err := runFig(t, `
		let x = try {
			return 7
		} onerror {
			return 0
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "7" {
		t.Fatalf("expected '7', got %q", out)
	}
}

// --- 28. Try with guarded block that falls through (no return) gives null ---

func TestTryGuardedBlockNoReturn(t *testing.T) {
	out, err := runFig(t, `
		let x = try {
			let a = 1
		} onerror {
			return 0
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "null" {
		t.Fatalf("expected 'null', got %q", out)
	}
}

// --- 29. Try with guarded block that errors and uses onerror(e) ---

func TestTryGuardedBlockErrorAndOnerror(t *testing.T) {
	out, err := runFig(t, `
		let x = try {
			let y = 10 / 0
		} onerror(e) {
			print(e)
			return 123
		}
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) < 2 {
		t.Fatalf("expected 2 lines, got: %q", out)
	}
	if !strings.Contains(lines[0], "division by zero") {
		t.Fatalf("expected error message about division by zero, got: %q", lines[0])
	}
	if lines[1] != "123" {
		t.Fatalf("expected '123', got %q", lines[1])
	}
}

// --- 30. Try with guarded block using continue in loop ---

func TestTryGuardedBlockContinue(t *testing.T) {
	out, err := runFig(t, `
		use "arrays"
		let results = []
		let inputs = [1, 2, 3]
		for x in inputs {
			let n = try {
				if x == 1 { continue }
				return x
			} onerror {
				return -1
			}
			arrays.push(results, n)
		}
		print(results)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "[2, 3]" {
		t.Fatalf("expected '[2, 3]', got %q", out)
	}
}

// --- 31. Try with guarded block using break in loop ---

func TestTryGuardedBlockBreak(t *testing.T) {
	out, err := runFig(t, `
		use "arrays"
		let results = []
		let inputs = [2, 4, 0, 8]
		for x in inputs {
			let n = try {
				if x == 0 { break }
				return 10 / x
			} onerror {
				return -1
			}
			arrays.push(results, n)
		}
		print(results)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "[5, 2.5]" {
		t.Fatalf("expected '[5, 2.5]', got %q", out)
	}
}
