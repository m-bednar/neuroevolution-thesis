package main

import (
	"fmt"
	"math"
)

const (
	INITIAL_FITNESS         = 0.0
	MOVE_EAST_NEURON_INDEX  = 0
	MOVE_WEST_NEURON_INDEX  = 1
	MOVE_SOUTH_NEURON_INDEX = 2
	MOVE_NORTH_NEURON_INDEX = 3
	NO_ACTION_INDEX         = 4
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

func GetSliceMaxValue(values []float64) float64 {
	var max = -math.MaxFloat64
	for _, value := range values {
		max = math.Max(max, value)
	}
	return max
}

func SoftMax(values []float64) []float64 {
	var max = GetSliceMaxValue(values)

	var sum = 0.0
	var result = make([]float64, len(values))
	for i, value := range values {
		var x = math.Exp(value - max)
		result[i] = x
		sum += x
	}

	if sum == 0 {
		sum = math.SmallestNonzeroFloat64
	}
	for i, value := range result {
		result[i] = value / sum
	}

	return result
}

func ProbabilitySelect(probabilities []float64) int {
	var rnd = rng.Float64()
	for i, probability := range probabilities {
		if rnd <= probability {
			return i
		}
		rnd -= probability
	}
	panic("Probability selection was done on slice with <1 sum")
}

func Activation(outputs []float64) (int, int) {
	var probabilities = SoftMax(outputs)
	var actionIndex = ProbabilitySelect(probabilities)

	switch actionIndex {
	case MOVE_EAST_NEURON_INDEX:
		return 1, 0
	case MOVE_WEST_NEURON_INDEX:
		return -1, 0
	case MOVE_SOUTH_NEURON_INDEX:
		return 0, 1
	case MOVE_NORTH_NEURON_INDEX:
		return 0, -1
	case NO_ACTION_INDEX:
		return 0, 0
	}

	panic("No programmed action was selected")
}

func (microbe *Microbe) Process(inputs []float64) Position {
	var outputs = microbe.neuralNetwork.Process(inputs)
	var moveX, moveY = Activation(outputs)
	return microbe.position.Add(moveX, moveY)
}

func (microbe *Microbe) MoveTo(position Position) {
	microbe.position = position
}

func (microbe *Microbe) Print() {
	fmt.Printf("position: (%2d, %2d) fitness: %2.3f\n", microbe.position.x, microbe.position.y, microbe.fitness)
}
