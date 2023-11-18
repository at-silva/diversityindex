// shannondsdi provides an implementation of the Shannons's Diversity Index
// see: https://en.wikipedia.org/wiki/Diversity_index#Shannon_index
package shannonsdi

import "math"

// CalculateSlice calculates the Shannon's Diversity Index for a given slice of values.
// It counts occurrences of each unique value and then computes the diversity index.
func CalculateSlice[T comparable](values []T) float64 {
	// Create a map to count occurrences of each unique value
	countByValue := make(map[T]uint64)
	for _, value := range values {
		countByValue[value] += 1
	}

	// Calculate diversity index using the countByValue map
	return CalculateMap(countByValue)
}

// CalculateMap calculates the Shannon's Diversity Index for a given map of values and their counts.
// It takes a map where keys represent values and values represent their counts.
func CalculateMap[TKey comparable, TCount uint | uint16 | uint32 | uint64](values map[TKey]TCount) float64 {
	var totalCount, ret, pi float64

	// Calculate the total count of all values
	for _, count := range values {
		totalCount += float64(count)
	}

	// Calculate Shannon's Diversity Index using the counts and total count
	for _, count := range values {
		pi = float64(count) / totalCount
		ret += pi * math.Log(pi)
	}

	if ret != 0 {
		ret *= -1
	}

	return ret
}
