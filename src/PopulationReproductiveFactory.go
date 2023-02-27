package main

import (
	"math/rand"
)

const TOURNAMENT_SIZE = 4

type PopulationReproductiveFactory struct {
	positionGenerator    *PositionGenerator
	neuralNetworkFactory *NNReproductionFactory
	selector             *Selector
	rng                  *rand.Rand
}

func NewPopulationReproductiveFactory(positionGenerator *PositionGenerator, neuralNetworkFactory *NNReproductionFactory, selector *Selector) *PopulationReproductiveFactory {
	return &PopulationReproductiveFactory{
		positionGenerator:    positionGenerator,
		neuralNetworkFactory: neuralNetworkFactory,
		selector:             selector,
		rng:                  NewUnixTimeRng(),
	}
}

func (factory *PopulationReproductiveFactory) Make(population []*Microbe, count int) []*Microbe {
	var new = make([]*Microbe, count)
	for i := 0; i < count; i++ {
		var parent1 = factory.selector.SelectOneByTournament(population, TOURNAMENT_SIZE)
		var parent2 = factory.selector.SelectOneByTournament(population, TOURNAMENT_SIZE)
		var position = factory.positionGenerator.Make()
		var neuralNetwork = factory.neuralNetworkFactory.Make(parent1, parent2)
		var reproduced = NewMicrobe(position, neuralNetwork)
		new[i] = reproduced
	}
	return new
}
