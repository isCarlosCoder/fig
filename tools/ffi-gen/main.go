// ffi-gen reads an ffi.def.toml definition file and generates Fig (.fig) wrapper code
// that calls ffi.load / ffi.sym / ffi.call for each declared function and struct.
//
// Usage:
//
//	go run ./tools/ffi-gen -input mylib.ffi.def.toml -output mylib.fig
//	go run ./tools/ffi-gen -input mylib.ffi.def.toml              # prints to stdout
//	go run ./tools/ffi-gen -init myproject                         # scaffolds a new FFI project
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

// SupportedTypes are the types accepted by the FFI protocol.
var SupportedTypes = map[string]bool{
	"int":    true,
	"double": true,
	"string": true,
	"void":   true,
}

// IsSupportedType checks if a type string is valid (including struct:Name).
func IsSupportedType(t string) bool {
	if SupportedTypes[t] {
		return true
	}
	if strings.HasPrefix(t, "struct:") && len(t) > 7 {
		return true
	}
	return false
}

// ---------- ffi.def.toml schema ----------

// DefFile is the top-level structure of an ffi.def.toml file.
type DefFile struct {
	Library   LibraryDef    `toml:"library"`
	Structs   []StructDef   `toml:"structs"`
	Functions []FunctionDef `toml:"functions"`
}

// LibraryDef describes the shared library to load.
type LibraryDef struct {
	Name string `toml:"name"` // base name, e.g. "mymath"
	Path string `toml:"path"` // optional explicit path (overrides name)
}

// StructDef declares a struct schema.
type StructDef struct {
	Name   string     `toml:"name"`
	Fields []FieldDef `toml:"fields"`
}

// FieldDef is a single struct field.
type FieldDef struct {
	Name string `toml:"name"`
	Type string `toml:"type"` // "int", "double", "string", "struct:X"
}

// FunctionDef declares a function binding.
type FunctionDef struct {
	Name   string   `toml:"name"`   // Fig function name
	Symbol string   `toml:"symbol"` // C symbol (defaults to Name)
	Return string   `toml:"return"` // return type
	Args   []string `toml:"args"`   // argument types
}

// ---------- code generation ----------

// ValidateDef checks all types in a DefFile against supported types.
// Returns a slice of human-readable error strings (empty = valid).
func ValidateDef(def *DefFile) []string {
	var errs []string
	for _, s := range def.Structs {
		for _, f := range s.Fields {
			if !IsSupportedType(f.Type) {
				errs = append(errs, fmt.Sprintf("struct %s, field %s: unknown type: %s. Supported: int, double, string, void, struct:Name", s.Name, f.Name, f.Type))
			}
		}
	}
	for _, f := range def.Functions {
		if f.Return != "" && !IsSupportedType(f.Return) {
			errs = append(errs, fmt.Sprintf("function %s: unknown return type: %s. Supported: int, double, string, void, struct:Name", f.Name, f.Return))
		}
		for i, a := range f.Args {
			if !IsSupportedType(a) {
				errs = append(errs, fmt.Sprintf("function %s, arg %d: unknown type: %s. Supported: int, double, string, void, struct:Name", f.Name, i+1, a))
			}
		}
	}
	return errs
}

// FigTomlFFI represents the [ffi] section of fig.toml.
type FigTomlFFI struct {
	Enabled     bool   `toml:"enabled"`
	Helper      string `toml:"helper"`
	CallTimeout int    `toml:"call_timeout"`
	APIVersion  string `toml:"api_version"`
}

// FigToml represents the top-level fig.toml configuration.
type FigToml struct {
	FFI FigTomlFFI `toml:"ffi"`
}

// ValidateFigToml reads and validates a fig.toml file for FFI readiness.
// Returns a slice of human-readable error strings (empty = valid).
func ValidateFigToml(path string) []string {
	var errs []string

	data, err := os.ReadFile(path)
	if err != nil {
		return []string{fmt.Sprintf("cannot read fig.toml: %v", err)}
	}

	var cfg FigToml
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return []string{fmt.Sprintf("cannot parse fig.toml: %v", err)}
	}

	if !cfg.FFI.Enabled {
		errs = append(errs, "fig.toml: ffi.enabled is not true")
	}

	if cfg.FFI.Helper == "" {
		errs = append(errs, "fig.toml: ffi.helper is not set")
	} else {
		if _, err := os.Stat(cfg.FFI.Helper); os.IsNotExist(err) {
			errs = append(errs, fmt.Sprintf("fig.toml: ffi.helper binary not found: %s", cfg.FFI.Helper))
		}
	}

	return errs
}

