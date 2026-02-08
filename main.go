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
