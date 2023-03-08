package main

import (
	"math/rand"
)

type PopulationFactory struct {
	populationSize       int
	positionGenerator    *PositionGenerator
	neuralNetworkFactory *NeuralNetworkFactory
	selector             *ParentSelector
	rng                  *rand.Rand
}

func NewPopulationFactory(populationSize int, positionGenerator *PositionGenerator, neuralNetworkFactory *NeuralNetworkFactory, selector *ParentSelector) *PopulationFactory {
	return &PopulationFactory{
		populationSize:       populationSize,
		positionGenerator:    positionGenerator,
		neuralNetworkFactory: neuralNetworkFactory,
		selector:             selector,
		rng:                  NewUnixTimeRng(),
	}
}

func (factory *PopulationFactory) ReproduceFrom(population []*Microbe) []*Microbe {
	var new = make([]*Microbe, factory.populationSize)
	for i := 0; i < factory.populationSize; i++ {
		var parent1 = factory.selector.SelectOneByTournament(population)
		var parent2 = factory.selector.SelectOneByTournament(population)
		var position = factory.positionGenerator.Make()
		var neuralNetwork = factory.neuralNetworkFactory.Reproduce(parent1, parent2)
		new[i] = NewMicrobe(position, neuralNetwork)
	}
	return new
}

func (factory *PopulationFactory) MakeRandom() []*Microbe {
	var population = make([]*Microbe, factory.populationSize)
	for i := 0; i < factory.populationSize; i++ {
		var position = factory.positionGenerator.Make()
		var neuralNetwork = factory.neuralNetworkFactory.MakeRandom()
		population[i] = NewMicrobe(position, neuralNetwork)
	}
	return population
}
