package cis193stats

// data is a custom type for easier naming of data sets
type data []float64

// get corresponding index
func (dta data) Get(i int) float64 { return dta[i] }

// Min finds the minimum of a data set
func (dta data) Min() (float64, error) { return Min(dta) }

// Max finds the maximum of a data set
func (dta data) Max() (float64, error) { return Max(dta) }

// SumAll finds the sum of all the data set
func (dta data) SumAll() (float64, error) { return SumAll(dta) }

// Mean finds the mean of all the data set
func (dta data) Mean() (float64, error) { return Mean(dta) }

// GeoMean will find the geometric mean of the data set
func (dta data) GeoMean() (float64, error) { return GeoMean(dta) }

// Median finds the median of the data set
func (dta data) Median() (float64, error) { return Median(dta) }

// Mode finds the mode of the data set
func (dta data) Mode() ([]float64, error) { return Mode(dta) }

// SD finds the standard deviation of the data set
func (dta data) SD() (float64, error) {
	return SD(dta)
}

// SDPopulation finds the standard deviation from the population of the data set
func (dta data) SDPopulation() (float64, error) {
	return SDPopulation(dta)
}

// SDSample finds the standard deviaton from a sample
func (dta data) SDSample() (float64, error) {
	return SDSample(dta)
}

// Corr Calculates the Correlation between two data sets
func (dta data) Corr(d data) (float64, error) {
	return Corr(dta, d)
}

// InterQuartileRange finds the range between Q1 and Q3
func (dta data) InterQuartileRange() (float64, error) {
	return InterQuartileRange(dta)
}

// Var finds the amount of variation in the dataset
func (dta data) Var() (float64, error) {
	return Var(dta)
}

// PopulationVar finds the amount of variance within a population in the data set
func (dta data) PopulationVar() (float64, error) {
	return PopulationVar(dta)
}

// SampleVar finds the amount of variance within a sample in the data set
func (dta data) SampleVar() (float64, error) {
	return SampleVar(dta)
}

// Covar finds the covariance between two data set
func (dta data) Covar(d data) (float64, error) {
	return Covar(dta, d)
}

// CovarPopulation finds the population covariance between two data set.
func (dta data) CovarPopulation(d data) (float64, error) {
	return CovarPopulation(dta, d)
}

// returns length
func (dta data) Len() int { return len(dta) }

// Less returns if one number is less than another
func (dta data) Less(i, j int) bool { return dta[i] < dta[j] }

// Swap switches out two numbers in slice
func (dta data) Swap(i, j int) { dta[i], dta[j] = dta[j], dta[i] }
