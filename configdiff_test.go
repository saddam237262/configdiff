package configdiff

import "testing"

func TestChangeTypeString(t *testing.T) {
	tests := []struct {
		ct   ChangeType
		want string
	}{
		{ChangeTypeAdd, "add"},
		{ChangeTypeRemove, "remove"},
		{ChangeTypeModify, "modify"},
		{ChangeTypeMove, "move"},
	}

	for _, tt := range tests {
		t.Run(string(tt.ct), func(t *testing.T) {
			if got := string(tt.ct); got != tt.want {
				t.Errorf("ChangeType = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions(t *testing.T) {
	opts := Options{
		IgnorePaths:  []string{"metadata.timestamp"},
		ArraySetKeys: map[string]string{"spec.containers": "name"},
		Coercions: Coercions{
			NumericStrings: true,
			BoolStrings:    true,
		},
		StableOrder: true,
	}

	if len(opts.IgnorePaths) != 1 {
		t.Errorf("IgnorePaths len = %v, want 1", len(opts.IgnorePaths))
	}
	if opts.ArraySetKeys["spec.containers"] != "name" {
		t.Errorf("ArraySetKeys = %v, want 'name'", opts.ArraySetKeys["spec.containers"])
	}
	if !opts.Coercions.NumericStrings {
		t.Error("NumericStrings = false, want true")
	}
	if !opts.StableOrder {
		t.Error("StableOrder = false, want true")
	}
}
