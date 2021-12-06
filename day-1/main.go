package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	depths, err := parseDepthsFromFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(depthsToString(depths))
	fmt.Printf("%d depth increases\n", countDepthIncreases(depths))
}

func parseDepthsFromFile(filename string) (depths []int, err error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return depths, err
	}

	return parseDepths(string(data))
}

func parseDepths(input string) (depths []int, err error) {
	splits := strings.Split(input, "\n")

	for _, split := range splits {
		if split == "" {
			continue
		}

		value, err := strconv.Atoi(split)

		if value < 0 {
			return nil, fmt.Errorf("cannot have a negative depth")
		}

		if err != nil {
			return nil, err
		}

		depths = append(depths, value)
	}

	return depths, nil
}

func depthsToString(depths []int) (output string) {
	for i, depth := range depths {
		var message string
		if i == 0 {
			message = "N/A - no previous measurement"
		} else if depth > depths[i-1] {
			message = "increased"
		} else {
			message = "decreased"
		}

		output = fmt.Sprintf("%s%d (%s)\n", output, depth, message)
	}

	return output
}

func countDepthIncreases(depths []int) (increases int) {
	for i, depth := range depths {
		if i == 0 {
			continue
		}
		if depth > depths[i-1] {
			increases++
		}
	}

	return increases
}
