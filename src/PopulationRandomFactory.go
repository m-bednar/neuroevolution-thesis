package main

type PopulationRandomFactory struct {
	positionGenerator    *PositionGenerator
	neuralNetworkFactory *NeuralNetworkFactory
}

func NewPopulationRandomFactory(positionGenerator *PositionGenerator, neuralNetworkFactory *NeuralNetworkFactory) *PopulationRandomFactory {
	return &PopulationRandomFactory{
		positionGenerator:    positionGenerator,
		neuralNetworkFactory: neuralNetworkFactory,
	}
}

func (factory *PopulationRandomFactory) Make(size int) []*Microbe {
	var population = make([]*Microbe, size)
	for i := 0; i < size; i++ {
		var position = factory.positionGenerator.Make()
		var neuralNetwork = factory.neuralNetworkFactory.MakeRandom()
		population[i] = NewMicrobe(position, neuralNetwork)
	}
	return population
}
