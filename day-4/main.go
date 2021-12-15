package main

import (
	"fmt"
	"github.com/braydend/advent-of-code/day-4/bingo"
	"github.com/braydend/advent-of-code/day-4/io"
	"log"
	"os"
)

func main() {
	game, err := parseFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	numbersToCall := game.numbers
	var calledNumbers []int

	for _, number := range numbersToCall {
		calledNumbers = append(calledNumbers, number)
		for j, board := range game.boards {
			if bingo.CheckBoard(board, calledNumbers) {
				fmt.Printf("Player %d is the winner!\nTheir final score is: %d", j+1, bingo.CalculateScore(board, calledNumbers))
				return
			}
		}
	}

	fmt.Printf("Everyone loses!\n")
}

func parseFile(filename string) (output struct {
	numbers []int
	boards  []bingo.Bingo
}, err error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return output, err
	}

	output.boards, err = io.ParseBoards(data)

	if err != nil {
		return output, err
	}

	output.numbers, err = io.ParseNumbers(data)

	if err != nil {
		return output, err
	}

	return output, nil
}
