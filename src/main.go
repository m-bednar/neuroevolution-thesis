package main

const PRINT_EVERY_NTH_GEN = 10

func main() {
	var arguments = ParseProgramArguments()
	var tiles = ReadEnviromentFile(arguments.enviromentFile)

	// Setup
	var enviroment = NewEnviroment(tiles)
	var evaluator = NewFitnessEvaluator(enviroment)
	var actionSelector = NewActionSelector()
	var renderer = NewRenderer(enviroment)
	var capturer = NewVideoCapturer("out.avi", renderer)
	var executor = NewTaskExecutor(enviroment, capturer, evaluator, actionSelector, arguments.steps)
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
		executor.ExecuteTask(generation, population)

		var safe = stats.CountMicrobesInSafeZone(population)
		var successRate = float64(safe) / float64(arguments.popSize)
		if successRate >= arguments.minSuccessRate {
			break
		}
		if generation >= arguments.maxGenerations {
			break
		}

		// Create new generation
		population = populationRepFactory.Make(population, arguments.popSize)
		mutator.MutatePopulation(population)
		generation++
	}

	capturer.SaveAndClose()
}
