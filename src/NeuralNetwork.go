package main

import (
	"math"
)

const (
	NN_WEIGHT_LIMIT = 8.0
)

type NeuralNetwork struct {
	structure *NeuralNetworkStructure
	weights   []float64
}

func ClampWeight(weight float64) float64 {
	return math.Max(-NN_WEIGHT_LIMIT, math.Min(NN_WEIGHT_LIMIT, weight))
}

func ClampWeights(weights []float64) {
	for i, weight := range weights {
		weights[i] = ClampWeight(weight)
	}
}

func NewNeuralNetwork(structure *NeuralNetworkStructure, weights []float64) NeuralNetwork {
	ClampWeights(weights)
	return NeuralNetwork{structure, weights}
}

func WeightedSum(weights []float64, inputs []float64) float64 {
	var sum = 0.0
	for i, value := range inputs {
		sum += value * weights[i]
	}
	return sum
}

func ReLU(value float64) float64 {
	if value <= 0.0 {
		return 0.0
	}
	return value
}

func (neuralNetwork *NeuralNetwork) Process(inputs []float64) []float64 {
	var widths = neuralNetwork.structure.GetLayersWidths()
	var buffer = inputs

	// Traverse layer by layer
	for layer := 1; layer < len(widths); layer++ {
		var previousWidth = widths[layer-1]
		var currentWidth = widths[layer]
		var offset = neuralNetwork.structure.GetLayerIndexOffset(layer)

		// Values of each neuron
		var neurons = make([]float64, currentWidth)

		// Traverse neurons and compute it's value
		for i := 0; i < currentWidth; i++ {
			var from = offset + (i * previousWidth)
			var to = from + previousWidth
			var weights = neuralNetwork.weights[from:to]
			var sum = WeightedSum(weights, buffer)
			neurons[i] = ReLU(sum)
		}

		buffer = make([]float64, currentWidth)
		copy(buffer, neurons)
	}

	return buffer
}