// CheckHelper verifies that the helper binary exists and is executable.
func CheckHelper(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("helper binary not found: %s", path)
	}
	if info.IsDir() {
		return fmt.Errorf("helper path is a directory: %s", path)
	}
	// Check executable bit (unix)
	if info.Mode()&0111 == 0 {
		return fmt.Errorf("helper binary is not executable: %s", path)
	}
	return nil
}

// buildCSignature creates a C-style signature comment like "# C: int add(int a, int b)"
func buildCSignature(f FunctionDef) string {
	sym := f.Symbol
	if sym == "" {
		sym = f.Name
	}
	ret := f.Return
	if ret == "" {
		ret = "int"
	}
	params := make([]string, len(f.Args))
	for i, a := range f.Args {
		params[i] = fmt.Sprintf("%s %s", a, paramName(i))
	}
	return fmt.Sprintf("# C: %s %s(%s)", ret, sym, strings.Join(params, ", "))
}

// Generate produces Fig source code from a DefFile.
func Generate(def *DefFile) string {
	var b strings.Builder

	libName := def.Library.Name
	if libName == "" {
		libName = "unknown"
	}

	b.WriteString(fmt.Sprintf("# Auto-generated FFI bindings for %s\n", libName))
	b.WriteString("# Generated by: fig ffi-gen — do not edit manually\n\n")
	b.WriteString("use \"ffi\"\n\n")

	// Library load
	if def.Library.Path != "" {
		b.WriteString(fmt.Sprintf("let __lib = ffi.load(\"%s\" + ffi.lib_ext())\n\n", def.Library.Path))
	} else {
		b.WriteString(fmt.Sprintf("let __lib = ffi.load(\"./\" + ffi.lib_name(\"%s\"))\n\n", libName))
	}

	// Struct definitions (use high-level wrapper API)
	if len(def.Structs) > 0 {
		b.WriteString("# --- Struct definitions ---\n\n")
		for _, s := range def.Structs {
			// declare wrapper variable with same name as struct
			b.WriteString(fmt.Sprintf("let %s = ffi.struct_(\"%s\", [\n", s.Name, s.Name))
			for i, f := range s.Fields {
				comma := ","
				if i == len(s.Fields)-1 {
					comma = ""
				}
				b.WriteString(fmt.Sprintf("    { name: \"%s\", type: \"%s\" }%s\n", f.Name, f.Type, comma))
			}
			b.WriteString("])\n\n")
		}
	}

	// Function wrappers
	if len(def.Functions) > 0 {
		b.WriteString("# --- Function wrappers ---\n\n")
		for _, f := range def.Functions {
			sym := f.Symbol
			if sym == "" {
				sym = f.Name
			}
			ret := f.Return
			if ret == "" {
				ret = "int"
			}

			// C signature comment
			b.WriteString(buildCSignature(f) + "\n")

			// symbol variable
			symVar := "__sym_" + f.Name

			// Always generate with argTypes for type safety; if an arg is a struct
			// and we defined a wrapper variable, reference it directly instead of a
			// literal string.
			argTypesLit := "["
			for i, a := range f.Args {
				if i > 0 {
					argTypesLit += ", "
				}
				if strings.HasPrefix(a, "struct:") {
					sn := a[7:]
					argTypesLit += sn // uses variable name
				} else {
					argTypesLit += fmt.Sprintf("\"%s\"", a)
				}
			}
			argTypesLit += "]"
			b.WriteString(fmt.Sprintf("let %s = ffi.sym(__lib, \"%s\", \"%s\", %s)\n", symVar, sym, ret, argTypesLit))

			// function wrapper
			params := buildParamList(len(f.Args))
			b.WriteString(fmt.Sprintf("fn %s(%s) {\n", f.Name, params))
			callArgs := symVar
			for i := range f.Args {
				callArgs += fmt.Sprintf(", %s", paramName(i))
			}
			b.WriteString(fmt.Sprintf("    return ffi.call(%s)\n", callArgs))
			b.WriteString("}\n\n")
		}
	}

	return b.String()
}

