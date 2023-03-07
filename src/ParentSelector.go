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

func (selector *ParentSelector) SelectOneRandomly(population []*Microbe) *Microbe {
	var max = len(population)
	var rnd = selector.rng.Intn(max)
	return population[rnd]
}

func (selector *ParentSelector) SelectOneWithHighestFitness(population []*Microbe) *Microbe {
	var highest = population[0]
	for _, microbe := range population {
		if microbe.fitness > highest.fitness {
			highest = microbe
		}
	}
	return highest
}

func (selector *ParentSelector) SelectOneByTournament(population []*Microbe) *Microbe {
	var selected = make([]*Microbe, selector.tournamentSize)
	for i := 0; i < selector.tournamentSize; i++ {
		selected[i] = selector.SelectOneRandomly(population)
	}
	return selector.SelectOneWithHighestFitness(selected)
}
