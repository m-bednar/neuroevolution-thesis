package main

const (
	LAYER_WIDTH = 4
	N_LAYERS    = 2

	N_OUTPUTS = 4
	N_INPUTS  = 2
)

type NeuralNetwork struct {
	weights []float64
}

func ComputeNumberOfWeights() int {
	return (N_INPUTS * LAYER_WIDTH) + (N_OUTPUTS * LAYER_WIDTH) + ((LAYER_WIDTH * LAYER_WIDTH) * N_LAYERS)
}

func (neuralNetwork *NeuralNetwork) Process(inputs []float64) []float64 {
	var buffer = make([]float64, LAYER_WIDTH)
	var widths = []int{N_INPUTS, LAYER_WIDTH, LAYER_WIDTH, LAYER_WIDTH, N_OUTPUTS}
	var lastShift = 0

	for i := 1; i < len(widths); i++ {
		var currWidth = widths[i]
		var prevWidth = widths[i-1]
		for n := 1; n < currWidth; n++ {
			var shift = lastShift + n*prevWidth
			var sum = neuralNetwork.WeightedSum(shift, neuralNetwork.weights, inputs)
			buffer[n] = neuralNetwork.Activation(sum)
		}
		lastShift = lastShift + (currWidth-1)*prevWidth
	}

	return buffer[0:N_OUTPUTS]
}

func (neuralNetwork *NeuralNetwork) Activation(value float64) float64 {
	if value <= 0.0 {
		return 0.0
	}
	return value
}

func (neuralNetwork *NeuralNetwork) WeightedSum(shift int, weights []float64, values []float64) float64 {
	var sum = 0.0
	for i, value := range values {
		sum += value * weights[shift+i]
	}
	return sum
}
