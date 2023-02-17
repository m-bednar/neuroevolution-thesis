package main

type MicrobeFactory struct {
	neuralNetworkFactory NeuralNetworkFactory
}

func NewMicrobeFactory(neuralNetworkFactory NeuralNetworkFactory) *MicrobeFactory {
	return &MicrobeFactory{neuralNetworkFactory}
}

func (factory *MicrobeFactory) Make(x uint, y uint) Microbe {
	return Microbe{
		position:      NewPosition(x, y),
		neuralNetwork: factory.neuralNetworkFactory.Make(),
	}
}
