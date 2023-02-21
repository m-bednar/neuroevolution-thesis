package main

type Evaluator struct {
	enviroment Enviroment
}

func NewEvaluator(enviroment Enviroment) Evaluator {
	return Evaluator{
		enviroment: enviroment,
	}
}

func (evaluator *Evaluator) DidMicrobeStartedAtSafeZone(microbe *Microbe) bool {
	return evaluator.enviroment.GetTile(microbe.startPosition).IsSafeZone()
}

func (evaluator *Evaluator) IsMicrobeAtSafeZone(microbe *Microbe) bool {
	return evaluator.enviroment.GetTile(microbe.position).IsSafeZone()
}

func (evaluator *Evaluator) Evaluate(microbe *Microbe) {
	var evaluation = 0.0
	/*
		if evaluator.DidMicrobeStartedAtSafeZone(microbe) {
			evaluation -= 0.5
		}*/
	if evaluator.IsMicrobeAtSafeZone(microbe) {
		evaluation += 1.0
	}
	microbe.evaluation = evaluation
}

func (evaluator *Evaluator) EvaluatePopulation(population []*Microbe) {
	for i := range population {
		evaluator.Evaluate(population[i])
	}
}

func (evaluator *Evaluator) GetNumberOfMicrobesAtSafeZone(population []*Microbe) int {
	var count = 0
	for i := range population {
		if evaluator.IsMicrobeAtSafeZone(population[i]) {
			count++
		}
	}
	return count
}
