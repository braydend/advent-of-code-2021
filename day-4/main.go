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

	winner, playerNumber, calledNumbers := bingo.FindWinningBoard(game.boards, game.numbers)

	fmt.Printf("Player %d is the winner!\nTheir final score is: %d", playerNumber, bingo.CalculateScore(winner, calledNumbers))

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
