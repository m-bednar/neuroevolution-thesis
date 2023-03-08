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
	var capturer = NewVideoCapturer(arguments.videoOutputPath, renderer)
	var executor = NewTaskExecutor(enviroment, capturer, evaluator, actionSelector, arguments.steps)
	var selector = NewParentSelector(arguments.tournamentSize)
	var stats = NewStatsGatherer(enviroment, selector)
	var mutator = NewMutator(NewGaussMutationStrategy(arguments.mutationStrength))

	// Factories and generators
	var positionGenerator = NewPositionGenerator(enviroment)
	var neuralNetworkFactory = NewNeuralNetworkFactory(NewArithmeticCrossoverStrategy())
	var populationFactory = NewPopulationReproductionFactory(positionGenerator, neuralNetworkFactory, selector)

	// Main loop
	var population = populationFactory.MakeRandom(arguments.popSize)
	var generation = 0
	for {
		executor.ExecuteTask(generation, population)
		// fmt.Println(generation)

		var successRate = stats.GetSuccessRate(population)
		if successRate >= arguments.minSuccessRate {
			break
		}
		if generation >= arguments.maxGenerations {
			break
		}

		// Create new generation
		population = populationFactory.ReproduceFrom(population, arguments.popSize)
		mutator.MutatePopulation(population)
		generation++
	}

	capturer.SaveAndClose()
}
