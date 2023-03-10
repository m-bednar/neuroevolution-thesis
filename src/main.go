package main

const PRINT_EVERY_NTH_GEN = 10

func main() {
	var arguments = ParseProgramArguments()
	var tiles = ReadEnviromentFile(arguments.enviromentFile)

	// Setup
	var enviroment = NewEnviroment(tiles)
	var evaluator = NewFitnessEvaluator(enviroment)
	var renderer = NewRenderer(enviroment)
	var actionSelector = NewActionSelector()
	var parentSelector = NewParentSelector(arguments.tournamentSize)
	var gatherer = NewStatsGatherer(enviroment, parentSelector)
	var collector = NewStatsCollector(gatherer)
	var terminator = NewExecutionTerminator(gatherer, arguments)
	var mutator = NewMutator(NewGaussMutationStrategy(arguments.mutationStrength))

	var outputter = NewOutputter(arguments.outputPath, renderer)
	var executor = NewTaskExecutor(enviroment, outputter, evaluator, actionSelector, arguments.steps)

	// Factories and generators
	var positionGenerator = NewPositionGenerator(enviroment)
	var neuralNetworkFactory = NewNeuralNetworkFactory(NewArithmeticCrossoverStrategy())
	var populationFactory = NewPopulationFactory(arguments.popSize, positionGenerator, neuralNetworkFactory, parentSelector)

	Loop(populationFactory, executor, terminator, collector, mutator)
	// outputter.SaveAll()
	MakeChart(collector)
}

func Loop(populationFactory *PopulationFactory, executor *TaskExecutor, terminator *ExecutionTerminator, collector *StatsCollector, mutator *Mutator) {
	var population = populationFactory.MakeRandom()
	var generation = 0
	for {
		executor.ExecuteTask(generation, population)
		collector.Collect(population)
		if terminator.ShouldTerminate(generation, population) {
			return
		}

		population = populationFactory.ReproduceFrom(population)
		mutator.MutatePopulation(population)
		generation++
	}
}
