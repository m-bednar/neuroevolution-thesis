package main

type Position struct {
	x uint
	y uint
}

func NewPosition(x uint, y uint) Position {
	return Position{x, y}
}
