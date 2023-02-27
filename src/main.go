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
	var evaluator = NewFitnessEvaluator(enviroment)
	var executor = NewTaskExecutor(enviroment, evaluator, STEPS)
	var selector = NewSelector(enviroment)
	var mutator = NewMutator(0.25)

	// Factories and generators
	var positionGenerator = NewPositionGenerator(enviroment)
	var nnRandomFactory = NewNNRandomFactory()
	var nnReproductiveFactory = NewNNReproductionFactory()
	var populationRepFactory = NewPopulationReproductiveFactory(positionGenerator, nnReproductiveFactory, selector)
	var populationRndFactory = NewPopulationRandomFactory(positionGenerator, nnRandomFactory)

	// Main loop
	var firstPopulation = populationRndFactory.Make(POP_SIZE)
	MainLoop(firstPopulation, populationRepFactory, executor, selector, mutator)
}

func MainLoop(population []*Microbe, populationFactory *PopulationReproductiveFactory, executor *TaskExecutor, selector *Selector, mutator *Mutator) {
	var saved = 0
	var fitness = 0.0
	var generation = 1
	for saved < POP_SIZE {
		executor.ExecuteTask(population)

		// Print stats
		saved = selector.CountMicrobesInSafeZone(population)
		fitness = selector.GetAverageFitness(population)
		fmt.Printf("%5d.  |  %3d/%d  |  %2.2f\n", generation, saved, POP_SIZE, fitness)

		// Create new genration
		population = populationFactory.Make(population, POP_SIZE)
		mutator.MutatePopulation(population)
		generation++
	}
}
