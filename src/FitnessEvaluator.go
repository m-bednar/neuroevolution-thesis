/**
 * @project neuroevolution
 * @file    FitnessEvaluator.go
 * @author  Martin Bednář (xbedna77@fit.vut.cz)
 */

package main

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
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
		evaluation := evaluator.evaluationMap.GetEvaluation(microbe.GetPosition())
		microbe.SetFitness(evaluation)
	}
}
