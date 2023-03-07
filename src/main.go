package main

import (
	"fmt"
)

const POP_SIZE = 400
const ENV_SIZE = 15
const STEPS = ENV_SIZE * 2
const MUTATION_STRENGTH = 0.2
const PRINT_EVERY_NTH_GEN = 10

// TODO: Read from given file
var tiles = []TileType{
	None, Safe, Safe, Safe, None, None, None, None, None, None, None, None, None, None, None,
	None, None, Safe, None, None, None, None, None, None, None, None, None, None, None, None,
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
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
	None, None, None, None, None, None, None, None, None, None, None, None, None, None, None,
}

func main() {
	// Setup
	var enviroment = NewEnviroment(tiles, ENV_SIZE)
	var evaluator = NewFitnessEvaluator(enviroment)
	var actionSelector = NewActionSelector()
	var executor = NewTaskExecutor(enviroment, evaluator, actionSelector, STEPS)
	var selector = NewParentSelector()
	var stats = NewStatsGatherer(enviroment, selector)
	var mutator = NewMutator(NewGaussMutationStrategy(MUTATION_STRENGTH))

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
	var averageFitness = 0.0
	var highestFitness = 0.0
	var generation = 1
	for {
		// Execute task
		executor.ExecuteTask(population)

		// Print stats
		if (generation%PRINT_EVERY_NTH_GEN == 0) || saved == POP_SIZE {
			saved = stats.CountMicrobesInSafeZone(population)
			averageFitness = stats.GetAverageFitness(population)
			highestFitness = stats.GetHighestFitness(population)
			fmt.Printf("%5d.  |  %3d/%d  |  %2.2f |  %2.2f\n", generation, saved, POP_SIZE, averageFitness, highestFitness)
		}

		// Create new generation
		population = populationFactory.Make(population, POP_SIZE)
		mutator.MutatePopulation(population)
		generation++
	}
}
