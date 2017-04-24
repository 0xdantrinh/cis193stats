package cis193stats

import (
	"sort"
)

// copyslice copies a slice of float64s
func copyslice(input data) data {
	s := make(data, input.Len())
	copy(s, input)
	return s
}

// sortedCopy returns a sorted copy of float64s
func sortedCopy(input data) (copy data) {
	copy = copyslice(input)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
func sortedCopyDif(input data) (copy data) {
	if sort.Float64sAreSorted(input) {
		return input
	}
	copy = copyslice(input)
	sort.Float64s(copy)
	return
}
