package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	depthReport, err := ReportDepths("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(depthReport)
}

func ReportDepths(filename string) (output string, err error) {
	depthsFromFile, err := parseDepthsFromFile(filename)

	if err != nil {
		return "", err
	}
	windows := createDepthWindows(depthsFromFile)
	depths := sumDepthWindows(windows)

	output = fmt.Sprintf("%s%s\n", output, depthsToString(depths))
	output = fmt.Sprintf("%s%d depth increases\n", output, countDepthIncreases(depths))

	return output, nil
}

func parseDepthsFromFile(filename string) (depths []int, err error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return depths, err
	}

	return parseDepths(string(data))
}

func createDepthWindows(depths []int) (windows [][3]int) {
	windowCount := len(depths) - 2
	windows = make([][3]int, windowCount)

	for i := 0; i < windowCount; i++ {
		for j := range depths {
			if j < 3 {
				windows[i][j] = depths[i+j]
			}
		}

	}

	return windows
}

func sumDepthWindows(windows [][3]int) (sums []int) {
	for _, window := range windows {
		sums = append(sums, window[0]+window[1]+window[2])
	}

	return sums
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

func getAsciiLetterForIndex(index int) (output string) {
	for i := index; i >= 0; i = i - 26 {
		//if i > 26 {
		//	letter := rune(int(math.Floor(float64(i/26))) + 65)
		//	output = fmt.Sprintf("%s%c", output, letter)
		//} else {
		letter := rune(int(i%26) + 65)
		output = fmt.Sprintf("%s%c", output, letter)
		//}
	}

	return output
}

func depthsToString(depths []int) (output string) {
	for i, depth := range depths {
		var message string

		if i == 0 {
			message = "N/A - no previous sum"
		} else if depth == depths[i-1] {
			message = "no change"
		} else if depth > depths[i-1] {
			message = "increased"
		} else {
			message = "decreased"
		}

		output = fmt.Sprintf("%s%s: %d (%s)\n", output, getAsciiLetterForIndex(i), depth, message)
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
