package main

import "math/rand"

type GaussMutationStrategy struct {
	rng      *rand.Rand
	strength float64 // 0.0 - 1.0
}

func NewGaussMutationStrategy(strength float64) *GaussMutationStrategy {
	var rng = NewUnixTimeRng()
	return &GaussMutationStrategy{rng, strength}
}

func (strategy *GaussMutationStrategy) MutateWeight(weight float64) float64 {
	var mutation = strategy.rng.NormFloat64() * (WEIGHT_LIMIT / 2) * strategy.strength
	return weight + mutation
}

func (strategy *GaussMutationStrategy) Mutate(microbe *Microbe) {
	var nn = microbe.neuralNetwork
	for i, weight := range nn.weights {
		nn.weights[i] = ClampWeight(strategy.MutateWeight(weight))
	}
}
