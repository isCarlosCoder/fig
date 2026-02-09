package interpreter

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// PrettyErrorListener prints user-friendly, colored parsing errors with context.
// It highlights the offending token in red and prints the source line with a caret.
type PrettyErrorListener struct {
	filename string
	srcLines []string
	out      io.Writer
	Occurred bool
	Err      error
	// AbortOnError if true will cause the listener to panic after reporting
	// the first syntax error. Callers that want parsing to stop immediately
	// (e.g., interactive runs) should set this to true; tests should leave
	// it false to avoid panics.
	AbortOnError bool
}

// ANSI color codes
const (
	ansiReset   = "\x1b[0m"
	ansiRed     = "\x1b[1;31m"
	ansiDimCyan = "\x1b[2;36m"
	ansiDim     = "\x1b[2m"
)

func NewPrettyErrorListener(source string, filename string, out io.Writer) *PrettyErrorListener {
	if out == nil {
		out = os.Stderr
	}
	lines := strings.Split(source, "\n")
	return &PrettyErrorListener{filename: filename, srcLines: lines, out: out}
}

// SyntaxError is called by ANTLR when a syntax error occurs.
func (l *PrettyErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// Attempt to produce a friendlier message for common ANTLR outputs
	// e.g., "mismatched input '/' expecting {...}" -> "unexpected token '/'"
	tokText := ""
	switch t := offendingSymbol.(type) {
	case antlr.Token:
		tokText = t.GetText()
	case antlr.Tree:
		// nothing
	default:
		_ = t
	}
	if strings.HasPrefix(msg, "mismatched input") {
		if tokText != "" {
			msg = fmt.Sprintf("unexpected token %q", tokText)
		} else {
			msg = "syntax error"
		}
	}

	// Header: file:line:col: error: message
	fmt.Fprintf(l.out, "%s:%d:%d: %serror: %s%s\n", l.filename, line, column, ansiRed, msg, ansiReset)
	l.Occurred = true
	l.Err = fmt.Errorf("%s:%d:%d: %s", l.filename, line, column, msg)

	// Get the source line if available (ANTLR lines are 1-based)
	if line-1 >= 0 && line-1 < len(l.srcLines) {
		srcLine := l.srcLines[line-1]

		start := column
		if start < 0 {
			start = 0
		}
		// token length
		length := 1
		if tokText != "" {
			length = len(tokText)
			// don't overflow
			if start+length > len(srcLine) {
				length = len(srcLine) - start
			}
		}

		// Compose colored line
		left := srcLine
		mid := ""
		right := ""
		if start < len(srcLine) {
			left = srcLine[:start]
			if length > 0 && start+length <= len(srcLine) {
				mid = srcLine[start : start+length]
				right = srcLine[start+length:]
			} else {
				mid = srcLine[start:]
				right = ""
			}
		}

		// Print the line with context coloring
		fmt.Fprint(l.out, "    ") // simple gutter
		if left != "" {
			fmt.Fprint(l.out, ansiDimCyan)
			fmt.Fprint(l.out, left)
			fmt.Fprint(l.out, ansiReset)
		}
		if mid != "" {
			fmt.Fprint(l.out, ansiRed)
			fmt.Fprint(l.out, mid)
			fmt.Fprint(l.out, ansiReset)
		}
		if right != "" {
			fmt.Fprint(l.out, ansiDim)
			fmt.Fprint(l.out, right)
			fmt.Fprint(l.out, ansiReset)
		}
		fmt.Fprint(l.out, "\n")

		// Print caret line
		fmt.Fprint(l.out, "    ")
		for i := 0; i < start; i++ {
			fmt.Fprint(l.out, " ")
		}
		if length <= 0 {
			length = 1
		}
		fmt.Fprint(l.out, ansiRed)
		for i := 0; i < length; i++ {
			fmt.Fprint(l.out, "^")
		}
		fmt.Fprint(l.out, ansiReset)
		fmt.Fprint(l.out, "\n")
	}

	// Abort parsing early if requested by the caller
	if l.AbortOnError {
		panic(l.Err)
	}
}

// Other error listener methods are no-ops for now
func (l *PrettyErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}
func (l *PrettyErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}
func (l *PrettyErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}
