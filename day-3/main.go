package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const AsciiOne = 49

func main() {
	readings, err := parseReadingsFromFile("input.txt")

	if err != nil {
		log.Fatalln(err)
	}

	gamma := CalculateGammaRate(readings)
	epsilon := CalculateEpsilonRate(readings)

	fmt.Printf("gamma rate: %d\nepsilon rate: %d\n", gamma, epsilon)
	fmt.Printf("Total power consumption is: %d\n", gamma*epsilon)
}

func parseReadingsFromFile(filename string) (output [][]bool, err error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		var reading []bool

		for _, char := range line {
			reading = append(reading, char == AsciiOne)
		}
		output = append(output, reading)
	}

	return output, nil
}

func CalculateGammaRate(input [][]bool) uint {
	return binaryBoolSliceToUint(getMajorityBits(input))
}

func CalculateEpsilonRate(input [][]bool) uint {
	return binaryBoolSliceToUint(invertBinarySlice(getMajorityBits(input)))
}

func CalculateOxygenGeneratorRating(input [][]bool) uint {
	relevantReadings := input

	for i := 0; i < len(relevantReadings[0]); i++ {
		majorityBits := getMajorityBits(relevantReadings)
		if isLastReading := len(relevantReadings) == 1; isLastReading {
			return binaryBoolSliceToUint(getMajorityBits(relevantReadings))
		}

		relevantReadings = getReadingsThatMatchBitAtPosition(relevantReadings, majorityBits[i], i)
	}

	return binaryBoolSliceToUint(getMajorityBits(relevantReadings))
}

func CalculateCo2ScrubberRating(input [][]bool) uint {
	relevantReadings := input

	for i := 0; i < len(relevantReadings[0]); i++ {
		minorityBits := invertBinarySlice(getMajorityBits(relevantReadings))
		if isLastReading := len(relevantReadings) == 1; isLastReading {
			return binaryBoolSliceToUint(getMajorityBits(relevantReadings))
		}

		relevantReadings = getReadingsThatMatchBitAtPosition(relevantReadings, minorityBits[i], i)
	}

	return binaryBoolSliceToUint(getMajorityBits(relevantReadings))
}

func getReadingsThatMatchBitAtPosition(input [][]bool, bit bool, position int) (output [][]bool) {
	for _, reading := range input {
		if reading[position] == bit {
			output = append(output, reading)
		}
	}

	return output
}

func getMajorityBits(readings [][]bool) (output []bool) {

	for i := 0; i < len(readings[0]); i++ {
		var trueCount uint

		for _, reading := range readings {
			if reading[i] {
				trueCount++
				continue
			}
		}

		halfLength := uint(math.Ceil(float64(len(readings)) / float64(2)))

		if trueCount == halfLength {
			output = append(output, true)
			continue
		}

		output = append(output, trueCount > uint(len(readings)/2))
	}

	return output
}

func binaryBoolSliceToUint(slice []bool) (output uint) {
	for i, bit := range slice {
		exponent := len(slice) - 1 - i
		if bit {
			output += uint(math.Pow(2, float64(exponent)))
		}
	}

	return output
}

func invertBinarySlice(input []bool) (output []bool) {
	for _, value := range input {
		output = append(output, !value)
	}

	return output
}
