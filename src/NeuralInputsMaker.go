package main

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
)

const (
	WALL_SENSORY_RANGE = 4.0
)

type NeuralInputsMaker struct {
	enviroment *Enviroment
}

func NewNeuralInputsMaker(enviroment *Enviroment) *NeuralInputsMaker {
	return &NeuralInputsMaker{
		enviroment: enviroment,
	}
}

func (maker *NeuralInputsMaker) GetSignalForWallTileInDirection(origin Position, direction Direction) float64 {
	var hasWallInDirection, distanceToWall = maker.enviroment.GetDistanceToWallInDirection(origin, direction)
	if hasWallInDirection && distanceToWall <= WALL_SENSORY_RANGE {
		return (WALL_SENSORY_RANGE - (distanceToWall - 1)) / WALL_SENSORY_RANGE
	}
	return 0.0
}

func (maker *NeuralInputsMaker) MakeInputsFor(microbe *Microbe) []float64 {
	var position = microbe.GetPosition()
	var enviroment = maker.enviroment
	var enviromentSize = float64(enviroment.GetSize())
	var borderDistWest = float64(position.GetX()) / enviromentSize
	var borderDistNorth = float64(position.GetY()) / enviromentSize
	var wallNorth = maker.GetSignalForWallTileInDirection(position, North)
	var wallSouth = maker.GetSignalForWallTileInDirection(position, South)
	var wallWest = maker.GetSignalForWallTileInDirection(position, West)
	var wallEast = maker.GetSignalForWallTileInDirection(position, East)

	return []float64{borderDistWest, borderDistNorth, wallNorth, wallSouth, wallWest, wallEast}
}
