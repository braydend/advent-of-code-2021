package bingo

import "reflect"

type Bingo [][]int

func FindLastWinningBoard(boards []Bingo, numbers []int) (board Bingo, calledNumbers []int) {
	remainingBoards := boards

	for i, number := range numbers {
		calledNumbers = append(calledNumbers, number)
		boards, _, _ := FindWinningBoards(remainingBoards, numbers[:i])

		if len(boards) > 0 {
			for _, bingo := range boards {
				remainingBoards = removeBoardFromSlice(bingo, remainingBoards)
			}
		}

		if isLastBoard := len(remainingBoards) == 1; isLastBoard {
			return remainingBoards[0], calledNumbers
		}
	}

	return board, calledNumbers
}

func removeBoardFromSlice(board Bingo, boards []Bingo) []Bingo {
	for i, bingo := range boards {
		if reflect.DeepEqual(bingo, board) {
			return append(boards[:i], boards[i+1:]...)
		}
	}

	return boards
}

func FindWinningBoards(boards []Bingo, numbers []int) (winners []Bingo, playerNumbers []int, calledNumbers []int) {
	for _, number := range numbers {
		calledNumbers = append(calledNumbers, number)
		for j, board := range boards {
			if CheckBoard(board, calledNumbers) {
				winners = append(winners, board)
				playerNumbers = append(playerNumbers, j+1)
			}
		}

		if len(winners) > 0 {
			return winners, playerNumbers, calledNumbers
		}
	}

	return winners, playerNumbers, calledNumbers
}

func CheckBoard(board Bingo, calledNumbers []int) (isComplete bool) {
	for i, row := range board {
		if checkArray(row, calledNumbers) {
			return true
		}
		var column []int
		for j := 0; j < len(row); j++ {
			column = append(column, board[j][i])
		}
		if checkArray(column, calledNumbers) {
			return true
		}
	}
	return false
}

func checkArray(bingoSlice []int, calledNumbers []int) bool {
	var foundCount int

	for _, calledNumber := range calledNumbers {
		if doesValueExistInSlice(calledNumber, bingoSlice) {
			foundCount++
		}
	}

	return foundCount == len(bingoSlice)
}

func CalculateScore(board Bingo, numbers []int) (score int) {
	var valuesToCount []int

	checkValue := func(cell int) {
		if !doesValueExistInSlice(cell, numbers) {
			valuesToCount = append(valuesToCount, cell)
		}
	}

	board.forEachCell(checkValue)

	for _, value := range valuesToCount {
		score += value
	}

	score *= numbers[len(numbers)-1]

	return score
}

func (board Bingo) forEachCell(fn func(cell int)) {
	for _, row := range board {
		for _, cell := range row {
			fn(cell)
		}
	}
}

func doesValueExistInSlice(value int, slice []int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
