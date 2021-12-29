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

func (v Vent) GetCoveredCoordinates() (coveredPositions []position) {
	var highestX, highestY, lowestX, lowestY int

	if v.startingPosition.x >= v.endingPosition.x {
		highestX = v.startingPosition.x
		lowestX = v.endingPosition.x
	} else {
		highestX = v.endingPosition.x
		lowestX = v.startingPosition.x
	}

	if v.startingPosition.y >= v.endingPosition.y {
		highestY = v.startingPosition.y
		lowestY = v.endingPosition.y
	} else {
		highestY = v.endingPosition.y
		lowestY = v.startingPosition.y
	}

	for x := lowestX; x <= highestX; x++ {
		for y := lowestY; y <= highestY; y++ {
			coveredPositions = append(coveredPositions, position{x, y})
		}
	}

	return coveredPositions
}

func (v Vent) IsHorizontal() bool {
	return v.startingPosition.x == v.endingPosition.x || v.startingPosition.y == v.endingPosition.y
}
