//go:build !linux

package builtins

import (
	"syscall"
)

// prlimit is a no-op on non-Linux platforms.
// Resource limits are applied via environment variables or OS-specific means.
func prlimit(_ int, _ *syscall.Rlimit) error {
	// no-op: prlimit not available on this platform
	return nil
}
