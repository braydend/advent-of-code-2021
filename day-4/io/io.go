package io

import (
	"github.com/braydend/advent-of-code/day-4/bingo"
	"strconv"
	"strings"
)

const (
	numbersLineNumber     = 0
	boardsLineStartNumber = 2
	bingoSize             = 5
)

func ParseNumbers(input []byte) (numbers []int, err error) {
	lineSplits := strings.Split(string(input), "\n")
	numberLine := lineSplits[numbersLineNumber]

	return parseNumbersFromLine(numberLine, ",")
}

func ParseBoards(input []byte) (boards []bingo.Bingo, err error) {
	linesWithNumbers, err := parseRowsOfNumbersFromInput(string(input))

	if err != nil {
		return nil, err
	}

	board := newBingo()

	for i, line := range linesWithNumbers {
		numberList, err := parseNumbersFromLine(line, " ")
		if err != nil {
			return nil, err
		}

		board[i%5] = numberList

		if (i+1)%bingoSize == 0 {
			boards = append(boards, board)
			board = newBingo()
		}
	}

	return boards, nil
}

func parseRowsOfNumbersFromInput(input string) ([]string, error) {
	lineSplits := strings.Split(input, "\n")
	lineSplits = lineSplits[boardsLineStartNumber:]
	var linesWithNumbers []string

	for _, split := range lineSplits {
		if split != "" {
			linesWithNumbers = append(linesWithNumbers, split)
		}
	}

	return linesWithNumbers, nil
}

func parseNumbersFromLine(line string, delimiter string) ([]int, error) {
	var numberList []int
	numbers := strings.Split(line, delimiter)

	for _, number := range numbers {
		if number == "" {
			continue
		}
		value, err := strconv.Atoi(number)

		if err != nil {
			return nil, err
		}

		numberList = append(numberList, value)
	}

	return numberList, nil
}

func newBingo() bingo.Bingo {
	return bingo.Bingo{[]int{}, []int{}, []int{}, []int{}, []int{}}
}
