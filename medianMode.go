package cis193stats

import (
	"errors"
	"math"
)

// Median gets the median in a data set
func Median(inp data) (median float64, err error) {

	c := sortedCopy(inp)

	l := len(c)
	if l == 0 {
		return math.NaN(), errors.New("Empty Set")
	} else if l%2 == 0 {
		median, _ = Mean(c[l/2-1 : l/2+1])
	} else {
		median = float64(c[l/2])
	}

	return median, nil
}

// Mode gets the mode of a data set
func Mode(inp data) (mode []float64, err error) {
	l := inp.Len()
	if l == 1 {
		return inp, nil
	} else if l == 0 {
		return nil, errors.New("Empty Set")
	}

	c := sortedCopyDif(inp)

	mode = make([]float64, 5)
	count, max := 1, 1
	for i := 1; i < l; i++ {
		if c[i] == c[i-1] {
			count++
		} else if count == max && max != 1 {
			mode = append(mode, c[i-1])
			count = 1
		} else if count > max {
			mode = append(mode[:0], c[i-1])
			max, count = count, 1
		} else {
			count = 1
		}
	}

	if count == max {
		mode = append(mode, c[l-1])
	} else if count > max {
		mode = append(mode[:0], c[l-1])
		max = count
	}

	if max == 1 {
		return data{}, nil
	}

	return mode, nil
}
