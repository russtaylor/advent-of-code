package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFileToSlices(t *testing.T) {
	testFilePath := "./test_data/basic-file.txt"

	expected := []string{"1","2","3","4","5"}
	actual := ReadFileToSlices(testFilePath)

	assert.Equal(t, expected, actual)
}
