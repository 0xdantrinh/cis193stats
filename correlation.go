package cis193stats

import (
	"errors"
	"math"
)

// Corr Calculates the Correlations between two data set
func Corr(data1, data2 data) (float64, error) {

	l1 := data1.Len()
	l2 := data2.Len()

	if l1 == 0 || l2 == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	if l1 != l2 {
		return math.NaN(), errors.New("Dataset Size Mismatch")
	}

	sdev1, _ := StandardDeviationPopulation(data1)

	sdev2, _ := StandardDeviationPopulation(data2)

	if sdev1 == 0 || sdev2 == 0 {
		return 0, nil
	}

	covp, _ := CovariancePopulation(data1, data2)

	denom := (sdev1 * sdev2)

	return covp / denom, nil
}
