package main

import (
	"math/rand"
	"time"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	randIdx := rand.Intn(10000) 
	var data [10000]bool 

	for i := randIdx; i < 10000; i++ {
		data[i] = true
	}
	
	ans1 := TwoCrystalBalls(data)
	
	if ans1 != randIdx {
		t.Errorf("got %d, want %d", ans1, randIdx)
	}

	var falseArr [10000]bool
	ans2 := TwoCrystalBalls(falseArr)

	if ans2 != -1 {
		t.Errorf("got %d, want %d", ans2, -1)
	}
}
