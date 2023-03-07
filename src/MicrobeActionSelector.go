package main

import (
	"math"
	"math/rand"
)

const (
	MOVE_EAST_ACTION_INDEX = iota
	MOVE_WEST_ACTION_INDEX
	MOVE_SOUTH_ACTION_INDEX
	MOVE_NORTH_ACTION_INDEX
	NO_MOVE_ACTION_INDEX
)

type ActionSelector struct {
	rng *rand.Rand
}

func NewActionSelector() *ActionSelector {
	return &ActionSelector{
		rng: NewUnixTimeRng(),
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

func (selector *ActionSelector) ProbabilitySelect(probabilities []float64) int {
	var rnd = selector.rng.Float64()
	for i, probability := range probabilities {
		if rnd <= probability {
			return i
		}
		rnd -= probability
	}
	panic("Probability selection was done on slice with <1 sum")
}

func (selector *ActionSelector) SelectAction(neuronOutputs []float64) (int, int) {
	var probabilities = SoftMax(neuronOutputs)
	var action = selector.ProbabilitySelect(probabilities)

	switch action {
	case MOVE_EAST_ACTION_INDEX:
		return 1, 0
	case MOVE_WEST_ACTION_INDEX:
		return -1, 0
	case MOVE_SOUTH_ACTION_INDEX:
		return 0, 1
	case MOVE_NORTH_ACTION_INDEX:
		return 0, -1
	case NO_MOVE_ACTION_INDEX:
		return 0, 0
	}

	panic("No programmed action was selected")
}
