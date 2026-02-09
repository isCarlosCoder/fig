package builtins

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/iscarloscoder/fig/environment"
)

// FigtestCaller is set by the interpreter so that figtest functions can invoke
// user-defined Fig functions and capture errors without crashing the visitor.
// Returns (resultValue, error). If the fn panics/errors, error is non-nil.
var FigtestCaller func(fn environment.Value, args []environment.Value) (environment.Value, error)

// FigtestState holds the mutable state for a figtest session.
// Each test file executed by `fig test` gets its own state instance,
// but within a single file the module functions share this state.
type FigtestState struct {
	mu sync.Mutex

	tests   []testEntry
	groups  []string // describe stack
	hasOnly bool     // true if any only() was registered

	// hooks for current describe scope
	beforeAll  environment.Value
	afterAll   environment.Value
	beforeEach environment.Value
	afterEach  environment.Value

	// counters
	passCount int
	failCount int
	skipCount int

	// output lines for summary
	output []string
}

type testEntry struct {
	name  string
	fn    environment.Value
	group string
	skip  bool
	only  bool
}

// Global figtest state — reset per file via ResetFigtest().
var globalFigtest = &FigtestState{}

// ResetFigtest creates a fresh state (used by CLI `fig test` per file).
func ResetFigtest() {
	globalFigtest = &FigtestState{}
}

// GetFigtestState returns the current state (for CLI summary aggregation).
func GetFigtestState() *FigtestState {
	return globalFigtest
}

func (s *FigtestState) Passed() int      { return s.passCount }
func (s *FigtestState) Failed() int      { return s.failCount }
func (s *FigtestState) Skipped() int     { return s.skipCount }
func (s *FigtestState) Output() []string { return s.output }

func (s *FigtestState) currentGroup() string {
	if len(s.groups) == 0 {
		return ""
	}
	return strings.Join(s.groups, " > ")
}

func (s *FigtestState) runTest(name string, testFn environment.Value, skip bool) {
	if skip {
		s.skipCount++
		s.output = append(s.output, fmt.Sprintf("  ○ %s (skipped)", name))
		return
	}

	// run beforeEach hook
	if s.beforeEach.Type == environment.FunctionType || s.beforeEach.Type == environment.BuiltinFnType {
		if FigtestCaller != nil {
			_, err := FigtestCaller(s.beforeEach, nil)
			if err != nil {
				s.failCount++
				s.output = append(s.output, fmt.Sprintf("  ✗ %s (beforeEach failed: %s)", name, err.Error()))
				return
			}
		}
	}

	// run test function
	_, err := FigtestCaller(testFn, nil)

	// run afterEach hook
	if s.afterEach.Type == environment.FunctionType || s.afterEach.Type == environment.BuiltinFnType {
		if FigtestCaller != nil {
			FigtestCaller(s.afterEach, nil)
		}
	}

	if err != nil {
		s.failCount++
		s.output = append(s.output, fmt.Sprintf("  ✗ %s (%s)", name, err.Error()))
	} else {
		s.passCount++
		s.output = append(s.output, fmt.Sprintf("  ✓ %s", name))
	}
}

func (s *FigtestState) executeAll() {
	// determine if we run only "only" tests
	hasOnly := false
	for _, t := range s.tests {
		if t.only {
			hasOnly = true
			break
		}
	}

	for _, t := range s.tests {
		skip := t.skip
		if hasOnly && !t.only {
			skip = true
		}
		s.runTest(t.name, t.fn, skip)
	}
}

func figtestAssertFn(args []environment.Value, minArgs int, name string) error {
	if len(args) < minArgs {
		return fmt.Errorf("%s() expects at least %d argument(s), got %d", name, minArgs, len(args))
	}
	return nil
}

func figtestMsg(args []environment.Value, idx int, fallback string) string {
	if len(args) > idx && args[idx].Type == environment.StringType {
		return args[idx].Str
	}
	return fallback
}

