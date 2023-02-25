package main

type Microbe struct {
	position      Position
	startPosition Position
	neuralNetwork NeuralNetwork
	fitness       float64
}

func NewMicrobe(position Position, neuralNetwork NeuralNetwork) *Microbe {
	return &Microbe{
		position:      position,
		startPosition: position,
		neuralNetwork: neuralNetwork,
		fitness:       0.0,
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
	return Sign(moveX), Sign(moveY)
}

func (microbe *Microbe) Process(inputs []float64) Position {
	var outputs = microbe.neuralNetwork.Process(inputs)
	return microbe.position.Add(Activation(outputs))
}

func (microbe *Microbe) MoveTo(position Position) {
	microbe.position = position
}
