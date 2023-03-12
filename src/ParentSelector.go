package main

import "math/rand"

type ParentSelector struct {
	tournamentSize int
	rng            *rand.Rand
}

func NewParentSelector(tournamentSize int) *ParentSelector {
	return &ParentSelector{
		tournamentSize: tournamentSize,
		rng:            NewUnixTimeRng(),
	}
}

func (selector *ParentSelector) SelectOneRandomly(population Population) *Microbe {
	var max = len(population)
	var rnd = selector.rng.Intn(max)
	return population[rnd]
}

func (selector *ParentSelector) SelectOneByTournament(population Population) *Microbe {
	var selected = make(Population, selector.tournamentSize)
	for i := 0; i < selector.tournamentSize; i++ {
		selected[i] = selector.SelectOneRandomly(population)
	}
	return selected.SelectOneWithHighestFitness()
}
