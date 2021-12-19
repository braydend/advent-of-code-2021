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

	winners, playerNumbers, calledNumbers := bingo.FindWinningBoards(game.boards, game.numbers)

	if len(winners) > 0 {
		for i, winner := range winners {
			fmt.Printf("Player %d is the winner!\nTheir final score is: %d\n", playerNumbers[i], bingo.CalculateScore(winner, calledNumbers))
		}

		lastWinner, calledNumbers := bingo.FindLastWinningBoard(game.boards, game.numbers)

		fmt.Printf("The last winner's final score is: %d \n", bingo.CalculateScore(lastWinner, calledNumbers))
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
