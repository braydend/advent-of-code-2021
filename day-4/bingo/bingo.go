package bingo

type Bingo [][]int

func FindWinningBoard(boards []Bingo, numbers []int) (winner Bingo, playerNumber int, numbersCalled []int) {
	var calledNumbers []int

	for _, number := range numbers {
		calledNumbers = append(calledNumbers, number)
		for j, board := range boards {
			if CheckBoard(board, calledNumbers) {
				return board, j + 1, calledNumbers
			}
		}
	}

	return nil, 0, nil
}

func CheckBoard(board Bingo, numbers []int) (isComplete bool) {
	for i, row := range board {
		if checkArray(row, numbers) {
			return true
		}
		var column []int
		for j := 0; j < len(row); j++ {
			column = append(column, board[j][i])
		}
		return checkArray(column, numbers)
	}
	return false
}

func checkArray(input []int, numbers []int) bool {
	var foundCount int

	for _, i := range input {
		for _, n := range numbers {
			if i == n {
				foundCount++
			}
		}
	}

	return foundCount == len(input)
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
