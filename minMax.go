package cis193stats

import (
	"errors"
	"math"
)

// Min finds the lowest number in a set of data (uses classic algorithm to find min value of an array)
func Min(inp data) (min float64, err error) {

	length := inp.Len()

	if length == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	min = inp.Get(0)

	for i := 1; i < length; i++ {
		if inp.Get(i) < min {
			min = inp.Get(i)
		}
	}
	return min, nil
}

// Max finds the highest number in a slice (uses classic algorithm to find max value of an array)
func Max(inp data) (max float64, err error) {

	if inp.Len() == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	max = inp.Get(0)

	for i := 1; i < inp.Len(); i++ {
		if inp.Get(i) > max {
			max = inp.Get(i)
		}
	}

	return max, nil
}
