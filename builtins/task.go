package builtins

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/iscarloscoder/fig/environment"
)

// TaskSpawner is set by the interpreter so that task.spawn can execute
// user-defined Fig functions in separate goroutines.
// The callback MUST start a new goroutine, execute fn, and send the result to resultCh.
var TaskSpawner func(fn environment.Value, resultCh chan<- TaskResult)

// TaskResult holds the return value or error from a completed task.
type TaskResult struct {
	Value environment.Value
	Err   error
}

// taskHandle is the internal representation of a spawned task.
type taskHandle struct {
	id     int64
	done   chan struct{} // closed when the task completes
	result TaskResult   // populated before done is closed
}

var (
	taskCounter  int64
	taskRegistry sync.Map // int64 → *taskHandle
)

// ResetTaskState clears all task state (for testing).
func ResetTaskState() {
	taskRegistry = sync.Map{}
	atomic.StoreInt64(&taskCounter, 0)
}

// getTaskHandle extracts a *taskHandle from a Value (expected: object with __task_id).
func getTaskHandle(v environment.Value) (*taskHandle, error) {
	if v.Type != environment.ObjectType || v.Obj == nil {
		return nil, fmt.Errorf("expected a task handle")
	}
	idVal, ok := v.Obj.Entries["__task_id"]
	if !ok || idVal.Type != environment.NumberType {
		return nil, fmt.Errorf("expected a task handle")
	}
	id := int64(idVal.Num)
	raw, ok := taskRegistry.Load(id)
	if !ok {
		return nil, fmt.Errorf("task not found (id %d)", id)
	}
	return raw.(*taskHandle), nil
}

func init() {
	register(newModule("task",
		// spawn(fn) → task handle
		fn("spawn", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("spawn() expects 1 argument, got %d", len(args))
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("spawn() argument must be a function")
			}
			if TaskSpawner == nil {
				return environment.NewNil(), fmt.Errorf("task module not properly initialized")
			}

			id := atomic.AddInt64(&taskCounter, 1)
			h := &taskHandle{
				id:   id,
				done: make(chan struct{}),
			}
			taskRegistry.Store(id, h)

			resultCh := make(chan TaskResult, 1)

			// Launch the spawner (interpreter creates goroutine)
			TaskSpawner(args[0], resultCh)

			// Bridge goroutine: waits for result and signals done
			go func() {
				r := <-resultCh
				h.result = r
				close(h.done)
			}()

			// Return task handle as object
			entries := map[string]environment.Value{
				"__task_id": environment.NewNumber(float64(id)),
			}
			return environment.NewObject(entries, []string{"__task_id"}), nil
		}),

		// await(handle) → value (blocks until task completes)
		fn("await", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("await() expects 1 argument, got %d", len(args))
			}
			h, err := getTaskHandle(args[0])
			if err != nil {
				return environment.NewNil(), fmt.Errorf("await(): %v", err)
			}
			<-h.done
			if h.result.Err != nil {
				return environment.NewNil(), h.result.Err
			}
			return h.result.Value, nil
		}),

		// awaitTimeout(handle, ms) → value (blocks with timeout)
		fn("awaitTimeout", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("awaitTimeout() expects 2 arguments, got %d", len(args))
			}
			h, err := getTaskHandle(args[0])
			if err != nil {
				return environment.NewNil(), fmt.Errorf("awaitTimeout(): %v", err)
			}
			if args[1].Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("awaitTimeout() timeout must be a number")
			}
			ms := int64(args[1].Num)

			select {
			case <-h.done:
				if h.result.Err != nil {
					return environment.NewNil(), h.result.Err
				}
				return h.result.Value, nil
			case <-time.After(time.Duration(ms) * time.Millisecond):
				return environment.NewNil(), fmt.Errorf("task timed out after %dms", ms)
			}
		}),

		// race(tasks) → value (returns first completed result)
		fn("race", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("race() expects 1 argument (array of tasks), got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("race() argument must be an array of task handles")
			}

			tasks := *args[0].Arr
			if len(tasks) == 0 {
				return environment.NewNil(), fmt.Errorf("race() requires at least one task")
			}

			handles := make([]*taskHandle, len(tasks))
			for i, t := range tasks {
				h, err := getTaskHandle(t)
				if err != nil {
					return environment.NewNil(), fmt.Errorf("race(): element %d: %v", i, err)
				}
				handles[i] = h
			}

			// Wait for the first one to complete
			winCh := make(chan TaskResult, len(handles))
			for _, h := range handles {
				go func(h *taskHandle) {
					<-h.done
					select {
					case winCh <- h.result:
					default:
					}
				}(h)
			}

			r := <-winCh
			if r.Err != nil {
				return environment.NewNil(), r.Err
			}
			return r.Value, nil
		}),
	))
}
