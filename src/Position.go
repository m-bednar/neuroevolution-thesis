package main

import "fmt"

const MIN_COORD = 0

type Position struct {
	x int
	y int
}

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

func (position *Position) Move(x int, y int) {
	position.x += x
	position.y += y
	fmt.Println(position)
}
