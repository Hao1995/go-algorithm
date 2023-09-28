package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

// TestQuickSort tests QuickSort
func TestQuickSort(t *testing.T) {

	tests := []struct {
		data   []int
		expRes []int
	}{
		{
			data:   []int{},
			expRes: []int{},
		},
		{
			data:   []int{1},
			expRes: []int{1},
		},
		{
			data:   []int{3, 2},
			expRes: []int{2, 3},
		},
		{
			data:   []int{3, 2, 1},
			expRes: []int{1, 2, 3},
		},
		{
			data:   []int{1, 2, 3},
			expRes: []int{1, 2, 3},
		},
		{
			data:   []int{3, 3, 1},
			expRes: []int{1, 3, 3},
		},
		{
			data:   []int{20, 10, 30, 40, 70, 60, 50},
			expRes: []int{10, 20, 30, 40, 50, 60, 70},
		},
		{
			data:   []int{33, 7, 24, 13, 12, 27, 12, 24, 5, 23, 29, 48, 30, 31},
			expRes: []int{5, 7, 12, 12, 13, 23, 24, 24, 27, 29, 30, 31, 33, 48},
		},
	}

	for _, test := range tests {
		QuickSort(test.data)
		assert.DeepEqual(t, test.expRes, test.data)
	}
}
