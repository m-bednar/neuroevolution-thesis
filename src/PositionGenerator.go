package main

import (
	"math/rand"
	"time"
)

type PositionGenerator struct {
	maxX int
	maxY int
	generator rand.Rand
}

func NewPositionGenerator(maxX int, maxY int) PositionGenerator {
	var seed = time.Now().UnixNano()
	var source = rand.NewSource(seed)
	var generator = rand.New(source)
	return PositionGenerator { maxX, maxY, *generator }
}

func (generator *PositionGenerator) Make() Position {
	var x = generator.generator.Intn(generator.maxX)
	var y = generator.generator.Intn(generator.maxY)
	return NewPosition(uint(x), uint(y))
}
