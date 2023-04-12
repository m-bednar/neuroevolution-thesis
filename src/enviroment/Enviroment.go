package enviroment

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
	size := ComputeEnviromentSize(tiles)
	return &Enviroment{
		tiles: tiles,
		size:  size,
	}
}

func ComputeEnviromentSize(tiles []TileType) int {
	count := float64(len(tiles))
	squared := math.Sqrt(count)
	if squared != math.Trunc(squared) {
		log.Fatal("Enviroment size must be NxN tiles.")
	}
	return int(squared)
}

func (enviroment *Enviroment) GetTiles() []TileType {
	return enviroment.tiles
}

func (enviroment *Enviroment) GetNumberOfTiles() int {
	return len(enviroment.tiles)
}

func (enviroment *Enviroment) GetSize() int {
	return enviroment.size
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

func (enviroment *Enviroment) GetTilePosition(index int) Position {
	return NewPosition(index%enviroment.size, index/enviroment.size)
}

func (enviroment *Enviroment) GetTile(position Position) TileType {
	index := enviroment.GetTileIndex(position)
	return enviroment.tiles[index]
}

func (enviroment *Enviroment) GetAllTilesOfType(tileType TileType) []Position {
	positions := make([]Position, 0, len(enviroment.tiles))
	for i, tile := range enviroment.tiles {
		if tile == tileType {
			position := enviroment.GetTilePosition(i)
			positions = append(positions, position)
		}
	}
	return positions
}

func (enviroment *Enviroment) GetAllPassableTiles() []Position {
	positions := make([]Position, 0, len(enviroment.tiles))
	for i, tile := range enviroment.tiles {
		if tile.IsPassable() {
			position := enviroment.GetTilePosition(i)
			positions = append(positions, position)
		}
	}
	return positions
}

func (enviroment *Enviroment) GetDistanceToWallInDirection(origin Position, direction Direction, max int) (bool, float64) {
	tries := 0
	current := origin
	for tries < max && enviroment.IsInsideBorders(current) {
		if enviroment.GetTile(current).IsWall() {
			return true, origin.DistanceTo(current)
		}
		current = current.Add(direction.x, direction.y)
		tries++
	}
	return false, 0
}

func (enviroment *Enviroment) Neighbours(position Position) []Position {
	neighbours := []Position{
		position.Add(1, 0), position.Add(-1, 0),
		position.Add(0, 1), position.Add(0, -1),
	}
	passable := make([]Position, 0, len(neighbours))
	for _, neighbour := range neighbours {
		if enviroment.IsPassable(neighbour) {
			passable = append(passable, neighbour)
		}
	}
	return passable
}
