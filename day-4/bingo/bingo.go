package bingo

type Bingo [][]int

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

	for _, row := range board {
		for _, cell := range row {
			isValue := false
			for _, number := range numbers {
				if number == cell {
					isValue = true
				}
			}
			if !isValue {
				valuesToCount = append(valuesToCount, cell)
			}
		}
	}

	for _, value := range valuesToCount {
		score += value
	}

	score *= numbers[len(numbers)-1]

	return score
}
