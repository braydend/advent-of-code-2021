package main

import "math"

func main() {

}

func CalculateGammaRate(input [][5]bool) (result uint) {
	var bits [5]bool

	for i := 0; i < 4; i++ {
		var trueCount uint
		var falseCount uint

		for _, reading := range input {
			if reading[i] {
				trueCount++
				continue
			}
			falseCount++
		}

		if trueCount > falseCount {
			bits[i] = true
			continue
		}
		bits[i] = false
	}

	for i, bit := range bits {
		if bit {
			result += uint(math.Pow(2, float64(len(bits)-1-i)))
		}
	}

	return result
}
