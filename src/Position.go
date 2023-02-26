package main

import "math"

type Position struct {
	x int
	y int
}

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

func (position *Position) Add(x int, y int) Position {
	return Position{position.x + x, position.y + y}
}

func (origin *Position) DistanceTo(position Position) float64 {
	var x1, y1 = float64(origin.x), float64(origin.y)
	var x2, y2 = float64(position.x), float64(position.y)
	var dx = x2 - x1
	var dy = y2 - y1
	return math.Sqrt((dx * dx) + (dy * dy))
}
