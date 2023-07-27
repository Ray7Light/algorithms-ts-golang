package main

import ( 
	"fmt"
	"testing"
)

func TestBinarySearchList(t *testing.T) {
	var testedArr = []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	var tests = []struct {
		value int
		want bool
	} {
		{69, true},
		{1336, false},
		{69420, true},
		{69421, false},
		{1, true},
		{0, false},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.value)
		t.Run(testname, func(t *testing.T) {
			got := BinarySearchList(testedArr, tt.value)
			if got != tt.want {
				t.Errorf("got %t want %t", got, tt.want)
			}
		})
	}
}
