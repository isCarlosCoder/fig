package interpreter

import (
	"fmt"
	"strings"
)

// RuntimeError represents an error that occurred during interpretation.
type RuntimeError struct {
	Line        int
	Column      int
	Message     string
	Snippet     string // the source line where the error happened
	ColumnStart int    // start column in the Snippet (0-based)
	Length      int    // length of the highlighted span
}

// Error returns a plain-text representation (no ANSI colors) for use in Go error chains.
func (e *RuntimeError) Error() string {
	if e.Line > 0 {
		base := fmt.Sprintf("%d:%d: runtime error: %s", e.Line, e.Column, e.Message)
		if e.Snippet != "" {
			caret := strings.Repeat(" ", e.ColumnStart)
			if e.Length <= 0 {
				e.Length = 1
			}
			caret = caret + strings.Repeat("^", e.Length)
			return fmt.Sprintf("%s\n    %s\n    %s", base, e.Snippet, caret)
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
		sb.WriteString(fmt.Sprintf("%d:%d: %sruntime error: %s%s\n", e.Line, e.Column, ansiRed, e.Message, ansiReset))
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

	return sb.String()
}
