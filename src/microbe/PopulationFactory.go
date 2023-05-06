/**
 * @project neuroevolution/microbe
 * @file    PopulationFactory.go
 * @author  Martin Bednář (xbedna77@fit.vut.cz)
 */

package microbe

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/neuralnet"
)

type PositionSelector interface {
	GetPosition() Position
}

type ParentSelector interface {
	Select(population Population) *Microbe
}

type PopulationFactory struct {
	populationSize       int
	positionSelector     PositionSelector
	parentSelector       ParentSelector
	neuralNetworkFactory *NeuralNetworkFactory
}

func NewPopulationFactory(populationSize int, positionSelector PositionSelector,
	neuralNetworkFactory *NeuralNetworkFactory, selector ParentSelector) *PopulationFactory {
	return &PopulationFactory{
		populationSize:       populationSize,
		positionSelector:     positionSelector,
		neuralNetworkFactory: neuralNetworkFactory,
		parentSelector:       selector,
	}
}

func (factory *PopulationFactory) ReproduceFrom(population Population) Population {
	new := make(Population, factory.populationSize)
	for i := 0; i < factory.populationSize; i++ {
		parent1 := factory.parentSelector.Select(population)
		parent2 := factory.parentSelector.Select(population)
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
