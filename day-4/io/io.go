package io

import (
	"github.com/braydend/advent-of-code/day-4/bingo"
	"strconv"
	"strings"
)

const (
	NumbersLineNumber     = 0
	BoardsLineStartNumber = 2
)

func ParseNumbers(input []byte) (numbers []int, err error) {
	lineSplits := strings.Split(string(input), "\n")

	numberLine := lineSplits[NumbersLineNumber]

	for _, number := range strings.Split(numberLine, ",") {
		n, err := strconv.Atoi(number)

		if err != nil {
			return nil, err
		}

		numbers = append(numbers, n)
	}

	return numbers, nil
}

func ParseBoards(input []byte) (boards []bingo.Bingo, err error) {
	lineSplits := strings.Split(string(input), "\n")
	lineSplits = lineSplits[BoardsLineStartNumber:]
	var linesWithNumbers []string

	for _, split := range lineSplits {
		if split != "" {
			linesWithNumbers = append(linesWithNumbers, split)
		}
	}

	var rows [][]int
	for i := 0; i < 5; i++ {
		rows = append(rows, []int{})
	}
	board := rows

	for i, line := range linesWithNumbers {
		if i > 1 && i%5 == 0 {
			var rows [][]int
			for i := 0; i < 5; i++ {
				rows = append(rows, []int{})
			}
			board = rows
		}

		var numberList []int
		numbers := strings.Split(line, " ")

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

		board[i%5] = numberList

		if err != nil {
			return nil, err
		}

		if (i+1)%5 == 0 {
			boards = append(boards, board)
		}
	}

	return boards, nil
}
