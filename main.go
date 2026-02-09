package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/interpreter"
	"github.com/pelletier/go-toml/v2"
)

const Version = "0.1.0"

func printHelp(out io.Writer) {
	fmt.Fprintln(out, "FigLang - a small interpreted language")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Usage:")
	fmt.Fprintln(out, "  fig <file>       Run a .fig source file")
	fmt.Fprintln(out, "  fig run [file]   Run a .fig source file or project main")
	fmt.Fprintln(out, "  fig init <dir>   Create a new Fig project")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Flags:")
	fmt.Fprintln(out, "  -h, --help       Show this help")
	fmt.Fprintln(out, "  -v, --version    Show version")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp(os.Stdout)
		os.Exit(0)
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
			// run subcommand: optional file argument
			var path string
			if i+1 < len(args) {
				path = args[i+1]
			}
			resolved, err := resolveRunTarget(path, os.Stderr)
			if err != nil {
				os.Exit(1)
			}
			path = resolved
			if err := runFile(path, os.Stdout, os.Stderr); err != nil {
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
		default:
			// allow running a file directly: fig <file>
			if isFigFile(s) {
				resolved, err := resolveRunTarget(s, os.Stderr)
				if err != nil {
					os.Exit(1)
				}
				if err := runFile(resolved, os.Stdout, os.Stderr); err != nil {
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

func runFile(path string, out io.Writer, errOut io.Writer) error {
	// read file
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(errOut, "cannot read file %s: %v\n", path, err)
		return err
	}

	// use filename as provided (but normalize to absolute for import resolution)
	filename, _ := filepath.Abs(path)

	// execute
	if err := interpreter.Run(string(data), filename, environment.NewEnv(nil), out, errOut); err != nil {
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

func isFigFile(path string) bool {
	return filepath.Ext(path) == ".fig"
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

type figTomlConfig struct {
	Project figProjectConfig               `toml:"project"`
	Authors figAuthorsConfig               `toml:"authors"`
	Deps    map[string]figDependencyConfig `toml:"dependencies"`
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
}
