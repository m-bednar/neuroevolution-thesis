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

func (neuralNetwork *NeuralNetwork) Process(inputs []float64) []float64 {
	structure := neuralNetwork.structure
	widths := structure.GetLayersWidths()
	maxWidth := structure.GetMaxLayerWidth()
	lastWidth := widths[len(widths)-1]

	buffer := make([]float64, maxWidth)
	neurons := make([]float64, maxWidth)

	copy(buffer, inputs)

	// Traverse layer by layer
	for layer := 1; layer < len(widths); layer++ {
		previous := widths[layer-1]
		current := widths[layer]
		offset := structure.GetLayerOffset(layer)

		// Values of neurons in previous layer
		values := buffer[:previous]

		// Traverse neurons in current layer and compute it's value
		for i := 0; i < current; i++ {
			from := offset + (i * previous)
			to := from + previous
			weights := neuralNetwork.weights[from:to]
			sum := WeightedSum(weights, values)
			neurons[i] = ReLU(sum)
		}

		copy(buffer, neurons)
	}

	return buffer[:lastWidth]
}
