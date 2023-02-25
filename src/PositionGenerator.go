package main

import "math/rand"

type PositionGenerator struct {
	maxCoord int
	rng      *rand.Rand
}

func NewPositionGenerator(maxCoord int) *PositionGenerator {
	var rng = NewUnixTimeRng()
	return &PositionGenerator{maxCoord, rng}
}

func (generator *PositionGenerator) Make() Position {
	var x = generator.rng.Intn(generator.maxCoord)
	var y = generator.rng.Intn(generator.maxCoord)
	return NewPosition(x, y)
}