// needsArgTypes returns true if the arg types list contains mixed types
// (not all the same) which requires explicit arg_types metadata.
func needsArgTypes(args []string) bool {
	// Always need arg_types for struct types
	for _, a := range args {
		if strings.HasPrefix(a, "struct:") {
			return true
		}
	}
	if len(args) <= 1 {
		return false
	}
	first := args[0]
	for _, a := range args[1:] {
		if a != first {
			return true
		}
	}
	return false
}

func buildParamList(n int) string {
	parts := make([]string, n)
	for i := range parts {
		parts[i] = paramName(i)
	}
	return strings.Join(parts, ", ")
}

func paramName(i int) string {
	// a, b, c, ..., z, a1, b1, ...
	base := string(rune('a' + i%26))
	if i >= 26 {
		return fmt.Sprintf("%s%d", base, i/26)
	}
	return base
}

// ---------- project scaffolding ----------

func scaffoldProject(name string) error {
	dir := name
	// Use only the base directory name for filenames
	baseName := filepath.Base(name)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir %s: %v", dir, err)
	}

	// fig.toml
	figToml := `[project]
name = "` + baseName + `"
type = 'application'
main = './main.fig'

[ffi]
enabled = true
helper = "ffi-helper" # Set this to the path of your ffi-helper binary
# call_timeout in milliseconds (0 = unlimited, omit for default 3000)
call_timeout = 5000
`
	if err := os.WriteFile(filepath.Join(dir, "fig.toml"), []byte(figToml), 0644); err != nil {
		return err
	}

	// ffi.def.toml (example definition)
	defToml := `# FFI binding definition for ` + baseName + `
# Edit this file and run: go run ./tools/ffi-gen -input ` + baseName + `.ffi.def.toml -output bindings.fig

[library]
name = "` + baseName + `"

[[structs]]
name = "Point"
fields = [
    { name = "x", type = "int" },
    { name = "y", type = "int" }
]

[[functions]]
name = "add"
symbol = "add"
return = "int"
args = ["int", "int"]

[[functions]]
name = "greet"
symbol = "greet"
return = "string"
args = ["string"]
`
	if err := os.WriteFile(filepath.Join(dir, baseName+".ffi.def.toml"), []byte(defToml), 0644); err != nil {
		return err
	}

	// example C source
	cSrc := `#include <stdio.h>
#include <string.h>

int add(int a, int b) {
    return a + b;
}

const char* greet(const char* name) {
    static char buf[256];
    snprintf(buf, sizeof(buf), "Hello, %s!", name);
    return buf;
}
`
	if err := os.WriteFile(filepath.Join(dir, baseName+".c"), []byte(cSrc), 0644); err != nil {
		return err
	}

	// Makefile
	makefile := `# Build shared library and generate Fig bindings
LIB_EXT := $(shell uname -s | grep -qi darwin && echo .dylib || (uname -s | grep -qi mingw && echo .dll || echo .so))

all: lib` + baseName + `$(LIB_EXT) bindings.fig

lib` + baseName + `$(LIB_EXT): ` + baseName + `.c
	gcc -shared -fPIC -o $@ $<

bindings.fig: ` + baseName + `.ffi.def.toml
	go run ./tools/ffi-gen -input $< -output $@

clean:
	rm -f lib` + baseName + `$(LIB_EXT) bindings.fig

.PHONY: all clean
`
	if err := os.WriteFile(filepath.Join(dir, "Makefile"), []byte(makefile), 0644); err != nil {
		return err
	}

	// main.fig (example usage)
	mainFig := `use "ffi"

# Load auto-generated bindings
# (run: go run ./tools/ffi-gen -input ` + baseName + `.ffi.def.toml -output bindings.fig)
# Then: use "bindings"

# Or use FFI directly:
let lib = ffi.load("./" + ffi.lib_name("` + baseName + `"))
let add_sym = ffi.sym(lib, "add", "int")
print(ffi.call(add_sym, 10, 20))    # 30

let greet_sym = ffi.sym(lib, "greet", "string")
print(ffi.call(greet_sym, "Fig"))    # Hello, Fig!
`
	if err := os.WriteFile(filepath.Join(dir, "main.fig"), []byte(mainFig), 0644); err != nil {
		return err
	}

	// README.md
	readme := "# " + baseName + ` — FFI Project

## Quick start

1. Build the shared library:
` + "```bash\ngcc -shared -fPIC -o lib" + baseName + "$(ffi.lib_ext) " + baseName + ".c\n```" + `

2. Set your ffi-helper path in fig.toml:
` + "```toml\n[ffi]\nhelper = \"path/to/ffi-helper\"\n```\n" + `

3. Generate Fig bindings:
` + "```bash\ngo run ./tools/ffi-gen -input " + baseName + ".ffi.def.toml -output bindings.fig\n```" + `

4. Run:
` + "```bash\nfig run main.fig\n```" + `

## Files

| File | Description |
|------|-------------|
| ` + "`fig.toml`" + ` | Project configuration |
| ` + "`" + baseName + ".ffi.def.toml`" + ` | FFI binding definitions (IDL) |
| ` + "`" + baseName + ".c`" + ` | Example C library source |
| ` + "`bindings.fig`" + ` | Auto-generated Fig bindings |
| ` + "`main.fig`" + ` | Example usage |
`
	if err := os.WriteFile(filepath.Join(dir, "README.md"), []byte(readme), 0644); err != nil {
		return err
	}

	return nil
}

