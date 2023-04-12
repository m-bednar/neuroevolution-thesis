package microbe

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
)

type Population []*Microbe

func (population Population) SelectOneWithHighestFitness() *Microbe {
	highest := population[0]
	for _, microbe := range population {
		if microbe.fitness > highest.fitness {
			highest = microbe
		}
	}
	return highest
}

func (population Population) CollectPositions() []Position {
	positions := make([]Position, len(population))
	for i, microbe := range population {
		positions[i] = microbe.position
	}
	return positions
}

func (population Population) CollectGenomes() [][]float64 {
	genomes := make([][]float64, len(population))
	for i, microbe := range population {
		weights := microbe.neuralNetwork.GetWeights()
		genomes[i] = make([]float64, len(weights))
		copy(genomes[i], weights)
	}
	return genomes
}

/*
Removes microbe at index *i* from the population.
Does not preserve order of microbes, does not reduce the capacity of population slice.
*/
func (population *Population) RemoveAt(i int) {
	size := len(*population)
	(*population)[i] = (*population)[size-1]
	(*population)[size-1] = nil
	*population = (*population)[:size-1]
}
