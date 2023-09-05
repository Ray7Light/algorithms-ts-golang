package BTSearch

import (
	"reflect"
	"testing"
)

func TestBTInOrder(t *testing.T) {
	result := BTInOrder(Tree)
	want := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("want: %d, got %d", want, result)
	}
}
