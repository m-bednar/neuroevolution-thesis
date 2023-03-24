package neuralnet

import "math"

const (
	NN_WEIGHT_LIMIT = 2.0
)

func ClampWeight(weight float64) float64 {
	return math.Max(-NN_WEIGHT_LIMIT, math.Min(NN_WEIGHT_LIMIT, weight))
}

func ClampWeights(weights []float64) {
	for i, weight := range weights {
		weights[i] = ClampWeight(weight)
	}
}
