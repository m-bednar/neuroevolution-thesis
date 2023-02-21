package main

import (
	"math/rand"
	"time"
)

type Mutator struct {
	rng        *rand.Rand
	strength   float64 // > 0.0
	occurrence float64 // 0.0 - 1.0
}

func NewMutator(strength float64, occurrence float64) Mutator {
	var seed = time.Now().UnixNano()
	var source = rand.NewSource(seed)
	var rng = rand.New(source)
	return Mutator{rng, strength, occurrence}
}

func (mutator *Mutator) Mutate(microbe *Microbe) {
	var nn = microbe.neuralNetwork
	for i, weight := range nn.weights {
		if mutator.ShouldMutationOccur() {
			nn.weights[i] = weight + mutator.GenerateMutationWeight()
		}
	}
}

func (mutator *Mutator) ShouldMutationOccur() bool {
	return mutator.rng.Float64() <= mutator.occurrence
}

func (mutator *Mutator) GenerateMutationWeight() float64 {
	return (mutator.rng.Float64() - 0.5) * mutator.strength
}
