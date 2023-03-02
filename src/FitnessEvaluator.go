package main

const NON_PASSABLE_PENALTY = -0.1
const SAFEZONE_FINAL_REWARD = 10.0
const MOVE_EVALUATION_FITNESS_COEF = 0.4

type FitnessEvaluator struct {
	enviroment *Enviroment
}

func NewFitnessEvaluator(enviroment *Enviroment) *FitnessEvaluator {
	return &FitnessEvaluator{
		enviroment: enviroment,
	}
}

func (evaluator *FitnessEvaluator) FindClosestSafeTile(position Position) Position {
	/* TODO: Uncomment
	var closest = NewPosition(0, 0)
	for x := 0; x < evaluator.enviroment.size; x++ {
		for y := 0; y < evaluator.enviroment.size; y++ {
			var selected = NewPosition(x, y)
			if evaluator.enviroment.GetTile(selected).IsSafe() {
				if position.DistanceTo(selected) < position.DistanceTo(closest) {
					closest = selected
				}
			}
		}
	}
	*/
	return NewPosition(2, 0)
}

func (evaluator *FitnessEvaluator) EvaluateMove(origin Position, next Position) float64 {
	if !evaluator.enviroment.IsPassable(next) {
		return NON_PASSABLE_PENALTY
	}

	var closestTileOrigin = evaluator.FindClosestSafeTile(origin)
	var closestTileNext = evaluator.FindClosestSafeTile(next)
	var value = closestTileOrigin.DistanceTo(origin) - closestTileNext.DistanceTo(next)
	return value * MOVE_EVALUATION_FITNESS_COEF
}

func (evaluator *FitnessEvaluator) GetFinalEvaluation(microbe *Microbe) float64 {
	if evaluator.enviroment.GetTile(microbe.position).IsSafe() {
		return SAFEZONE_FINAL_REWARD
	}
	return 0
}
