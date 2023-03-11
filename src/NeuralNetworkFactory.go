package main

import "math/rand"

type CrossoverStrategy interface {
	Crossover(weights1 []float64, weights2 []float64) []float64
}

type NeuralNetworkFactory struct {
	rng       *rand.Rand
	structure *NeuralNetworkStructure
	strategy  CrossoverStrategy
}

func NewNeuralNetworkFactory(structure *NeuralNetworkStructure, strategy CrossoverStrategy) *NeuralNetworkFactory {
	return &NeuralNetworkFactory{
		rng:       NewUnixTimeRng(),
		structure: structure,
		strategy:  strategy,
	}
}

func (factory *NeuralNetworkFactory) GenerateFloat(min float64, max float64) float64 {
	var rand = factory.rng.Float64()
	var size = max - min
	return (rand * size) + min
}

func (factory *NeuralNetworkFactory) Reproduce(nn1 NeuralNetwork, nn2 NeuralNetwork) NeuralNetwork {
	var weights = factory.strategy.Crossover(nn1.weights, nn2.weights)
	return NewNeuralNetwork(factory.structure, weights)
}

func (factory *NeuralNetworkFactory) MakeRandom() NeuralNetwork {
	var size = factory.structure.ComputeNumberOfWeights()
	var weights = make([]float64, size)
	for i := 0; i < size; i++ {
		weights[i] = factory.GenerateFloat(-NN_WEIGHT_LIMIT, NN_WEIGHT_LIMIT)
	}
	return NewNeuralNetwork(factory.structure, weights)
}
