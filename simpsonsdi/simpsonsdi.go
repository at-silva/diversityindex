// simpsonsdi provides an implementation of the Simpson's Diversity Index
// This algorithm measures the diversity of a dataset based on the abundance of different elements.
// see: https://en.wikipedia.org/wiki/Diversity_index#Simpson_index
package simpsonsdi

import "math"

// CalculateSlice calculates the Simpson's Diversity Index for a given slice of values.
// It counts occurrences of each unique value and then computes the diversity index.
func CalculateSlice[T comparable](values []T) float64 {
	// Create a map to count occurrences of each unique value
	countByValue := make(map[T]uint64)
	for _, value := range values {
		countByValue[value] += 1
	}

	// Calculate diversity index using the CalculateMap function
	return CalculateMap(countByValue)
}

// CalculateMap calculates the Simpson's Diversity Index for a given map of values and their counts.
// It takes a map where keys represent values and values represent their counts.
func CalculateMap[TKey comparable, TCount uint | uint16 | uint32 | uint64](values map[TKey]TCount) float64 {
	var totalCount, ret float64

	// Calculate the total count of all values
	for _, count := range values {
		totalCount += float64(count)
	}

	// Calculate Simpson's Diversity Index using the counts and total count
	for _, count := range values {
		// Formula for Simpson's Diversity Index: 1 - Σ(pi^2), where pi = count of each value / total count
		ret += math.Pow(float64(count)/totalCount, 2)
	}

	// Finalize the calculation and return the Simpson's Diversity Index
	return 1 - ret
}
