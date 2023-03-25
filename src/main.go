package main

import (
	"fmt"

	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
	. "github.com/m-bednar/neuroevolution-thesis/src/neuralnet"
	. "github.com/m-bednar/neuroevolution-thesis/src/output"
	. "github.com/m-bednar/neuroevolution-thesis/src/strategies"
)

func main() {
	args := ParseProgramArguments()
	tiles := ReadEnviromentFile(args.enviromentFile)

	// Arguments
	populationSize := args.populationSize
	tournamentSize := args.tournamentSize
	maxGenerations := args.maxGenerations
	stepsCount := args.stepsCount
	captureModifier := args.captureModifier
	mutationStrength := args.mutationStrength
	outputPath := args.outputPath
	neuralNetworkScheme := args.neuralNetworkScheme

	// Setup
	enviroment := NewEnviroment(tiles)
	evaluationMap := NewEvaluationMap(enviroment)
	evaluator := NewFitnessEvaluator(enviroment, evaluationMap)
	renderer := NewRenderer(enviroment)
	parentSelector := NewTournamentParentSelector(tournamentSize)
	gatherer := NewStatsGatherer(enviroment)
	collector := NewDataCollector(gatherer, maxGenerations, stepsCount, captureModifier)
	mutationStrategy := NewGaussMutationStrategy(mutationStrength)
	mutator := NewMutator(mutationStrategy)

	outputter := NewOutputter(collector, renderer)
	executor := NewTaskExecutor(enviroment, collector, evaluator, stepsCount)

	// Factories and generators
	crossoverStrategy := NewArithmeticCrossoverStrategy()
	spawnSelector := NewSpawnSelector(enviroment)
	neuralNetworkStructure := NewNeuralNetworkStructure(neuralNetworkScheme)
	neuralNetworkFactory := NewNeuralNetworkFactory(neuralNetworkStructure, crossoverStrategy)
	populationFactory := NewPopulationFactory(populationSize, spawnSelector, neuralNetworkFactory, parentSelector)

	Loop(maxGenerations, populationFactory, executor, mutator)
	outputter.MakeOutput(outputPath)
	fmt.Println("Done.")
}

func Loop(maxGenerations int, populationFactory *PopulationFactory, executor *TaskExecutor, mutator *Mutator) {
	population := populationFactory.MakeRandom()
	for generation := 0; generation <= maxGenerations; generation++ {
		fmt.Printf("Simulating %d/%d\n", generation, maxGenerations)
		executor.ExecuteTask(generation, population)
		population = populationFactory.ReproduceFrom(population)
		mutator.MutatePopulation(population)
	}
}
