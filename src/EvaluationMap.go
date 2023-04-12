package main

import (
	"fmt"

	"github.com/fzipp/astar"
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
	"github.com/myfantasy/mft/im"
)

const (
	NO_EVALUATION = -1
)

type EvaluationMap struct {
	evaluations []float64
	enviroment  *Enviroment
}

func EvaluateDistance(distance, max int) float64 {
	value := float64(-(distance - max))
	evaluation := value / float64(max)
	return evaluation + (evaluation / 100.0)
}

func GetMinDistance(current int, new int) int {
	if current == NO_EVALUATION || new < current {
		return new
	}
	return current
}

func GetMinPathDistanceForTile(enviroment *Enviroment, tile Position) int {
	safeTiles := enviroment.GetAllTilesOfType(Safe)
	findPath := astar.FindPath[Position]
	distFunc := Position.DistanceTo
	distance := NO_EVALUATION

	for _, safeTile := range safeTiles {
		path := findPath(enviroment, tile, safeTile, distFunc, distFunc)
		if path != nil {
			distance = GetMinDistance(distance, len(path))
		}
	}

	return distance
}

func GetTilesPathDistances(enviroment *Enviroment) []int {
	size := enviroment.GetNumberOfTiles()
	tiles := enviroment.GetAllPassableTiles()
	distances := make([]int, size)

	// Evaluate each passable tile by finding
	// it's shortest path to closest safe tile
	AsyncFor(tiles, func(_ int, tile Position) {
		distance := GetMinPathDistanceForTile(enviroment, tile)
		index := enviroment.GetTileIndex(tile)
		distances[index] = distance
	})

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
				fmt.Printf("xxxx ")
			} else {
				fmt.Printf("%4.2f ", evaluation)
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
