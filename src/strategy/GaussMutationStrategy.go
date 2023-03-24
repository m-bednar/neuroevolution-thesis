package strategy

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/neuralnet"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
)

type GaussMutationStrategy struct {
	rng      *Rng
	strength float64
}

func NewGaussMutationStrategy(strength float64) *GaussMutationStrategy {
	var rng = NewTimeSeedRng()
	return &GaussMutationStrategy{rng, strength}
}

func (strategy *GaussMutationStrategy) MutateWeight(weight float64) float64 {
	var halfLimit = NN_WEIGHT_LIMIT / 2
	var deviation = halfLimit * strategy.strength
	var mutation = strategy.rng.NormFloat64(deviation)
	return weight + mutation
}

func (strategy *GaussMutationStrategy) Mutate(weights []float64) {
	for i, weight := range weights {
		weights[i] = ClampWeight(strategy.MutateWeight(weight))
	}
}
