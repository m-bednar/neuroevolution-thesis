package main

type Population []*Microbe

func (population Population) SelectOneWithHighestFitness() *Microbe {
	var highest = population[0]
	for _, microbe := range population {
		if microbe.fitness > highest.fitness {
			highest = microbe
		}
	}
	return highest
}

func (population Population) CollectPositions() []Position {
	var positions = make([]Position, len(population))
	for i, microbe := range population {
		positions[i] = microbe.position
	}
	return positions
}
