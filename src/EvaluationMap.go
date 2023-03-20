package main

import (
	"fmt"
	"math"
	"sync"

	as "github.com/fzipp/astar"
	"github.com/myfantasy/mft/im"
)

type EvaluationMap struct {
	evaluations []float64
	enviroment  *Enviroment
}

func EvaluateDistance(distance, max int) float64 {
	var value = float64(-(distance - max)) + math.SmallestNonzeroFloat64
	return value / float64(max)
}

func CreateEvaluations(enviroment *Enviroment) []float64 {
	var evaluations = make([]float64, len(enviroment.tiles))
	var distances = make([]int, len(enviroment.tiles))
	var safeTiles = enviroment.GetAllTilesOfType(Safe)
	var passableTiles = enviroment.GetAllPassableTiles()
	var distFunc = Position.DistanceTo
	var findPath = as.FindPath[Position]
	var mutex = sync.Mutex{}

	// Evaluate each passable tile by finding
	// it's shortest path to closest safe tile
	LoopAsync(safeTiles, func(index int, safeTile Position) {
		for _, passableTile := range passableTiles {
			var start = safeTile
			var end = passableTile
			var path = findPath(enviroment, start, end, distFunc, distFunc)
			if path != nil {
				var index = enviroment.GetTileIndex(end)
				var curr = len(path)
				mutex.Lock()
				var prev = distances[index]
				if prev == 0 || curr < prev {
					distances[index] = curr
				}
				mutex.Unlock()
			}
		}
	})

	// Transform distances to evaluations
	var max = im.MaxS(distances...)
	for i, distance := range distances {
		if distance != 0 {
			evaluations[i] = EvaluateDistance(distance, max)
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
				fmt.Printf("xxxxx ")
			} else {
				fmt.Printf("%5.2f ", evaluation)
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
