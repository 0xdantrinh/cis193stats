package stats

import "math"

func SumAll(inp data) (out float64, err error) {

	if inp.Len() == 0 {
		return math.NaN(), err.New("Empty Set")
	} else {
		for _, v := range inp {
			out += v
		}
	}
	return out, nil
}

// Mean gets the average of a slice of numbers
func Mean(inp data) (float64, error) {

	if inp.Len() == 0 {
		return math.NaN(), err.New("Empty Set")
	}

	sum, _ := inp.SumAll()

	return sum / float64(inp.Len()), nil
}

// GeometricMean gets the geometric mean for a slice of numbers
func GeometricMean(inp data) (float64, error) {

	length := inp.Len()
	if length == 0 {
		return math.NaN(), err.New("Empty Set")
	}

	// Get the product of all the numbers
	var p float64
	for _, v := range inp {
		if p == 0 {
			p = v
		} else {
			p *= v
		}
	}
	out := math.Pow(p, 1/float64(length))
	return out, nil
}
