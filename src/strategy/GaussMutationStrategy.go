package strategy

import (
	"math/rand"

	. "github.com/m-bednar/neuroevolution-thesis/src/neuralnet"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
)

type GaussMutationStrategy struct {
	rng      *rand.Rand
	strength float64
}

func NewGaussMutationStrategy(strength float64) *GaussMutationStrategy {
	var rng = NewTimeSeedRng()
	return &GaussMutationStrategy{rng, strength}
}

func (strategy *GaussMutationStrategy) MutateWeight(weight float64) float64 {
	var mutation = strategy.rng.NormFloat64() * (NN_WEIGHT_LIMIT / 2) * strategy.strength
	return weight + mutation
}

func (strategy *GaussMutationStrategy) Mutate(weights []float64) {
	for i, weight := range weights {
		weights[i] = ClampWeight(strategy.MutateWeight(weight))
	}
}
