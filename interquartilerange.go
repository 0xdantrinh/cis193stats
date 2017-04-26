package cis193stats

import (
	"errors"
	"math"
)

// InterQuartileRange finds the range between Q1 and Q3 for a data set
func InterQuartileRange(input data) (float64, error) {

	length := input.Len()
	if length == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	copy := sortedCopy(input)

	var c1 int
	var c2 int
	if length%2 == 0 {
		c1 = length / 2
		c2 = length / 2
	} else {
		c1 = (length - 1) / 2
		c2 = c1 + 1
	}

	Q1, _ := Median(copy[:c1])
	Q3, _ := Median(copy[c2:])

	iqr := Q3 - Q1
	return iqr, nil
}
