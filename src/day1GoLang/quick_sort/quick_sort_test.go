package quickSort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	got := QuickSort(arr)
	want := []int{3, 4, 7, 9, 42, 69, 420}

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("got %d, want %d", got, want)
	}
}
