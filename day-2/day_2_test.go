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
		input string
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

// Test scenarios in which the contestant wins
func TestPlayRPSRoundWinScenarios(t *testing.T) {
	scenarios := [][]string{
		{"A", "Y"},
		{"B", "Z"},
		{"C", "X"},
	}

	expected := 6 // A 'win' counts for 6 points
	for _, scenario := range scenarios {
		result := PlayRPSRound(scenario[0], scenario[1])
		assert.Equal(t, expected, result)
	}
}

// Test scenarios which result in a draw
func TestPlayRPSRoundDrawScenarios(t *testing.T) {
	scenarios := [][]string{
		{"A", "X"},
		{"B", "Y"},
		{"C", "Z"},
	}

	expected := 3 // A 'draw' counts for 3 points
	for _, scenario := range scenarios {
		result := PlayRPSRound(scenario[0], scenario[1])
		assert.Equal(t, expected, result)
	}
}

// Test scenarios which result in a loss
func TestPlayRPSRoundLossScenarios(t *testing.T) {
	scenarios := [][]string{
		{"A", "Z"},
		{"B", "X"},
		{"C", "Y"},
	}

	expected := 0 // A 'loss' counts for 3 points
	for _, scenario := range scenarios {
		result := PlayRPSRound(scenario[0], scenario[1])
		assert.Equal(t, expected, result)
	}
}