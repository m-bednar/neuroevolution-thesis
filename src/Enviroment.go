package main

import (
	"log"
	"math"
)

type TileType uint8

const (
	None TileType = iota
	Safe
	Wall
	Spawn
)

type Direction Position

var (
	North = Direction(Position{0, -1})
	South = Direction(Position{0, 1})
	West  = Direction(Position{-1, 0})
	East  = Direction(Position{1, 0})
)

func (tile TileType) IsSafe() bool {
	return tile == Safe
}

func (tile TileType) IsWall() bool {
	return tile == Wall
}

func (tile TileType) IsPassable() bool {
	return tile != Wall
}

type Enviroment struct {
	tiles []TileType
	size  int
}

func NewEnviroment(tiles []TileType) *Enviroment {
	var size = ComputeEnviromentSize(tiles)
	return &Enviroment{
		tiles: tiles,
		size:  size,
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

func (enviroment *Enviroment) GetTileIndex(position Position) int {
	return (position.y * enviroment.size) + position.x
}

func (enviroment *Enviroment) GetTile(position Position) TileType {
	var index = enviroment.GetTileIndex(position)
	return enviroment.tiles[index]
}

func (enviroment *Enviroment) GetAllTilesOfType(tileType TileType) []Position {
	var positions = make([]Position, 0, len(enviroment.tiles))
	for i, tile := range enviroment.tiles {
		if tile == tileType {
			var position = NewPosition(i%enviroment.size, i/enviroment.size)
			positions = append(positions, position)
		}
	}
	return positions
}

func (enviroment *Enviroment) GetAllPassableTiles() []Position {
	var positions = make([]Position, 0, len(enviroment.tiles))
	for i, tile := range enviroment.tiles {
		if tile.IsPassable() {
			var position = NewPosition(i%enviroment.size, i/enviroment.size)
			positions = append(positions, position)
		}
	}
	return positions
}

func (enviroment *Enviroment) GetDistanceToWallInDirection(origin Position, direction Direction) (bool, float64) {
	var current = origin
	for enviroment.IsInsideBorders(current) {
		if enviroment.GetTile(current).IsWall() {
			return true, origin.DistanceTo(current)
		}
		current = current.Add(direction.x, direction.y)
	}
	return false, 0
}

func (enviroment *Enviroment) Neighbours(position Position) []Position {
	var neighbours = []Position{
		position.Add(1, 0), position.Add(-1, 0),
		position.Add(0, 1), position.Add(0, -1),
	}
	var passable = make([]Position, 0, len(neighbours))
	for _, neighbour := range neighbours {
		if enviroment.IsPassable(neighbour) {
			passable = append(passable, neighbour)
		}
	}
	return passable
}
