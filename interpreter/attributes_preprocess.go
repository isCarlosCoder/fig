package interpreter

import (
	"regexp"
	"strings"
)

// preprocessAttributes rewrites a source-level attribute syntax like:
//
//	@native fn name(...) { ... }
//	@[native(...)] fn name(...) { ... }
//
// into a top-level metadata var so the existing parser doesn't need grammar
// changes. Example transformation:
//
//	let __native_meta_name = "fallback=true"\nfn name(...) { ... }
//
// The visitor inspects that metadata var when visiting the function declaration.
// Supported forms: @native, @native(), @native(k=v,...), @[native], @[native(...)].
func preprocessAttributes(src string) string {
	// Regex captures: 1 = optional args inside (), 2 = function name
	re := regexp.MustCompile(`(?m)@\[?native(?:\(([^)]*)\))?\]?\s*fn\s+([A-Za-z_][A-Za-z0-9_]*)`)
	return re.ReplaceAllStringFunc(src, func(m string) string {
		sub := re.FindStringSubmatch(m)
		if len(sub) < 3 {
			return m
		}
		args := strings.TrimSpace(sub[1])
		name := sub[2]
		// encode args as-is (visitor will parse key=value pairs)
		meta := "let __native_meta_" + name + " = \"" + strings.ReplaceAll(args, "\"", "\\\"") + "\"\n"
		// remove the attribute prefix from the matched text: keep only the `fn name(` part
		repl := re.ReplaceAllString(m, "fn "+name)
		return meta + repl
	})
}
