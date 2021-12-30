package vent

type Vent struct {
	startingPosition coordinates
	endingPosition   coordinates
}

type coordinates struct {
	x int
	y int
}

func NewVent(x1, y1, x2, y2 int) Vent {
	return Vent{coordinates{x1, y1}, coordinates{x2, y2}}
}

func (v Vent) findHighestXAndY() (highest coordinates) {
	if v.startingPosition.x >= v.endingPosition.x {
		highest.x = v.startingPosition.x
	} else {
		highest.x = v.endingPosition.x
	}

	if v.startingPosition.y >= v.endingPosition.y {
		highest.y = v.startingPosition.y
	} else {
		highest.y = v.endingPosition.y
	}

	return highest
}

func (v Vent) findLowestXAndY() (lowest coordinates) {
	if v.startingPosition.x >= v.endingPosition.x {
		lowest.x = v.endingPosition.x
	} else {
		lowest.x = v.startingPosition.x
	}

	if v.startingPosition.y >= v.endingPosition.y {
		lowest.y = v.endingPosition.y
	} else {
		lowest.y = v.startingPosition.y
	}

	return lowest
}

func (v Vent) GetCoveredCoordinates() (coveredCoordinates []coordinates) {
	highest := v.findHighestXAndY()
	lowest := v.findLowestXAndY()

	for x := lowest.x; x <= highest.x; x++ {
		for y := lowest.y; y <= highest.y; y++ {
			coveredCoordinates = append(coveredCoordinates, coordinates{x, y})
		}
	}

	return coveredCoordinates
}

func FindOverlappingVents(vents []Vent) (ventLocations map[coordinates]int) {
	ventLocations = make(map[coordinates]int)
	for _, vent := range vents {
		if !vent.IsHorizontal() {
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
	return v.startingPosition.x == v.endingPosition.x || v.startingPosition.y == v.endingPosition.y
}
