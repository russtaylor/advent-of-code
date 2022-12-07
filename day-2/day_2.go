package day2

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func Part1() {
	score := ScoreRPSFile("./day-2/full-data.txt")
	fmt.Printf("Full contestant score for rock/paper/scissors tournament: %d\n", score)
}

func Part2() {
	fmt.Println("TODO")
}

const opponentRock = "A"
const opponentPaper = "B"
const opponentScissors = "C"

const contestantRock = "X"
const contestantPaper = "Y"
const contestantScissors = "Z"

const contestantRockPoints = 1
const contestantPaperPoints = 2
const contestantScissorsPoints = 3

const drawPoints = 3
const winPoints = 6
const lossPoints = 0

func ScoreRPSFile(path string) int {
	rounds := utils.ReadFileToSlices(path)
	totalScore := 0
	for _, round := range rounds {
		totalScore += ScoreRPSRound(round)
	}
	return totalScore
}

func ScoreRPSRound(round string) int {
	splitRound := strings.Split(round, " ")
	opponentPlay := splitRound[0]
	contestantPlay := splitRound[1]
	roundWinScore := PlayRPSRound(opponentPlay, contestantPlay)

	var roundPlayScore int
	switch contestantPlay {
	case contestantRock:
		roundPlayScore = contestantRockPoints
	case contestantPaper:
		roundPlayScore = contestantPaperPoints
	case contestantScissors:
		roundPlayScore = contestantScissorsPoints
	}
	return roundWinScore + roundPlayScore
}

func PlayRPSRound(opponentPlay string, contestantPlay string) int {
	switch opponentPlay {
	case opponentRock:
		switch contestantPlay {
		case contestantRock:
			return drawPoints
		case contestantPaper:
			return winPoints
		case contestantScissors:
			return lossPoints
		}
	case opponentPaper:
		switch contestantPlay {
		case contestantRock:
			return lossPoints
		case contestantPaper:
			return drawPoints
		case contestantScissors:
			return winPoints
		}
	case opponentScissors:
		switch contestantPlay {
		case contestantRock:
			return winPoints
		case contestantPaper:
			return lossPoints
		case contestantScissors:
			return drawPoints
		}
	}
	panic("Invalid play")
}
