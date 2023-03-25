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
	new := make(Population, factory.populationSize)
	for i := 0; i < factory.populationSize; i++ {
		parent1 := factory.selector.SelectParent(population)
		parent2 := factory.selector.SelectParent(population)
		position := factory.positionSelector.GetPosition()
		neuralNetwork1 := parent1.neuralNetwork
		neuralNetwork2 := parent2.neuralNetwork
		neuralNetwork := factory.neuralNetworkFactory.Reproduce(neuralNetwork1, neuralNetwork2)
		new[i] = NewMicrobe(position, neuralNetwork)
	}
	return new
}

func (factory *PopulationFactory) MakeRandom() Population {
	population := make(Population, factory.populationSize)
	for i := 0; i < factory.populationSize; i++ {
		position := factory.positionSelector.GetPosition()
		neuralNetwork := factory.neuralNetworkFactory.MakeRandom()
		population[i] = NewMicrobe(position, neuralNetwork)
	}
	return population
}
