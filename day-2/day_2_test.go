package day2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreRPSFile(t *testing.T) {
	testFilename := "./test-data.txt"

	expected := 15
	result := ScoreRPSFile(testFilename)

	assert.Equal(t, expected, result)
}

func TestScoreRPSRound(t *testing.T) {
	scenarios := []struct {
		input    string
		expected int
	}{
		{"A Y", 8},
		{"B X", 1},
		{"C Z", 6},
	}

	for _, scenario := range scenarios {
		testName := fmt.Sprintf("%v", scenario.input)
		t.Run(testName, func(t *testing.T) {
			result := ScoreRPSRound(scenario.input)
			assert.Equal(t, scenario.expected, result)
		})
	}
}

// Test scenarios!
func TestPlayRPSRoundScenarios(t *testing.T) {
	scenarios := []struct {
		elfPlay       string
		myPlay        string
		expectedScore int
	}{
		{"A", "X", 3},
		{"A", "Y", 6},
		{"A", "Z", 0},
		{"B", "X", 0},
		{"B", "Y", 3},
		{"B", "Z", 6},
		{"C", "X", 6},
		{"C", "Y", 0},
		{"C", "Z", 3},
	}

	for _, scenario := range scenarios {
		testName := fmt.Sprintf("%v_%v", scenario.elfPlay, scenario.myPlay)
		t.Run(testName, func(t *testing.T) {
			result := PlayRPSRound(scenario.elfPlay, scenario.myPlay)
			assert.Equal(t, scenario.expectedScore, result)
		})
	}
}

func TestScoreRoundPart2(t *testing.T) {
	scenarios := []struct {
		input    string
		expected int
	}{
		{"A Y", 4},
		{"B X", 1},
		{"C Z", 7},
	}

	for _, scenario := range scenarios {
		testName := fmt.Sprintf("%v", scenario.input)
		t.Run(testName, func(t *testing.T) {
			result := ScoreRoundPart2(scenario.input)
			assert.Equal(t, scenario.expected, result)
		})
	}
}
