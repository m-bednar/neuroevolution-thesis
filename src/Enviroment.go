package main

import (
	"log"
	"math"
)

type TileType uint8

const (
	None TileType = iota
	Safe
)

func (tile TileType) IsSafe() bool {
	return tile == Safe
}

func (tile TileType) IsPassable() bool {
	return tile == None || tile == Safe
}

type Enviroment struct {
	tiles []TileType
	size  int
}

func NewEnviroment(tiles []TileType) *Enviroment {
	return &Enviroment{
		tiles: tiles,
		size:  ComputeEnviromentSize(tiles),
	}
}

func ComputeEnviromentSize(tiles []TileType) int {
	var count = float64(len(tiles))
	var squared = math.Sqrt(count)
	if squared != math.Trunc(squared) {
		log.Fatal("Enviroment size must be NxN tiles.")
	}
	return int(squared)
}

func (enviroment *Enviroment) IsPassable(position Position) bool {
	if !enviroment.IsInsideBorders(position) {
		return false
	}
	if !enviroment.GetTile(position).IsPassable() {
		return false
	}
	return true
}

func (enviroment *Enviroment) IsInsideBorders(position Position) bool {
	return position.x >= 0 && position.y >= 0 && position.x < enviroment.size && position.y < enviroment.size
}

func (enviroment *Enviroment) GetTile(position Position) TileType {
	var index = (position.y * enviroment.size) + position.x
	return enviroment.tiles[index]
}
