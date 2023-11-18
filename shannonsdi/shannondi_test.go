package shannonsdi_test

import (
	"fmt"
	"testing"

	"github.com/at-silva/diversityindex/shannonsdi"
)

// TestCalculateSlice verifies the correctness of CalculateSlice function.
// It tests the CalculateSlice function against various input slices and expected output values.
func TestCalculateSlice(t *testing.T) {
	testCases := []struct {
		desc   string        // Description of the test case
		input  []interface{} // Input slice for diversity calculation
		output float64       // Expected output for diversity index
	}{
		{"int slices", []interface{}{1, 1, 1, 2, 2}, .67},
		{"string slices", []interface{}{"a", "a", "a", "b", "b"}, 0.67},
		{"float slices", []interface{}{1.1, 1.1, 1.1, 2.2, 2.2}, 0.67},
		{"empty slices", []interface{}{}, 0},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			v := shannonsdi.CalculateSlice(tC.input)
			if fmt.Sprintf("%.2f", tC.output) != fmt.Sprintf("%.2f", v) {
				t.Fatalf("shannonsdi for %s is wrong, expecting %.2f, got %.2f", tC.desc, tC.output, v)
			}
		})
	}
}

// TestCalculateMap verifies the correctness of CalculateMap function.
// It tests the CalculateMap function against various input maps and expected output values.
func TestCalculateMap(t *testing.T) {
	testCases := []struct {
		desc   string                 // Description of the test case
		input  map[interface{}]uint64 // Input map for diversity calculation
		output float64                // Expected output for diversity index
	}{
		{"string keyed maps", map[interface{}]uint64{"1": 40, "2": 20, "3": 15, "4": 8, "5": 22}, 1.49},
		{"int keyed maps", map[interface{}]uint64{1: 40, 2: 20, 3: 15, 4: 8, 5: 22}, 1.49},
		{"float keyed maps", map[interface{}]uint64{1.1: 40, 2.2: 20, 3.3: 15, 4.4: 8, 5.5: 22}, 1.49},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			v := shannonsdi.CalculateMap(tC.input)
			if fmt.Sprintf("%.2f", tC.output) != fmt.Sprintf("%.2f", v) {
				t.Fatalf("shannonsdi for %s is wrong, expecting diff %v, got %.2f", tC.desc, tC.output, v)
			}
		})
	}
}
