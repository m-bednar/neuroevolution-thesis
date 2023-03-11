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

func (microbe *Microbe) Process(selector *ActionSelector, inputs []float64) Position {
	var outputs = microbe.neuralNetwork.Process(inputs)
	var moveX, moveY = selector.SelectAction(outputs)
	return microbe.position.Add(moveX, moveY)
}

func (microbe *Microbe) MoveTo(position Position) {
	microbe.position = position
}
