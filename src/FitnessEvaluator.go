package main

const MOVE_EVALUATION_FITNESS_COEF = 0.5

type FitnessEvaluator struct {
	enviroment *Enviroment
}

func NewFitnessEvaluator(enviroment *Enviroment) *FitnessEvaluator {
	return &FitnessEvaluator{
		enviroment: enviroment,
	}
}

func (evaluator *FitnessEvaluator) FindClosestSafeTile(position Position) Position {
	var closest = NewPosition(0, 0)
	for x := 0; x < evaluator.enviroment.size; x++ {
		for y := 0; y < evaluator.enviroment.size; y++ {
			var selected = NewPosition(x, y)
			if evaluator.enviroment.GetTile(selected).IsSafeZone() {
				if position.DistanceTo(selected) < position.DistanceTo(closest) {
					closest = selected
				}
			}
		}
	}
	return closest
}

func (evaluator *FitnessEvaluator) EvaluateMove(origin Position, next Position) float64 {
	var closestTileOrigin = evaluator.FindClosestSafeTile(origin)
	var closestTileNext = evaluator.FindClosestSafeTile(next)
	var value = closestTileOrigin.DistanceTo(origin) - closestTileNext.DistanceTo(next)
	return value * MOVE_EVALUATION_FITNESS_COEF
}
