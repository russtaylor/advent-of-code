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
	score := ScoreRPSFilePart2("./day-2/full-data.txt")
	fmt.Printf("Full score for rock/paper/scissors tournament, part 2: %d\n", score)
}

const elfRock = "A"
const elfPaper = "B"
const elfScissors = "C"

const myRock = "X"
const myPaper = "Y"
const myScissors = "Z"

const myRockPoints = 1
const myPaperPoints = 2
const myScissorsPoints = 3

const desireLoss = "X"
const desireDraw = "Y"
const desireWin = "Z"

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

func ScoreRPSFilePart2(path string) int {
	rounds := utils.ReadFileToSlices(path)
	totalScore := 0
	for _, round := range rounds {
		totalScore += ScoreRoundPart2(round)
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
	case myRock:
		roundPlayScore = myRockPoints
	case myPaper:
		roundPlayScore = myPaperPoints
	case myScissors:
		roundPlayScore = myScissorsPoints
	}
	return roundWinScore + roundPlayScore
}

// I was gonna use this to be clever, but it's taking too long
func playToInt(play string) int {
	playRune := int(rune(play[0]))
	if playRune >= 65 && playRune <= 67 {
		return playRune - 65
	}
	if playRune >= 88 && playRune <= 90 {
		return playRune - 88
	}
	panic("Invalid play")
}

func PlayRPSRound(elfPlay string, myPlay string) int {
	switch elfPlay {
	case elfRock:
		switch myPlay {
		case myRock:
			return drawPoints
		case myPaper:
			return winPoints
		case myScissors:
			return lossPoints
		}
	case elfPaper:
		switch myPlay {
		case myRock:
			return lossPoints
		case myPaper:
			return drawPoints
		case myScissors:
			return winPoints
		}
	case elfScissors:
		switch myPlay {
		case myRock:
			return winPoints
		case myPaper:
			return lossPoints
		case myScissors:
			return drawPoints
		}
	}
	panic("Invalid play")
}

// Just gonna hardcode the options rather than being clever
func ScoreRoundPart2(round string) int {
	splitRound := strings.Split(round, " ")
	elfPlay := splitRound[0]
	desiredResult := splitRound[1]

	switch elfPlay {
	case elfRock:
		switch desiredResult {
		case desireLoss:
			return lossPoints + myScissorsPoints
		case desireDraw:
			return drawPoints + myRockPoints
		case desireWin:
			return winPoints + myPaperPoints
		}
	case elfPaper:
		switch desiredResult {
		case desireLoss:
			return lossPoints + myRockPoints
		case desireDraw:
			return drawPoints + myPaperPoints
		case desireWin:
			return winPoints + myScissorsPoints
		}
	case elfScissors:
		switch desiredResult {
		case desireLoss:
			return lossPoints + myPaperPoints
		case desireDraw:
			return drawPoints + myScissorsPoints
		case desireWin:
			return winPoints + myRockPoints
		}
	}
	panic("Invalid arguments")
}
