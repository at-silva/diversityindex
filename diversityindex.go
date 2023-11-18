package diversityindex

import (
	"errors"

	"github.com/at-silva/diversityindex/shannonsdi"
	"github.com/at-silva/diversityindex/simpsonsdi"
)

type (
	// DiversityIndex represents different types of diversity indices.
	DiversityIndex uint16
)

const (
	// DiversityIndexShannons represents the Shannon's Diversity Index.
	DiversityIndexShannons DiversityIndex = iota

	// DiversityIndexSimpsons represents the Simpson's Diversity Index.
	DiversityIndexSimpsons
)

var (
	// ErrInvalidDiversityIndex is an error for an invalid diversity index.
	ErrInvalidDiversityIndex = errors.New("invalid diversity index")
)

// CalculateMap calculates the diversity index for a given map of values based on the specified diversity index type.
// It accepts a map where keys represent values and values represent their counts.
// Returns the diversity index value or an error if the diversity index type is invalid.
func CalculateMap[TKey comparable, TCount uint | uint16 | uint32 | uint64](f DiversityIndex, values map[TKey]TCount) (float64, error) {
	switch f {
	case DiversityIndexShannons:
		return shannonsdi.CalculateMap(values), nil // Calculate Shannon's Diversity Index
	case DiversityIndexSimpsons:
		return simpsonsdi.CalculateMap(values), nil // Calculate Simpson's Diversity Index
	default:
		return 0, ErrInvalidDiversityIndex // Return error for an invalid diversity index type
	}
}

// CalculateSlice calculates the diversity index for a given slice of values based on the specified diversity index type.
// It accepts a slice of values and computes the diversity index accordingly.
// Returns the diversity index value or an error if the diversity index type is invalid.
func CalculateSlice[T comparable](f DiversityIndex, values []T) (float64, error) {
	switch f {
	case DiversityIndexShannons:
		return shannonsdi.CalculateSlice(values), nil // Calculate Shannon's Diversity Index
	case DiversityIndexSimpsons:
		return simpsonsdi.CalculateSlice(values), nil // Calculate Simpson's Diversity Index
	default:
		return 0, ErrInvalidDiversityIndex // Return error for an invalid diversity index type
	}
}
