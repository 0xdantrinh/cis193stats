package cis193stats

import (
	"errors"
	"math"
)

type Series []Coordinate

type Coordinate struct {
	X, Y float64
}

// LinearRegression uses the least squares regression on a data series
func LinearRegression(s Series) (regs Series, err error) {

	if len(s) == 0 {
		return nil, errors.New("Empty Set")
	}

	var sum [5]float64

	i := 0
	for ; i < len(s); i++ {
		sum[0] += s[i].X
		sum[1] += s[i].Y
		sum[2] += s[i].X * s[i].X
		sum[3] += s[i].X * s[i].Y
		sum[4] += s[i].Y * s[i].Y
	}

	f := float64(i)
	grad := (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	intercept := (sum[1] / f) - (grad * sum[0] / f)

	for j := 0; j < len(s); j++ {
		regs = append(regs, Coordinate{
			X: s[j].X,
			Y: s[j].X*grad + intercept,
		})
	}

	return regs, nil

}

// ExponentialRegression uses the exponential regression on a data series
func ExponentialRegression(s Series) (regs Series, err error) {

	if len(s) == 0 {
		return nil, errors.New("Empty Set")
	}

	var sum [6]float64

	for i := 0; i < len(s); i++ {
		sum[0] += s[i].X
		sum[1] += s[i].Y
		sum[2] += s[i].X * s[i].X * s[i].Y
		sum[3] += s[i].Y * math.Log(s[i].Y)
		sum[4] += s[i].X * s[i].Y * math.Log(s[i].Y)
		sum[5] += s[i].X * s[i].Y
	}

	denominator := (sum[1]*sum[2] - sum[5]*sum[5])
	a := math.Pow(math.E, (sum[2]*sum[3]-sum[5]*sum[4])/denominator)
	b := (sum[1]*sum[4] - sum[5]*sum[3]) / denominator

	for j := 0; j < len(s); j++ {
		regs = append(regs, Coordinate{
			X: s[j].X,
			Y: a * math.Exp(b*s[j].X),
		})
	}

	return regs, nil

}

// LogarithmicRegression uses the logarithmic regression on a data series
func LogarithmicRegression(s Series) (regs Series, err error) {

	if len(s) == 0 {
		return nil, errors.New("Empty Set")
	}

	var sum [4]float64

	i := 0
	for ; i < len(s); i++ {
		sum[0] += math.Log(s[i].X)
		sum[1] += s[i].Y * math.Log(s[i].X)
		sum[2] += s[i].Y
		sum[3] += math.Pow(math.Log(s[i].X), 2)
	}

	f := float64(i)
	a := (f*sum[1] - sum[2]*sum[0]) / (f*sum[3] - sum[0]*sum[0])
	b := (sum[2] - a*sum[0]) / f

	for j := 0; j < len(s); j++ {
		regs = append(regs, Coordinate{
			X: s[j].X,
			Y: b + a*math.Log(s[j].X),
		})
	}

	return regs, nil

}
