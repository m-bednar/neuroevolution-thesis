package main

import (
	"math/rand"
	"time"
)

type NeuralNetworkRandomFactory struct {
	layerWidth uint
	generator  rand.Rand
}

func NewNeuralNetworkRandomizedFactory(layerWidth uint) NeuralNetworkRandomFactory {
	var seed = time.Now().UnixNano()
	var source = rand.NewSource(seed)
	var generator = rand.New(source)
	return NeuralNetworkRandomFactory{layerWidth, *generator}
}

func (factory *NeuralNetworkRandomFactory) GenerateFloat(min float64, max float64) float64 {
	var rand = factory.generator.Float64()
	var size = max - min
	return (rand * size) + min
}

func (factory *NeuralNetworkRandomFactory) Make() NeuralNetwork {
	var size = ComputeNumberOfWeights(factory.layerWidth)
	var weights = make([]float64, size)
	for i := uint(0); i < size; i++ {
		weights[i] = factory.GenerateFloat(MIN_RAND_WEIGHT, MAX_RAND_WEIGHT)
	}
	return NeuralNetwork{ weights }
}
