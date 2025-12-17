package main

import (
	"math/rand/v2"
	"testing"
)

func TestRandomFailure(t *testing.T) {
	// Randomly fails with a 10% chance
	if rand.Float64() < 0.10 {
		t.Fatalf("Test failed randomly!")
	}
	
	t.Log("Test passed!")
}