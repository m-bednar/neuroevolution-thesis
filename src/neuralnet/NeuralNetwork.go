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

func (neuralNetwork *NeuralNetwork) GetWeights() []float64 {
	return neuralNetwork.weights
}

func (neuralNetwork *NeuralNetwork) Process(inputs []float64) []float64 {
	var structure = neuralNetwork.structure
	var widths = structure.GetLayersWidths()
	var maxWidth = structure.GetMaxLayerWidth()
	var lastWidth = widths[len(widths)-1]

	var buffer = make([]float64, maxWidth)
	var neurons = make([]float64, maxWidth)

	copy(buffer, inputs)

	// Traverse layer by layer
	for layer := 1; layer < len(widths); layer++ {
		var previous = widths[layer-1]
		var current = widths[layer]
		var offset = structure.GetLayerOffset(layer)

		// Values of neurons in previous layer
		var values = buffer[:previous]

		// Traverse neurons in current layer and compute it's value
		for i := 0; i < current; i++ {
			var from = offset + (i * previous)
			var to = from + previous
			var weights = neuralNetwork.weights[from:to]
			var sum = WeightedSum(weights, values)
			neurons[i] = ReLU(sum)
		}

		copy(buffer, neurons)
	}

	return buffer[:lastWidth]
}
