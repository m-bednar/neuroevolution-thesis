package main

import "math/rand"

type OnePointCrossoverStrategy struct {
	rng *rand.Rand
}

func NewOnePointCrossoverStrategy() *OnePointCrossoverStrategy {
	return &OnePointCrossoverStrategy{
		rng: NewUnixTimeRng(),
	}
}

func (strategy *OnePointCrossoverStrategy) Crossover(weights1 []float64, weights2 []float64) []float64 {
	var min = 1
	var max = len(weights1) - 1
	var ext = max - min
	var rnd = strategy.rng.Intn(ext) + min
	var part1 = weights1[:rnd]
	var part2 = weights2[rnd:]
	var whole = append(part1, part2...)
	return whole
}
