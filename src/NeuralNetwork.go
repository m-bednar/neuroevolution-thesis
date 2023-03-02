package main

import "math"

const (
	LAYER_WIDTH  = 8
	N_LAYERS     = 2
	N_OUTPUTS    = 4
	N_INPUTS     = 2
	WEIGHT_LIMIT = 12.0
)

type NeuralNetwork struct {
	weights []float64
}

func ClampWeight(weight float64) float64 {
	return math.Max(-WEIGHT_LIMIT, math.Min(WEIGHT_LIMIT, weight))
}

func ClampWeights(weights []float64) {
	for i, weight := range weights {
		weights[i] = ClampWeight(weight)
	}
}

func NewNeuralNetwork(weights []float64) NeuralNetwork {
	ClampWeights(weights)
	return NeuralNetwork{weights}
}

func ComputeNumberOfWeights() int {
	// TODO: implement through GetLayersWidths()
	return (N_INPUTS * LAYER_WIDTH) + (N_OUTPUTS * LAYER_WIDTH) + ((LAYER_WIDTH * LAYER_WIDTH) * (N_LAYERS - 1))
}

func GetLayersWidths() []int {
	// TODO Make automatic
	return []int{N_INPUTS, LAYER_WIDTH, LAYER_WIDTH, N_OUTPUTS}
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

func GetLayerIndexOffset(layer int) int {
	var layersWidths = GetLayersWidths()
	var sum = 0
	for i := 0; i < (layer - 1); i++ {
		sum += layersWidths[i] * layersWidths[i+1]
	}
	return sum
}

func (neuralNetwork *NeuralNetwork) Process(inputs []float64) []float64 {
	var widths = GetLayersWidths()
	var buffer = inputs

	// Traverse layer by layer
	for layer := 1; layer < len(widths); layer++ {
		var previousWidth = widths[layer-1]
		var currentWidth = widths[layer]
		var offset = GetLayerIndexOffset(layer)

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
