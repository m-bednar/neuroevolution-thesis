package main

const MIN_COORD = 0

type Position struct {
	x int
	y int
}

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (position *Position) Move(x int, y int) {
	position.x = Max(position.x+x, MIN_COORD)
	position.y = Max(position.y+y, MIN_COORD)
}
