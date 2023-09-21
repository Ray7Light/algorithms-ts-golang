package CompareBT

import (
	"testing"
)

func TestCompareBT(t *testing.T) {
	var tests = []struct {
		tree  BinaryNode
		tree2 BinaryNode
		want  bool
	}{
		{Tree, Tree, true},
		{Tree, Tree2, false},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := CompareBT(&tt.tree, &tt.tree2)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}
