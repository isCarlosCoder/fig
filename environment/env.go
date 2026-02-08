package environment

import (
	"fmt"
	"sync"
)

type Env struct {
	mu     sync.RWMutex
	vars   map[string]Value
	parent *Env
}

func NewEnv(parent *Env) *Env {
	return &Env{
		vars:   make(map[string]Value),
		parent: parent,
	}
}

func (e *Env) Get(name string) (Value, bool) {
	e.mu.RLock()
	val, ok := e.vars[name]
	e.mu.RUnlock()
	if !ok && e.parent != nil {
		return e.parent.Get(name)
	}
	return val, ok
}

func (e *Env) Has(name string) bool {
	e.mu.RLock()
	_, ok := e.vars[name]
	e.mu.RUnlock()
	if !ok && e.parent != nil {
		return e.parent.Has(name)
	}
	return ok
}

// Define declares a new variable in the current scope.
// Returns an error if the variable is already defined in the same scope.
func (e *Env) Define(name string, val Value) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, exists := e.vars[name]; exists {
		return fmt.Errorf("variable '%s' already defined in this scope", name)
	}
	e.vars[name] = val
	return nil
}

// Assign updates an existing variable in this scope chain.
// Returns an error if the variable is not found.
func (e *Env) Assign(name string, val Value) error {
	e.mu.Lock()
	_, ok := e.vars[name]
	if ok {
		e.vars[name] = val
		e.mu.Unlock()
		return nil
	}
	e.mu.Unlock()
	if e.parent != nil {
		return e.parent.Assign(name, val)
	}
	return fmt.Errorf("variable '%s' not defined", name)
}

// Set performs an unguarded set on the current scope (for internal use).
func (e *Env) Set(name string, val Value) {
	e.mu.Lock()
	e.vars[name] = val
	e.mu.Unlock()
}

func (e *Env) Update(name string, val Value) bool {
	e.mu.Lock()
	_, ok := e.vars[name]
	if ok {
		e.vars[name] = val
		e.mu.Unlock()
		return true
	}
	e.mu.Unlock()
	if e.parent != nil {
		return e.parent.Update(name, val)
	}
	return false
}

// Parent returns the parent environment (can be nil).
func (e *Env) Parent() *Env { return e.parent }
