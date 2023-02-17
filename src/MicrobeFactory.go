package main

type MicrobeFactory struct {
	positionGenerator    PositionGenerator
	neuralNetworkFactory NeuralNetworkFactory
}

func NewMicrobeFactory(positionGenerator PositionGenerator, neuralNetworkFactory NeuralNetworkFactory) MicrobeFactory {
	return MicrobeFactory{positionGenerator, neuralNetworkFactory}
}

func (factory *MicrobeFactory) Make() Microbe {
	return Microbe{
		position:      factory.positionGenerator.Make(),
		neuralNetwork: factory.neuralNetworkFactory.Make(),
	}
}
