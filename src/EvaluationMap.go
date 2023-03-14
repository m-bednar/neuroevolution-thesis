package main

import (
	"log"
)

type EvaluationMap struct {
	evaluations []float64
	enviroment  *Enviroment
}

func FindClosestSafeTile(enviroment *Enviroment, position Position) Position {
	var safeTiles = enviroment.GetAllTilesOfType(Safe)
	if len(safeTiles) == 0 {
		log.Fatal("No safe tiles found in enviroment.")
	}

	var closest = safeTiles[0]
	for _, safeTile := range safeTiles {
		if position.DistanceTo(safeTile) < position.DistanceTo(closest) {
			closest = safeTile
		}
	}

	return closest
}

func CreateEvaluation(enviroment *Enviroment, x int, y int) float64 {
	var position = NewPosition(x, y)

	if !enviroment.GetTile(position).IsPassable() {
		return WALL_TILE_EVALUATION
	}

	var closestSafeTile = FindClosestSafeTile(enviroment, position)
	var distance = position.DistanceTo(closestSafeTile) + 1

	return EVALUATION_DISTANCE_MULT * (1 / distance)
}

func CreateEvaluations(enviroment *Enviroment) []float64 {
	var size = enviroment.size
	var evaluations = make([]float64, size*size)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			var i = y*size + x
			evaluations[i] = CreateEvaluation(enviroment, x, y)
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

func (evaluationMap *EvaluationMap) GetEvaluation(position Position) float64 {
	var size = evaluationMap.enviroment.size
	var i = position.y*size + position.x
	return evaluationMap.evaluations[i]
}
