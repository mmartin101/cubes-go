package main

import "testing"

func TestUsing3MaxPermutations(t *testing.T) {
	pa := compute(3)
	expectedBases := []float64{345, 384, 405}
	expectedValues := []float64{41063625, 56623104, 66430125}
	for i := 0; i < 3; i++ {
		if pa.bases[i] != expectedBases[i] {
			t.Error("Expected base %d, but got %d", expectedBases[i], pa.bases[i])
		}
		if pa.values[i] != expectedValues[i] {
			t.Error("Expected base %d, but got %d", expectedValues[i], pa.values[i])
		}
	}
}
