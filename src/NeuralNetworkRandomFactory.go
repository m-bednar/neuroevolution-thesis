package main

import "math/rand"

const RAND_WEIGHT = 4

type NeuralNetworkRandomFactory struct {
	rng *rand.Rand
}

func NewNeuralNetworkRandomFactory() *NeuralNetworkRandomFactory {
	var rng = NewUnixTimeRng()
	return &NeuralNetworkRandomFactory{rng}
}

func (factory *NeuralNetworkRandomFactory) GenerateFloat(min float64, max float64) float64 {
	var rand = factory.rng.Float64()
	var size = max - min
	return (rand * size) + min
}

func (factory *NeuralNetworkRandomFactory) Make() NeuralNetwork {
	var size = ComputeNumberOfWeights()
	var weights = make([]float64, size)
	for i := 0; i < size; i++ {
		weights[i] = factory.GenerateFloat(-RAND_WEIGHT, RAND_WEIGHT)
	}
	return NeuralNetwork{weights}
}
