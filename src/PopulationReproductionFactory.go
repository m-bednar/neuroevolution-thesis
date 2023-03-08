package main

import (
	"math/rand"
)

type PopulationReproductionFactory struct {
	positionGenerator    *PositionGenerator
	neuralNetworkFactory *NeuralNetworkFactory
	selector             *ParentSelector
	rng                  *rand.Rand
}

func NewPopulationReproductionFactory(positionGenerator *PositionGenerator, neuralNetworkFactory *NeuralNetworkFactory, selector *ParentSelector) *PopulationReproductionFactory {
	return &PopulationReproductionFactory{
		positionGenerator:    positionGenerator,
		neuralNetworkFactory: neuralNetworkFactory,
		selector:             selector,
		rng:                  NewUnixTimeRng(),
	}
}

func (factory *PopulationReproductionFactory) Make(population []*Microbe, count int) []*Microbe {
	var new = make([]*Microbe, count)
	for i := 0; i < count; i++ {
		var parent1 = factory.selector.SelectOneByTournament(population)
		var parent2 = factory.selector.SelectOneByTournament(population)
		var position = factory.positionGenerator.Make()
		var neuralNetwork = factory.neuralNetworkFactory.Reproduce(parent1, parent2)
		new[i] = NewMicrobe(position, neuralNetwork)
	}
	return new
}
