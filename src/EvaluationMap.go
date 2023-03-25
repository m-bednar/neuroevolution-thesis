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
	value := float64(-(distance - max)) + math.SmallestNonzeroFloat64
	return value / float64(max)
}

func GetBetterPathDistance(previous int, current int) int {
	if previous == 0 || current < previous {
		return current
	}
	return previous
}

func GetTilesPathDistances(enviroment *Enviroment) []int {
	size := enviroment.GetNumberOfTiles()
	distances := make([]int, size)
	safeTiles := enviroment.GetAllTilesOfType(Safe)
	passableTiles := enviroment.GetAllPassableTiles()
	distFunc := Position.DistanceTo
	findPath := astar.FindPath[Position]
	mutex := sync.Mutex{}

	// Evaluate each passable tile by finding
	// it's shortest path to closest safe tile
	for _, start := range safeTiles {
		ConcurrentLoop(passableTiles, func(_ int, end Position) {
			path := findPath(enviroment, start, end, distFunc, distFunc)
			if path != nil {
				index := enviroment.GetTileIndex(end)
				curr := len(path)
				mutex.Lock()
				prev := distances[index]
				distances[index] = GetBetterPathDistance(prev, curr)
				mutex.Unlock()
			}
		})
	}

	return distances
}

func CreateEvaluations(enviroment *Enviroment) []float64 {
	size := enviroment.GetNumberOfTiles()
	distances := GetTilesPathDistances(enviroment)

	// Transform distances to evaluations
	evaluations := make([]float64, size)
	max := im.MaxS(distances...)
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
			evaluation := evaluationMap.GetEvaluation(NewPosition(x, y))
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
	size := evaluationMap.enviroment.GetSize()
	i := position.GetY()*size + position.GetX()
	return evaluationMap.evaluations[i]
}
