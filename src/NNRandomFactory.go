package main

import "math/rand"

type NNRandomFactory struct {
	rng *rand.Rand
}

func NewNNRandomFactory() *NNRandomFactory {
	var rng = NewUnixTimeRng()
	return &NNRandomFactory{rng}
}

func (factory *NNRandomFactory) GenerateFloat(min float64, max float64) float64 {
	var rand = factory.rng.Float64()
	var size = max - min
	return (rand * size) + min
}

func (factory *NNRandomFactory) Make() NeuralNetwork {
	var size = ComputeNumberOfWeights()
	var weights = make([]float64, size)
	for i := 0; i < size; i++ {
		weights[i] = factory.GenerateFloat(-WEIGHT_LIMIT, WEIGHT_LIMIT)
	}
	return NewNeuralNetwork(weights)
}
