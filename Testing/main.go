package main

import (
	"fmt"

	"github.com/dantrinh/cis193stats"
)

func main() {

	a, _ := cis193stats.Mean([]float64{1, 2, 3, 4, 5})
	fmt.Println(a) // 3

	a, _ = cis193stats.Median([]float64{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(a) // 4

	a, _ = cis193stats.Min([]float64{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(a) // 1.1

	a, _ = cis193stats.Max([]float64{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(a) // 5

	a, _ = cis193stats.SumAll([]float64{1.1, 2.2, 3.3})
	fmt.Println(a) // 6.6

	m, _ := cis193stats.Mode([]float64{5, 5, 3, 3, 4, 2, 1})
	fmt.Println(m) // [5 3]

	a, _ = cis193stats.PopulationVariance([]float64{1, 2, 3, 4, 5})
	fmt.Println(a) // 2

	a, _ = cis193stats.SampleVariance([]float64{1, 2, 3, 4, 5})
	fmt.Println(a) // 2.5

	a, _ = cis193stats.StandardDeviationPopulation([]float64{1, 2, 3})
	fmt.Println(a) // 0.816496580927726

	a, _ = cis193stats.StandardDeviationSample([]float64{1, 2, 3})
	fmt.Println(a) // 1

	a, _ = cis193stats.Correlation([]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 4, 5})
	fmt.Println(a) // ~1

	a, _ = cis193stats.Covariance([]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 4, 5})
	fmt.Println(a) // 2.5

	iqr, _ := cis193stats.InterQuartileRange([]float64{102, 104, 105, 107, 108, 109, 110, 112, 115, 116, 118})
	fmt.Println(iqr) // 10
}
