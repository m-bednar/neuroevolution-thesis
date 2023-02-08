package main

type MicrobeFactory struct {
	neuralNetworkFactory NeuralNetworkFactory
}

func (factory *MicrobeFactory) Make(x uint, y uint) Microbe {
	return Microbe{
		position:      Position{x, y},
		neuralNetwork: factory.neuralNetworkFactory.Make(),
	}
}

func NewMicrobeFactory(neuralNetworkFactory NeuralNetworkFactory) *MicrobeFactory {
	return &MicrobeFactory{neuralNetworkFactory}
}
