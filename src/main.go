package main

import (
	"fmt"
)

const PRINT_EVERY_NTH_GEN = 10

func main() {
	var arguments = ParseProgramArguments()
	var tiles = ReadEnviromentFile(arguments.enviromentFile)

	// Setup
	var enviroment = NewEnviroment(tiles)
	var evaluator = NewFitnessEvaluator(enviroment)
	var actionSelector = NewActionSelector()
	var executor = NewTaskExecutor(enviroment, evaluator, actionSelector, arguments.steps)
	var selector = NewParentSelector()
	var stats = NewStatsGatherer(enviroment, selector)
	var mutator = NewMutator(NewGaussMutationStrategy(arguments.mutationStrength))

	// Factories and generators
	var positionGenerator = NewPositionGenerator(enviroment)
	var nnRandomFactory = NewNNRandomFactory()
	var nnReproductionFactory = NewNNReproductionFactory(NewArithmeticCrossoverStrategy())
	var populationRepFactory = NewPopulationReproductionFactory(positionGenerator, nnReproductionFactory, selector)
	var populationRndFactory = NewPopulationRandomFactory(positionGenerator, nnRandomFactory)
	var firstPopulation = populationRndFactory.Make(arguments.popSize)

	// Main loop
	var population = firstPopulation
	var saved = 0
	var averageFitness = 0.0
	var highestFitness = 0.0
	var generation = 1
	for {
		// Execute task
		executor.ExecuteTask(population)

		// Print stats
		if (generation%PRINT_EVERY_NTH_GEN == 0) || saved == arguments.popSize {
			saved = stats.CountMicrobesInSafeZone(population)
			averageFitness = stats.GetAverageFitness(population)
			highestFitness = stats.GetHighestFitness(population)
			fmt.Printf("%5d.  |  %3d/%d  |  %2.2f |  %2.2f\n", generation, saved, arguments.popSize, averageFitness, highestFitness)
		}

		// Create new generation
		population = populationRepFactory.Make(population, arguments.popSize)
		mutator.MutatePopulation(population)
		generation++
	}
}
