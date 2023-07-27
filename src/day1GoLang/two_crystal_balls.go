package main

import (
	"math"
)

func TwoCrystalBalls(arr [10000]bool) int {
	jumpAmount := int(math.Sqrt(float64(len(arr)))) 
	var i int 

	for i = jumpAmount; i < len(arr); i += jumpAmount {
		if arr[i] == true {
			break
		}
	}

	for j := i - jumpAmount; j < len(arr); j++ {
		if arr[j] == true {
			return j
		}
	}

	return -1 
}
