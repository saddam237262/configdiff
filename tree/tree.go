// Package tree provides a normalized tree representation for configuration data.
//
// All configuration formats (YAML, JSON, HCL) are parsed into this common tree structure,
// enabling format-agnostic diff operations.
package tree

import (
	"fmt"
	"sort"
	"strings"
)

// NodeKind represents the type of a tree node.
type NodeKind int

const (
	// KindNull represents a null/nil value.
	KindNull NodeKind = iota

	// KindBool represents a boolean value.
	KindBool

	// KindNumber represents a numeric value.
	KindNumber

	// KindString represents a string value.
	KindString

	// KindObject represents a key-value mapping.
	KindObject

	// KindArray represents an ordered list.
	KindArray
)

// String returns the string representation of a NodeKind.
func (k NodeKind) String() string {
	switch k {
	case KindNull:
		return "null"
	case KindBool:
		return "bool"
	case KindNumber:
		return "number"
	case KindString:
		return "string"
	case KindObject:
		return "object"
	case KindArray:
		return "array"
	default:
		return "unknown"
	}
}

// Node represents a single node in the normalized configuration tree.
type Node struct {
	// Kind is the type of this node.
	Kind NodeKind

	// Value holds the scalar value for null, bool, number, or string nodes.
	Value interface{}

	// Object holds key-value pairs for object nodes.
	Object map[string]*Node

	// Array holds elements for array nodes.
	Array []*Node

	// Path is the canonical path to this node from the root.
	// Example: "/spec/template/spec/containers[0]/image"
	Path string
}

// NewNull creates a null node.
func NewNull() *Node {
	return &Node{Kind: KindNull, Value: nil}
}

// NewBool creates a boolean node.
func NewBool(v bool) *Node {
	return &Node{Kind: KindBool, Value: v}
}

// NewNumber creates a numeric node.
func NewNumber(v float64) *Node {
	return &Node{Kind: KindNumber, Value: v}
}

// NewString creates a string node.
func NewString(v string) *Node {
	return &Node{Kind: KindString, Value: v}
}

// NewObject creates an object node with the given key-value pairs.
func NewObject(kvs map[string]*Node) *Node {
	return &Node{Kind: KindObject, Object: kvs}
}

// NewArray creates an array node with the given elements.
func NewArray(elements []*Node) *Node {
	return &Node{Kind: KindArray, Array: elements}
}

// Clone creates a deep copy of the node.
func (n *Node) Clone() *Node {
	if n == nil {
		return nil
	}

	cloned := &Node{
		Kind:  n.Kind,
		Value: n.Value,
		Path:  n.Path,
	}

	if n.Object != nil {
		cloned.Object = make(map[string]*Node, len(n.Object))
		for k, v := range n.Object {
			cloned.Object[k] = v.Clone()
		}
	}

	if n.Array != nil {
		cloned.Array = make([]*Node, len(n.Array))
		for i, elem := range n.Array {
			cloned.Array[i] = elem.Clone()
		}
	}

	return cloned
}

// Equal checks if two nodes are equal.
func (n *Node) Equal(other *Node) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if n.Kind != other.Kind {
		return false
	}

	switch n.Kind {
	case KindNull:
		return true
	case KindBool, KindNumber, KindString:
		return n.Value == other.Value
	case KindObject:
		if len(n.Object) != len(other.Object) {
			return false
		}
		for k, v := range n.Object {
			otherV, exists := other.Object[k]
			if !exists || !v.Equal(otherV) {
				return false
			}
		}
		return true
	case KindArray:
		if len(n.Array) != len(other.Array) {
			return false
		}
		for i := range n.Array {
			if !n.Array[i].Equal(other.Array[i]) {
				return false
			}
		}
		return true
	}

	return false
}

// SortedKeys returns the sorted keys of an object node.
// Returns nil for non-object nodes.
func (n *Node) SortedKeys() []string {
	if n.Kind != KindObject {
		return nil
	}

	keys := make([]string, 0, len(n.Object))
	for k := range n.Object {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// SetPaths recursively sets the canonical path for all nodes in the tree.
func (n *Node) SetPaths(basePath string) {
	if n == nil {
		return
	}

	n.Path = basePath

	switch n.Kind {
	case KindObject:
		for k, v := range n.Object {
			childPath := joinPath(basePath, k)
			v.SetPaths(childPath)
		}
	case KindArray:
		for i, elem := range n.Array {
			childPath := fmt.Sprintf("%s[%d]", basePath, i)
			elem.SetPaths(childPath)
		}
	}
}

// joinPath joins path segments with proper formatting.
func joinPath(base, key string) string {
	if base == "" || base == "/" {
		return "/" + key
	}
	return base + "/" + key
}

// ParsePath parses a canonical path into segments.
// Example: "/spec/containers[0]/name" -> ["spec", "containers[0]", "name"]
func ParsePath(path string) []string {
	if path == "" || path == "/" {
		return nil
	}

	trimmed := strings.TrimPrefix(path, "/")
	if trimmed == "" {
		return nil
	}

	return strings.Split(trimmed, "/")
}

// GetByPath retrieves a node at the given path.
// Returns nil if the path doesn't exist.
func (n *Node) GetByPath(path string) *Node {
	if n == nil {
		return nil
	}

	segments := ParsePath(path)
	if len(segments) == 0 {
		return n
	}

	current := n
	for _, segment := range segments {
		if current == nil {
			return nil
		}

		// Check for array index notation like "containers[0]"
		if baseName, idx, isArray := parseArrayNotation(segment); isArray {
			// First navigate to the object key if there's a base name
			if baseName != "" {
				if current.Kind != KindObject {
					return nil
				}
				var exists bool
				current, exists = current.Object[baseName]
				if !exists {
					return nil
				}
			}

			// Then index into the array
			if current.Kind != KindArray || idx >= len(current.Array) {
				return nil
			}
			current = current.Array[idx]
		} else {
			// Regular object key
			if current.Kind != KindObject {
				return nil
			}
			var exists bool
			current, exists = current.Object[segment]
			if !exists {
				return nil
			}
		}
	}

	return current
}

// parseArrayNotation extracts the base name and index from array notation.
// Examples:
//   - "containers[0]" -> ("containers", 0, true)
//   - "[0]" -> ("", 0, true)
//   - "name" -> ("", 0, false)
func parseArrayNotation(segment string) (baseName string, idx int, isArray bool) {
	start := strings.Index(segment, "[")
	if start == -1 {
		return "", 0, false
	}

	end := strings.Index(segment, "]")
	if end == -1 || end <= start+1 {
		return "", 0, false
	}

	baseName = segment[:start]
	_, err := fmt.Sscanf(segment[start+1:end], "%d", &idx)
	if err != nil {
		return "", 0, false
	}

	return baseName, idx, true
}
