package cis193stats

import (
	"errors"
	"math"
)

// Min finds the lowest number in a set of data
func Min(inp data) (min float64, err error) {

	// Get the count of numbers in the slice
	l := inp.Len()

	// Return an error if there are no numbers
	if l == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	// Get the first value as the starting point
	min = inp.Get(0)

	// Iterate until done checking for a lower value
	for i := 1; i < l; i++ {
		if inp.Get(i) < min {
			min = inp.Get(i)
		}
	}
	return min, nil
}

// Max finds the highest number in a slice
func Max(inp data) (max float64, err error) {

	// Return an error if there are no numbers
	if inp.Len() == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	// Get the first value as the starting point
	max = inp.Get(0)

	// Loop and replace higher values
	for i := 1; i < inp.Len(); i++ {
		if inp.Get(i) > max {
			max = inp.Get(i)
		}
	}

	return max, nil
}
