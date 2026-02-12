//go:build linux

package builtins

import (
	"syscall"
	"unsafe"
)

// prlimit sets resource limits using the prlimit64 syscall on Linux.
func prlimit(resource int, lim *syscall.Rlimit) error {
	// SYS_PRLIMIT64 = 302 on x86_64, 369 on arm64
	_, _, e := syscall.RawSyscall6(
		syscall.SYS_PRLIMIT64,
		0, // pid 0 = current process
		uintptr(resource),
		uintptr(unsafe.Pointer(lim)),
		0,
		0,
		0,
	)
	if e != 0 {
		return e
	}
	return nil
}
