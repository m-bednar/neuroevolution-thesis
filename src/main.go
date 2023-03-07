package main

const PRINT_EVERY_NTH_GEN = 10

func main() {
	var arguments = ParseProgramArguments()
	var tiles = ReadEnviromentFile(arguments.enviromentFile)

	// Setup
	var enviroment = NewEnviroment(tiles)
	var evaluator = NewFitnessEvaluator(enviroment)
	var actionSelector = NewActionSelector()
	var executor = NewTaskExecutor(enviroment, evaluator, actionSelector, arguments.steps)
	var selector = NewParentSelector(arguments.tournamentSize)
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
	var generation = 1
	for {
		// Execute task
		executor.ExecuteTask(population)

		// TODO: Output generation status (eg. hash values of microbes)

		var safe = stats.CountMicrobesInSafeZone(population)
		var successRate = float64(safe) / float64(arguments.popSize)

		if successRate >= arguments.minSuccessRate {
			break
		}
		if generation >= arguments.maxGenerations {
			break
		}

		// TODO: Output generation stats (eg. avg/max fitness, success rate)

		// Create new generation
		population = populationRepFactory.Make(population, arguments.popSize)
		mutator.MutatePopulation(population)
		generation++
	}
}
