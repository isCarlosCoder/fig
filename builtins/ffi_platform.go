package builtins

import "runtime"

// LibExt returns the shared library extension for the current OS.
// Linux: ".so", macOS: ".dylib", Windows: ".dll"
func LibExt() string {
	switch runtime.GOOS {
	case "darwin":
		return ".dylib"
	case "windows":
		return ".dll"
	default: // linux, freebsd, etc.
		return ".so"
	}
}

// LibName returns a conventional shared library filename for the given base name.
// Linux: "lib<base>.so", macOS: "lib<base>.dylib", Windows: "<base>.dll"
func LibName(base string) string {
	switch runtime.GOOS {
	case "windows":
		return base + ".dll"
	default: // linux, darwin, freebsd
		return "lib" + base + LibExt()
	}
}
