package microbe

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/env"
	. "github.com/m-bednar/neuroevolution-thesis/src/neuralnet"
)

const (
	INITIAL_FITNESS = 0.0
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

func (microbe *Microbe) Process(inputs []float64) []float64 {
	return microbe.neuralNetwork.Process(inputs)
}

func (microbe *Microbe) MoveTo(position Position) {
	microbe.position = position
}

func (microbe *Microbe) GetPosition() Position {
	return microbe.position
}

func (microbe *Microbe) GetNeuralNetwork() *NeuralNetwork {
	return &microbe.neuralNetwork
}

func (microbe *Microbe) GetFitness() float64 {
	return microbe.fitness
}

func (microbe *Microbe) SetFitness(fitness float64) {
	microbe.fitness = fitness
}
