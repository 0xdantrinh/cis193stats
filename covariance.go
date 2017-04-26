package cis193stats

import (
	"errors"
	"math"
)

// Covar is the covariance of two data sets
func Covar(data1, data2 data) (float64, error) {

	length1 := data1.Len()
	length2 := data2.Len()

	if length1 == 0 || length2 == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	if length1 != length2 {
		return math.NaN(), errors.New("Dataset Size Mismatch")
	}

	m1, _ := Mean(data1)
	m2, _ := Mean(data2)

	var ss float64
	for i := 0; i < length1; i++ {
		delta1 := (data1.Get(i) - m1)
		delta2 := (data2.Get(i) - m2)
		ss += (delta1*delta2 - ss) / float64(i+1)
	}

	return ss * float64(length1) / float64(length1-1), nil
}

// CovarPopulation is the population covariance.
func CovarPopulation(data1, data2 data) (float64, error) {

	length1 := data1.Len()
	length2 := data2.Len()

	if length1 == 0 || length2 == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	if length1 != length2 {
		return math.NaN(), errors.New("Dataset Size Mismatch")
	}

	m1, _ := Mean(data1)
	m2, _ := Mean(data2)

	var s float64
	for i := 0; i < length1; i++ {
		delta1 := (data1.Get(i) - m1)
		delta2 := (data2.Get(i) - m2)
		s += delta1 * delta2
	}

	return s / float64(length1), nil
}
