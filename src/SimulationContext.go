/**
 * @project neuroevolution
 * @file    SimulationContext.go
 * @author  Martin Bednář (xbedna77@fit.vut.cz)
 */

package main

import (
	"fmt"

	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
	. "github.com/m-bednar/neuroevolution-thesis/src/output"
)

type SimulationContext struct {
	populationFactory *PopulationFactory
	executor          *TaskExecutor
	evaluator         *FitnessEvaluator
	collector         *DataCollector
	mutation          *GaussMutation
}

func NewSimulationContext(populationFactory *PopulationFactory, executor *TaskExecutor,
	evaluator *FitnessEvaluator, collector *DataCollector, mutation *GaussMutation) *SimulationContext {
	return &SimulationContext{populationFactory, executor, evaluator, collector, mutation}
}

func (context *SimulationContext) Run(maxGenerations int) {
	population := context.populationFactory.MakeRandom()
	for generation := 0; generation <= maxGenerations; generation++ {
		fmt.Printf("Simulating %d/%d\n", generation, maxGenerations)
		context.executor.ExecuteTask(generation, population)
		context.evaluator.Evaluate(population)
		context.collector.CollectGeneration(generation, population)
		population = context.populationFactory.ReproduceFrom(population)
		context.mutation.MutatePopulation(population)
	}
}
