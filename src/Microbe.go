package main

import "fmt"

const (
	INITIAL_FITNESS         = 0.0
	MOVE_TRESHOLD           = 0.1
	MOVE_EAST_NEURON_INDEX  = 0
	MOVE_WEST_NEURON_INDEX  = 1
	MOVE_SOUTH_NEURON_INDEX = 2
	MOVE_NORTH_NEURON_INDEX = 3
)

type Microbe struct {
	position      Position
	neuralNetwork NeuralNetwork
	fitness       float64
}

func NewMicrobe(position Position, neuralNetwork NeuralNetwork) *Microbe {
	return &Microbe{
		position:      position,
		neuralNetwork: neuralNetwork,
		fitness:       INITIAL_FITNESS,
	}
}

func TresholdedSign(value float64) int {
	if value >= MOVE_TRESHOLD {
		return 1
	}
	if value <= -MOVE_TRESHOLD {
		return -1
	}
	return 0
}

func Activation(outputs []float64) (int, int) {
	var moveX = outputs[MOVE_EAST_NEURON_INDEX] - outputs[MOVE_WEST_NEURON_INDEX]
	var moveY = outputs[MOVE_SOUTH_NEURON_INDEX] - outputs[MOVE_NORTH_NEURON_INDEX]
	return TresholdedSign(moveX), TresholdedSign(moveY)
}

func (microbe *Microbe) Process(inputs []float64) Position {
	var outputs = microbe.neuralNetwork.Process(inputs)
	var moveX, moveY = Activation(outputs)
	var proposed = microbe.position.Add(moveX, moveY)
	return proposed
}

func (microbe *Microbe) MoveTo(position Position) {
	microbe.position = position
}

func (microbe *Microbe) Print() {
	fmt.Printf("position: (%2d, %2d) fitness: %2.3f\n", microbe.position.x, microbe.position.y, microbe.fitness)
}
