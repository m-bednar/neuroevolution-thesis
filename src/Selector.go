package main

import "math/rand"

type Selector struct {
	rng *rand.Rand
}

func NewSelector() *Selector {
	return &Selector{
		rng: NewUnixTimeRng(),
	}
}

func (selector *Selector) SelectOneRandomly(population Population) *Microbe {
	var max = len(population)
	var rnd = selector.rng.Intn(max)
	return population[rnd]
}

func (selector *Selector) SelectOneWithHighestFitness(population Population) *Microbe {
	var highest = population[0]
	for _, microbe := range population {
		if microbe.fitness > highest.fitness {
			highest = microbe
		}
	}
	return highest
}

func (selector *Selector) SelectOneByTournament(population Population, tournamentSize int) *Microbe {
	var selected = make([]*Microbe, tournamentSize)
	for i := 0; i < tournamentSize; i++ {
		selected[i] = selector.SelectOneRandomly(population)
	}
	return selector.SelectOneWithHighestFitness(selected)
}
