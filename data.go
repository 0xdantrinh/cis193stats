package stats

// data is a named type for []float64 with helper methods
type data []float64

// Get item in slice
func (dta data) Get(i int) float64 { return dta[i] }

// Min returns the minimum number in the data
func (dta data) Min() (float64, error) { return Min(dta) }

// Max returns the maximum number in the data
func (dta data) Max() (float64, error) { return Max(dta) }

// Sum returns the total of all the numbers in the data
func (dta data) SumAll() (float64, error) { return SumAll(dta) }

// Mean returns the mean of the data
func (dta data) Mean() (float64, error) { return Mean(dta) }

// GeometricMean returns the median of the data
func (dta data) GeometricMean() (float64, error) { return GeometricMean(dta) }

/*// Median returns the median of the data
func (dta data) Median() (float64, error) { return Median(dta) }

// Mode returns the mode of the data
func (dta data) Mode() ([]float64, error) { return Mode(dta) }


// MedianAbsoluteDeviation the median of the absolute deviations from the dataset median
func (dta data) MedianAbsoluteDeviation() (float64, error) {
	return MedianAbsoluteDeviation(dta)
}

// MedianAbsoluteDeviationPopulation finds the median of the absolute deviations from the population median
func (dta data) MedianAbsoluteDeviationPopulation() (float64, error) {
	return MedianAbsoluteDeviationPopulation(dta)
}

// StandardDeviation the amount of variation in the dataset
func (dta data) StandardDeviation() (float64, error) {
	return StandardDeviation(dta)
}

// StandardDeviationPopulation finds the amount of variation from the population
func (dta data) StandardDeviationPopulation() (float64, error) {
	return StandardDeviationPopulation(dta)
}

// StandardDeviationSample finds the amount of variation from a sample
func (dta data) StandardDeviationSample() (float64, error) {
	return StandardDeviationSample(dta)
}

// Percentile finds the relative standing in a slice of floats
func (dta data) Percentile(p float64) (float64, error) {
	return Percentile(dta, p)
}

// Correlation describes the degree of relationship between two sets of data
func (dta data) Correlation(d data) (float64, error) {
	return Correlation(dta, d)
}

// Quartile returns the three quartile points from a slice of data
func (dta data) Quartile(d data) (Quartiles, error) {
	return Quartile(d)
}

// InterQuartileRange finds the range between Q1 and Q3
func (dta data) InterQuartileRange() (float64, error) {
	return InterQuartileRange(dta)
}

// Sample returns sample from input with replacement or without
func (dta data) Sample(n int, r bool) ([]float64, error) {
	return Sample(dta, n, r)
}

// Variance the amount of variation in the dataset
func (dta data) Variance() (float64, error) {
	return Variance(dta)
}

// PopulationVariance finds the amount of variance within a population
func (dta data) PopulationVariance() (float64, error) {
	return PopulationVariance(dta)
}

// SampleVariance finds the amount of variance within a sample
func (dta data) SampleVariance() (float64, error) {
	return SampleVariance(dta)
}

// Covariance is a measure of how much two sets of data change
func (dta data) Covariance(d data) (float64, error) {
	return Covariance(dta, d)
}

// CovariancePopulation computes covariance for entire population between two variables.
func (dta data) CovariancePopulation(d data) (float64, error) {
	return CovariancePopulation(dta, d)
}*/

// Len returns length of slice
func (dta data) Len() int { return len(dta) }

// Less returns if one number is less than another
func (dta data) Less(i, j int) bool { return dta[i] < dta[j] }

// Swap switches out two numbers in slice
func (dta data) Swap(i, j int) { dta[i], dta[j] = dta[j], dta[i] }
