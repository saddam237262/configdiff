package tree

import (
	"testing"
)

func TestNodeKindString(t *testing.T) {
	tests := []struct {
		kind NodeKind
		want string
	}{
		{KindNull, "null"},
		{KindBool, "bool"},
		{KindNumber, "number"},
		{KindString, "string"},
		{KindObject, "object"},
		{KindArray, "array"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.kind.String(); got != tt.want {
				t.Errorf("NodeKind.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNodes(t *testing.T) {
	t.Run("NewNull", func(t *testing.T) {
		n := NewNull()
		if n.Kind != KindNull {
			t.Errorf("NewNull() kind = %v, want %v", n.Kind, KindNull)
		}
		if n.Value != nil {
			t.Errorf("NewNull() value = %v, want nil", n.Value)
		}
	})

	t.Run("NewBool", func(t *testing.T) {
		n := NewBool(true)
		if n.Kind != KindBool {
			t.Errorf("NewBool() kind = %v, want %v", n.Kind, KindBool)
		}
		if n.Value != true {
			t.Errorf("NewBool() value = %v, want true", n.Value)
		}
	})

	t.Run("NewNumber", func(t *testing.T) {
		n := NewNumber(42.5)
		if n.Kind != KindNumber {
			t.Errorf("NewNumber() kind = %v, want %v", n.Kind, KindNumber)
		}
		if n.Value != 42.5 {
			t.Errorf("NewNumber() value = %v, want 42.5", n.Value)
		}
	})

	t.Run("NewString", func(t *testing.T) {
		n := NewString("hello")
		if n.Kind != KindString {
			t.Errorf("NewString() kind = %v, want %v", n.Kind, KindString)
		}
		if n.Value != "hello" {
			t.Errorf("NewString() value = %v, want hello", n.Value)
		}
	})

	t.Run("NewObject", func(t *testing.T) {
		obj := map[string]*Node{
			"key": NewString("value"),
		}
		n := NewObject(obj)
		if n.Kind != KindObject {
			t.Errorf("NewObject() kind = %v, want %v", n.Kind, KindObject)
		}
		if len(n.Object) != 1 {
			t.Errorf("NewObject() len = %v, want 1", len(n.Object))
		}
	})

	t.Run("NewArray", func(t *testing.T) {
		arr := []*Node{NewString("item")}
		n := NewArray(arr)
		if n.Kind != KindArray {
			t.Errorf("NewArray() kind = %v, want %v", n.Kind, KindArray)
		}
		if len(n.Array) != 1 {
			t.Errorf("NewArray() len = %v, want 1", len(n.Array))
		}
	})
}

func TestNodeEqual(t *testing.T) {
	tests := []struct {
		name string
		a    *Node
		b    *Node
		want bool
	}{
		{
			name: "both nil",
			a:    nil,
			b:    nil,
			want: true,
		},
		{
			name: "one nil",
			a:    NewNull(),
			b:    nil,
			want: false,
		},
		{
			name: "different kinds",
			a:    NewString("test"),
			b:    NewNumber(42),
			want: false,
		},
		{
			name: "equal null",
			a:    NewNull(),
			b:    NewNull(),
			want: true,
		},
		{
			name: "equal bool",
			a:    NewBool(true),
			b:    NewBool(true),
			want: true,
		},
		{
			name: "unequal bool",
			a:    NewBool(true),
			b:    NewBool(false),
			want: false,
		},
		{
			name: "equal number",
			a:    NewNumber(42),
			b:    NewNumber(42),
			want: true,
		},
		{
			name: "unequal number",
			a:    NewNumber(42),
			b:    NewNumber(43),
			want: false,
		},
		{
			name: "equal string",
			a:    NewString("hello"),
			b:    NewString("hello"),
			want: true,
		},
		{
			name: "unequal string",
			a:    NewString("hello"),
			b:    NewString("world"),
			want: false,
		},
		{
			name: "equal object",
			a: NewObject(map[string]*Node{
				"key": NewString("value"),
			}),
			b: NewObject(map[string]*Node{
				"key": NewString("value"),
			}),
			want: true,
		},
		{
			name: "unequal object - different values",
			a: NewObject(map[string]*Node{
				"key": NewString("value1"),
			}),
			b: NewObject(map[string]*Node{
				"key": NewString("value2"),
			}),
			want: false,
		},
		{
			name: "unequal object - different keys",
			a: NewObject(map[string]*Node{
				"key1": NewString("value"),
			}),
			b: NewObject(map[string]*Node{
				"key2": NewString("value"),
			}),
			want: false,
		},
		{
			name: "equal array",
			a:    NewArray([]*Node{NewString("a"), NewString("b")}),
			b:    NewArray([]*Node{NewString("a"), NewString("b")}),
			want: true,
		},
		{
			name: "unequal array - different length",
			a:    NewArray([]*Node{NewString("a")}),
			b:    NewArray([]*Node{NewString("a"), NewString("b")}),
			want: false,
		},
		{
			name: "unequal array - different elements",
			a:    NewArray([]*Node{NewString("a"), NewString("b")}),
			b:    NewArray([]*Node{NewString("a"), NewString("c")}),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Equal(tt.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeClone(t *testing.T) {
	t.Run("nil node", func(t *testing.T) {
		var n *Node
		cloned := n.Clone()
		if cloned != nil {
			t.Errorf("Clone() = %v, want nil", cloned)
		}
	})

	t.Run("scalar node", func(t *testing.T) {
		n := NewString("test")
		cloned := n.Clone()
		if !n.Equal(cloned) {
			t.Error("Clone() not equal to original")
		}
		if n == cloned {
			t.Error("Clone() returned same pointer")
		}
	})

	t.Run("object node", func(t *testing.T) {
		n := NewObject(map[string]*Node{
			"key": NewString("value"),
		})
		cloned := n.Clone()
		if !n.Equal(cloned) {
			t.Error("Clone() not equal to original")
		}
		if n == cloned {
			t.Error("Clone() returned same pointer")
		}
		if n.Object["key"] == cloned.Object["key"] {
			t.Error("Clone() did not deep copy object values")
		}
	})

	t.Run("array node", func(t *testing.T) {
		n := NewArray([]*Node{NewString("item")})
		cloned := n.Clone()
		if !n.Equal(cloned) {
			t.Error("Clone() not equal to original")
		}
		if n == cloned {
			t.Error("Clone() returned same pointer")
		}
		if n.Array[0] == cloned.Array[0] {
			t.Error("Clone() did not deep copy array elements")
		}
	})
}

func TestNodeSortedKeys(t *testing.T) {
	t.Run("object node", func(t *testing.T) {
		n := NewObject(map[string]*Node{
			"z": NewString("last"),
			"a": NewString("first"),
			"m": NewString("middle"),
		})
		keys := n.SortedKeys()
		want := []string{"a", "m", "z"}
		if len(keys) != len(want) {
			t.Fatalf("SortedKeys() len = %v, want %v", len(keys), len(want))
		}
		for i := range keys {
			if keys[i] != want[i] {
				t.Errorf("SortedKeys()[%d] = %v, want %v", i, keys[i], want[i])
			}
		}
	})

	t.Run("non-object node", func(t *testing.T) {
		n := NewString("test")
		keys := n.SortedKeys()
		if keys != nil {
			t.Errorf("SortedKeys() = %v, want nil", keys)
		}
	})
}

func TestSetPaths(t *testing.T) {
	root := NewObject(map[string]*Node{
		"spec": NewObject(map[string]*Node{
			"containers": NewArray([]*Node{
				NewObject(map[string]*Node{
					"name": NewString("nginx"),
				}),
			}),
		}),
	})

	root.SetPaths("/")

	tests := []struct {
		path string
		want string
	}{
		{"/", "/"},
		{"/spec", "/spec"},
		{"/spec/containers", "/spec/containers"},
		{"/spec/containers[0]", "/spec/containers[0]"},
		{"/spec/containers[0]/name", "/spec/containers[0]/name"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			node := root.GetByPath(tt.path)
			if node == nil {
				t.Fatalf("GetByPath(%q) returned nil", tt.path)
			}
			if node.Path != tt.want {
				t.Errorf("Path = %q, want %q", node.Path, tt.want)
			}
		})
	}
}

func TestParsePath(t *testing.T) {
	tests := []struct {
		path string
		want []string
	}{
		{"", nil},
		{"/", nil},
		{"/key", []string{"key"}},
		{"/spec/containers", []string{"spec", "containers"}},
		{"/spec/containers[0]/name", []string{"spec", "containers[0]", "name"}},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := ParsePath(tt.path)
			if len(got) != len(tt.want) {
				t.Fatalf("ParsePath() len = %v, want %v", len(got), len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("ParsePath()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestGetByPath(t *testing.T) {
	root := NewObject(map[string]*Node{
		"spec": NewObject(map[string]*Node{
			"replicas": NewNumber(3),
			"containers": NewArray([]*Node{
				NewObject(map[string]*Node{
					"name":  NewString("nginx"),
					"image": NewString("nginx:latest"),
				}),
			}),
		}),
	})

	root.SetPaths("/")

	tests := []struct {
		path    string
		wantNil bool
		check   func(*testing.T, *Node)
	}{
		{
			path:    "/",
			wantNil: false,
			check: func(t *testing.T, n *Node) {
				if n.Kind != KindObject {
					t.Errorf("Root node kind = %v, want %v", n.Kind, KindObject)
				}
			},
		},
		{
			path:    "/spec/replicas",
			wantNil: false,
			check: func(t *testing.T, n *Node) {
				if n.Kind != KindNumber || n.Value != 3.0 {
					t.Errorf("Node = %v, want number 3", n.Value)
				}
			},
		},
		{
			path:    "/spec/containers[0]/name",
			wantNil: false,
			check: func(t *testing.T, n *Node) {
				if n.Kind != KindString || n.Value != "nginx" {
					t.Errorf("Node = %v, want string 'nginx'", n.Value)
				}
			},
		},
		{
			path:    "/nonexistent",
			wantNil: true,
		},
		{
			path:    "/spec/containers[99]",
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := root.GetByPath(tt.path)
			if tt.wantNil {
				if got != nil {
					t.Errorf("GetByPath() = %v, want nil", got)
				}
			} else {
				if got == nil {
					t.Fatal("GetByPath() = nil, want non-nil")
				}
				if tt.check != nil {
					tt.check(t, got)
				}
			}
		})
	}
}
