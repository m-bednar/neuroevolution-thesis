package main

type Evaluator struct {
	enviroment Enviroment
}

func NewEvaluator(enviroment Enviroment) Evaluator {
	return Evaluator{
		enviroment: enviroment,
	}
}

func (evaluator *Evaluator) Evaluate(microbe Microbe) float64 {
	if evaluator.enviroment.GetTile(microbe.position).IsSafeZone() {
		return 1.0
	}
	return 0.0
}
