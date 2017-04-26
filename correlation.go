package cis193stats

import (
	"errors"
	"math"
)

// Corr Calculates the Correlations between two data set
func Corr(data1, data2 data) (float64, error) {

	length1 := data1.Len()
	length2 := data2.Len()

	if length1 == 0 || length2 == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	if length1 != length2 {
		return math.NaN(), errors.New("Dataset Size Mismatch")
	}

	sd_1, _ := SDPopulation(data1)

	sd_2, _ := SDPopulation(data2)

	if sd_1 == 0 || sd_2 == 0 {
		return 0, nil
	}

	covp, _ := CovarPopulation(data1, data2)

	denom := (sd_1 * sd_2)

	return covp / denom, nil
}
