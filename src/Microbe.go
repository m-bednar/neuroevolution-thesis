package main

import "fmt"

type Microbe struct {
	position      Position
	neuralNetwork NeuralNetwork
}

func NewMicrobe(position Position, neuralNetwork NeuralNetwork) Microbe {
	return Microbe{position, neuralNetwork}
}

func (microbe *Microbe) Process(inputs []float64) {
	var outputs = microbe.neuralNetwork.Process(inputs)
	fmt.Println(outputs)
	microbe.position.Move(0, 0) // TODO: Make decision based on outputs
}
