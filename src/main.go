package main

import "fmt"

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
	var collector = NewDataCollector(gatherer, arguments.maxGenerations)
	var terminator = NewExecutionTerminator(gatherer, arguments)
	var mutator = NewMutator(NewGaussMutationStrategy(arguments.mutationStrength))

	var outputter = NewOutputter(collector, renderer)
	var executor = NewTaskExecutor(enviroment, outputter, collector, evaluator, actionSelector, arguments.steps)

	// Factories and generators
	var crossoverStrategy = NewArithmeticCrossoverStrategy()
	var positionGenerator = NewPositionGenerator(enviroment)
	var neuralNetworkStructure = NewNeuralNetworkStructure(arguments.neuralNetworkScheme)
	var neuralNetworkFactory = NewNeuralNetworkFactory(neuralNetworkStructure, crossoverStrategy)
	var populationFactory = NewPopulationFactory(arguments.popSize, positionGenerator, neuralNetworkFactory, parentSelector)

	Loop(populationFactory, executor, terminator, collector, mutator)
	outputter.MakeOutput(arguments.outputPath, arguments.captureModifier)
	fmt.Println("Done.")
}

func Loop(populationFactory *PopulationFactory, executor *TaskExecutor, terminator *ExecutionTerminator, collector *DataCollector, mutator *Mutator) {
	var population = populationFactory.MakeRandom()
	var generation = 0
	for {
		fmt.Printf("Simulating %d/%d\n", generation, terminator.arguments.maxGenerations)
		executor.ExecuteTask(generation, population)
		collector.CollectStats(generation, population)
		if terminator.ShouldTerminate(generation, population) {
			return
		}

		population = populationFactory.ReproduceFrom(population)
		mutator.MutatePopulation(population)
		generation++
	}
}
