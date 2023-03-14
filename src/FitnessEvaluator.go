package main

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

func (evaluator *FitnessEvaluator) EvaluateMove(origin Position, next Position) float64 {
	if !evaluator.enviroment.IsInsideBorders(next) {
		return OUT_OF_BORDER_MOVE_EVALUATION
	}
	var originEvaluation = evaluator.evaluationMap.GetEvaluation(origin)
	var nextEvaluation = evaluator.evaluationMap.GetEvaluation(next)
	var evaluation = nextEvaluation - originEvaluation
	return evaluation * MOVE_EVALUATION_FITNESS_COEF
}

func (evaluator *FitnessEvaluator) GetFinalEvaluation(microbe *Microbe) float64 {
	if evaluator.enviroment.GetTile(microbe.position).IsSafe() {
		return SAFEZONE_FINAL_REWARD
	}
	return 0
}
