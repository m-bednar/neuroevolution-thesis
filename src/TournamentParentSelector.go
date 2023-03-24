package main

import (
	"math/rand"

	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
)

type TournamentParentSelector struct {
	tournamentSize int
	rng            *rand.Rand
}

func NewTournamentParentSelector(tournamentSize int) *TournamentParentSelector {
	return &TournamentParentSelector{
		tournamentSize: tournamentSize,
		rng:            NewTimeSeedRng(),
	}
}

func (selector *TournamentParentSelector) SelectOneRandomly(population Population) *Microbe {
	var max = len(population)
	var rnd = selector.rng.Intn(max)
	return population[rnd]
}

func (selector *TournamentParentSelector) SelectParent(population Population) *Microbe {
	var selected = make(Population, selector.tournamentSize)
	for i := 0; i < selector.tournamentSize; i++ {
		selected[i] = selector.SelectOneRandomly(population)
	}
	return selected.SelectOneWithHighestFitness()
}
