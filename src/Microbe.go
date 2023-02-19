package main

import "fmt"

type Microbe struct {
	position      Position
	neuralNetwork NeuralNetwork
}

func NewMicrobe(position Position, neuralNetwork NeuralNetwork) Microbe {
	return Microbe{
		position:      position,
		neuralNetwork: neuralNetwork,
	}
}

func Sign(value float64) int {
	if value >= 0 {
		return 1
	}
	if value <= 0 {
		return -1
	}
	return 0
}

func Activation(outputs []float64) (int, int) {
	var moveX = outputs[0] - outputs[1]
	var moveY = outputs[2] - outputs[3]
	if moveX > moveY {
		return Sign(moveX), 0
	} else {
		return 0, Sign(moveY)
	}
}

func (microbe *Microbe) Process(inputs []float64) Position {
	var outputs = microbe.neuralNetwork.Process(inputs)
	return microbe.position.Add(Activation(outputs))
}

func (microbe *Microbe) MoveTo(position Position) {
	fmt.Println(microbe.position, "->", position)
	microbe.position = position
}
