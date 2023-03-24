package microbe

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/neuralnet"
)

type PositionSelector interface {
	GetPosition() Position
}

type ParentSelector interface {
	SelectParent(population Population) *Microbe
}

type PopulationFactory struct {
	populationSize       int
	positionSelector     PositionSelector
	selector             ParentSelector
	neuralNetworkFactory *NeuralNetworkFactory
}

func NewPopulationFactory(populationSize int, positionSelector PositionSelector, neuralNetworkFactory *NeuralNetworkFactory, selector ParentSelector) *PopulationFactory {
	return &PopulationFactory{
		populationSize:       populationSize,
		positionSelector:     positionSelector,
		neuralNetworkFactory: neuralNetworkFactory,
		selector:             selector,
	}
}

func (factory *PopulationFactory) ReproduceFrom(population Population) Population {
	var new = make(Population, factory.populationSize)
	for i := 0; i < factory.populationSize; i++ {
		var parent1 = factory.selector.SelectParent(population)
		var parent2 = factory.selector.SelectParent(population)
		var position = factory.positionSelector.GetPosition()
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
		var position = factory.positionSelector.GetPosition()
		var neuralNetwork = factory.neuralNetworkFactory.MakeRandom()
		population[i] = NewMicrobe(position, neuralNetwork)
	}
	return population
}
