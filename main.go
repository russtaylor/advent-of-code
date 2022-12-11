package main

import (
	"advent-of-code/day-1"
	day2 "advent-of-code/day-2"
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
	case 2:
		switch part {
		case 1:
			day2.Part1()
		case 2:
			day2.Part2()
		}
	} // Okay, now I'm switching away from Go for now.
}
