package environment

import "fmt"

// ExitSignal is used to signal a controlled process exit
// from Fig code (e.g. system.exit()). It replaces a direct
// os.Exit() call so the signal can be intercepted by the
// test runner or other callers.
type ExitSignal struct {
	Code int
}

func (e ExitSignal) Error() string {
	return fmt.Sprintf("exit(%d)", e.Code)
}

// IsExitSignal checks if an error is an ExitSignal and returns its code.
func IsExitSignal(err error) (int, bool) {
	if es, ok := err.(ExitSignal); ok {
		return es.Code, true
	}
	return 0, false
}
