package vent

import "fmt"

type Vent struct {
	StartingPosition Coordinates
	EndingPosition   Coordinates
}

type Coordinates struct {
	X int
	Y int
}

func NewVent(x1, y1, x2, y2 int) Vent {
	return Vent{Coordinates{x1, y1}, Coordinates{x2, y2}}
}

func (v Vent) findHighestXAndY() (highest Coordinates) {
	if v.StartingPosition.X >= v.EndingPosition.X {
		highest.X = v.StartingPosition.X
	} else {
		highest.X = v.EndingPosition.X
	}

	if v.StartingPosition.Y >= v.EndingPosition.Y {
		highest.Y = v.StartingPosition.Y
	} else {
		highest.Y = v.EndingPosition.Y
	}

	return highest
}

func (v Vent) findLowestXAndY() (lowest Coordinates) {
	if v.StartingPosition.X >= v.EndingPosition.X {
		lowest.X = v.EndingPosition.X
	} else {
		lowest.X = v.StartingPosition.X
	}

	if v.StartingPosition.Y >= v.EndingPosition.Y {
		lowest.Y = v.EndingPosition.Y
	} else {
		lowest.Y = v.StartingPosition.Y
	}

	return lowest
}

func (v Vent) GetCoveredCoordinates() (coveredCoordinates []Coordinates) {
	highest := v.findHighestXAndY()
	lowest := v.findLowestXAndY()

	for x := lowest.X; x <= highest.X; x++ {
		for y := lowest.Y; y <= highest.Y; y++ {
			coveredCoordinates = append(coveredCoordinates, Coordinates{x, y})
		}
	}

	return coveredCoordinates
}

func FindOverlappingVents(vents []Vent, onlyHorizontal bool) (ventLocations map[Coordinates]int) {
	ventLocations = make(map[Coordinates]int)
	for _, vent := range vents {
		if onlyHorizontal && !vent.IsHorizontal() {
			continue
		}
		coveredCoordinates := vent.GetCoveredCoordinates()

		for _, coordinate := range coveredCoordinates {
			count, exists := ventLocations[coordinate]

			if exists {
				ventLocations[coordinate] = count + 1
				continue
			}

			ventLocations[coordinate] = 1
		}
	}

	return ventLocations
}

func (v Vent) IsHorizontal() bool {
	return v.StartingPosition.X == v.EndingPosition.X || v.StartingPosition.Y == v.EndingPosition.Y
}

func CountOverlappingVents(ventLocations map[Coordinates]int, threshold int) (overlapCount int) {
	for _, count := range ventLocations {
		if count >= threshold {
			overlapCount++
		}
	}

	return overlapCount
}

func FindHighestXAndYFromCoordinates(ventCoordinates map[Coordinates]int) (highest Coordinates) {
	for coordinates, _ := range ventCoordinates {
		if coordinates.X > highest.X {
			highest.X = coordinates.X
		}

		if coordinates.Y > highest.Y {
			highest.Y = coordinates.Y
		}
	}

	return highest
}

func BuildVentLocationMap(ventLocations map[Coordinates]int) (output string) {
	highestCoordinates := FindHighestXAndYFromCoordinates(ventLocations)
	for y := 0; y < highestCoordinates.Y; y++ {
		fmt.Printf("%s\n", output)
		for x := 0; x < highestCoordinates.X; x++ {
			if value, exists := ventLocations[Coordinates{x, y}]; exists {
				fmt.Printf("%s%d", output, value)
				continue
			}
			fmt.Printf("%s%s", output, ".")
		}
	}
	return output
}
