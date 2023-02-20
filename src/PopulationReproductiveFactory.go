package main

import "math/rand"

type PopulationReproductiveFactory struct {
	positionGenerator    PositionGenerator
	neuralNetworkFactory NeuralNetworkReproductionFactory
	mutator              Mutator
}

func NewPopulationReproductiveFactory(positionGenerator PositionGenerator, neuralNetworkFactory NeuralNetworkReproductionFactory, mutator Mutator) PopulationReproductiveFactory {
	return PopulationReproductiveFactory{
		positionGenerator:    positionGenerator,
		neuralNetworkFactory: neuralNetworkFactory,
		mutator:              mutator,
	}
}

func SelectParentRandomly(parents []*Microbe) *Microbe {
	var i = rand.Intn(len(parents))
	return parents[i]
}

func (factory *PopulationReproductiveFactory) Make(selected []*Microbe, size int) []Microbe {
	var population = make([]Microbe, size)
	for i := 0; i < size; i++ {
		var parent1 = SelectParentRandomly(selected)
		var parent2 = SelectParentRandomly(selected)
		population[i] = Microbe{
			position:      factory.positionGenerator.Make(),
			neuralNetwork: factory.neuralNetworkFactory.Make(parent1, parent2),
		}
	}
	return population
}
