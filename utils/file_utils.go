package utils

import (
	"bufio"
	"os"
)

func ReadFileToSlices(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}
	return lines
}