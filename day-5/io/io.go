package io

import (
	"github.com/braydend/advent-of-code/day-5/vent"
	"os"
	"strconv"
	"strings"
)

func ParseVentsFromFile(filename string) (vents []vent.Vent, err error) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return parseVents(bytes)
}

func parseVents(data []byte) (vents []vent.Vent, err error) {
	lines := parseLinesFromBytes(data)
	for _, line := range lines {
		if line == "" {
			continue
		}
		ventFromLine, err := parseVentFromLine(line)

		if err != nil {
			return nil, err
		}

		vents = append(vents, ventFromLine)
	}

	return vents, nil
}

func parseLinesFromBytes(data []byte) (lines []string) {
	return strings.Split(string(data), "\n")
}

func parseVentFromLine(line string) (vent.Vent, error) {
	coordinates := strings.Split(line, "->")
	startingPosition, err := parseCoordinatesFromString(coordinates[0])

	if err != nil {
		return vent.Vent{}, err
	}

	endingPosition, err := parseCoordinatesFromString(coordinates[1])

	if err != nil {
		return vent.Vent{}, err
	}

	return vent.Vent{startingPosition, endingPosition}, nil
}

func parseCoordinatesFromString(string string) (vent.Coordinates, error) {
	coords := strings.Split(strings.ReplaceAll(string, " ", ""), ",")

	x, err := strconv.Atoi(coords[0])

	if err != nil {
		return vent.Coordinates{}, err
	}

	y, err := strconv.Atoi(coords[1])

	if err != nil {
		return vent.Coordinates{}, err
	}

	return vent.Coordinates{x, y}, nil
}
