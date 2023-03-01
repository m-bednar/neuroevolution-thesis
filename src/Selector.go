package main

import "math/rand"

const TOURNAMENT_SIZE = 4

type Selector struct {
	rng *rand.Rand
}

func NewSelector() *Selector {
	return &Selector{
		rng: NewUnixTimeRng(),
	}
}

func (selector *Selector) SelectOneRandomly(population []*Microbe) *Microbe {
	var max = len(population)
	var rnd = selector.rng.Intn(max)
	return population[rnd]
}

func (selector *Selector) SelectOneWithHighestFitness(population []*Microbe) *Microbe {
	var highest = population[0]
	for _, microbe := range population {
		if microbe.fitness > highest.fitness {
			highest = microbe
		}
	}
	return highest
}

func (selector *Selector) SelectOneByTournament(population []*Microbe) *Microbe {
	var selected = make([]*Microbe, TOURNAMENT_SIZE)
	for i := 0; i < TOURNAMENT_SIZE; i++ {
		selected[i] = selector.SelectOneRandomly(population)
	}
	return selector.SelectOneWithHighestFitness(selected)
}
