package enviroment

import "math"

type Position struct {
	x int
	y int
}

type Direction Position

var (
	Origin = Direction(Position{0, 0})
	North  = Direction(Position{0, -1})
	South  = Direction(Position{0, 1})
	West   = Direction(Position{-1, 0})
	East   = Direction(Position{1, 0})
)

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

func (position Position) Add(x int, y int) Position {
	return Position{position.x + x, position.y + y}
}

func (position Position) AddToDirection(direction Direction) Position {
	return position.Add(direction.x, direction.y)
}

func (origin Position) DistanceTo(position Position) float64 {
	var x1, y1 = float64(origin.x), float64(origin.y)
	var x2, y2 = float64(position.x), float64(position.y)
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt((dx * dx) + (dy * dy))
}

func (position Position) GetX() int {
	return position.x
}

func (position Position) GetY() int {
	return position.y
}
