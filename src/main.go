package main

import (
	"fmt"
)

const POP_SIZE = 200
const ENV_SIZE = 15
const STEPS = ENV_SIZE * 2

var tiles = []TileType{
	Safe, Safe, Safe, Safe, None, None, None, None, None, None, None, None, None, None, None,
	Safe, Safe, Safe, Safe, None, None, None, None, None, None, None, None, None, None, None,
	Safe, Safe, Safe, Safe, None, None, None, None, None, None, None, None, None, None, None,
	Safe, Safe, Safe, Safe, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
}

func main() {
	var enviroment = NewEnviroment(tiles, ENV_SIZE)
	var positionGenerator = NewPositionGenerator(ENV_SIZE)
	var neuralNetworkRandomFactory = NewNeuralNetworkRandomFactory()
	var populationRandomFactory = NewPopulationRandomFactory(positionGenerator, neuralNetworkRandomFactory)
	var mutator = NewMutator(0.25)
	var selector = NewSelector(enviroment)
	var neuralNetworkReproductiveFactory = NewNeuralNetworkReproductionFactory()
	var populationReproductiveFactory = NewPopulationReproductiveFactory(positionGenerator, neuralNetworkReproductiveFactory, mutator, selector)
	var evaluator = NewFitnessEvaluator(enviroment)
	var executor = NewTaskExecutor(enviroment, evaluator, STEPS)

	MainLoop(executor, selector, populationRandomFactory, populationReproductiveFactory)
}

func MainLoop(executor *TaskExecutor, selector *Selector, populationRandomFactory *PopulationRandomFactory, populationReproductiveFactory *PopulationReproductiveFactory) {
	var population = populationRandomFactory.Make(POP_SIZE)
	var generation = 1
	var saved = 0

	for saved < POP_SIZE {
		executor.ExecuteTask(population)

		saved = selector.CountMicrobesInSafeZone(population)
		var averageFitness = selector.GetAverageFitness(population)
		fmt.Printf("%5d.  |  %3d/%d  |  %2.2f\n", generation, saved, POP_SIZE, averageFitness)

		population = populationReproductiveFactory.Make(population, POP_SIZE)
		generation++
	}
}
