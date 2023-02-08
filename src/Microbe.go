package main

type Position struct {
	x uint
	y uint
}

type Microbe struct {
	position      Position
	neuralNetwork NeuralNetwork
}
