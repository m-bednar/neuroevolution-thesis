package main

import (
	"math/rand"
	"time"
)

type PositionGenerator struct {
	maxCoord int
	rng      *rand.Rand
}

func NewPositionGenerator(maxCoord int) PositionGenerator {
	var seed = time.Now().UnixNano()
	var source = rand.NewSource(seed)
	var rng = rand.New(source)
	return PositionGenerator{maxCoord, rng}
}

func (generator *PositionGenerator) Make() Position {
	var x = generator.rng.Intn(generator.maxCoord)
	var y = generator.rng.Intn(generator.maxCoord)
	return NewPosition(x, y)
}
