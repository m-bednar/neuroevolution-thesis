package main

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
	. "github.com/m-bednar/neuroevolution-thesis/src/neuralnet"
	. "github.com/m-bednar/neuroevolution-thesis/src/output"
)

func main() {
	// Inputs
	arguments := ParseProgramArguments()
	tiles := ReadEnviromentFile(arguments.enviromentFile)

	// Arguments
	populationSize := arguments.populationSize
	tournamentSize := arguments.tournamentSize
	maxGenerations := arguments.maxGenerations
	stepsCount := arguments.stepsCount
	captureModifier := arguments.captureModifier
	mutationStrength := arguments.mutationStrength
	outputPath := arguments.outputPath
	nnScheme := arguments.neuralNetworkScheme

	// Neuroevolution setup
	enviroment := NewEnviroment(tiles)
	evaluationMap := NewEvaluationMap(enviroment)
	evaluator := NewFitnessEvaluator(enviroment, evaluationMap)
	parentSelector := NewTournamentSelector(tournamentSize)
	neuralNetworkStructure := NewNeuralNetworkStructure(nnScheme)
	mutation := NewGaussMutation(mutationStrength)
	crossover := NewArithmeticCrossover()

	// Task and output setup
	spawnSelector := NewSpawnSelector(enviroment)
	renderer := NewRenderer(enviroment)
	gatherer := NewStatsGatherer(enviroment)
	collector := NewDataCollector(gatherer, maxGenerations, stepsCount, captureModifier)
	outputter := NewOutputter(collector, renderer)
	executor := NewTaskExecutor(enviroment, collector, stepsCount)

	// Factories
	neuralNetworkFactory := NewNeuralNetworkFactory(neuralNetworkStructure, crossover)
	populationFactory := NewPopulationFactory(populationSize, spawnSelector, neuralNetworkFactory, parentSelector)

	// Run simulation
	simulationContext := NewSimulationContext(populationFactory, executor, evaluator, collector, mutation)
	simulationContext.Run(maxGenerations)
	outputter.MakeOutput(outputPath)
}
