package day1

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestReadFileToSlices(t *testing.T) {
	fileName := "./test-data.txt"

	expected := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}

	actual := ReadDay1FileToSlices(fileName)
	assert.Equal(t, expected, actual)
}

