package main

import (
	"math/rand"
	"time"
)

type PositionGenerator struct {
	maxCoord  int
	generator *rand.Rand
}

func NewPositionGenerator(maxCoord int) PositionGenerator {
	var seed = time.Now().UnixNano()
	var source = rand.NewSource(seed)
	var generator = rand.New(source)
	return PositionGenerator{maxCoord, generator}
}

func (generator *PositionGenerator) Make() Position {
	var x = generator.generator.Intn(generator.maxCoord)
	var y = generator.generator.Intn(generator.maxCoord)
	return NewPosition(x, y)
}
