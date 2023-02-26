package main

import "math/rand"

type PopulationSelector struct {
	enviroment *Enviroment
	rng        *rand.Rand
}

func NewPopulationSelector(enviroment *Enviroment) *PopulationSelector {
	return &PopulationSelector{
		enviroment: enviroment,
		rng:        NewUnixTimeRng(),
	}
}

func (selector *PopulationSelector) IsMicrobeInsideSafeZone(microbe *Microbe) bool {
	return selector.enviroment.GetTile(microbe.position).IsSafeZone()
}

func (selector *PopulationSelector) CountMicrobesInSafeZone(population []*Microbe) int {
	var count = 0
	for _, microbe := range population {
		if selector.IsMicrobeInsideSafeZone(microbe) {
			count++
		}
	}
	return count
}

func (selector *PopulationSelector) GetAverageFitness(population []*Microbe) float64 {
	var sum = 0.0
	for _, microbe := range population {
		sum += microbe.fitness
	}
	return sum / float64(len(population))
}

func (selector *PopulationSelector) SelectOneRandomly(population []*Microbe) *Microbe {
	var max = len(population)
	var rnd = selector.rng.Intn(max)
	return population[rnd]
}

func (selector *PopulationSelector) SelectOneWithHighestFitness(population []*Microbe) *Microbe {
	var highest = population[0]
	for _, microbe := range population {
		if microbe.fitness > highest.fitness {
			highest = microbe
		}
	}
	return highest
}

func (selector *PopulationSelector) SelectOneByTournament(population []*Microbe, tournamentSize int) *Microbe {
	var selected = make([]*Microbe, tournamentSize)
	for i := 0; i < tournamentSize; i++ {
		selected[i] = selector.SelectOneRandomly(population)
	}
	return selector.SelectOneWithHighestFitness(selected)
}
