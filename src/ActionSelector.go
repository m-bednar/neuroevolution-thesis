package main

import (
	"math"

	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
)

const (
	MOVE_NORTH_ACTION_INDEX = iota
	MOVE_SOUTH_ACTION_INDEX
	MOVE_WEST_ACTION_INDEX
	MOVE_EAST_ACTION_INDEX
	NO_MOVE_ACTION_INDEX
)

type ActionSelector struct {
	rng *MutexedRand
}

func NewActionSelector() *ActionSelector {
	return &ActionSelector{
		rng: NewMutexedRand(),
	}
}

func GetSliceMaxValue(values []float64) float64 {
	max := -math.MaxFloat64
	for _, value := range values {
		max = math.Max(max, value)
	}
	return max
}

func SoftMax(values []float64) []float64 {
	max := GetSliceMaxValue(values)

	sum := 0.0
	result := make([]float64, len(values))
	for i, value := range values {
		x := math.Exp(value - max)
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
	rnd := selector.rng.Float64()

	for i, probability := range probabilities {
		if rnd <= probability {
			return i
		}
		rnd -= probability
	}
	panic("Probability selection was done on slice with <1 sum")
}

func (selector *ActionSelector) SelectMoveAction(neuronOutputs []float64) Direction {
	probabilities := SoftMax(neuronOutputs)
	action := selector.ProbabilitySelect(probabilities)

	switch action {
	case MOVE_NORTH_ACTION_INDEX:
		return North
	case MOVE_SOUTH_ACTION_INDEX:
		return South
	case MOVE_WEST_ACTION_INDEX:
		return West
	case MOVE_EAST_ACTION_INDEX:
		return East
	case NO_MOVE_ACTION_INDEX:
		return Origin
	}

	panic("No programmed action was selected")
}
