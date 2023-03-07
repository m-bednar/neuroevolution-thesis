package main

import "math"

const (
	NN_LAYER_WIDTH  = 8
	NN_NUM_LAYERS   = 2
	NN_NUM_OUTPUTS  = 5
	NN_NUM_INPUTS   = 2
	NN_WEIGHT_LIMIT = 6.0
)

type NeuralNetwork struct {
	weights []float64
}

func ClampWeight(weight float64) float64 {
	return math.Max(-NN_WEIGHT_LIMIT, math.Min(NN_WEIGHT_LIMIT, weight))
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
	return (NN_NUM_INPUTS * NN_LAYER_WIDTH) + (NN_NUM_OUTPUTS * NN_LAYER_WIDTH) + ((NN_LAYER_WIDTH * NN_LAYER_WIDTH) * (NN_NUM_LAYERS - 1))
}

func GetLayersWidths() []int {
	// TODO Make automatic
	return []int{NN_NUM_INPUTS, NN_LAYER_WIDTH, NN_LAYER_WIDTH, NN_NUM_OUTPUTS}
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
