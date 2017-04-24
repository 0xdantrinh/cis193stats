package cis193stats

import (
	"errors"
	"math"
)

func SumAll(inp data) (out float64, err error) {

	if inp.Len() == 0 {
		return math.NaN(), errors.New("Empty Set")
	} else {
		for _, v := range inp {
			out += v
		}
	}
	return out, nil
}

// Mean gets the average of a slice of numbers for data set
func Mean(inp data) (float64, error) {

	if inp.Len() == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	sum, _ := inp.SumAll()

	return sum / float64(inp.Len()), nil
}

// GeoMean gets the geometric mean for a slice of numbers for data set
func GeomMean(inp data) (float64, error) {

	length := inp.Len()
	if length == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	var p float64
	for _, v := range inp {
		if p == 0 {
			p = v
		} else {
			p *= v
		}
	}
	out := math.Pow(p, 1/float64(length))
	return out, nil
}
