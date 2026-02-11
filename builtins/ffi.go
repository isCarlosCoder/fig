package builtins

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iscarloscoder/fig/environment"
	"github.com/pelletier/go-toml/v2"
)

func findProjectTomlFrom(startDir string) (string, error) {
	dir := startDir
	for {
		candidate := filepath.Join(dir, "fig.toml")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("fig.toml not found")
		}
		dir = parent
	}
}

func readFfiEnabled() (bool, string, error) {
	// find project toml relative to cwd
	cwd, _ := os.Getwd()
	p, err := findProjectTomlFrom(cwd)
	if err != nil {
		return false, "", nil
	}
	var cfg struct {
		Ffi struct {
			Enabled bool   `toml:"enabled"`
			Helper  string `toml:"helper"`
		} `toml:"ffi"`
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return false, "", err
	}
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return false, "", err
	}
	return cfg.Ffi.Enabled, cfg.Ffi.Helper, nil
}

func init() {
	register(newModule("ffi",
		fn("enabled", func(args []environment.Value) (environment.Value, error) {
			en, _, err := readFfiEnabled()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			return environment.NewBool(en), nil
		}),

		fn("load", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("load(path) expects 1 argument")
			}
			path, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("load(path) expects a string")
			}
			en, helper, err := readFfiEnabled()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			if !en {
				return environment.NewNil(), fmt.Errorf("FFI is not enabled for this project; run 'fig setup-ffi' and enable in fig.toml")
			}
			// For v1 we just check helper exists; real loading will be implemented later
			if helper == "" {
				return environment.NewNil(), fmt.Errorf("no ffi.helper configured in fig.toml; run 'fig setup-ffi'")
			}
			if _, statErr := os.Stat(helper); statErr != nil {
				return environment.NewNil(), fmt.Errorf("ffi helper not found at '%s' (run 'fig setup-ffi')", helper)
			}
			return environment.NewNil(), fmt.Errorf("ffi.load not implemented yet; helper present: %s (requested %s)", helper, path)
		}),
	))
}
