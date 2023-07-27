package main

import (
	"testing"
	"reflect"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 43}

	got := BubbleSort(arr)
	want := []int{3, 4, 7, 9, 43, 69, 420}

	if !reflect.DeepEqual(got,  want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
