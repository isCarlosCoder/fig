package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/interpreter"
)

const Version = "0.1.0"

func printHelp(out io.Writer) {
	fmt.Fprintln(out, "FigLang - a small interpreted language")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Usage:")
	fmt.Fprintln(out, "  fig run <file>   Run a .fig source file")
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
			// run subcommand: expect next arg as file
			if i+1 >= len(args) {
				fmt.Fprintln(os.Stderr, "run requires a file argument")
				os.Exit(1)
			}
			path := args[i+1]
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

	figToml := fmt.Sprintf(`[project]
name = "%s"
version = "0.1.0"
description = ""
type = "application"
main = "src/main.fig"

[authors]
name = ""

[dependencies]
`, projName)
	if err := os.WriteFile(filepath.Join(absPath, "fig.toml"), []byte(figToml), 0644); err != nil {
		fmt.Fprintf(errOut, "cannot write fig.toml: %v\n", err)
		return err
	}

	fmt.Fprintf(out, "Project initialized at %s\n", absPath)
	return nil
}
