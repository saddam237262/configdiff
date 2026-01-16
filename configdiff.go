// Package configdiff provides semantic, human-grade diffs for YAML/JSON/HCL configuration files.
//
// It parses configuration files into a normalized tree representation, applies customizable
// diff rules, and generates both machine-readable patches and human-friendly reports.
package configdiff

import (
	"github.com/pfrederiksen/configdiff/tree"
)

// Options configures how diffs are computed.
type Options struct {
	// IgnorePaths specifies paths to ignore in the diff (JSONPath-like syntax with wildcards).
	// Example: []string{"metadata.creationTimestamp", "status.*"}
	IgnorePaths []string

	// ArraySetKeys maps array paths to their key field names for set-based comparison.
	// When specified, arrays are treated as sets keyed by the given field.
	// Example: map[string]string{"spec.containers": "name"}
	ArraySetKeys map[string]string

	// Coercions configures type coercion rules for comparison.
	Coercions Coercions

	// StableOrder ensures deterministic, stable ordering in output.
	StableOrder bool
}

// Coercions defines rules for type coercion during comparison.
type Coercions struct {
	// NumericStrings allows comparing string numbers with numeric values.
	// Example: "1" can equal 1
	NumericStrings bool

	// BoolStrings allows comparing string booleans with boolean values.
	// Example: "true" can equal true
	BoolStrings bool
}

// Result contains the output of a diff operation.
type Result struct {
	// Changes is the list of detected changes.
	Changes []Change

	// Patch is the machine-readable patch representation.
	Patch Patch

	// Report is the human-friendly pretty report.
	Report string
}

// Change represents a single detected change.
type Change struct {
	// Type is the kind of change (add, remove, modify, move).
	Type ChangeType

	// Path is the location of the change in the tree.
	Path string

	// OldValue is the previous value (nil for additions).
	OldValue *tree.Node

	// NewValue is the new value (nil for removals).
	NewValue *tree.Node
}

// ChangeType categorizes the kind of change.
type ChangeType string

const (
	// ChangeTypeAdd indicates a new value was added.
	ChangeTypeAdd ChangeType = "add"

	// ChangeTypeRemove indicates a value was removed.
	ChangeTypeRemove ChangeType = "remove"

	// ChangeTypeModify indicates a value was changed.
	ChangeTypeModify ChangeType = "modify"

	// ChangeTypeMove indicates a value was moved (array reordering).
	ChangeTypeMove ChangeType = "move"
)

// Patch represents a machine-readable set of operations.
type Patch struct {
	// Operations is the list of patch operations.
	Operations []Operation
}

// Operation is a single patch operation (JSON Patch-like).
type Operation struct {
	// Op is the operation type (add, remove, replace, move).
	Op string `json:"op"`

	// Path is the target path for the operation.
	Path string `json:"path"`

	// Value is the value for add/replace operations.
	Value interface{} `json:"value,omitempty"`

	// From is the source path for move operations.
	From string `json:"from,omitempty"`
}

// DiffBytes compares two configuration byte slices and returns the diff result.
//
// Supported formats: "yaml", "json", "hcl"
func DiffBytes(a []byte, aFormat string, b []byte, bFormat string, opts Options) (*Result, error) {
	// TODO: implement
	return nil, nil
}

// DiffTrees compares two normalized tree nodes and returns the diff result.
func DiffTrees(a, b *tree.Node, opts Options) (*Result, error) {
	// TODO: implement
	return nil, nil
}

// DiffYAML is a convenience function for comparing two YAML byte slices.
func DiffYAML(a, b []byte, opts Options) (*Result, error) {
	return DiffBytes(a, "yaml", b, "yaml", opts)
}

// DiffJSON is a convenience function for comparing two JSON byte slices.
func DiffJSON(a, b []byte, opts Options) (*Result, error) {
	return DiffBytes(a, "json", b, "json", opts)
}
