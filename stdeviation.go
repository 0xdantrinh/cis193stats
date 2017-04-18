package cis193stats

import (
	"errors"
	"math"
)

// _variance finds the variance for both population and sampleSize data
func _variance(inp data, sampleSize int) (variance float64, err error) {

	if inp.Len() == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	// Sum the square of the mean subtracted from each number
	m, _ := Mean(inp)

	for _, n := range inp {
		variance += (float64(n) - m) * (float64(n) - m)
	}

	// When getting the mean of the squared differences
	// "sampleSize" will allow us to know if it's a sampleSize
	// or population and wether to subtract by one or not
	denom := float64((inp.Len() - (1 * sampleSize)))
	return variance / denom, nil
}

// Variance the amount of variation in the dataset
func Variance(inp data) (sdev float64, err error) {
	return PopulationVariance(inp)
}

// PopulationVariance finds the amount of variance within a population
func PopulationVariance(inp data) (pvar float64, err error) {

	v, err := _variance(inp, 0)
	if err != nil {
		return math.NaN(), err
	}

	return v, nil
}

// SampleVariance finds the amount of variance within a sampleSize
func SampleVariance(inp data) (svar float64, err error) {

	v, err := _variance(inp, 1)
	if err != nil {
		return math.NaN(), err
	}

	return v, nil
}

// StandardDeviation the amount of variation in the dataset
func StandardDeviation(inp data) (std float64, err error) {
	return StandardDeviationPopulation(inp)
}

// StandardDeviationPopulation finds the amount of variation from the population
func StandardDeviationPopulation(inp data) (std float64, err error) {

	if inp.Len() == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	// Get the population variance
	vp, _ := PopulationVariance(inp)

	// Return the population standard deviation
	return math.Pow(vp, 0.5), nil
}

// StandardDeviationSample finds the amount of variation from a sampleSize
func StandardDeviationSample(inp data) (std float64, err error) {

	if inp.Len() == 0 {
		return math.NaN(), errors.New("Empty Set")
	}

	// Get the sampleSize variance
	vs, _ := SampleVariance(inp)

	// Return the sampleSize standard deviation
	return math.Pow(vs, 0.5), nil
}
