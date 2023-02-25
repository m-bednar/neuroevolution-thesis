package main

import (
	"fmt"
	"math"
)

const POP_SIZE = 200
const ENV_SIZE = 15
const STEPS = ENV_SIZE * 2
const MIN_SUCCESSFULNESS = 0.99
const SELECTION_COEF = 1.45

func main() {
	var positionGenerator = NewPositionGenerator(ENV_SIZE)
	var neuralNetworkRandomFactory = NewNeuralNetworkRandomFactory()
	var populationRandomFactory = NewPopulationRandomFactory(positionGenerator, neuralNetworkRandomFactory)

	var mutator = NewMutator(1, 0.1)
	var neuralNetworkReproductiveFactory = NewNeuralNetworkReproductionFactory()
	var populationReproductiveFactory = NewPopulationReproductiveFactory(positionGenerator, neuralNetworkReproductiveFactory, mutator)

	var tiles = []TileType{
		SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
	}

	var enviroment = NewEnviroment(tiles, ENV_SIZE)
	var executor = NewTaskExecutor(enviroment, STEPS)
	var selector = NewPopulationSelector(enviroment)
	MainLoop(executor, selector, populationRandomFactory, populationReproductiveFactory)
}

func MainLoop(executor *TaskExecutor, selector *PopulationSelector, populationRandomFactory *PopulationRandomFactory, populationReproductiveFactory *PopulationReproductiveFactory) {
	var generation = 0
	var population = populationRandomFactory.Make(POP_SIZE)
	var successfulness = 0.0

	for successfulness < MIN_SUCCESSFULNESS {
		executor.ExecuteTask(population)

		var selected = selector.SelectFrom(population)
		successfulness = float64(len(selected)) / float64(POP_SIZE)

		var n = int(math.Min(float64(len(selected))*SELECTION_COEF, POP_SIZE))
		population = populationReproductiveFactory.Make(selected, n)
		population = append(population, populationRandomFactory.Make(POP_SIZE-n)...)

		fmt.Printf("%5d: %.2f%% (%d)\n", generation, successfulness*100.0, len(selected))
		generation++
	}
}
