package main

import "math/rand"

type Mutator struct {
	rng      *rand.Rand
	strength float64 // 0.0 - 1.0
}

func NewMutator(strength float64) *Mutator {
	if strength < 0 || strength > 1 {
		panic("Mutator strength outside of range 0-1")
	}
	var rng = NewUnixTimeRng()
	return &Mutator{rng, strength}
}

func (mutator *Mutator) MutateMicrobe(microbe *Microbe) {
	var nn = microbe.neuralNetwork
	for i, weight := range nn.weights {
		nn.weights[i] = ClampWeight(weight + mutator.GenerateMutationWeight())
	}
}

func (mutator *Mutator) MutatePopulation(population Population) {
	for i := range population {
		mutator.MutateMicrobe(population[i])
	}
}

func (mutator *Mutator) GenerateMutationWeight() float64 {
	return mutator.rng.NormFloat64() * (WEIGHT_LIMIT / 2) * mutator.strength
}
