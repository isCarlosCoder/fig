package tests

import (
	"testing"
	"time"
)

func useTask(code string) string {
	return `use "task"; use "system"; ` + code
}

func TestTaskSpawnAwait(t *testing.T) {
	src := useTask(`let t = task.spawn(fn() { return 42 }); print(task.await(t));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "42" {
		t.Fatalf("expected '42', got %q", out)
	}
}

func TestTaskSpawnReturnValue(t *testing.T) {
	src := useTask(`let t = task.spawn(fn() { return 21 * 2 }); let v = task.await(t); print(v);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "42" {
		t.Fatalf("expected '42', got %q", out)
	}
}

func TestTaskSpawnWithClosure(t *testing.T) {
	src := useTask(`let x = 10; let t = task.spawn(fn() { return x * 5 }); print(task.await(t));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "50" {
		t.Fatalf("expected '50', got %q", out)
	}
}

func TestTaskSpawnCallsFunction(t *testing.T) {
	src := useTask(`fn work(n) { return n * 10 } let t = task.spawn(fn() { return work(5) }); print(task.await(t));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "50" {
		t.Fatalf("expected '50', got %q", out)
	}
}

func TestTaskFireAndForget(t *testing.T) {
	src := useTask(`task.spawn(fn() { return 1 }); print("ok");`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "ok" {
		t.Fatalf("expected 'ok', got %q", out)
	}
}

func TestTaskAwaitTimeout(t *testing.T) {
	src := useTask(`let t = task.spawn(fn() { return 99 }); print(task.awaitTimeout(t, 1000));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "99" {
		t.Fatalf("expected '99', got %q", out)
	}
}

func TestTaskAwaitTimeoutExpires(t *testing.T) {
	src := useTask(`let t = task.spawn(fn() { system.sleep(2000); return 99 }); let v = try task.awaitTimeout(t, 100) onerror { return -1 }; print(v);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "-1" {
		t.Fatalf("expected '-1', got %q", out)
	}
}

func TestTaskRace(t *testing.T) {
	src := useTask(`let t1 = task.spawn(fn() { system.sleep(500); return "slow" }); let t2 = task.spawn(fn() { return "fast" }); print(task.race([t1, t2]));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "fast" {
		t.Fatalf("expected 'fast', got %q", out)
	}
}

func TestTaskRaceEmptyArray(t *testing.T) {
	src := useTask(`task.race([]);`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for empty race array")
	}
}

func TestTaskParallel(t *testing.T) {
	start := time.Now()
	src := useTask(`let t1 = task.spawn(fn() { system.sleep(100); return 1 }); let t2 = task.spawn(fn() { system.sleep(100); return 2 }); let t3 = task.spawn(fn() { system.sleep(100); return 3 }); print(task.await(t1)); print(task.await(t2)); print(task.await(t3));`)
	out, err := runFig(t, src)
	elapsed := time.Since(start)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "1\n2\n3" {
		t.Fatalf("expected '1\\n2\\n3', got %q", out)
	}
	if elapsed > 300*time.Millisecond {
		t.Fatalf("tasks should run in parallel; took %v", elapsed)
	}
}

func TestTaskPipeline(t *testing.T) {
	src := useTask(`let t1 = task.spawn(fn() { return 10 }); let t2 = task.spawn(fn() { return task.await(t1) * 3 }); print(task.await(t2));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "30" {
		t.Fatalf("expected '30', got %q", out)
	}
}

func TestTaskErrorPropagation(t *testing.T) {
	src := `use "task"; use "debug"; let t = task.spawn(fn() { debug.panic("task failed") }); let v = try task.await(t) onerror(e) { return "caught" }; print(v);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "caught" {
		t.Fatalf("expected 'caught', got %q", out)
	}
}

func TestTaskSpawnNoArgs(t *testing.T) {
	src := useTask(`task.spawn();`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for missing args")
	}
}

func TestTaskSpawnNonFunction(t *testing.T) {
	src := useTask(`task.spawn("not_fn");`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for non-function argument")
	}
}

func TestTaskAwaitNonTask(t *testing.T) {
	src := useTask(`task.await(42);`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for non-task argument")
	}
}

func TestTaskAwaitTimeoutNoArgs(t *testing.T) {
	src := useTask(`task.awaitTimeout();`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for missing args")
	}
}

func TestTaskAwaitTimeoutWrongType(t *testing.T) {
	src := useTask(`let t = task.spawn(fn() { return 1 }); task.awaitTimeout(t, "abc");`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for non-number timeout")
	}
}

func TestTaskRaceNonArray(t *testing.T) {
	src := useTask(`task.race(42);`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for non-array argument")
	}
}

func TestTaskSpawnReturnsNull(t *testing.T) {
	src := useTask(`let t = task.spawn(fn() {}); print(task.await(t));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "null" {
		t.Fatalf("expected 'null', got %q", out)
	}
}
