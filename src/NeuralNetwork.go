package main

const (
	N_LAYERS  = 3
	N_OUTPUTS = 3
	N_INPUTS  = 3
)

type NeuralNetwork struct {
	layerWidth uint
	weights    []float64
}

func ComputeNumberOfWeights(width uint) uint {
	return (N_INPUTS * width) + (N_OUTPUTS * width) + ((width * width) * N_LAYERS)
}
