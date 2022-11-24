package main

type Position struct {
	x uint8
	y uint8
}

type Microbe struct {
	position      Position
	neuralNetwork NeuralNetwork
}

func NewMicrobe(x uint8, y uint8, in uint8, out uint8, width uint8, depth uint8, genome Genome) Microbe {
	return Microbe{
		position:      Position{x, y},
		neuralNetwork: NewNeuralNetwork(in, out, width, depth, genome),
	}
}
