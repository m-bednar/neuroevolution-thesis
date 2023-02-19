package main

import "fmt"

type Evaluator struct {
	executor   Executor
	enviroment Enviroment
}

func NewEvaluator(executor Executor, enviroment Enviroment) Evaluator {
	return Evaluator{
		executor:   executor,
		enviroment: enviroment,
	}
}

func (evaluator *Evaluator) Evaluate(population []Microbe) {
	evaluator.executor.Execute(population)
	var evaluations = evaluator.GetEvaluations(population)
	fmt.Println(evaluations) // TODO
}

func (evaluator *Evaluator) GetEvaluations(population []Microbe) []float64 {
	var evaluations = make([]float64, len(population))
	for i, microbe := range population {
		evaluations[i] = evaluator.GetEvaluation(microbe)
	}
	return evaluations
}

func (evaluator *Evaluator) GetEvaluation(microbe Microbe) float64 {
	if evaluator.enviroment.GetTile(microbe.position).IsSafeZone() {
		return 1.0
	}
	return 0.0
}
