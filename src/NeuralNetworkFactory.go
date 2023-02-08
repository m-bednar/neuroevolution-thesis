package main

import (
	"math/rand"
	"time"
)

const (
	LAYER_WIDTH = 6
	MIN_RAND_WEIGHT = -100
	MAX_RAND_WEIGHT = 100
)

type NeuralNetworkFactory interface {
	Make() NeuralNetwork
}

type RandomNeuralNetworkFactory struct {
	layerWidth uint
	generator *rand.Rand
}

type ReproductiveNeuralNetworkFactory struct {
	parent1 *Microbe
	parent2 *Microbe
}

func NewRandomNeuralNetworkFactory(layerWidth uint) *RandomNeuralNetworkFactory {
	var seed = time.Now().UnixNano()
	var source = rand.NewSource(seed)
	var generator = rand.New(source)
	return &RandomNeuralNetworkFactory{ layerWidth, generator }
}

func (factory *RandomNeuralNetworkFactory) GenerateFloat(min float64, max float64) float64 {
	var rand = factory.generator.Float64()
	var size = max - min
	return (rand * size) + min
}

func (factory *RandomNeuralNetworkFactory) Make() NeuralNetwork {
	var size = ComputeNumberOfWeights(factory.layerWidth)
	var weights = make([]float64, size)
	for i := uint(0); i < size; i++ {
		weights[i] = factory.GenerateFloat(MIN_RAND_WEIGHT, MAX_RAND_WEIGHT)
	}
	return NeuralNetwork{weights: weights, layerWidth: factory.layerWidth}
}

func (factory *ReproductiveNeuralNetworkFactory) Make() NeuralNetwork {
	// TODO
	return NeuralNetwork{}
}
