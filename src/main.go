package main

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
	. "github.com/m-bednar/neuroevolution-thesis/src/neuralnet"
	. "github.com/m-bednar/neuroevolution-thesis/src/output"
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
	nnScheme := args.neuralNetworkScheme

	// Setup
	enviroment := NewEnviroment(tiles)
	evaluationMap := NewEvaluationMap(enviroment)
	evaluator := NewFitnessEvaluator(enviroment, evaluationMap)
	renderer := NewRenderer(enviroment)
	parentSelector := NewTournamentSelector(tournamentSize)
	gatherer := NewStatsGatherer(enviroment)
	collector := NewDataCollector(gatherer, maxGenerations, stepsCount, captureModifier)
	outputter := NewOutputter(collector, renderer)
	executor := NewTaskExecutor(enviroment, collector, stepsCount)
	spawnSelector := NewSpawnSelector(enviroment)
	neuralNetworkStructure := NewNeuralNetworkStructure(nnScheme)
	mutation := NewGaussMutation(mutationStrength)
	crossover := NewArithmeticCrossover()

	// Factories
	neuralNetworkFactory := NewNeuralNetworkFactory(neuralNetworkStructure, crossover)
	populationFactory := NewPopulationFactory(populationSize, spawnSelector, neuralNetworkFactory, parentSelector)

	// Run simulation
	simulationContext := NewSimulationContext(populationFactory, executor, evaluator, collector, mutation)
	simulationContext.Run(maxGenerations)
	outputter.MakeOutput(outputPath)
}
