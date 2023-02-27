package main

import "math/rand"

type Selector struct {
	enviroment *Enviroment
	rng        *rand.Rand
}

func NewSelector(enviroment *Enviroment) *Selector {
	return &Selector{
		enviroment: enviroment,
		rng:        NewUnixTimeRng(),
	}
}

func (selector *Selector) IsMicrobeInsideSafeZone(microbe *Microbe) bool {
	return selector.enviroment.GetTile(microbe.position).IsSafe()
}

func (selector *Selector) CountMicrobesInSafeZone(population Population) int {
	var count = 0
	for _, microbe := range population {
		if selector.IsMicrobeInsideSafeZone(microbe) {
			count++
		}
	}
	return count
}

func (selector *Selector) GetAverageFitness(population Population) float64 {
	var sum = 0.0
	for _, microbe := range population {
		sum += microbe.fitness
	}
	return sum / float64(len(population))
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
