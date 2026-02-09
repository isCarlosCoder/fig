package interpreter

import (
	"fmt"
	"strings"
)

// RuntimeError represents an error that occurred during interpretation.
type StackFrame struct {
	Kind   string // "function" or "module"
	Name   string
	File   string
	Line   int
	Column int
}

// RuntimeError represents an error that occurred during interpretation.
type RuntimeError struct {
	File        string // optional file path where the error occurred
	Line        int
	Column      int
	Message     string
	Snippet     string // the source line where the error happened
	ColumnStart int    // start column in the Snippet (0-based)
	Length      int    // length of the highlighted span
	Frames      []StackFrame
}

// Error returns a plain-text representation (no ANSI colors) for use in Go error chains.
func (e *RuntimeError) Error() string {
	if e.Line > 0 {
		// include file path when available: file:line:col: runtime error: msg
		var base string
		if e.File != "" {
			base = fmt.Sprintf("%s:%d:%d: runtime error: %s", e.File, e.Line, e.Column, e.Message)
		} else {
			base = fmt.Sprintf("%d:%d: runtime error: %s", e.Line, e.Column, e.Message)
		}
		// If we have frames, annotate header with the most-recent frame for more direct context
		if len(e.Frames) > 0 {
			f := e.Frames[len(e.Frames)-1]
			if f.Name != "" {
				base = base + fmt.Sprintf(" — in %s (%s:%d:%d)", f.Name, f.File, f.Line, f.Column)
			} else {
				base = base + fmt.Sprintf(" — in %s:%d:%d", f.File, f.Line, f.Column)
			}
		}

		if e.Snippet != "" {
			caret := strings.Repeat(" ", e.ColumnStart)
			if e.Length <= 0 {
				e.Length = 1
			}
			caret = caret + strings.Repeat("^", e.Length)
			base = fmt.Sprintf("%s\n    %s\n    %s", base, e.Snippet, caret)
		}
		// Append stack frames if available
		if len(e.Frames) > 0 {
			base = base + "\nstack trace:"
			for _, f := range e.Frames {
				base = base + fmt.Sprintf("\n    at %s (%s:%d:%d)", f.Name, f.File, f.Line, f.Column)
			}
		}
		return base
	}
	return fmt.Sprintf("runtime error: %s", e.Message)
}

// PrettyError returns a colorized version of the error with ANSI escape codes,
// matching the style of syntax errors from PrettyErrorListener.
func (e *RuntimeError) PrettyError() string {
	var sb strings.Builder

	if e.Line > 0 {
		// Header: line:col: runtime error: message (with red highlight)
		if e.File != "" {
			sb.WriteString(fmt.Sprintf("%s:%d:%d: %sruntime error: %s%s\n", e.File, e.Line, e.Column, ansiRed, e.Message, ansiReset))
		} else {
			sb.WriteString(fmt.Sprintf("%d:%d: %sruntime error: %s%s\n", e.Line, e.Column, ansiRed, e.Message, ansiReset))
		}
		if e.Snippet != "" {
			length := e.Length
			if length <= 0 {
				length = 1
			}
			start := e.ColumnStart

			// Build colored snippet: dim cyan before, red on error, dim after
			left := e.Snippet
			mid := ""
			right := ""
			if start < len(e.Snippet) {
				left = e.Snippet[:start]
				if start+length <= len(e.Snippet) {
					mid = e.Snippet[start : start+length]
					right = e.Snippet[start+length:]
				} else {
					mid = e.Snippet[start:]
				}
			}

			sb.WriteString("    ")
			if left != "" {
				sb.WriteString(ansiDimCyan)
				sb.WriteString(left)
				sb.WriteString(ansiReset)
			}
			if mid != "" {
				sb.WriteString(ansiRed)
				sb.WriteString(mid)
				sb.WriteString(ansiReset)
			}
			if right != "" {
				sb.WriteString(ansiDim)
				sb.WriteString(right)
				sb.WriteString(ansiReset)
			}
			sb.WriteString("\n")

			// Caret line
			sb.WriteString("    ")
			sb.WriteString(strings.Repeat(" ", start))
			sb.WriteString(ansiRed)
			sb.WriteString(strings.Repeat("^", length))
			sb.WriteString(ansiReset)
			sb.WriteString("\n")
		}
	} else {
		sb.WriteString(fmt.Sprintf("%sruntime error: %s%s\n", ansiRed, e.Message, ansiReset))
	}

	// Append stack trace if available
	if len(e.Frames) > 0 {
		sb.WriteString("\n")
		sb.WriteString("Stack trace:\n")
		for _, f := range e.Frames {
			sb.WriteString(fmt.Sprintf("    at %s (%s:%d:%d)\n", f.Name, f.File, f.Line, f.Column))
		}
	}

	return sb.String()
}
