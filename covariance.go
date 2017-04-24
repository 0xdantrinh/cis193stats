package cis193stats

import (
	"errors"
	"math"
)

// Covar is a measure of how much two sets of data change
func Covar(data1, data2 data) (float64, error) {

	l1 := data1.Len()
	l2 := data2.Len()

	if l1 == 0 || l2 == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	if l1 != l2 {
		return math.NaN(), errors.New("Dataset Size Mismatch")
	}

	m1, _ := Mean(data1)
	m2, _ := Mean(data2)

	// Calculate sum of squares
	var ss float64
	for i := 0; i < l1; i++ {
		delta1 := (data1.Get(i) - m1)
		delta2 := (data2.Get(i) - m2)
		ss += (delta1*delta2 - ss) / float64(i+1)
	}

	return ss * float64(l1) / float64(l1-1), nil
}

// CovarPopulation computes covariance for entire population between two variables.
func CovarPopulation(data1, data2 data) (float64, error) {

	l1 := data1.Len()
	l2 := data2.Len()

	if l1 == 0 || l2 == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	if l1 != l2 {
		return math.NaN(), errors.New("Dataset Size Mismatch")
	}

	m1, _ := Mean(data1)
	m2, _ := Mean(data2)

	var s float64
	for i := 0; i < l1; i++ {
		delta1 := (data1.Get(i) - m1)
		delta2 := (data2.Get(i) - m2)
		s += delta1 * delta2
	}

	return s / float64(l1), nil
}
