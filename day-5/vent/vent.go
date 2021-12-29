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

func (v Vent) IsHorizontal() bool {
	return v.startingPosition.x == v.endingPosition.x || v.startingPosition.y == v.endingPosition.y
}
