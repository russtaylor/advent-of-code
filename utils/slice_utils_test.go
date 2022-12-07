package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumSlice(t *testing.T) {
	input := []int{1, 2, 3}

	expected := 6
	actual := SumSlice(input)
	assert.Equal(t, expected, actual)
}

func TestFindMax(t *testing.T) {
	input := []int{2, 4, 8, 4}

	expected := 8
	actual := FindMax(input)

	assert.Equal(t, expected, actual)
}

func TestSumSubslices(t *testing.T) {
	input := [][]int{
		{10, 20, 30},
		{40},
		{50, 60},
		{70, 80, 90},
		{100},
	}

	expected := []int{60, 40, 110, 240, 100}
	actual := SumSubslices(input)

	assert.Equal(t, expected, actual)
}

func TestSortSliceDesc(t *testing.T) {
	input := []int{10, 50, 5, 30, 60, 8, 32}

	expected := []int{60, 50, 32, 30, 10, 8, 5}
	actual := SortSliceDesc(input)

	assert.Equal(t, expected, actual)
}