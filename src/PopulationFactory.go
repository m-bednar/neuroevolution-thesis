package main

import (
	"math/rand"
)

type PopulationFactory struct {
	populationSize       int
	spawnSelector        *SpawnSelector
	neuralNetworkFactory *NeuralNetworkFactory
	selector             *ParentSelector
	rng                  *rand.Rand
}

func NewPopulationFactory(populationSize int, spawnSelector *SpawnSelector, neuralNetworkFactory *NeuralNetworkFactory, selector *ParentSelector) *PopulationFactory {
	return &PopulationFactory{
		populationSize:       populationSize,
		spawnSelector:        spawnSelector,
		neuralNetworkFactory: neuralNetworkFactory,
		selector:             selector,
		rng:                  NewUnixTimeRng(),
	}
}

func (factory *PopulationFactory) ReproduceFrom(population Population) Population {
	var new = make(Population, factory.populationSize)
	for i := 0; i < factory.populationSize; i++ {
		var parent1 = factory.selector.SelectOneByTournament(population)
		var parent2 = factory.selector.SelectOneByTournament(population)
		var position = factory.spawnSelector.GetRandomSpawnPosition()
		var neuralNetwork1 = parent1.neuralNetwork
		var neuralNetwork2 = parent2.neuralNetwork
		var neuralNetwork = factory.neuralNetworkFactory.Reproduce(neuralNetwork1, neuralNetwork2)
		new[i] = NewMicrobe(position, neuralNetwork)
	}
	return new
}

func (factory *PopulationFactory) MakeRandom() Population {
	var population = make(Population, factory.populationSize)
	for i := 0; i < factory.populationSize; i++ {
		var position = factory.spawnSelector.GetRandomSpawnPosition()
		var neuralNetwork = factory.neuralNetworkFactory.MakeRandom()
		population[i] = NewMicrobe(position, neuralNetwork)
	}
	return population
}
