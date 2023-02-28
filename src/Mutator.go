package main

type Mutator struct {
	strategy MutationStrategy
}

type MutationStrategy interface {
	Mutate(microbe *Microbe)
}

func NewMutator(strategy MutationStrategy) *Mutator {
	return &Mutator{strategy}
}

func (mutator *Mutator) MutatePopulation(population Population) {
	for i := range population {
		mutator.strategy.Mutate(population[i])
	}
}