func init() {
	register(newModule("figtest",
		// test(name, fn) — register and immediately execute a test
		fn("test", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 2 {
				return environment.NewNil(), fmt.Errorf("test() expects 2 arguments (name, fn)")
			}
			name := args[0].String()
			testFn := args[1]
			if !isCallable(testFn) {
				return environment.NewNil(), fmt.Errorf("test() second argument must be a function")
			}
			s := globalFigtest
			s.mu.Lock()
			group := s.currentGroup()
			s.mu.Unlock()
			if group != "" {
				name = group + " > " + name
			}
			s.runTest(name, testFn, false)
			return environment.NewNil(), nil
		}),

		// describe(name, fn) — group tests
		fn("describe", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 2 {
				return environment.NewNil(), fmt.Errorf("describe() expects 2 arguments (name, fn)")
			}
			groupName := args[0].String()
			groupFn := args[1]
			if !isCallable(groupFn) {
				return environment.NewNil(), fmt.Errorf("describe() second argument must be a function")
			}
			s := globalFigtest
			s.mu.Lock()
			s.groups = append(s.groups, groupName)
			s.output = append(s.output, "")
			s.output = append(s.output, groupName)

			// save hooks from parent scope
			prevBE := s.beforeEach
			prevAE := s.afterEach
			prevBA := s.beforeAll
			prevAA := s.afterAll
			s.beforeEach = environment.NewNil()
			s.afterEach = environment.NewNil()
			s.beforeAll = environment.NewNil()
			s.afterAll = environment.NewNil()
			s.mu.Unlock()

			// Execute the group function — this registers hooks and runs tests immediately.
			// beforeAll(fn) executes the fn immediately when called within the group.
			// afterAll(fn) stores the fn to be called after the group finishes.
			if FigtestCaller != nil {
				FigtestCaller(groupFn, nil)
			}

			// run afterAll hook after group finishes
			s.mu.Lock()
			aa := s.afterAll
			s.mu.Unlock()
			if isCallable(aa) && FigtestCaller != nil {
				FigtestCaller(aa, nil)
			}

			// restore hooks and pop group
			s.mu.Lock()
			s.beforeEach = prevBE
			s.afterEach = prevAE
			s.beforeAll = prevBA
			s.afterAll = prevAA
			s.afterAll = prevAA
			if len(s.groups) > 0 {
				s.groups = s.groups[:len(s.groups)-1]
			}
			s.mu.Unlock()

			return environment.NewNil(), nil
		}),

		// skip(name, fn) — register a skipped test
		fn("skip", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 2 {
				return environment.NewNil(), fmt.Errorf("skip() expects 2 arguments (name, fn)")
			}
			name := args[0].String()
			s := globalFigtest
			s.mu.Lock()
			group := s.currentGroup()
			s.mu.Unlock()
			if group != "" {
				name = group + " > " + name
			}
			s.runTest(name, environment.NewNil(), true)
			return environment.NewNil(), nil
		}),

		// only(name, fn) — run only this test
		fn("only", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 2 {
				return environment.NewNil(), fmt.Errorf("only() expects 2 arguments (name, fn)")
			}
			name := args[0].String()
			testFn := args[1]
			if !isCallable(testFn) {
				return environment.NewNil(), fmt.Errorf("only() second argument must be a function")
			}
			// For simplicity with immediate execution, only() just runs the test
			// and marks everything else after it as skipped is complex.
			// Instead, we run it immediately.
			s := globalFigtest
			s.mu.Lock()
			group := s.currentGroup()
			s.mu.Unlock()
			if group != "" {
				name = group + " > " + name
			}
			s.runTest(name, testFn, false)
			return environment.NewNil(), nil
		}),

		// assert(cond, msg?) — fail if cond is falsy
		fn("assert", func(args []environment.Value) (environment.Value, error) {
			if err := figtestAssertFn(args, 1, "assert"); err != nil {
				return environment.NewNil(), err
			}
			if !isTruthy(args[0]) {
				msg := figtestMsg(args, 1, "assertion failed")
				return environment.NewNil(), fmt.Errorf("%s", msg)
			}
			return environment.NewBool(true), nil
		}),

		// assertEq(a, b, msg?) — fail if a != b
		fn("assertEq", func(args []environment.Value) (environment.Value, error) {
			if err := figtestAssertFn(args, 2, "assertEq"); err != nil {
				return environment.NewNil(), err
			}
			if args[0].String() != args[1].String() {
				msg := figtestMsg(args, 2, fmt.Sprintf("expected %s, got %s", args[1].String(), args[0].String()))
				return environment.NewNil(), fmt.Errorf("%s", msg)
			}
			return environment.NewBool(true), nil
		}),

		// assertNeq(a, b, msg?) — fail if a == b
		fn("assertNeq", func(args []environment.Value) (environment.Value, error) {
			if err := figtestAssertFn(args, 2, "assertNeq"); err != nil {
				return environment.NewNil(), err
			}
			if args[0].String() == args[1].String() {
				msg := figtestMsg(args, 2, fmt.Sprintf("expected values to differ, both are %s", args[0].String()))
				return environment.NewNil(), fmt.Errorf("%s", msg)
			}
			return environment.NewBool(true), nil
		}),

		// assertError(fn, msg?) — pass if fn throws error
		fn("assertError", func(args []environment.Value) (environment.Value, error) {
			if err := figtestAssertFn(args, 1, "assertError"); err != nil {
				return environment.NewNil(), err
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("assertError() first argument must be a function")
			}
			if FigtestCaller == nil {
				return environment.NewNil(), fmt.Errorf("figtest not properly initialized")
			}
			_, callErr := FigtestCaller(args[0], nil)
			if callErr == nil {
				msg := figtestMsg(args, 1, "expected function to throw an error, but it did not")
				return environment.NewNil(), fmt.Errorf("%s", msg)
			}
			return environment.NewBool(true), nil
		}),

		// assertNoError(fn, msg?) — fail if fn throws error
		fn("assertNoError", func(args []environment.Value) (environment.Value, error) {
			if err := figtestAssertFn(args, 1, "assertNoError"); err != nil {
				return environment.NewNil(), err
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("assertNoError() first argument must be a function")
			}
			if FigtestCaller == nil {
				return environment.NewNil(), fmt.Errorf("figtest not properly initialized")
			}
			_, callErr := FigtestCaller(args[0], nil)
			if callErr != nil {
				msg := figtestMsg(args, 1, fmt.Sprintf("expected no error, got: %s", callErr.Error()))
				return environment.NewNil(), fmt.Errorf("%s", msg)
			}
			return environment.NewBool(true), nil
		}),

		// assertNear(a, b, epsilon, msg?) — compare floats within epsilon
		fn("assertNear", func(args []environment.Value) (environment.Value, error) {
			if err := figtestAssertFn(args, 3, "assertNear"); err != nil {
				return environment.NewNil(), err
			}
			if args[0].Type != environment.NumberType || args[1].Type != environment.NumberType || args[2].Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("assertNear() expects 3 numbers")
			}
			diff := math.Abs(args[0].Num - args[1].Num)
			if diff > args[2].Num {
				msg := figtestMsg(args, 3, fmt.Sprintf("expected %g ≈ %g (epsilon %g), diff is %g", args[0].Num, args[1].Num, args[2].Num, diff))
				return environment.NewNil(), fmt.Errorf("%s", msg)
			}
			return environment.NewBool(true), nil
		}),

		// assertContains(container, item, msg?) — check if array/string contains item
		fn("assertContains", func(args []environment.Value) (environment.Value, error) {
			if err := figtestAssertFn(args, 2, "assertContains"); err != nil {
				return environment.NewNil(), err
			}
			container := args[0]
			item := args[1]
			found := false
			if container.Type == environment.StringType && item.Type == environment.StringType {
				found = strings.Contains(container.Str, item.Str)
			} else if container.Type == environment.ArrayType && container.Arr != nil {
				for _, elem := range *container.Arr {
					if elem.String() == item.String() {
						found = true
						break
					}
				}
			} else {
				return environment.NewNil(), fmt.Errorf("assertContains() first argument must be a string or array")
			}
			if !found {
				msg := figtestMsg(args, 2, fmt.Sprintf("expected %s to contain %s", container.String(), item.String()))
				return environment.NewNil(), fmt.Errorf("%s", msg)
			}
			return environment.NewBool(true), nil
		}),

		// assertType(val, typeName, msg?) — check if val has specific type
		fn("assertType", func(args []environment.Value) (environment.Value, error) {
			if err := figtestAssertFn(args, 2, "assertType"); err != nil {
				return environment.NewNil(), err
			}
			expected := args[1].String()
			actual := args[0].TypeName()
			if actual != expected {
				msg := figtestMsg(args, 2, fmt.Sprintf("expected type %s, got %s", expected, actual))
				return environment.NewNil(), fmt.Errorf("%s", msg)
			}
			return environment.NewBool(true), nil
		}),

		// assertLength(val, len, msg?) — check array/string length
		fn("assertLength", func(args []environment.Value) (environment.Value, error) {
			if err := figtestAssertFn(args, 2, "assertLength"); err != nil {
				return environment.NewNil(), err
			}
			if args[1].Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("assertLength() second argument must be a number")
			}
			expectedLen := int(args[1].Num)
			var actualLen int
			switch args[0].Type {
			case environment.StringType:
				actualLen = len([]rune(args[0].Str))
			case environment.ArrayType:
				if args[0].Arr != nil {
					actualLen = len(*args[0].Arr)
				}
			default:
				return environment.NewNil(), fmt.Errorf("assertLength() first argument must be a string or array")
			}
			if actualLen != expectedLen {
				msg := figtestMsg(args, 2, fmt.Sprintf("expected length %d, got %d", expectedLen, actualLen))
				return environment.NewNil(), fmt.Errorf("%s", msg)
			}
			return environment.NewBool(true), nil
		}),

		// beforeEach(fn)
		fn("beforeEach", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 || !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("beforeEach() expects a function")
			}
			s := globalFigtest
			s.mu.Lock()
			s.beforeEach = args[0]
			s.mu.Unlock()
			return environment.NewNil(), nil
		}),

		// afterEach(fn)
		fn("afterEach", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 || !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("afterEach() expects a function")
			}
			s := globalFigtest
			s.mu.Lock()
			s.afterEach = args[0]
			s.mu.Unlock()
			return environment.NewNil(), nil
		}),

		// beforeAll(fn) — execute immediately within a describe scope
		fn("beforeAll", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 || !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("beforeAll() expects a function")
			}
			s := globalFigtest
			s.mu.Lock()
			s.beforeAll = args[0]
			s.mu.Unlock()
			// Execute immediately so setup runs before any tests in the group
			if FigtestCaller != nil {
				FigtestCaller(args[0], nil)
			}
			return environment.NewNil(), nil
		}),

		// afterAll(fn)
		fn("afterAll", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 || !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("afterAll() expects a function")
			}
			s := globalFigtest
			s.mu.Lock()
			s.afterAll = args[0]
			s.mu.Unlock()
			return environment.NewNil(), nil
		}),

		// summary() — return summary text (CLI prints this automatically)
		fn("summary", func(args []environment.Value) (environment.Value, error) {
			s := globalFigtest
			s.mu.Lock()
			defer s.mu.Unlock()
			var sb strings.Builder
			for _, line := range s.output {
				sb.WriteString(line)
				sb.WriteByte('\n')
			}
			total := s.passCount + s.failCount + s.skipCount
			sb.WriteString(fmt.Sprintf("\n%d passed, %d failed, %d skipped (total: %d)\n", s.passCount, s.failCount, s.skipCount, total))
			return environment.NewString(sb.String()), nil
		}),

		// reset() — clear all test state
		fn("reset", func(args []environment.Value) (environment.Value, error) {
			ResetFigtest()
			return environment.NewNil(), nil
		}),

		// count() — total test count
		fn("count", func(args []environment.Value) (environment.Value, error) {
			s := globalFigtest
			s.mu.Lock()
			defer s.mu.Unlock()
			return environment.NewNumber(float64(s.passCount + s.failCount + s.skipCount)), nil
		}),

		// passed() — passed count
		fn("passed", func(args []environment.Value) (environment.Value, error) {
			s := globalFigtest
			return environment.NewNumber(float64(s.passCount)), nil
		}),

		// failed() — failed count
		fn("failed", func(args []environment.Value) (environment.Value, error) {
			s := globalFigtest
			return environment.NewNumber(float64(s.failCount)), nil
		}),

		// skipped() — skipped count
		fn("skipped", func(args []environment.Value) (environment.Value, error) {
			s := globalFigtest
			return environment.NewNumber(float64(s.skipCount)), nil
		}),
	))
}
