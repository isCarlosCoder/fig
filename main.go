package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/iscarloscoder/fig/builtins"
	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/interpreter"
	"github.com/pelletier/go-toml/v2"
)

const Version = "0.1.0"

func printHelp(out io.Writer) {
	fmt.Fprintln(out, "FigLang - a small interpreted language")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Usage:")
	fmt.Fprintln(out, "  fig <file>          Run a .fig source file")
	fmt.Fprintln(out, "  fig -i <files...>   Load one or more .fig files into the REPL session and open REPL")
	fmt.Fprintln(out, "  fig run [file]      Run a .fig source file or project main")
	fmt.Fprintln(out, "  fig test [pattern]  Run test files (tests/*.fig, *_test.fig)")
	fmt.Fprintln(out, "  fig init <dir>      Create a new Fig project")
	fmt.Fprintln(out, "  fig install [mods]  Install modules (no args = sync from fig.toml)")
	fmt.Fprintln(out, "  fig remove <mods>   Remove installed modules (--force to skip dependency check)")
	fmt.Fprintln(out, "  fig setup-ffi       Setup FFI (Foreign Function Interface) for modules")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Flags:")
	fmt.Fprintln(out, "  -h, --help          Show this help")
	fmt.Fprintln(out, "  -v, --version       Show version")
}

func main() {
	args := os.Args[1:]
	// -i <files...> : preload files into REPL env then open REPL
	if len(args) > 0 && args[0] == "-i" {
		files := args[1:]
		if len(files) == 0 {
			fmt.Fprintln(os.Stderr, "-i requires at least one .fig file path")
			os.Exit(1)
		}
		env := environment.NewEnv(nil)
		// set ScriptCwd/Args for preload runs (consistent with runFile)
		prevArgs := builtins.ScriptArgs
		prevCwd := builtins.ScriptCwd
		builtins.ScriptArgs = []string{}
		cwd, _ := os.Getwd()
		builtins.ScriptCwd = cwd
		defer func() {
			builtins.ScriptArgs = prevArgs
			builtins.ScriptCwd = prevCwd
		}()

		for _, f := range files {
			abs, err := filepath.Abs(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot resolve %s: %v\n", f, err)
				continue
			}
			data, err := os.ReadFile(abs)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot read %s: %v\n", f, err)
				continue
			}
			if err := interpreter.RunInEnv(string(data), abs, env, os.Stdout, os.Stderr); err != nil {
				fmt.Fprintf(os.Stderr, "error loading %s: %v\n", f, err)
			}
		}

		runRepl(os.Stdin, os.Stdout, os.Stderr, env)
		return
	}

	if len(args) == 0 {
		runRepl(os.Stdin, os.Stdout, os.Stderr, nil)
		return
	}

	// handle global flags
	for i := 0; i < len(args); i++ {
		s := args[i]
		switch s {
		case "-h", "--help":
			printHelp(os.Stdout)
			return
		case "-v", "--version":
			fmt.Println(Version)
			return
		case "run":
			// run subcommand: optional file argument and optional script args after file
			var path string
			var scriptArgs []string
			if i+1 < len(args) {
				path = args[i+1]
				if i+2 < len(args) {
					scriptArgs = args[i+2:]
				}
			} else {
				// no explicit path: all remaining args are script args
				scriptArgs = args[i+1:]
			}
			resolved, err := resolveRunTarget(path, os.Stderr)
			if err != nil {
				os.Exit(1)
			}
			path = resolved
			if err := runFile(path, scriptArgs, os.Stdout, os.Stderr); err != nil {
				// Pretty output already printed to stderr by interpreter; exit non-zero
				os.Exit(1)
			}
			return
		case "init":
			// init subcommand: expect target dir
			if i+1 >= len(args) {
				fmt.Fprintln(os.Stderr, "init requires a directory argument")
				os.Exit(1)
			}
			target := args[i+1]
			if err := initProject(target, os.Stdout, os.Stderr); err != nil {
				os.Exit(1)
			}
			return
		case "install":
			mods := args[i+1:]
			if len(mods) == 0 {
				// No arguments: sync missing dependencies from fig.toml
				if err := installFromToml(os.Stdout, os.Stderr); err != nil {
					os.Exit(1)
				}
			} else {
				hasErr := false
				for _, mod := range mods {
					if strings.Contains(mod, "/") {
						// treat as direct GitHub spec owner/repo
						if err := installModule(mod, "", os.Stdout, os.Stderr); err != nil {
							hasErr = true
						}
					} else {
						// alias lookup via registry
						if err := installAlias(mod, os.Stdout, os.Stderr); err != nil {
							hasErr = true
						}
					}
				}
				if hasErr {
					os.Exit(1)
				}
			}
			return
		case "remove":
			remaining := args[i+1:]
			force := false
			var mods []string
			for _, a := range remaining {
				if a == "--force" || a == "-f" {
					force = true
				} else {
					mods = append(mods, a)
				}
			}
			if len(mods) == 0 {
				fmt.Fprintln(os.Stderr, "remove requires at least one module argument (alias)")
				os.Exit(1)
			}
			hasErr := false
			for _, mod := range mods {
				if strings.Contains(mod, "/") {
					fmt.Fprintf(os.Stderr, "invalid module spec '%s': expected alias name (no '/'),\n", mod)
					hasErr = true
					continue
				}
				if err := removeByAlias(mod, force, os.Stdout, os.Stderr); err != nil {
					hasErr = true
				}
			}
			if hasErr {
				os.Exit(1)
			}
			return
		case "test":
			// test subcommand: discover and run test files
			patterns := args[i+1:]
			verbose := false
			var filtered []string
			for _, p := range patterns {
				if p == "--verbose" || p == "-V" {
					verbose = true
				} else {
					filtered = append(filtered, p)
				}
			}
			if err := runTests(filtered, verbose, os.Stdout, os.Stderr); err != nil {
				os.Exit(1)
			}
			return
		case "setup-ffi":
			// setup-ffi: initialize FFI helper in the project (creates .fig/ffi and updates fig.toml)
			if err := setupFFI(os.Stdout, os.Stderr); err != nil {
				fmt.Fprintf(os.Stderr, "setup-ffi failed: %v\n", err)
				os.Exit(1)
			}
			return
		default:
			// allow running a file directly: fig <file>
			if isFigFile(s) {
				resolved, err := resolveRunTarget(s, os.Stderr)
				if err != nil {
					os.Exit(1)
				}
				scriptArgs := args[i+1:]
				if err := runFile(resolved, scriptArgs, os.Stdout, os.Stderr); err != nil {
					os.Exit(1)
				}
				return
			}

			// unknown top-level token -> show help
			fmt.Fprintf(os.Stderr, "unknown command or flag: %s\n", s)
			printHelp(os.Stderr)
			os.Exit(1)
		}
	}
}

