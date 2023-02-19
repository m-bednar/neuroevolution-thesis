package main

type PopulationSelector struct {
	evaluator Evaluator
}

func NewPopulationSelector(evaluator Evaluator) PopulationSelector {
	return PopulationSelector{evaluator}
}

func (selector *PopulationSelector) SelectFrom(population []Microbe) []*Microbe {
	var size = len(population)
	var selected = make([]*Microbe, 0, size)
	for i := 0; i < size; i++ {
		if selector.evaluator.Evaluate(population[i]) == 1.0 {
			selected = append(selected, &population[i])
		}
	}
	return selected
}
