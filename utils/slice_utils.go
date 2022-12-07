package utils

import (
	"sort"
)

func SumSlice(slice []int) int {
	result := 0
	for _, value := range slice {
		result += value
	}
	return result
}

func FindMax(slice []int) int {
	max := 0
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}

func SumSubslices(slice [][]int) []int {
	var result []int
	for _, value := range slice {
		result = append(result, SumSlice(value))
	}
	return result
}

func SortSliceDesc(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return slice
}
