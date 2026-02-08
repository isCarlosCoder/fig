package builtins

import "github.com/iscarloscoder/fig/environment"

// Module represents a built-in module with its functions and constants.
type Module struct {
	Name    string
	Entries map[string]environment.Value
	Keys    []string // insertion order
}

// registry holds all built-in modules keyed by name.
var registry = map[string]*Module{}

// register adds a module to the global registry.
func register(m *Module) {
	registry[m.Name] = m
}

// Get returns the built-in module with the given name, or nil if not found.
func Get(name string) *Module {
	return registry[name]
}

// ToObject converts a Module to an environment.Value of ObjectType,
// so it can be assigned to a variable and accessed via dot notation.
func (m *Module) ToObject() environment.Value {
	entries := make(map[string]environment.Value, len(m.Entries))
	keys := make([]string, len(m.Keys))
	for k, v := range m.Entries {
		entries[k] = v
	}
	copy(keys, m.Keys)
	return environment.NewObject(entries, keys)
}

// helper to build a module from a list of name/value pairs (preserves order).
func newModule(name string, pairs ...struct {
	k string
	v environment.Value
}) *Module {
	m := &Module{
		Name:    name,
		Entries: make(map[string]environment.Value, len(pairs)),
		Keys:    make([]string, 0, len(pairs)),
	}
	for _, p := range pairs {
		m.Entries[p.k] = p.v
		m.Keys = append(m.Keys, p.k)
	}
	return m
}

// entry is a shortcut to create a name/value pair for newModule.
type entry = struct {
	k string
	v environment.Value
}

func fn(name string, f environment.BuiltinFn) entry {
	return entry{k: name, v: environment.NewBuiltinFn(name, f)}
}

func num(name string, n float64) entry {
	return entry{k: name, v: environment.NewNumber(n)}
}

func str(name string, s string) entry {
	return entry{k: name, v: environment.NewString(s)}
}
