package main

import (
	"advent-of-code/day-1"
	"fmt"
	"os"
	"strconv"
)

func main() {
	day, _ := strconv.Atoi(os.Args[1])
	part, _ := strconv.Atoi(os.Args[2])

	fmt.Printf("Running Day %v, Part %v\n", day, part)
	switch day {
	case 1:
		switch part {
		case 1:
			day1.Part1()
		case 2:
			day1.Part2()
		}
	}
}
