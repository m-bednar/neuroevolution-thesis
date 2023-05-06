/**
 * @project neuroevolution
 * @file    GaussMutation.go
 * @author  Martin Bednář (xbedna77@fit.vut.cz)
 */

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

func (mutation *GaussMutation) MutateWeight(weight float64) float64 {
	deviation := NN_WEIGHT_LIMIT * mutation.strength
	mutationWeight := mutation.rng.NormFloat64(deviation)
	return weight + mutationWeight
}

func (mutation *GaussMutation) MutateNeuralNetwork(neuralNetwork *NeuralNetwork) {
	weights := neuralNetwork.GetWeights()
	for i, weight := range weights {
		weights[i] = ClampWeight(mutation.MutateWeight(weight))
	}
}

func (mutation *GaussMutation) MutatePopulation(population Population) {
	for _, microbe := range population {
		neuralNetwork := microbe.GetNeuralNetwork()
		mutation.MutateNeuralNetwork(neuralNetwork)
	}
}
