package main

import "math/rand"

type CrossoverStrategy interface {
	Crossover(weights1 []float64, weights2 []float64) []float64
}

type NeuralNetworkFactory struct {
	rng      *rand.Rand
	strategy CrossoverStrategy
}

func NewNeuralNetworkFactory(strategy CrossoverStrategy) *NeuralNetworkFactory {
	return &NeuralNetworkFactory{
		rng:      NewUnixTimeRng(),
		strategy: strategy,
	}
}

func (factory *NeuralNetworkFactory) GenerateFloat(min float64, max float64) float64 {
	var rand = factory.rng.Float64()
	var size = max - min
	return (rand * size) + min
}

func (factory *NeuralNetworkFactory) Reproduce(parent1 *Microbe, parent2 *Microbe) NeuralNetwork {
	var nn1 = parent1.neuralNetwork
	var nn2 = parent2.neuralNetwork
	var weights = factory.strategy.Crossover(nn1.weights, nn2.weights)
	return NewNeuralNetwork(weights)
}

func (factory *NeuralNetworkFactory) MakeRandom() NeuralNetwork {
	var size = ComputeNumberOfWeights()
	var weights = make([]float64, size)
	for i := 0; i < size; i++ {
		weights[i] = factory.GenerateFloat(-NN_WEIGHT_LIMIT, NN_WEIGHT_LIMIT)
	}
	return NewNeuralNetwork(weights)
}
