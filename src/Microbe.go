package main

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
