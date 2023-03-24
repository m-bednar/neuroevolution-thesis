package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/fzipp/astar"
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
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

func GetBetterPathDistance(previous int, current int) int {
	if previous == 0 || current < previous {
		return current
	}
	return previous
}

func GetTilesPathDistances(enviroment *Enviroment) []int {
	var size = enviroment.GetNumberOfTiles()
	var distances = make([]int, size)
	var safeTiles = enviroment.GetAllTilesOfType(Safe)
	var passableTiles = enviroment.GetAllPassableTiles()
	var distFunc = Position.DistanceTo
	var findPath = astar.FindPath[Position]
	var mutex = sync.Mutex{}

	// Evaluate each passable tile by finding
	// it's shortest path to closest safe tile
	for _, start := range safeTiles {
		ConcurrentLoop(passableTiles, func(_ int, end Position) {
			var path = findPath(enviroment, start, end, distFunc, distFunc)
			if path != nil {
				var index = enviroment.GetTileIndex(end)
				var curr = len(path)
				mutex.Lock()
				var prev = distances[index]
				distances[index] = GetBetterPathDistance(prev, curr)
				mutex.Unlock()
			}
		})
	}

	return distances
}

func CreateEvaluations(enviroment *Enviroment) []float64 {
	var size = enviroment.GetNumberOfTiles()
	var distances = GetTilesPathDistances(enviroment)

	// Transform distances to evaluations
	var evaluations = make([]float64, size)
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
	for y := 0; y < evaluationMap.enviroment.GetSize(); y++ {
		for x := 0; x < evaluationMap.enviroment.GetSize(); x++ {
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
	var size = evaluationMap.enviroment.GetSize()
	var i = position.GetY()*size + position.GetX()
	return evaluationMap.evaluations[i]
}
