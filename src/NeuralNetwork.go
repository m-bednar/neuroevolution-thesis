package main

const (
	LAYER_WIDTH = 6
	N_LAYERS    = 2
	N_OUTPUTS   = 4
	N_INPUTS    = 2
)

type NeuralNetwork struct {
	weights []float64
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
	var layersWidths = GetLayersWidths()
	var in = inputs

	// Traverse layer by layer
	for layer := 1; layer < len(layersWidths); layer++ {
		var currLayerWidth = layersWidths[layer]
		var prevLayerWidth = layersWidths[layer-1]
		var offset = GetLayerIndexOffset(layer)

		// Inner value of each neuron
		var neurons = make([]float64, currLayerWidth)

		// Traverse neurons and compute it's value
		for i := 0; i < currLayerWidth; i++ {
			var from = offset + (i * prevLayerWidth)
			var to = from + prevLayerWidth
			var weights = neuralNetwork.weights[from:to]
			var sum = WeightedSum(weights, in)
			neurons[i] = ReLU(sum)
		}

		in = make([]float64, currLayerWidth)
		copy(in, neurons)
	}

	return in
}
