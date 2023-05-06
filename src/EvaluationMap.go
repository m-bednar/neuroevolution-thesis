/**
 * @project neuroevolution
 * @file    EvaluationMap.go
 * @author  Martin Bednář (xbedna77@fit.vut.cz)
 */

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

// Used to store pre-computed evaluations
type EvaluationMap struct {
	evaluations []float64
	enviroment  *Enviroment
}

func EvaluateDistance(distance int, max int) float64 {
	if distance == NO_EVALUATION {
		return NO_EVALUATION
	}
	value := float64(-(distance - max))
	evaluation := value / float64(max-1)
	return evaluation
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

	// Fill distances with default value
	for i := range distances {
		distances[i] = NO_EVALUATION
	}

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
	distances := GetTilesPathDistances(enviroment)
	evaluations := make([]float64, len(distances))
	max := im.MaxS(distances...)

	// Transform distances to evaluations
	for i, distance := range distances {
		evaluations[i] = EvaluateDistance(distance, max)
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
			if evaluation == NO_EVALUATION {
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
