package main

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/env"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
)

const (
	SAFEZONE_FINAL_REWARD = 1
)

type FitnessEvaluator struct {
	enviroment    *Enviroment
	evaluationMap *EvaluationMap
}

func NewFitnessEvaluator(enviroment *Enviroment, evaluationMap *EvaluationMap) *FitnessEvaluator {
	return &FitnessEvaluator{
		enviroment:    enviroment,
		evaluationMap: evaluationMap,
	}
}

func (evaluator *FitnessEvaluator) Evaluate(population Population) {
	for _, microbe := range population {
		var evaluation = evaluator.evaluationMap.GetEvaluation(microbe.GetPosition())
		microbe.SetFitness(evaluation)
	}
}
