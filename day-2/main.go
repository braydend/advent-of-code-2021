package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Down
	Forward
)

type Position struct {
	horizontal uint
	depth      uint
}

type Instruction struct {
	direction Direction
	amount    uint
}

func main() {
	instructions, err := parseDirectionsFromFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	position := Navigate(instructions)

	fmt.Printf("Sub is at %d depth and %d horizontal", position.depth, position.horizontal)
}

func parseDirectionsFromFile(filename string) (instructions []Instruction, err error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		splits := strings.Split(line, " ")
		var direction Direction
		directionString := splits[0]
		amount, err := strconv.Atoi(splits[1])

		if err != nil {
			return nil, err
		}

		switch directionString {
		case "up":
			direction = Up
			break

		case "down":
			direction = Down
			break

		case "forward":
			direction = Forward
			break

		default:
			return nil, fmt.Errorf("Unknown direction: %s", directionString)
		}

		instructions = append(instructions, Instruction{direction, uint(amount)})
	}

	return instructions, nil
}

func Navigate(instructions []Instruction) (position Position) {
	var aim uint
	for _, instruction := range instructions {
		switch instruction.direction {
		case Up:
			aim -= instruction.amount
			break

		case Down:
			aim += instruction.amount
			break

		case Forward:
			position = Position{position.horizontal + instruction.amount, position.depth + (aim * instruction.amount)}
			break
		}
	}

	return position
}
