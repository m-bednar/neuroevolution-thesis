package main

import "math/rand"

const TOURNAMENT_SIZE = 6

type ParentSelector struct {
	rng *rand.Rand
}

func NewParentSelector() *ParentSelector {
	return &ParentSelector{
		rng: NewUnixTimeRng(),
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
	var selected = make([]*Microbe, TOURNAMENT_SIZE)
	for i := 0; i < TOURNAMENT_SIZE; i++ {
		selected[i] = selector.SelectOneRandomly(population)
	}
	return selector.SelectOneWithHighestFitness(selected)
}
