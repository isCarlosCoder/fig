package builtins

import (
	"bufio"
	"encoding/csv"
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
			err := os.WriteFile(path, []byte(content), 0o644)
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
			f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
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

		// isDir(path) — returns true if path exists and is a directory
		fn("isDir", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isDir() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("isDir() path must be a string")
			}
			path := args[0].Str
			st, err := os.Stat(path)
			if err != nil {
				if os.IsNotExist(err) {
					return environment.NewBool(false), nil
				}
				return environment.NewNil(), fmt.Errorf("isDir(): %v", err)
			}
			return environment.NewBool(st.IsDir()), nil
		}),

		// mkdir(path) — create a directory (error if exists)
		fn("mkdir", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("mkdir() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("mkdir() path must be a string")
			}
			path := args[0].Str
			err := os.Mkdir(path, 0o755)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("mkdir(): %v", err)
			}
			return environment.NewNil(), nil
		}),

		// mkdirAll(path) — create directory and parents (like mkdir -p)
		fn("mkdirAll", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("mkdirAll() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("mkdirAll() path must be a string")
			}
			path := args[0].Str
			err := os.MkdirAll(path, 0o755)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("mkdirAll(): %v", err)
			}
			return environment.NewNil(), nil
		}),

		// readDir(path) — returns array of entries (names) in a directory
		fn("readDir", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("readDir() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("readDir() path must be a string")
			}
			path := args[0].Str
			entries, err := os.ReadDir(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("readDir(): %v", err)
			}
			arr := make([]environment.Value, 0, len(entries))
			for _, e := range entries {
				arr = append(arr, environment.NewString(e.Name()))
			}
			return environment.NewArray(arr), nil
		}),

		// rmdir(path) — remove empty directory
		fn("rmdir", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("rmdir() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("rmdir() path must be a string")
			}
			path := args[0].Str
			err := os.Remove(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("rmdir(): %v", err)
			}
			return environment.NewNil(), nil
		}),

		// rmdirAll(path) — remove directory and its contents recursively
		fn("rmdirAll", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("rmdirAll() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("rmdirAll() path must be a string")
			}
			path := args[0].Str
			err := os.RemoveAll(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("rmdirAll(): %v", err)
			}
			return environment.NewNil(), nil
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

		fn("readCSV", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("readCSV() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("readCSV() path must be a string")
			}
			path := args[0].Str
			f, err := os.Open(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("readCSV(): %v", err)
			}
			defer f.Close()

			r := csv.NewReader(f)
			records, err := r.ReadAll()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("readCSV(): %v", err)
			}

			// we will return a array<object> format where each row is an object with keys as column names (from header) and values as cell values
			if len(records) == 0 {
				return environment.NewArray([]environment.Value{}), nil
			}

			header := records[0]
			rows := make([]environment.Value, 0, len(records)-1)
			for _, record := range records[1:] {
				obj := make(map[string]environment.Value)
				for i, cell := range record {
					if i < len(header) {
						obj[header[i]] = environment.NewString(cell)
					}
				}

				rows = append(rows, environment.NewObject(obj, header))
			}

			return environment.NewArray(rows), nil
		}),

		fn("writeCSV", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("writeCSV() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("writeCSV() path must be a string")
			}
			if args[1].Type != environment.ArrayType {
				return environment.NewNil(), fmt.Errorf("writeCSV() data must be an array of objects")
			}
			path := args[0].Str
			arr := args[1].Arr

			f, err := os.Create(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("writeCSV(): %v", err)
			}
			defer f.Close()

			w := csv.NewWriter(f)
			defer w.Flush()

			// we expect an array of objects where keys are column names and values are cell values
			if len(*arr) == 0 {
				return environment.NewNil(), nil
			}

			// get header from keys of the first object
			firstObj := (*arr)[0]
			if firstObj.Type != environment.ObjectType {
				return environment.NewNil(), fmt.Errorf("writeCSV() data must be an array of objects")
			}
			header := make([]string, 0, len(firstObj.Obj.Entries))
			for k := range firstObj.Obj.Entries {
				header = append(header, k)
			}

			err = w.Write(header)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("writeCSV(): %v", err)
			}

			for _, v := range *arr {
				if v.Type != environment.ObjectType {
					return environment.NewNil(), fmt.Errorf("writeCSV() data must be an array of objects")
				}
				row := make([]string, len(header))
				for i, col := range header {
					if cellVal, ok := v.Obj.Entries[col]; ok {
						row[i] = cellVal.String()
					} else {
						row[i] = ""
					}
				}
				err = w.Write(row)
				if err != nil {
					return environment.NewNil(), fmt.Errorf("writeCSV(): %v", err)
				}
			}

			return environment.NewNil(), nil
		}),
	))
}
