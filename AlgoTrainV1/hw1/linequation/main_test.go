package main

import (
	"testing"
)

func TestSolveSLAU(t *testing.T) {
	tabletests := []struct {
		in       []float64
		wantCode int
		wantAns  []float64
	}{
		{
			in:       []float64{1, 0, 0, 1, 3, 3},
			wantCode: 2,
			wantAns:  []float64{3, 3},
		},
		{
			in:       []float64{1, 1, 2, 2, 1, 2},
			wantCode: 1,
			wantAns:  []float64{-1, 1},
		},
		{
			in:       []float64{0, 2, 0, 4, 1, 2},
			wantCode: 4,
			wantAns:  []float64{0.5},
		},
		{
			in:       []float64{0, 2, 1, 4, 1, 2},
			wantCode: 2,
			wantAns:  []float64{-0, 0.5},
		},
	}
	for _, tt := range tabletests {
		code, ans := solveSLAU(tt.in[0], tt.in[1], tt.in[2], tt.in[3], tt.in[4], tt.in[5])
		if code != tt.wantCode {
			t.Errorf("solveSLAU code: %v, want code %v", code, tt.wantCode)
		}
		for i := range tt.wantAns {
			if tt.wantAns[i] != ans[i] {
				t.Errorf("solveSLAU answer: %v, want answer %v", ans, tt.wantAns)
			}
		}
	}
}

func FuzzSolveSLAU(f *testing.F) {
	// Seed the fuzzer with some known test cases
	testCases := [][]float64{
		{1, 0, 0, 1, 3, 3},
		{1, 1, 2, 2, 1, 2},
		{0, 2, 0, 4, 1, 2},
		{0, 0, 0, 0, 0, 0},                   // Edge case: All zeros
		{1, 2, 3, 4, 5, 6},                   // Random values
		{1e9, 1e9, 1e9, 1e9, 1e9, 1e9},       // Large values
		{1e-9, 1e-9, 1e-9, 1e-9, 1e-9, 1e-9}, // Small values
	}

	for _, tc := range testCases {
		f.Add(tc[0], tc[1], tc[2], tc[3], tc[4], tc[5])
	}

	f.Fuzz(func(t *testing.T, a, b, c, d, e, f float64) {
		code, ans := solveSLAU(a, b, c, d, e, f)

		// Check if the output is within expected range
		if code < 0 || code > 5 {
			t.Errorf("Unexpected return code: %d for input (%f, %f, %f, %f, %f, %f)", code, a, b, c, d, e, f)
		}

		// Validate the solution if there's one
		if code == 2 && len(ans) != 2 {
			// Solution should satisfy the equations
			x, y := ans[0], ans[1]
			left1 := a*x + b*y
			left2 := c*x + d*y
			if !approxEqual(left1, e) || !approxEqual(left2, f) {
				t.Errorf("Invalid solution: (x = %.5f, y = %.5f, eCalc = %.5f, fCalc = %.5f) for input (%.5f, %.5f, %.5f, %.5f, %.5f, %.5f)", x, y, left1, left2, a, b, c, d, e, f)
			}
		}
		if ((code == 1 || code == 2) && len(ans) != 2) || ((code == 0 || code == 5) && len(ans) != 0) || ((code == 3 || code == 4) && len(ans) != 1) {
			t.Errorf("Invalid solution for input (%.5f, %.5f, %.5f, %.5f, %.5f, %.5f)", a, b, c, d, e, f)
		}
	})
}

// Helper function to compare floating point values with a small tolerance
func approxEqual(a, b float64) bool {
	const epsilon = 100
	return (a-b) < epsilon && (b-a) < epsilon
}
