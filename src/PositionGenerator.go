package main

import "math/rand"

type PositionGenerator struct {
	enviroment *Enviroment
	rng        *rand.Rand
}

func NewPositionGenerator(enviroment *Enviroment) *PositionGenerator {
	var rng = NewUnixTimeRng()
	return &PositionGenerator{enviroment, rng}
}

func (generator *PositionGenerator) Make() Position {
	for {
		var x = generator.rng.Intn(generator.enviroment.size)
		var y = generator.rng.Intn(generator.enviroment.size)
		var position = NewPosition(x, y)
		if generator.enviroment.GetTile(position).IsPassable() {
			return position
		}
	}
}