func runFile(path string, scriptArgs []string, out io.Writer, errOut io.Writer) error {
	// read file
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(errOut, "cannot read file %s: %v\n", path, err)
		return err
	}

	// use filename as provided (but normalize to absolute for import resolution)
	filename, _ := filepath.Abs(path)

	// Before executing, set global runtime state for script args and cwd
	prevArgs := builtins.ScriptArgs
	prevCwd := builtins.ScriptCwd
	builtins.ScriptArgs = scriptArgs
	cwd, _ := os.Getwd()
	builtins.ScriptCwd = cwd
	defer func() {
		builtins.ScriptArgs = prevArgs
		builtins.ScriptCwd = prevCwd
	}()

	// execute
	if err := interpreter.Run(string(data), filename, environment.NewEnv(nil), out, errOut); err != nil {
		// If Fig code called system.exit(code), honour it.
		if code, ok := environment.IsExitSignal(err); ok {
			os.Exit(code)
		}
		return err
	}
	return nil
}

func resolveRunTarget(path string, errOut io.Writer) (string, error) {
	if path != "" {
		return path, nil
	}

	projectPath, err := findProjectToml()
	if err != nil {
		fmt.Fprintln(errOut, "run requires a file argument or a fig.toml in the current directory")
		return "", err
	}

	config, err := loadProjectToml(projectPath)
	if err != nil {
		fmt.Fprintf(errOut, "cannot read fig.toml: %v\n", err)
		return "", err
	}

	mainPath := config.Project.Main
	if mainPath == "" {
		mainPath = "src/main.fig"
	}

	baseDir := filepath.Dir(projectPath)
	return filepath.Join(baseDir, mainPath), nil
}

func findProjectToml() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path := filepath.Join(cwd, "fig.toml")
	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	return path, nil
}

