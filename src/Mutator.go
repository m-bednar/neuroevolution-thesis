package main

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
)

type Mutator struct {
	strategy MutationStrategy
}

type MutationStrategy interface {
	Mutate(weights []float64)
}

func NewMutator(strategy MutationStrategy) *Mutator {
	return &Mutator{strategy}
}

func (mutator *Mutator) MutatePopulation(population Population) {
	for i := range population {
		mutator.strategy.Mutate(population[i].GetNeuralNetwork().GetWeights())
	}
}
