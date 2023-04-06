package main

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
	. "github.com/m-bednar/neuroevolution-thesis/src/neuralnet"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
)

type GaussMutation struct {
	rng      *Rng
	strength float64
}

func NewGaussMutation(strength float64) *GaussMutation {
	rng := NewTimeSeedRng()
	return &GaussMutation{rng, strength}
}

func (strategy *GaussMutation) MutateWeight(weight float64) float64 {
	deviation := NN_WEIGHT_LIMIT * strategy.strength
	mutation := strategy.rng.NormFloat64(deviation)
	return weight + mutation
}

func (strategy *GaussMutation) MutateNeuralNetwork(neuralNetwork *NeuralNetwork) {
	weights := neuralNetwork.GetWeights()
	for i, weight := range weights {
		weights[i] = ClampWeight(strategy.MutateWeight(weight))
	}
}

func (strategy *GaussMutation) MutatePopulation(population Population) {
	for _, microbe := range population {
		neuralNetwork := microbe.GetNeuralNetwork()
		strategy.MutateNeuralNetwork(neuralNetwork)
	}
}
