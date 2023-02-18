package main

type PopulationReproductiveFactory struct {
	positionGenerator    PositionGenerator
	neuralNetworkFactory NeuralNetworkReproductiveFactory
}

func NewPopulationReproductiveFactory(positionGenerator PositionGenerator, neuralNetworkFactory NeuralNetworkReproductiveFactory) PopulationReproductiveFactory {
	return PopulationReproductiveFactory{positionGenerator, neuralNetworkFactory}
}

func (factory *PopulationReproductiveFactory) Make(oldPopulation []Microbe) []Microbe {
	var size = len(oldPopulation)
	var newPopulation = make([]Microbe, size)
	copy(newPopulation, oldPopulation) // TODO
	return newPopulation
}