func loadProjectToml(path string) (figTomlConfig, error) {
	var cfg figTomlConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func loadModuleToml(path string) (moduleTomlConfig, error) {
	var cfg moduleTomlConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func isFigFile(path string) bool {
	return filepath.Ext(path) == ".fig"
}

// runRepl starts a simple Read-Eval-Print loop using a persistent global environment.
// - input: reads lines from `in` and prints results / errors to `out`/`errOut`.
// - supports a minimal multiline mode when parentheses/braces/brackets are unbalanced.
// - special commands: `.exit` or `exit` to quit.
func runRepl(in io.Reader, out io.Writer, errOut io.Writer, preloadedEnv *environment.Env) {
	reader := bufio.NewReader(in)
	var env *environment.Env
	if preloadedEnv != nil {
		env = preloadedEnv
	} else {
		env = environment.NewEnv(nil)
	}
	// set script context for REPL
	prevArgs := builtins.ScriptArgs
	prevCwd := builtins.ScriptCwd
	builtins.ScriptArgs = []string{}
	cwd, _ := os.Getwd()
	builtins.ScriptCwd = cwd
	defer func() {
		builtins.ScriptArgs = prevArgs
		builtins.ScriptCwd = prevCwd
	}()

	fmt.Fprintln(out, "Fig REPL — type 'exit' or '.exit' to quit")
	for {
		// read lines until balanced or EOF
		var lines []string
		prompt := "fig> "
		for {
			fmt.Fprint(out, prompt)
			line, err := reader.ReadString('\n')
			if err != nil {
				// EOF — exit quietly
				fmt.Fprintln(out, "")
				return
			}
			line = strings.TrimRight(line, "\r\n")
			if line == "" && len(lines) == 0 {
				// empty input — continue
				break
			}
			if line == ".exit" || line == "exit" {
				return
			}
			lines = append(lines, line)
			// check simple balance for delimiters; if balanced, stop reading more
			if delimitersBalanced(strings.Join(lines, "\n")) {
				break
			}
			// continue with a continuation prompt
			prompt = "  > "
		}

		src := strings.Join(lines, "\n")
		if strings.TrimSpace(src) == "" {
			continue
		}

		// First try to evaluate as a single expression and echo the result
		if val, err := interpreter.EvalExpression(src, "<repl>", env, out, errOut); err == nil {
			// Don't echo `null` (print() already performs output and returns nil)
			if val.Type != environment.NilType {
				fmt.Fprintln(out, val.String())
			}
			continue
		} else if err != interpreter.ErrNotExpression {
			// error evaluating expression (parse/runtime) — already printed
			continue
		}
		// Not a bare expression — execute normally so statements work as before
		// Use RunInEnv so top-level declarations persist in the REPL environment.
		if err := interpreter.RunInEnv(src, "<repl>", env, out, errOut); err != nil {
			// show runtime/parse error (interpreter.RunInEnv already prints to errOut)
			continue
		}
	}
}

// delimitersBalanced performs a naive check for balanced (), {}, []
// It ignores strings/comments — sufficient for basic interactive usage.
func delimitersBalanced(s string) bool {
	p := 0
	b := 0
	q := 0
	for _, r := range s {
		switch r {
		case '(':
			p++
		case ')':
			p--
		case '{':
			b++
		case '}':
			b--
		case '[':
			q++
		case ']':
			q--
		}
		if p < 0 || b < 0 || q < 0 {
			return false
		}
	}
	return p == 0 && b == 0 && q == 0
}

func initProject(target string, out io.Writer, errOut io.Writer) error {
	path := target
	if target == "." {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(errOut, "cannot resolve current directory: %v\n", err)
			return err
		}
		path = cwd
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Fprintf(errOut, "cannot resolve path %s: %v\n", path, err)
		return err
	}

	info, statErr := os.Stat(absPath)
	if statErr == nil {
		if !info.IsDir() {
			fmt.Fprintf(errOut, "init target is not a directory: %s\n", absPath)
			return fmt.Errorf("init target is not a directory")
		}
		entries, readErr := os.ReadDir(absPath)
		if readErr != nil {
			fmt.Fprintf(errOut, "cannot read directory %s: %v\n", absPath, readErr)
			return readErr
		}
		if len(entries) != 0 {
			fmt.Fprintf(errOut, "directory not empty: %s\n", absPath)
			return fmt.Errorf("directory not empty")
		}
	} else if !os.IsNotExist(statErr) {
		fmt.Fprintf(errOut, "cannot access %s: %v\n", absPath, statErr)
		return statErr
	}

	if statErr != nil {
		if mkErr := os.MkdirAll(absPath, 0755); mkErr != nil {
			fmt.Fprintf(errOut, "cannot create directory %s: %v\n", absPath, mkErr)
			return mkErr
		}
	}

	projName := filepath.Base(absPath)

	if err := os.MkdirAll(filepath.Join(absPath, "_modules"), 0755); err != nil {
		fmt.Fprintf(errOut, "cannot create _modules: %v\n", err)
		return err
	}
	if err := os.MkdirAll(filepath.Join(absPath, "src"), 0755); err != nil {
		fmt.Fprintf(errOut, "cannot create src: %v\n", err)
		return err
	}

	mainFig := "print(\"Hello, Fig!\")\n"
	if err := os.WriteFile(filepath.Join(absPath, "src", "main.fig"), []byte(mainFig), 0644); err != nil {
		fmt.Fprintf(errOut, "cannot write src/main.fig: %v\n", err)
		return err
	}

	gitignore := "_modules/\n"
	if err := os.WriteFile(filepath.Join(absPath, ".gitignore"), []byte(gitignore), 0644); err != nil {
		fmt.Fprintf(errOut, "cannot write .gitignore: %v\n", err)
		return err
	}

	figTomlData := figTomlConfig{
		Project: figProjectConfig{
			Name:        projName,
			Version:     "0.1.0",
			Description: "",
			Type:        "application",
			Main:        "src/main.fig",
		},
		Authors: figAuthorsConfig{Name: ""},
		Deps:    map[string]figDependencyConfig{},
	}
	figTomlBytes, err := toml.Marshal(figTomlData)
	if err != nil {
		fmt.Fprintf(errOut, "cannot serialize fig.toml: %v\n", err)
		return err
	}
	if err := os.WriteFile(filepath.Join(absPath, "fig.toml"), figTomlBytes, 0644); err != nil {
		fmt.Fprintf(errOut, "cannot write fig.toml: %v\n", err)
		return err
	}

	fmt.Fprintf(out, "Project initialized at %s\n", absPath)
	return nil
}

// installFromToml reads fig.toml and installs any dependencies whose
// directories are missing from _modules/.
func installFromToml(out io.Writer, errOut io.Writer) error {
	projectToml, err := findProjectToml()
	if err != nil {
		fmt.Fprintln(errOut, "fig.toml not found in current directory")
		return err
	}
	projectRoot := filepath.Dir(projectToml)

	cfg, err := loadProjectToml(projectToml)
	if err != nil {
		fmt.Fprintf(errOut, "cannot read fig.toml: %v\n", err)
		return err
	}

	if len(cfg.Deps) == 0 {
		fmt.Fprintln(out, "no dependencies declared in fig.toml")
		return nil
	}

	installed := 0
	hasErr := false
	for _, dep := range cfg.Deps {
		// Extract owner/repo from source (e.g. "github.com/owner/repo")
		parts := strings.Split(dep.Source, "/")
		if len(parts) < 3 {
			continue
		}
		owner := parts[len(parts)-2]
		repo := parts[len(parts)-1]

		moduleDir := filepath.Join(projectRoot, "_modules", repo)
		if _, statErr := os.Stat(moduleDir); statErr == nil {
			// already installed
			continue
		}

		mod := owner + "/" + repo
		if err := installModule(mod, "", out, errOut); err != nil {
			hasErr = true
			continue
		}
		installed++
	}

	if installed == 0 && !hasErr {
		fmt.Fprintln(out, "all dependencies already installed")
	}

	if hasErr {
		return fmt.Errorf("some dependencies failed to install")
	}
	return nil
}

func installModule(mod string, alias string, out io.Writer, errOut io.Writer) error {
	// mod is expected to be owner/repo
	owner, repo, err := parseModuleSpec(mod)
	if err != nil {
		fmt.Fprintf(errOut, "invalid module: %v\n", err)
		return err
	}

	projectToml, err := findProjectToml()
	if err != nil {
		fmt.Fprintln(errOut, "fig.toml not found in current directory")
		return err
	}
	projectRoot := filepath.Dir(projectToml)

	modulesDir := filepath.Join(projectRoot, "_modules")
	if err := os.MkdirAll(modulesDir, 0755); err != nil {
		fmt.Fprintf(errOut, "cannot create _modules: %v\n", err)
		return err
	}

	moduleDir := filepath.Join(modulesDir, repo)
	if _, err := os.Stat(moduleDir); err == nil {
		fmt.Fprintf(errOut, "module already installed: %s\n", repo)
		return fmt.Errorf("module already installed")
	} else if !os.IsNotExist(err) {
		fmt.Fprintf(errOut, "cannot access module directory: %v\n", err)
		return err
	}

	// Allow overriding the GitHub host for tests via GITHUB_BASE env var
	githubBase := os.Getenv("GITHUB_BASE")
	if githubBase == "" {
		githubBase = "https://github.com"
	}
	zipURL := fmt.Sprintf("%s/%s/%s/archive/refs/heads/main.zip", githubBase, owner, repo)
	tmpZip, err := downloadZip(zipURL)
	if err != nil {
		fmt.Fprintf(errOut, "cannot download module: %v\n", err)
		return err
	}
	defer os.Remove(tmpZip)

	if err := extractZip(tmpZip, moduleDir); err != nil {
		fmt.Fprintf(errOut, "cannot extract module: %v\n", err)
		return err
	}

	moduleTomlPath := filepath.Join(moduleDir, "fig.toml")
	modCfg, err := loadModuleToml(moduleTomlPath)
	if err != nil {
		fmt.Fprintf(errOut, "module fig.toml missing or invalid: %v\n", err)
		return err
	}

	// ensure the downloaded project is a library (not an application or other)
	projCfg, err2 := loadProjectToml(moduleTomlPath)
	if err2 != nil {
		fmt.Fprintf(errOut, "cannot parse module fig.toml: %v\n", err2)
		return err2
	}
	if strings.ToLower(projCfg.Project.Type) != "library" {
		fmt.Fprintf(errOut, "module %s/%s has project.type=%q (must be \"library\")\n", owner, repo, projCfg.Project.Type)
		return fmt.Errorf("module is not a library")
	}

	depName := modCfg.Project.Name
	if depName == "" {
		depName = repo
	}
	depVersion := modCfg.Project.Version

	// update parent project fig.toml with new dependency
	projCfg, err = loadProjectToml(projectToml)
	if err != nil {
		fmt.Fprintf(errOut, "cannot read fig.toml: %v\n", err)
		return err
	}
	if projCfg.Deps == nil {
		projCfg.Deps = map[string]figDependencyConfig{}
	}
	projCfg.Deps[depName] = figDependencyConfig{
		Version:  depVersion,
		Source:   fmt.Sprintf("github.com/%s/%s", owner, repo),
		Location: filepath.ToSlash(filepath.Join("_modules", repo)),
		Alias:    alias,
	}

	encoded, err := toml.Marshal(projCfg)
	if err != nil {
		fmt.Fprintf(errOut, "cannot write fig.toml: %v\n", err)
		return err
	}
	if err := os.WriteFile(projectToml, encoded, 0644); err != nil {
		fmt.Fprintf(errOut, "cannot write fig.toml: %v\n", err)
		return err
	}

	fmt.Fprintf(out, "Installed %s/%s (%s)\n", owner, repo, depVersion)
	return nil
}

// resolveAlias looks up an alias in the configured registry and returns owner/repo
func resolveAlias(alias string) (string, string, error) {
	base := os.Getenv("FIGREPO_BASE")
	if base == "" {
		base = "https://figrepo.vercel.app"
	}
	u, err := url.Parse(base)
	if err != nil {
		return "", "", fmt.Errorf("invalid FIGREPO_BASE: %v", err)
	}
	u.Path = "/api/resolve"
	q := u.Query()
	q.Set("alias", alias)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return "", "", fmt.Errorf("alias not found: %s", alias)
	}
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("bad status from registry: %s", resp.Status)
	}
	var body struct {
		Ok      bool `json:"ok"`
		Package struct {
			RepoFullName string `json:"repo_full_name"`
		} `json:"package"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", "", err
	}
	if !body.Ok || body.Package.RepoFullName == "" {
		return "", "", fmt.Errorf("alias not found: %s", alias)
	}
	parts := strings.Split(body.Package.RepoFullName, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid repo name from registry: %s", body.Package.RepoFullName)
	}
	return parts[0], parts[1], nil
}

func installAlias(alias string, out io.Writer, errOut io.Writer) error {
	owner, repo, err := resolveAlias(alias)
	if err != nil {
		fmt.Fprintf(errOut, "cannot resolve alias: %v\n", err)
		return err
	}
	mod := owner + "/" + repo
	// Proceed to install and record alias in fig.toml
	if err := installModule(mod, alias, out, errOut); err != nil {
		return err
	}
	fmt.Fprintf(out, "Installed %s (%s) via alias '%s'\n", mod, alias, alias)
	return nil
}

func removeByAlias(alias string, force bool, out io.Writer, errOut io.Writer) error {
	owner, repo, err := resolveAlias(alias)
	if err != nil {
		fmt.Fprintf(errOut, "cannot resolve alias: %v\n", err)
		return err
	}
	mod := owner + "/" + repo
	return removeModule(mod, force, out, errOut)
}

func removeModule(mod string, force bool, out io.Writer, errOut io.Writer) error {
	owner, repo, err := parseModuleSpec(mod)
	if err != nil {
		fmt.Fprintf(errOut, "invalid module: %v\n", err)
		return err
	}

	projectToml, err := findProjectToml()
	if err != nil {
		fmt.Fprintln(errOut, "fig.toml not found in current directory")
		return err
	}
	projectRoot := filepath.Dir(projectToml)

	moduleDir := filepath.Join(projectRoot, "_modules", repo)
	if _, statErr := os.Stat(moduleDir); os.IsNotExist(statErr) {
		fmt.Fprintf(errOut, "module not installed: %s\n", repo)
		return fmt.Errorf("module not installed")
	}

	// Check if any other installed module depends on the one being removed
	dependents := findDependents(projectRoot, repo)
	if len(dependents) > 0 && !force {
		fmt.Fprintf(errOut, "cannot remove %s/%s: required by other modules:\n", owner, repo)
		for _, d := range dependents {
			fmt.Fprintf(errOut, "  - %s\n", d)
		}
		fmt.Fprintln(errOut, "\nuse --force to remove anyway")
		return fmt.Errorf("module is a dependency of other modules")
	}

	if len(dependents) > 0 && force {
		fmt.Fprintf(out, "warning: %s/%s is required by: %s (forced removal)\n", owner, repo, strings.Join(dependents, ", "))
	}

	// Remove the module directory
	if err := os.RemoveAll(moduleDir); err != nil {
		fmt.Fprintf(errOut, "cannot remove module directory: %v\n", err)
		return err
	}

	// Remove the dependency from fig.toml
	projCfg, err := loadProjectToml(projectToml)
	if err != nil {
		fmt.Fprintf(errOut, "cannot read fig.toml: %v\n", err)
		return err
	}

	found := false
	if projCfg.Deps != nil {
		source := fmt.Sprintf("github.com/%s/%s", owner, repo)
		for name, dep := range projCfg.Deps {
			if dep.Source == source || name == repo {
				delete(projCfg.Deps, name)
				found = true
				break
			}
		}
	}

	if !found {
		fmt.Fprintf(out, "warning: module %s/%s not found in fig.toml dependencies\n", owner, repo)
	}

	encoded, err := toml.Marshal(projCfg)
	if err != nil {
		fmt.Fprintf(errOut, "cannot serialize fig.toml: %v\n", err)
		return err
	}
	if err := os.WriteFile(projectToml, encoded, 0644); err != nil {
		fmt.Fprintf(errOut, "cannot write fig.toml: %v\n", err)
		return err
	}

	fmt.Fprintf(out, "Removed %s/%s\n", owner, repo)
	return nil
}

// findDependents scans all installed modules' fig.toml to find which ones
// depend on the given repo name.
func findDependents(projectRoot, repo string) []string {
	modulesDir := filepath.Join(projectRoot, "_modules")
	entries, err := os.ReadDir(modulesDir)
	if err != nil {
		return nil
	}

	var dependents []string
	for _, entry := range entries {
		if !entry.IsDir() || entry.Name() == repo {
			continue
		}
		modToml := filepath.Join(modulesDir, entry.Name(), "fig.toml")
		data, err := os.ReadFile(modToml)
		if err != nil {
			continue
		}
		// Parse as project config to check dependencies
		var cfg figTomlConfig
		if err := toml.Unmarshal(data, &cfg); err != nil {
			continue
		}
		for name, dep := range cfg.Deps {
			// Match by dependency name or by location containing the repo
			if name == repo || strings.HasSuffix(dep.Location, "/"+repo) || strings.Contains(dep.Source, "/"+repo) {
				dependents = append(dependents, entry.Name())
				break
			}
		}
	}
	return dependents
}

func parseModuleSpec(spec string) (string, string, error) {
	parts := strings.Split(spec, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("expected <owner>/<repo>")
	}
	return parts[0], parts[1], nil
}

func downloadZip(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}
	tmp, err := os.CreateTemp("", "fig-mod-*.zip")
	if err != nil {
		return "", err
	}
	defer tmp.Close()
	if _, err := io.Copy(tmp, resp.Body); err != nil {
		return "", err
	}
	return tmp.Name(), nil
}

func extractZip(zipPath, destDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	for _, f := range r.File {
		parts := strings.SplitN(f.Name, "/", 2)
		if len(parts) < 2 {
			continue
		}
		rel := parts[1]
		if rel == "" {
			continue
		}
		target := filepath.Join(destDir, rel)
		cleanTarget := filepath.Clean(target)
		if !strings.HasPrefix(cleanTarget, filepath.Clean(destDir)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid zip path: %s", f.Name)
		}
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(cleanTarget, 0755); err != nil {
				return err
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(cleanTarget), 0755); err != nil {
			return err
		}
		in, err := f.Open()
		if err != nil {
			return err
		}
		out, err := os.OpenFile(cleanTarget, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			in.Close()
			return err
		}
		if _, err := io.Copy(out, in); err != nil {
			out.Close()
			in.Close()
			return err
		}
		out.Close()
		in.Close()
	}

	return nil
}

// runTests discovers and executes Fig test files, printing results and a summary.
func runTests(patterns []string, verbose bool, out io.Writer, errOut io.Writer) error {
	projectToml, findErr := findProjectToml()
	var projectRoot string
	if findErr == nil {
		projectRoot = filepath.Dir(projectToml)
	} else {
		// no fig.toml — use cwd
		cwd, _ := os.Getwd()
		projectRoot = cwd
	}

	files, err := discoverTestFiles(projectRoot, patterns)
	if err != nil {
		fmt.Fprintf(errOut, "error discovering test files: %v\n", err)
		return err
	}

	if len(files) == 0 {
		fmt.Fprintln(out, "no test files found")
		return nil
	}

	totalPassed := 0
	totalFailed := 0
	totalSkipped := 0
	anyFailed := false

	for _, file := range files {
		rel, _ := filepath.Rel(projectRoot, file)
		if rel == "" {
			rel = file
		}

		// Reset figtest state per file
		builtins.ResetFigtest()

		data, readErr := os.ReadFile(file)
		if readErr != nil {
			fmt.Fprintf(errOut, "cannot read %s: %v\n", file, readErr)
			anyFailed = true
			continue
		}

		absPath, _ := filepath.Abs(file)
		env := environment.NewEnv(nil)

		// Execute the test file in isolation
		runErr := interpreter.Run(string(data), absPath, env, io.Discard, errOut)

		state := builtins.GetFigtestState()
		passed := state.Passed()
		failed := state.Failed()
		skipped := state.Skipped()

		totalPassed += passed
		totalFailed += failed
		totalSkipped += skipped

		if failed > 0 || runErr != nil {
			anyFailed = true
		}

		// Print per-file results
		fmt.Fprintf(out, "\n%s\n", rel)
		for _, line := range state.Output() {
			fmt.Fprintln(out, line)
		}

		if runErr != nil && failed == 0 {
			// runtime error outside of a test
			fmt.Fprintf(out, "  ✗ runtime error: %v\n", runErr)
		}
	}

	// Final summary
	fmt.Fprintln(out)
	total := totalPassed + totalFailed + totalSkipped
	if totalFailed > 0 {
		fmt.Fprintf(out, "\x1b[1;31m%d passed, %d failed", totalPassed, totalFailed)
	} else {
		fmt.Fprintf(out, "\x1b[1;32m%d passed, %d failed", totalPassed, totalFailed)
	}
	if totalSkipped > 0 {
		fmt.Fprintf(out, ", %d skipped", totalSkipped)
	}
	fmt.Fprintf(out, " (total: %d)\x1b[0m\n", total)

	if verbose {
		fmt.Fprintln(out)
		fmt.Fprintf(out, "Ran %d file(s)\n", len(files))
	}

	if anyFailed {
		return fmt.Errorf("tests failed")
	}
	return nil
}

// discoverTestFiles finds .fig test files from patterns or default locations.
func discoverTestFiles(projectRoot string, patterns []string) ([]string, error) {
	if len(patterns) > 0 {
		var files []string
		for _, pat := range patterns {
			// If pattern is absolute, use as-is; otherwise relative to projectRoot
			if !filepath.IsAbs(pat) {
				pat = filepath.Join(projectRoot, pat)
			}
			matches, err := filepath.Glob(pat)
			if err != nil {
				return nil, fmt.Errorf("invalid pattern %q: %v", pat, err)
			}
			files = append(files, matches...)
		}
		return files, nil
	}

	// Default discovery: tests/*.fig + **/*_test.fig
	var files []string

	// 1) tests/ directory
	testsDir := filepath.Join(projectRoot, "tests")
	if info, err := os.Stat(testsDir); err == nil && info.IsDir() {
		matches, _ := filepath.Glob(filepath.Join(testsDir, "*.fig"))
		files = append(files, matches...)
	}

	// 2) *_test.fig recursively
	filepath.Walk(projectRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			base := info.Name()
			if base == "_modules" || base == ".git" || base == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(info.Name(), "_test.fig") {
			// avoid duplicates from tests/ dir
			for _, f := range files {
				if f == path {
					return nil
				}
			}
			files = append(files, path)
		}
		return nil
	})

	return files, nil
}

type figTomlConfig struct {
	Project figProjectConfig               `toml:"project"`
	Authors figAuthorsConfig               `toml:"authors"`
	Deps    map[string]figDependencyConfig `toml:"dependencies"`
	Ffi     figFfiConfig                   `toml:"ffi,omitempty"`
}

type figProjectConfig struct {
	Name        string `toml:"name"`
	Version     string `toml:"version"`
	Description string `toml:"description"`
	Type        string `toml:"type"`
	Main        string `toml:"main"`
}

type figAuthorsConfig struct {
	Name string `toml:"name"`
}

type figDependencyConfig struct {
	Version  string `toml:"version"`
	Source   string `toml:"source"`
	Location string `toml:"location"`
	Alias    string `toml:"alias,omitempty"`
}

type moduleTomlConfig struct {
	Project moduleProjectConfig `toml:"project"`
}

type moduleProjectConfig struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
	Main    string `toml:"main"`
}
type figFfiConfig struct {
	Enabled   bool     `toml:"enabled,omitempty"`
	Helper    string   `toml:"helper,omitempty"`
	Whitelist []string `toml:"whitelist,omitempty"`
}

func setupFFI(out io.Writer, errOut io.Writer) error {
	projectToml, err := findProjectToml()
	if err != nil {
		fmt.Fprintln(errOut, "fig.toml not found in current directory")
		return err
	}
	projectRoot := filepath.Dir(projectToml)

	cfg, err := loadProjectToml(projectToml)
	if err != nil {
		fmt.Fprintf(errOut, "cannot read fig.toml: %v\n", err)
		return err
	}

	// set defaults
	if cfg.Deps == nil {
		cfg.Deps = map[string]figDependencyConfig{}
	}

	// default helper path (local to project) if empty
	defaultHelper := filepath.Join(projectRoot, ".fig", "ffi", "ffi-helper")
	if cfg.Ffi.Helper == "" {
		cfg.Ffi.Helper = defaultHelper
	}

	// If a global `ffi-helper` is available in PATH, prefer and record its absolute path
	if lp, err := exec.LookPath("ffi-helper"); err == nil {
		if abs, aerr := filepath.Abs(lp); aerr == nil {
			cfg.Ffi.Helper = abs
		}
		fmt.Fprintf(out, "Found global 'ffi-helper' at %s — configuring fig.toml to use it\n", cfg.Ffi.Helper)
	} else {
		fmt.Fprintln(out, "FFI helper not found in your PATH.")
		fmt.Fprintln(out, "To compile and install the helper globally, run:")
		fmt.Fprintln(out, "  cd tools/ffi-helper && go build -o ffi-helper .")
		fmt.Fprintln(out, "  sudo mv ffi-helper /usr/local/bin/   # or mv to $HOME/.local/bin and add to PATH")
		fmt.Fprintln(out)
		fmt.Fprintf(out, "Or build the helper into this project and keep it local: %s\n", defaultHelper)
		fmt.Fprintln(out)
		fmt.Fprintln(out, "See: https://github.com/isCarlosCoder/fig/docs/ffi/setup.md")
	}

	cfg.Ffi.Enabled = true

	b, err := toml.Marshal(cfg)
	if err != nil {
		fmt.Fprintf(errOut, "cannot serialize fig.toml: %v\n", err)
		return err
	}
	if err := os.WriteFile(projectToml, b, 0644); err != nil {
		fmt.Fprintf(errOut, "cannot write fig.toml: %v\n", err)
		return err
	}

	// create .fig/ffi dir
	ffiDir := filepath.Join(projectRoot, ".fig", "ffi")
	if err := os.MkdirAll(ffiDir, 0755); err != nil {
		fmt.Fprintf(errOut, "cannot create %s: %v\n", ffiDir, err)
		return err
	}

	fmt.Fprintf(out, "FFI enabled; helper path: %s\n", cfg.Ffi.Helper)
	return nil
}
