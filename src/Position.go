package main

type Position struct {
	x int
	y int
}

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

func (origin *Position) Add(x int, y int) Position {
	return Position{origin.x + x, origin.y + y}
}