// ---------- main ----------

func main() {
	input := flag.String("input", "", "path to ffi.def.toml input file")
	output := flag.String("output", "", "path to write generated .fig file (stdout if omitted)")
	init_ := flag.String("init", "", "scaffold a new FFI project with given name")
	validate := flag.Bool("validate", false, "validate the .ffi.def.toml without generating code")
	checkHelper := flag.String("check-helper", "", "verify the helper binary exists and is executable")
	figToml := flag.String("fig-toml", "", "path to fig.toml to validate FFI configuration")
	flag.Parse()

	// --check-helper: verify helper binary
	if *checkHelper != "" {
		if err := CheckHelper(*checkHelper); err != nil {
			fmt.Fprintf(os.Stderr, "❌ %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "✅ Helper binary OK: %s\n", *checkHelper)
		return
	}

	if *init_ != "" {
		if err := scaffoldProject(*init_); err != nil {
			fmt.Fprintf(os.Stderr, "error scaffolding project: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "✅ Project '%s' created. See %s/README.md for next steps.\n", filepath.Base(*init_), *init_)
		return
	}

	if *input == "" {
		fmt.Fprintln(os.Stderr, "usage: ffi-gen -input <file.ffi.def.toml> [-output <file.fig>]")
		fmt.Fprintln(os.Stderr, "       ffi-gen -init <project-name>")
		fmt.Fprintln(os.Stderr, "       ffi-gen -input <file> -validate")
		fmt.Fprintln(os.Stderr, "       ffi-gen -check-helper <path>")
		os.Exit(1)
	}

	data, err := os.ReadFile(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read %s: %v\n", *input, err)
		os.Exit(1)
	}
	var def DefFile
	if err := toml.Unmarshal(data, &def); err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse %s: %v\n", *input, err)
		os.Exit(1)
	}

	// Validate types
	if validationErrs := ValidateDef(&def); len(validationErrs) > 0 {
		for _, e := range validationErrs {
			fmt.Fprintf(os.Stderr, "❌ %s\n", e)
		}
		os.Exit(1)
	}

	// Validate fig.toml if provided
	if *figToml != "" {
		if tomlErrs := ValidateFigToml(*figToml); len(tomlErrs) > 0 {
			for _, e := range tomlErrs {
				fmt.Fprintf(os.Stderr, "⚠️  %s\n", e)
			}
			// Warnings, don't exit — continue with generation
		}
	}

	// --validate mode: just validate, no output
	if *validate {
		fmt.Fprintf(os.Stderr, "✅ %s is valid\n", *input)
		return
	}

	code := Generate(&def)

	if *output != "" {
		if err := os.WriteFile(*output, []byte(code), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "cannot write %s: %v\n", *output, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "✅ Generated %s\n", *output)
	} else {
		fmt.Print(code)
	}
}
