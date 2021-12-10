package main

import "math"

const ReadingLength = 5

func main() {

}

func CalculateGammaRate(input [][ReadingLength]bool) (result uint) {
	var majorityBits [ReadingLength]bool

	for i := 0; i < ReadingLength; i++ {
		var trueCount uint

		for _, reading := range input {
			if reading[i] {
				trueCount++
				continue
			}
		}

		majorityBits[i] = trueCount > uint(len(input)/2)
	}

	for i, bit := range majorityBits {
		exponent := len(majorityBits) - 1 - i
		if bit {
			result += uint(math.Pow(2, float64(exponent)))
		}
	}

	return result
}

func CalculateEpsilonRate(input [][ReadingLength]bool) (result uint) {
	var minorityBits [ReadingLength]bool

	for i := 0; i < ReadingLength; i++ {
		var trueCount uint

		for _, reading := range input {
			if reading[i] {
				trueCount++
				continue
			}
		}

		minorityBits[i] = trueCount < uint(len(input)/2)
	}

	for i, bit := range minorityBits {
		exponent := len(minorityBits) - 1 - i
		if bit {
			result += uint(math.Pow(2, float64(exponent)))
		}
	}

	return result
}
