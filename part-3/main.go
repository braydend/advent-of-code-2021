package main

import "math"

const ReadingLength = 5

func main() {

}

func CalculateGammaRate(input [][ReadingLength]bool) uint {
	return binaryBoolSliceToUint(getMajorityBits(input))
}

func CalculateEpsilonRate(input [][ReadingLength]bool) uint {
	return binaryBoolSliceToUint(invertBinarySlice(getMajorityBits(input)))
}

func getMajorityBits(readings [][ReadingLength]bool) (output [ReadingLength]bool) {

	for i := 0; i < ReadingLength; i++ {
		var trueCount uint

		for _, reading := range readings {
			if reading[i] {
				trueCount++
				continue
			}
		}

		output[i] = trueCount > uint(len(readings)/2)
	}

	return output
}

func binaryBoolSliceToUint(slice [ReadingLength]bool) (output uint) {
	for i, bit := range slice {
		exponent := len(slice) - 1 - i
		if bit {
			output += uint(math.Pow(2, float64(exponent)))
		}
	}

	return output
}

func invertBinarySlice(input [ReadingLength]bool) (output [ReadingLength]bool) {
	for i, value := range input {
		output[i] = !value
	}

	return output
}
