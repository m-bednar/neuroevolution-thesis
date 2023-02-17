package main

type Microbe struct {
	position      Position
	neuralNetwork NeuralNetwork
}

func NewMicrobe(position Position, neuralNetwork NeuralNetwork) Microbe {
	return Microbe{position, neuralNetwork}
}
