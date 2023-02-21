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

func (evaluator *Evaluator) Evaluate(microbe *Microbe) float64 {
	if evaluator.IsMicrobeAtSafeZone(microbe) {
		return 1.0
	}
	return 0.0
}

func (evaluator *Evaluator) EvaluatePopulation(population []*Microbe) {
	for i := range population {
		population[i].fitness = evaluator.Evaluate(population[i])
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
