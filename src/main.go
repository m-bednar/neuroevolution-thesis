package main

import (
	"fmt"
)

const POP_SIZE = 200
const ENV_SIZE = 15
const STEPS = ENV_SIZE * 2

// TODO: Read from given file
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
	var stats = NewStatsGatherer(enviroment)
	var mutator = NewMutator(NewGaussMutationStrategy(0.25))
	var selector = NewSelector()

	// Factories and generators
	var positionGenerator = NewPositionGenerator(enviroment)
	var nnRandomFactory = NewNNRandomFactory()
	var nnReproductionFactory = NewNNReproductionFactory(NewArithmeticCrossoverStrategy())
	var populationRepFactory = NewPopulationReproductionFactory(positionGenerator, nnReproductionFactory, selector)
	var populationRndFactory = NewPopulationRandomFactory(positionGenerator, nnRandomFactory)

	// Main loop
	var firstPopulation = populationRndFactory.Make(POP_SIZE)
	MainLoop(firstPopulation, populationRepFactory, executor, stats, mutator)
}

func MainLoop(population []*Microbe, populationFactory *PopulationReproductionFactory, executor *TaskExecutor, stats *StatsGatherer, mutator *Mutator) {
	var saved = 0
	var fitness = 0.0
	var generation = 1
	for saved < POP_SIZE {
		// Execute task
		executor.ExecuteTask(population)

		// Print stats
		saved = stats.CountMicrobesInSafeZone(population)
		fitness = stats.GetAverageFitness(population)
		fmt.Printf("%5d.  |  %3d/%d  |  %2.2f\n", generation, saved, POP_SIZE, fitness)

		// Create new generation
		population = populationFactory.Make(population, POP_SIZE)
		mutator.MutatePopulation(population)
		generation++
	}
}
