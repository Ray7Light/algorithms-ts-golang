package BTSearch

import (
	"reflect"
	"testing"
)

func TestBTPostOrder(t *testing.T) {
	result := BTPostOrder(Tree)
	want := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("want: %d, got %d", want, result)
	}
}
