package diversityindex_test

import (
	"fmt"
	"testing"

	"github.com/at-silva/diversityindex"
)

// TestCalculateSlice verifies the correctness of diversity index calculations for slices.
// It tests CalculateSlice function for Simpson's and Shannon's diversity indices against various input slices.
func TestCalculateSlice(t *testing.T) {
	// Test Simpson's Diversity Index calculation for a uint64 slice
	expectedSimpsons := 0.48
	calculatedSimpsons, err := diversityindex.CalculateSlice(diversityindex.DiversityIndexSimpsons, []uint64{1, 1, 1, 2, 2})
	if err != nil {
		t.Fatal(err)
	}
	if calculatedSimpsons != expectedSimpsons {
		t.Fatalf("diversityindex for slices is wrong, expecting %v, got %v", expectedSimpsons, calculatedSimpsons)
	}

	// Test Shannon's Diversity Index calculation for a uint64 slice
	expectedShannons := 0.67
	calculatedShannons, err := diversityindex.CalculateSlice(diversityindex.DiversityIndexShannons, []uint64{1, 1, 1, 2, 2})
	if err != nil {
		t.Fatal(err)
	}
	if fmt.Sprintf("%.2f", expectedShannons) != fmt.Sprintf("%.2f", calculatedShannons) {
		t.Fatalf("diversityindex for slices is wrong, expecting %v, got %.2f", expectedShannons, calculatedShannons)
	}
}

// TestCalculateInvalidIndex verifies the error message when an invalid diversity index is provided.
// It tests CalculateSlice and CalculateMap functions for an invalid diversity index.
func TestCalculateInvalidIndex(t *testing.T) {
	expectedError := diversityindex.ErrInvalidDiversityIndex

	// Test CalculateSlice with an invalid diversity index
	_, calculatedSliceError := diversityindex.CalculateSlice(diversityindex.DiversityIndex(100), []uint{})
	if expectedError != calculatedSliceError {
		t.Fatalf("diversityindex error message for invalid index is wrong, expecting '%v', got '%v'", expectedError, calculatedSliceError)
	}

	// Test CalculateMap with an invalid diversity index
	_, calculatedMapError := diversityindex.CalculateMap(diversityindex.DiversityIndex(100), map[string]uint{})
	if expectedError != calculatedMapError {
		t.Fatalf("diversityindex error message for invalid index is wrong, expecting '%v', got '%v'", expectedError, calculatedMapError)
	}
}

// TestCalculateMap verifies the correctness of diversity index calculations for maps.
// It tests CalculateMap function for Simpson's and Shannon's diversity indices against various input maps.
func TestCalculateMap(t *testing.T) {
	// Test Simpson's Diversity Index calculation for a map with integer keys
	expectedSimpsons := 0.48
	calculatedSimpsons, err := diversityindex.CalculateMap(diversityindex.DiversityIndexSimpsons, map[int]uint{1: 3, 2: 2})
	if err != nil {
		t.Fatal(err)
	}
	if calculatedSimpsons != expectedSimpsons {
		t.Fatalf("diversityindex for maps is wrong, expecting %v, got %v", expectedSimpsons, calculatedSimpsons)
	}

	// Test Shannon's Diversity Index calculation for a map with string keys
	expectedShannons := 1.49
	calculatedShannons, err := diversityindex.CalculateMap(diversityindex.DiversityIndexShannons, map[string]uint{"1": 40, "2": 20, "3": 15, "4": 8, "5": 22})
	if err != nil {
		t.Fatal(err)
	}
	if fmt.Sprintf("%.2f", expectedShannons) != fmt.Sprintf("%.2f", calculatedShannons) {
		t.Fatalf("shannondi for string keyed maps is wrong, expecting diff %v, got %.2f", expectedShannons, calculatedShannons)
	}
}
