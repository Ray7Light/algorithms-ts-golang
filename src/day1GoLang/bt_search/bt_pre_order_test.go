package BTSearch

import (
	"reflect"
	"testing"
)

func TestBTPreOrder(t *testing.T) {
	result := BTPreOrder(Tree)
	want := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("want: %d, got %d", want, result)
	}
}
