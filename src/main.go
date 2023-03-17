package main

import "fmt"

const PRINT_EVERY_NTH_GEN = 10

func main() {
	var arguments = ParseProgramArguments()
	var tiles = ReadEnviromentFile(arguments.enviromentFile)

	// Setup
	var enviroment = NewEnviroment(tiles)
	var evaluationMap = NewEvaluationMap(enviroment)
	var evaluator = NewFitnessEvaluator(enviroment, evaluationMap)
	var renderer = NewRenderer(enviroment)
	var parentSelector = NewParentSelector(arguments.tournamentSize)
	var gatherer = NewStatsGatherer(enviroment, parentSelector)
	var collector = NewDataCollector(gatherer, arguments.maxGenerations)
	var mutationStrategy = NewGaussMutationStrategy(arguments.mutationStrength)
	var mutator = NewMutator(mutationStrategy)

	var outputter = NewOutputter(collector, renderer, arguments.captureModifier)
	var executor = NewTaskExecutor(enviroment, collector, evaluator, arguments.steps)

	// Factories and generators
	var crossoverStrategy = NewArithmeticCrossoverStrategy()
	var positionGenerator = NewSpawnSelector(enviroment)
	var neuralNetworkStructure = NewNeuralNetworkStructure(arguments.neuralNetworkScheme)
	var neuralNetworkFactory = NewNeuralNetworkFactory(neuralNetworkStructure, crossoverStrategy)
	var populationFactory = NewPopulationFactory(arguments.popSize, positionGenerator, neuralNetworkFactory, parentSelector)

	Loop(arguments, populationFactory, executor, collector, mutator)
	outputter.MakeOutput(arguments.outputPath)
	fmt.Println("Done.")
}

func Loop(args *ProgramArguments, populationFactory *PopulationFactory, executor *TaskExecutor, collector *DataCollector, mutator *Mutator) {
	var population = populationFactory.MakeRandom()
	var generation = 0
	for {
		fmt.Printf("Simulating %d/%d\n", generation, args.maxGenerations)
		executor.ExecuteTask(generation, population)
		collector.CollectStats(generation, population)
		if generation >= args.maxGenerations {
			return
		}

		population = populationFactory.ReproduceFrom(population)
		mutator.MutatePopulation(population)
		generation++
	}
}
