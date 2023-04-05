package main

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
)

type TournamentSelector struct {
	tournamentSize int
	rng            *Rng
}

func NewTournamentSelector(tournamentSize int) *TournamentSelector {
	return &TournamentSelector{
		tournamentSize: tournamentSize,
		rng:            NewTimeSeedRng(),
	}
}

func (selector *TournamentSelector) SelectRandomIndex(population Population) int {
	max := len(population)
	rnd := selector.rng.Intn(max)
	return rnd
}

func (selector *TournamentSelector) Select(population Population) *Microbe {
	selected := make(Population, selector.tournamentSize)
	remaining := make(Population, len(population))
	copy(remaining, population)

	for i := 0; i < selector.tournamentSize; i++ {
		index := selector.SelectRandomIndex(remaining)
		selected[i] = remaining[index]
		remaining.RemoveAt(index)
	}

	return selected.SelectOneWithHighestFitness()
}
