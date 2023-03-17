package main

import "fmt"

const PRINT_EVERY_NTH_GEN = 10

func main() {
	var args = ParseProgramArguments()
	var tiles = ReadEnviromentFile(args.enviromentFile)

	// Arguments
	var populationSize = args.populationSize
	var tournamentSize = args.tournamentSize
	var maxGenerations = args.maxGenerations
	var stepsCount = args.stepsCount
	var captureModifier = args.captureModifier
	var mutationStrength = args.mutationStrength
	var outputPath = args.outputPath
	var neuralNetworkScheme = args.neuralNetworkScheme

	// Setup
	var enviroment = NewEnviroment(tiles)
	var evaluationMap = NewEvaluationMap(enviroment)
	var evaluator = NewFitnessEvaluator(enviroment, evaluationMap)
	var renderer = NewRenderer(enviroment)
	var parentSelector = NewParentSelector(tournamentSize)
	var gatherer = NewStatsGatherer(enviroment, parentSelector)
	var collector = NewDataCollector(gatherer, maxGenerations, stepsCount, captureModifier)
	var mutationStrategy = NewGaussMutationStrategy(mutationStrength)
	var mutator = NewMutator(mutationStrategy)

	var outputter = NewOutputter(collector, renderer)
	var executor = NewTaskExecutor(enviroment, collector, evaluator, stepsCount)

	// Factories and generators
	var crossoverStrategy = NewArithmeticCrossoverStrategy()
	var spawnSelector = NewSpawnSelector(enviroment)
	var neuralNetworkStructure = NewNeuralNetworkStructure(neuralNetworkScheme)
	var neuralNetworkFactory = NewNeuralNetworkFactory(neuralNetworkStructure, crossoverStrategy)
	var populationFactory = NewPopulationFactory(populationSize, spawnSelector, neuralNetworkFactory, parentSelector)

	Loop(args, populationFactory, executor, mutator)
	outputter.MakeOutput(outputPath)
	fmt.Println("Done.")
}

func Loop(args *ProgramArguments, populationFactory *PopulationFactory, executor *TaskExecutor, mutator *Mutator) {
	var population = populationFactory.MakeRandom()
	var generation = 0
	for {
		fmt.Printf("Simulating %d/%d\n", generation, args.maxGenerations)
		executor.ExecuteTask(generation, population)
		if generation >= args.maxGenerations {
			return
		}

		population = populationFactory.ReproduceFrom(population)
		mutator.MutatePopulation(population)
		generation++
	}
}
