package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Part1() {
	fileName := "./day-1/full-data.txt"

	slices := readFileToSlices(fileName)
	subSlices := sumSubslices(slices)

	max := findMax(subSlices)
	fmt.Printf("Largest sub slice: %v\n", max)
}

func Part2() {
	fileName := "./day-1/full-data.txt"

	slices := readFileToSlices(fileName)
	subSlices := sumSubslices(slices)

	sorted := sortSliceDesc(subSlices)
	first3 := sorted[0:3]
	sumFirst3 := sumSlice(first3)

	fmt.Printf("Sum of 3 biggest subslices: %v\n", sumFirst3)
}

func readFileToSlices(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var elements [][]int
	var currentElement []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			elements = append(elements, currentElement)
			currentElement = nil
		} else {
			currentNum, _ := strconv.Atoi(text)
			currentElement = append(currentElement, currentNum)
		}
	}
	elements = append(elements, currentElement)

	return elements
}

func sumSlice(slice []int) int {
	result := 0
	for _, value := range slice {
		result += value
	}
	return result
}

func findMax(slice []int) int {
	max := 0
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}

func sumSubslices(slice [][]int) []int {
	var result []int
	for _, value := range slice {
		result = append(result, sumSlice(value))
	}
	return result
}

func sortSliceDesc(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return slice
}
