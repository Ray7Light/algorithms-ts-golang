package BFSearch

import ( 
	"testing"
)

func TestBTBFS(t *testing.T) {
	var tests = []struct{
		tree BinaryNode
		value int
		want bool
	}{
		{Tree, 45, true},
		{Tree, 7, true},
		{Tree, 69, false},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := BTBFS(tt.tree, tt.value)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}
