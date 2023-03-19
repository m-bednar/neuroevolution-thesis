package main

import (
	"fmt"
	"math"

	as "github.com/fzipp/astar"
)

type EvaluationMap struct {
	evaluations []float64
	enviroment  *Enviroment
}

func EvaluateDistance(distance int) float64 {
	return 1.0 / math.Sqrt(float64(distance))
}

func CreateEvaluations(enviroment *Enviroment) []float64 {
	var evaluations = make([]float64, len(enviroment.tiles))
	var distances = make([]int, len(enviroment.tiles))
	var safeTiles = enviroment.GetAllTilesOfType(Safe)
	var passableTiles = enviroment.GetAllPassableTiles()
	var distFunc = Position.DistanceTo
	var findPath = as.FindPath[Position]

	// Evaluate each passable tile by finding
	// it's shortest path to closest safe tile
	for _, safeTile := range safeTiles {
		var start = safeTile
		for _, passableTile := range passableTiles {
			var end = passableTile
			var path = findPath(enviroment, start, end, distFunc, distFunc)
			if path != nil {
				var index = enviroment.GetTileIndex(end)
				var prev = distances[index]
				var curr = len(path)
				if prev == 0 || curr < prev {
					distances[index] = curr
				}
			}
		}
	}

	// Transform distances to evaluations
	for i, distance := range distances {
		if distance != 0 {
			evaluations[i] = EvaluateDistance(distance)
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
			var evaluation = evaluationMap.GetEvaluation(NewPosition(x, y))
			if evaluation == 0.0 {
				fmt.Printf("xxxx ")
			} else {
				fmt.Printf("%4.2f ", evaluation)
			}
		}
		fmt.Println()
	}
}

func (evaluationMap *EvaluationMap) GetEvaluation(position Position) float64 {
	var size = evaluationMap.enviroment.size
	var i = position.y*size + position.x
	return evaluationMap.evaluations[i]
}
