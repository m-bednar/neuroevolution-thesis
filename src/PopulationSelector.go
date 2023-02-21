package main

import (
	"sort"
)

type PopulationSelector struct {
	evaluator Evaluator
}

func NewPopulationSelector(evaluator Evaluator) PopulationSelector {
	return PopulationSelector{evaluator}
}

func (selector *PopulationSelector) SelectFrom(population []*Microbe, percentage float64) []*Microbe {
	selector.evaluator.EvaluatePopulation(population)

	sort.Slice(population, func(i, j int) bool {
		return population[i].evaluation > population[j].evaluation
	})

	var selectSize = int(float64(len(population)) * percentage)
	var selected = population[:selectSize]

	return selected
}
