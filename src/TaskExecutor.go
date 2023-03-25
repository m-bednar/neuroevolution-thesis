package main

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
	. "github.com/m-bednar/neuroevolution-thesis/src/output"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
)

type TaskExecutor struct {
	enviroment  *Enviroment
	collector   *DataCollector
	evaluator   *FitnessEvaluator
	selector    *ActionSelector
	inputsMaker *NeuralInputsMaker
	steps       int
}

func NewTaskExecutor(enviroment *Enviroment, collector *DataCollector, evaluator *FitnessEvaluator, steps int) *TaskExecutor {
	selector := NewActionSelector()
	inputsMaker := NewNeuralInputsMaker(enviroment)
	return &TaskExecutor{enviroment, collector, evaluator, selector, inputsMaker, steps}
}

func (executor *TaskExecutor) ExecuteTask(generation int, population Population) {
	for i := 0; i < executor.steps; i++ {
		executor.ExecuteStep(population)
		executor.collector.CollectStep(i, population)
	}
	executor.evaluator.Evaluate(population)
	executor.collector.CollectGeneration(generation, population)
}

func (executor *TaskExecutor) ExecuteStep(population Population) {
	ConcurrentLoop(population, func(_ int, microbe *Microbe) {
		inputs := executor.inputsMaker.MakeInputsFor(microbe)
		output := microbe.Process(inputs)
		action := executor.selector.SelectMoveAction(output)
		result := microbe.GetPosition().AddToDirection(action)
		if executor.enviroment.IsPassable(result) {
			microbe.MoveTo(result)
		}
	})
}
