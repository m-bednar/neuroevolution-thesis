package main

import "fmt"

const (
	INITIAL_FITNESS = 0.0
)

// TODO: Move to more suitable location
var rng = NewUnixTimeRng()

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

func (microbe *Microbe) Process(selector *ActionSelector, inputs []float64) Position {
	var outputs = microbe.neuralNetwork.Process(inputs)
	var moveX, moveY = selector.SelectAction(outputs)
	return microbe.position.Add(moveX, moveY)
}

func (microbe *Microbe) MoveTo(position Position) {
	microbe.position = position
}

func (microbe *Microbe) Print() {
	fmt.Printf("position: (%2d, %2d) fitness: %2.3f\n", microbe.position.x, microbe.position.y, microbe.fitness)
}

func (microbe *Microbe) GetRGBHexCode() (byte, byte, byte) {
	// TODO: output hex-code
	return 0, 0, 0
}
