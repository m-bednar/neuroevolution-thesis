package main

import (
	"fmt"

	astar "github.com/fzipp/astar"
)

type EvaluationMap struct {
	evaluations []float64
	enviroment  *Enviroment
}

func CreateEvaluations(enviroment *Enviroment) []float64 {
	var evaluations = make([]float64, len(enviroment.tiles))
	var safeTiles = enviroment.GetAllTilesOfType(Safe)
	var passableTiles = enviroment.GetAllPassableTiles()
	var distanceTo = Position.DistanceTo

	for _, safeTile := range safeTiles {
		var start = safeTile
		for _, emptyTile := range passableTiles {
			var end = emptyTile
			var path = astar.FindPath[Position](enviroment, start, end, distanceTo, distanceTo)
			if path != nil {
				var index = enviroment.GetTileIndex(end)
				var prev = evaluations[index]
				var curr = float64(len(path))
				if prev == 0 || curr < prev {
					evaluations[index] = curr
				}
			}
		}
	}

	// Distance -> evaluation
	for i, evaluation := range evaluations {
		if evaluation != 0 {
			evaluations[i] = 1.0 / evaluation
		}
	}

	return evaluations
}

func NewEvaluationMap(enviroment *Enviroment) *EvaluationMap {
	return &EvaluationMap{
		evaluations: CreateEvaluations(enviroment),
		enviroment:  enviroment,
	}
}

func (evaluationMap *EvaluationMap) Print() {
	for y := 0; y < evaluationMap.enviroment.size; y++ {
		for x := 0; x < evaluationMap.enviroment.size; x++ {
			fmt.Printf("%5.2f ", evaluationMap.GetEvaluation(NewPosition(x, y)))
		}
		fmt.Println()
	}
}

func (evaluationMap *EvaluationMap) GetEvaluation(position Position) float64 {
	var size = evaluationMap.enviroment.size
	var i = position.y*size + position.x
	return evaluationMap.evaluations[i]
}
