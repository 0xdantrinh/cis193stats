package cis193stats

import (
	"sort"
	"time"
)

// unixnano returns nanoseconds from UTC epoch
func unixnano() int64 {
	return time.Now().UTC().UnixNano()
}

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
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func sortedCopyDif(input data) (copy data) {
	if sort.Float64sAreSorted(input) {
		return input
	}
	copy = copyslice(input)
	sort.Float64s(copy)
	return
}
