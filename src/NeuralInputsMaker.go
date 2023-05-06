/**
 * @project neuroevolution
 * @file    NeuralInputsMaker.go
 * @author  Martin Bednář (xbedna77@fit.vut.cz)
 */

package main

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
)

const (
	WALL_SENSORY_RANGE = 4 // Microbe wall sensory range (in tiles)
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
	hasWallInDirection, distanceToWall := maker.enviroment.GetDistanceToWallInDirection(origin, direction, WALL_SENSORY_RANGE)
	if hasWallInDirection && distanceToWall <= WALL_SENSORY_RANGE {
		return (WALL_SENSORY_RANGE - (distanceToWall - 1)) / WALL_SENSORY_RANGE
	}
	return 0.0
}

func (maker *NeuralInputsMaker) MakeInputsFor(microbe *Microbe) []float64 {
	position := microbe.GetPosition()
	enviroment := maker.enviroment
	enviromentSize := float64(enviroment.GetSize())
	borderDistWest := float64(position.GetX()) / enviromentSize
	borderDistNorth := float64(position.GetY()) / enviromentSize
	wallNorth := maker.GetSignalForWallTileInDirection(position, North)
	wallSouth := maker.GetSignalForWallTileInDirection(position, South)
	wallWest := maker.GetSignalForWallTileInDirection(position, West)
	wallEast := maker.GetSignalForWallTileInDirection(position, East)
	return []float64{borderDistWest, borderDistNorth, wallNorth, wallSouth, wallWest, wallEast}
}
