package main

import "testing"

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name       string
		in         []int
		wantResult bool
	}{
		{
			name:       "Small non symmetric array",
			in:         []int{1, 2, 3},
			wantResult: false,
		},
		{
			name:       "Empty array",
			in:         []int{},
			wantResult: true,
		},
		{
			name:       "Single element array",
			in:         []int{1},
			wantResult: true,
		},
		{
			name:       "Long symmetric array",
			in:         []int{1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1},
			wantResult: true,
		},
		{
			name:       "Short symmetric array",
			in:         []int{1, 2, 2, 1},
			wantResult: true,
		},
		{
			name:       "Super short symmetric array",
			in:         []int{2, 2},
			wantResult: true,
		},
		{
			name:       "Small non symmetric array",
			in:         []int{1, 2, 3, 4, 1},
			wantResult: false,
		},
		{
			name:       "Long symmetric 2",
			in:         []int{5, 4, 3, 2, 1, 2, 3, 4, 5},
			wantResult: true,
		},
		{
			name:       "Long non symmetric array",
			in:         []int{5, 4, 3, 2, 1, 2, 3, 4, 6},
			wantResult: false,
		},
		{
			name:       "Two different elements array",
			in:         []int{1, 2},
			wantResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSymmetric(tt.in)
			if result != tt.wantResult {
				t.Errorf("For input %v, expected %v but got %v", tt.in, tt.wantResult, result)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name       string
		in         []int
		wantResult []int
	}{
		{
			name:       "Small non symmetric array",
			in:         []int{1, 2, 3},
			wantResult: []int{3, 2, 1},
		},
		{
			name:       "Empty array",
			in:         []int{},
			wantResult: []int{},
		},
		{
			name:       "Single element array",
			in:         []int{1},
			wantResult: []int{1},
		},
		{
			name:       "Long symmetric array",
			in:         []int{1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1},
			wantResult: []int{1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:       "Short symmetric array",
			in:         []int{1, 2, 2, 1},
			wantResult: []int{1, 2, 2, 1},
		},
		{
			name:       "Super short symmetric array",
			in:         []int{2, 2},
			wantResult: []int{2, 2},
		},
		{
			name:       "Small non symmetric array",
			in:         []int{1, 2, 3, 4, 1},
			wantResult: []int{1, 4, 3, 2, 1},
		},
		{
			name:       "Long symmetric 2",
			in:         []int{5, 4, 3, 2, 1, 2, 3, 4, 5},
			wantResult: []int{5, 4, 3, 2, 1, 2, 3, 4, 5},
		},
		{
			name:       "Long non symmetric array",
			in:         []int{5, 4, 3, 2, 1, 2, 3, 4, 6},
			wantResult: []int{6, 4, 3, 2, 1, 2, 3, 4, 5},
		},
		{
			name:       "Two different elements array",
			in:         []int{1, 2},
			wantResult: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseArray(tt.in)
			for i := range result {
				if tt.wantResult[i] != result[i] {
					t.Errorf("For input %v, expected %v but got %v", tt.in, tt.wantResult, result)
				}
			}
		})
	}
}
