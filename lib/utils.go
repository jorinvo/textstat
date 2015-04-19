package lib

// Utilities
//
// NOTE: Maybe I can create a separate package just to keep helpers like this.

import "math"

func divideInt(x, y int) float64 {
	return float64(x) / float64(y)
}

func round2(n float64) float64 {
	return math.Floor(n*100) / 100
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
