package main

import (
	"fmt"
)

const POP_SIZE = 200
const ENV_SIZE = 15
const STEPS = ENV_SIZE * 2

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

func main() {
	var enviroment = NewEnviroment(tiles, ENV_SIZE)
	var positionGenerator = NewPositionGenerator(ENV_SIZE)
	var neuralNetworkRandomFactory = NewNeuralNetworkRandomFactory()
	var populationRandomFactory = NewPopulationRandomFactory(positionGenerator, neuralNetworkRandomFactory)

	var mutator = NewMutator(0, 0.0)
	var selector = NewPopulationSelector(enviroment)
	var neuralNetworkReproductiveFactory = NewNeuralNetworkReproductionFactory()
	var populationReproductiveFactory = NewPopulationReproductiveFactory(positionGenerator, neuralNetworkReproductiveFactory, mutator, selector)

	var evaluator = NewFitnessEvaluator(enviroment)
	var executor = NewTaskExecutor(enviroment, evaluator, STEPS)

	MainLoop(executor, selector, populationRandomFactory, populationReproductiveFactory)
}

func MainLoop(executor *TaskExecutor, selector *PopulationSelector, populationRandomFactory *PopulationRandomFactory, populationReproductiveFactory *PopulationReproductiveFactory) {
	var generation = 0
	var population = populationRandomFactory.Make(POP_SIZE)
	var safe = 0

	for safe < POP_SIZE {
		executor.ExecuteTask(population)

		safe = selector.CountMicrobesInSafeZone(population)
		population = populationReproductiveFactory.Make(population, POP_SIZE)

		fmt.Printf("%5d: %3d/%d\n", generation, safe, POP_SIZE)
		generation++
	}
}
