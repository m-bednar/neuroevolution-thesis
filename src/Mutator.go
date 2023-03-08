package main

type Mutator struct {
	strategy MutationStrategy
}

type MutationStrategy interface {
	Mutate(weights []float64)
}

func NewMutator(strategy MutationStrategy) *Mutator {
	return &Mutator{strategy}
}

func (mutator *Mutator) MutatePopulation(population []*Microbe) {
	for i := range population {
		mutator.strategy.Mutate(population[i].neuralNetwork.weights)
	}
}
