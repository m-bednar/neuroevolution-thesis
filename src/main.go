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
	// Setup
	var enviroment = NewEnviroment(tiles, ENV_SIZE)
	var positionGenerator = NewPositionGenerator(ENV_SIZE)
	var nnRandomFactory = NewNNRandomFactory()
	var nnReproductiveFactory = NewNNReproductionFactory()
	var evaluator = NewFitnessEvaluator(enviroment)
	var executor = NewTaskExecutor(enviroment, evaluator, STEPS)

	var mutator = NewMutator(0.25)
	var selector = NewSelector(enviroment)
	var population = NewPopulationRandomFactory(positionGenerator, nnRandomFactory).Make(POP_SIZE)
	var populationFactory = NewPopulationReproductiveFactory(positionGenerator, nnReproductiveFactory, selector)

	// Main loop
	var generation = 1
	var saved = 0
	for saved < POP_SIZE {
		executor.ExecuteTask(population)

		saved = selector.CountMicrobesInSafeZone(population)
		var averageFitness = selector.GetAverageFitness(population)
		fmt.Printf("%5d.  |  %3d/%d  |  %2.2f\n", generation, saved, POP_SIZE, averageFitness)

		population = populationFactory.Make(population, POP_SIZE)
		mutator.MutatePopulation(population)

		generation++
	}
}
