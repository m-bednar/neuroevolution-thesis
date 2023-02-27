package main

import "fmt"

const MOVE_TRESHOLD = 0.1

type Population []*Microbe

type Microbe struct {
	position      Position
	neuralNetwork NeuralNetwork
	fitness       float64
}

func NewMicrobe(position Position, neuralNetwork NeuralNetwork) *Microbe {
	return &Microbe{
		position:      position,
		neuralNetwork: neuralNetwork,
		fitness:       0.0,
	}
}

func TresholdedSign(value float64, treshold float64) int {
	if value >= treshold {
		return 1
	}
	if value <= -treshold {
		return -1
	}
	return 0
}

func Activation(outputs []float64) (int, int) {
	var moveX = TresholdedSign(outputs[0]-outputs[1], MOVE_TRESHOLD)
	var moveY = TresholdedSign(outputs[2]-outputs[3], MOVE_TRESHOLD)
	return moveX, moveY
}

func (microbe *Microbe) Process(inputs []float64) Position {
	var outputs = microbe.neuralNetwork.Process(inputs)
	var proposed = microbe.position.Add(Activation(outputs))
	return proposed
}

func (microbe *Microbe) MoveTo(position Position) {
	microbe.position = position
}

func (microbe *Microbe) Print() {
	fmt.Printf("position: (%2d, %2d) fitness: %2.3f\n", microbe.position.x, microbe.position.y, microbe.fitness)
}
