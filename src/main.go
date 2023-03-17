package main

import "fmt"

const PRINT_EVERY_NTH_GEN = 10

func main() {
	var args = ParseProgramArguments()
	var tiles = ReadEnviromentFile(args.enviromentFile)

	// Setup
	var enviroment = NewEnviroment(tiles)
	var evaluationMap = NewEvaluationMap(enviroment)
	var evaluator = NewFitnessEvaluator(enviroment, evaluationMap)
	var renderer = NewRenderer(enviroment)
	var parentSelector = NewParentSelector(args.tournamentSize)
	var gatherer = NewStatsGatherer(enviroment, parentSelector)
	var collector = NewDataCollector(gatherer, args.maxGenerations, args.steps, args.captureModifier)
	var mutationStrategy = NewGaussMutationStrategy(args.mutationStrength)
	var mutator = NewMutator(mutationStrategy)

	var outputter = NewOutputter(collector, renderer, args.captureModifier)
	var executor = NewTaskExecutor(enviroment, collector, evaluator, args.steps)

	// Factories and generators
	var crossoverStrategy = NewArithmeticCrossoverStrategy()
	var positionGenerator = NewSpawnSelector(enviroment)
	var neuralNetworkStructure = NewNeuralNetworkStructure(args.neuralNetworkScheme)
	var neuralNetworkFactory = NewNeuralNetworkFactory(neuralNetworkStructure, crossoverStrategy)
	var populationFactory = NewPopulationFactory(args.popSize, positionGenerator, neuralNetworkFactory, parentSelector)

	Loop(args, populationFactory, executor, mutator)
	outputter.MakeOutput(args.outputPath)
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
