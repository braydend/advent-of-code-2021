package vent

type Vent struct {
	startingPosition position
	endingPosition   position
}

type position struct {
	x int
	y int
}

func NewVent(x1, y1, x2, y2 int) Vent {
	return Vent{position{x1, y1}, position{x2, y2}}
}

func (v Vent) findHighestXAndY() (highest position) {
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

func (v Vent) findLowestXAndY() (lowest position) {
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

func (v Vent) GetCoveredCoordinates() (coveredPositions []position) {
	highest := v.findHighestXAndY()
	lowest := v.findLowestXAndY()

	for x := lowest.x; x <= highest.x; x++ {
		for y := lowest.y; y <= highest.y; y++ {
			coveredPositions = append(coveredPositions, position{x, y})
		}
	}

	return coveredPositions
}

func (v Vent) IsHorizontal() bool {
	return v.startingPosition.x == v.endingPosition.x || v.startingPosition.y == v.endingPosition.y
}
