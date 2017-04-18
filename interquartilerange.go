package cis193stats

import (
	"errors"
	"math"
)

/*// Quartiles holds the three quartile points
type Quartiles struct {
	Q1 float64
	Q2 float64
	Q3 float64
}*/

// InterQuartileRange finds the range between Q1 and Q3
func InterQuartileRange(input data) (float64, error) {

	length := input.Len()
	if length == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	// Start by sorting a copy of the slice
	copy := sortedCopy(input)

	// Find the cutoff places depeding on if
	// the input slice length is even or odd
	var c1 int
	var c2 int
	if length%2 == 0 {
		c1 = length / 2
		c2 = length / 2
	} else {
		c1 = (length - 1) / 2
		c2 = c1 + 1
	}

	// Find the Medians with the cutoff points
	Q1, _ := Median(copy[:c1])
	//Q2, _ := Median(copy)
	Q3, _ := Median(copy[c2:])

	iqr := Q3 - Q1
	return iqr, nil
}
