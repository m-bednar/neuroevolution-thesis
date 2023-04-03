package neuralnet

type NeuralNetwork struct {
	structure *NeuralNetworkStructure
	weights   []float64
}

func NewNeuralNetwork(structure *NeuralNetworkStructure, weights []float64) NeuralNetwork {
	ClampWeights(weights)
	return NeuralNetwork{structure, weights}
}

func WeightedSum(weights []float64, inputs []float64) float64 {
	sum := 0.0
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

func (neuralNetwork *NeuralNetwork) GetWeights() []float64 {
	return neuralNetwork.weights
}

func (neuralNetwork *NeuralNetwork) GetWeightsFor(layer, neuron int) []float64 {
	from, to := neuralNetwork.structure.GetWeightsIndexSpan(layer, neuron)
	return neuralNetwork.weights[from:to]
}

func (neuralNetwork *NeuralNetwork) Process(inputs []float64) []float64 {
	widths := neuralNetwork.structure.GetLayersWidths()
	maxWidth := neuralNetwork.structure.GetMaxLayerWidth()
	lastWidth := widths[len(widths)-1]

	buffer := make([]float64, maxWidth)
	values := make([]float64, maxWidth)
	copy(buffer, inputs)

	// Traverse layer by layer
	for layer := 1; layer < len(widths); layer++ {
		previous := widths[layer-1]
		current := widths[layer]

		// Traverse neurons in current layer and compute it's value
		for neuron := 0; neuron < current; neuron++ {
			weights := neuralNetwork.GetWeightsFor(layer, neuron)
			sum := WeightedSum(weights, buffer[:previous])
			values[neuron] = ReLU(sum)
		}

		copy(buffer, values)
	}

	return values[:lastWidth]
}
