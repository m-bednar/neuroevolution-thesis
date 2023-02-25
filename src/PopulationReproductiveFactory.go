package main

import "math/rand"

type PopulationReproductiveFactory struct {
	positionGenerator    *PositionGenerator
	neuralNetworkFactory *NeuralNetworkReproductionFactory
	mutator              *Mutator
	rng                  *rand.Rand
}

func NewPopulationReproductiveFactory(positionGenerator *PositionGenerator, neuralNetworkFactory *NeuralNetworkReproductionFactory, mutator *Mutator) *PopulationReproductiveFactory {
	return &PopulationReproductiveFactory{
		positionGenerator:    positionGenerator,
		neuralNetworkFactory: neuralNetworkFactory,
		mutator:              mutator,
		rng:                  NewUnixTimeRng(),
	}
}

func (factory *PopulationReproductiveFactory) SelectParentRandomly(parents []*Microbe) *Microbe {
	var i = factory.rng.Intn(len(parents))
	return parents[i]
}

func (factory *PopulationReproductiveFactory) Make(selected []*Microbe, count int) []*Microbe {
	var population = make([]*Microbe, count)
	for i := 0; i < count; i++ {
		var parent1 = factory.SelectParentRandomly(selected)
		var parent2 = factory.SelectParentRandomly(selected)
		var position = factory.positionGenerator.Make()
		var neuralNetwork = factory.neuralNetworkFactory.Make(parent1, parent2)
		population[i] = NewMicrobe(position, neuralNetwork)
	}
	return population
}
