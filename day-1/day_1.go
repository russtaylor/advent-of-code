package day1

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part1() {
	fileName := "./day-1/full-data.txt"

	slices := ReadDay1FileToSlices(fileName)
	subSlices := utils.SumSubslices(slices)

	max := utils.FindMax(subSlices)
	fmt.Printf("Largest sub slice: %v\n", max)
}

func Part2() {
	fileName := "./day-1/full-data.txt"

	slices := ReadDay1FileToSlices(fileName)
	subSlices := utils.SumSubslices(slices)

	sorted := utils.SortSliceDesc(subSlices)
	first3 := sorted[0:3]
	sumFirst3 := utils.SumSlice(first3)

	fmt.Printf("Sum of 3 biggest subslices: %v\n", sumFirst3)
}

func ReadDay1FileToSlices(filename string) [][]int {
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
