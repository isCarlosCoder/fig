package builtins

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("io",
		// input(prompt) — prints prompt and reads a line from stdin
		fn("input", func(args []environment.Value) (environment.Value, error) {
			if len(args) > 1 {
				return environment.NewNil(), fmt.Errorf("input() expects 0 or 1 argument, got %d", len(args))
			}
			if len(args) == 1 {
				if args[0].Type != environment.StringType {
					return environment.NewNil(), fmt.Errorf("input() prompt must be a string")
				}
				fmt.Print(args[0].Str)
			}
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				return environment.NewString(scanner.Text()), nil
			}
			return environment.NewString(""), nil
		}),

		// readLine() — reads a line from stdin (no prompt)
		fn("readLine", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("readLine() expects 0 arguments, got %d", len(args))
			}
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				return environment.NewString(scanner.Text()), nil
			}
			return environment.NewString(""), nil
		}),

		// readFile(path) — reads the entire file and returns its content
		fn("readFile", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("readFile() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("readFile() path must be a string")
			}
			path := args[0].Str
			data, err := os.ReadFile(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("readFile(): %v", err)
			}
			return environment.NewString(string(data)), nil
		}),

		// writeFile(path, data) — writes data to a file (creates or overwrites)
		fn("writeFile", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("writeFile() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("writeFile() path must be a string")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("writeFile() data must be a string")
			}
			path := args[0].Str
			content := args[1].Str
			err := os.WriteFile(path, []byte(content), 0644)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("writeFile(): %v", err)
			}
			return environment.NewNil(), nil
		}),

		// appendFile(path, data) — appends data to a file
		fn("appendFile", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("appendFile() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("appendFile() path must be a string")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("appendFile() data must be a string")
			}
			path := args[0].Str
			content := args[1].Str
			f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("appendFile(): %v", err)
			}
			defer f.Close()
			_, err = f.WriteString(content)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("appendFile(): %v", err)
			}
			return environment.NewNil(), nil
		}),

		// exists(path) — checks if a file or directory exists
		fn("exists", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("exists() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("exists() path must be a string")
			}
			path := args[0].Str
			_, err := os.Stat(path)
			return environment.NewBool(!os.IsNotExist(err)), nil
		}),

		// deleteFile(path) — deletes a file
		fn("deleteFile", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("deleteFile() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("deleteFile() path must be a string")
			}
			path := args[0].Str
			err := os.Remove(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("deleteFile(): %v", err)
			}
			return environment.NewNil(), nil
		}),
	))
}
