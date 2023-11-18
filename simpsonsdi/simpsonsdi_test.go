package simpsonsdi_test

import (
	"fmt"
	"testing"

	"github.com/at-silva/diversityindex/simpsonsdi"
)

// TestCalculateSlices verifies the correctness of CalculateSlice function.
// It tests the CalculateSlice function against various input slices and expected output values.
func TestCalculateSlices(t *testing.T) {
	s := make([]interface{}, 0, 100)
	for i := 0; i <= 100; i++ {
		s = append(s, uint64(i))
	}

	testCases := []struct {
		desc   string        // Description of the test case
		input  []interface{} // Input slice for diversity calculation
		output float64       // Expected output for Simpson's Diversity Index
	}{
		{"int slices", []interface{}{1, 1, 1, 2, 2}, 0.48},
		{"string slices", []interface{}{"1", "1", "1", "2", "2"}, 0.48},
		{"float slices", []interface{}{1.1, 1.1, 1.1, 2.2, 2.2}, 0.48},
		{"index min", []interface{}{1, 1, 1}, 0},
		{"index max", s, .99},
		{"empty slice", []interface{}{}, 1},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			v := simpsonsdi.CalculateSlice(tC.input)
			if fmt.Sprintf("%.2f", tC.output) != fmt.Sprintf("%.2f", v) {
				t.Fatalf("simpsonsdi for %s is wrong, expecting %v, got %v", tC.desc, tC.output, v)
			}
		})
	}
}

// TestCalculateMap verifies the correctness of CalculateMap function.
// It tests the CalculateMap function against various input maps and expected output values.
func Test(t *testing.T) {
	testCases := []struct {
		desc   string               // Description of the test case
		input  map[interface{}]uint // Input slice for diversity calculation
		output float64              // Expected output for Simpson's Diversity Index
	}{
		{"string keyed maps", map[interface{}]uint{"1": 81, "2": 2, "3": 2, "4": 2, "5": 2}, .17},
		{"int keyed maps", map[interface{}]uint{1: 3, 2: 2}, .48},
		{"float keyed maps", map[interface{}]uint{1.1: 3, 2.2: 2}, .48},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			v := simpsonsdi.CalculateMap(tC.input)
			if fmt.Sprintf("%.2f", tC.output) != fmt.Sprintf("%.2f", v) {
				t.Fatalf("simpsonsdi for %s is wrong, expecting diff %v, got %.2f", tC.desc, tC.output, v)
			}
		})
	}
}
