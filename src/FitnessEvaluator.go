package main

const MOVE_EVALUATION_FITNESS_COEF = 0.25

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
	var distance = position.DistanceTo(closest)
	for x := 0; x < evaluator.enviroment.size; x++ {
		for y := 0; y < evaluator.enviroment.size; y++ {
			var selected = NewPosition(x, y)
			if evaluator.enviroment.GetTile(selected).IsSafeZone() {
				var distanceToSelected = position.DistanceTo(selected)
				if distanceToSelected < distance {
					closest = selected
					distance = distanceToSelected
				}
			}
		}
	}
	return closest
}

func (evaluator *FitnessEvaluator) EvaluateMove(origin Position, next Position) float64 {
	var closestTileOrigin = evaluator.FindClosestSafeTile(origin)
	var closestTileNext = evaluator.FindClosestSafeTile(origin)
	var value = closestTileOrigin.DistanceTo(origin) - closestTileNext.DistanceTo(next)
	return value * MOVE_EVALUATION_FITNESS_COEF
}
