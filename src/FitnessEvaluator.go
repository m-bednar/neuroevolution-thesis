package main

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

func (evaluator *FitnessEvaluator) GetEvaluation(microbe *Microbe) float64 {
	return evaluator.evaluationMap.GetEvaluation(microbe.position)
}
