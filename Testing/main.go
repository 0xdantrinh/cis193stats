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

	c := []stats.Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := stats.LinearRegression(c)
	fmt.Println(r) // [{1 2.3800000000000026} {2 3.0800000000000014} {3 3.7800000000000002} {4 4.479999999999999} {5 5.179999999999998}]

	r, _ = stats.ExponentialRegression(c)
	fmt.Println(r) // [{1 2.5150181024736638} {2 3.032084111136781} {3 3.6554544271334493} {4 4.406984298281804} {5 5.313022222665875}]

	r, _ = stats.LogarithmicRegression(c)
	fmt.Println(r) // [{1 2.1520822363811702} {2 3.3305559222492214} {3 4.019918836568674} {4 4.509029608117273} {5 4.888413396683663}]

}
