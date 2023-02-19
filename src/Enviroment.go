package main

type TileType uint8

const (
	Empty TileType = iota
	SafeZone
	Wall
)

func (tile TileType) IsSafeZone() bool {
	return tile == SafeZone
}

func (tile TileType) IsPassable() bool {
	return tile == Empty || tile == SafeZone
}

type Enviroment struct {
	tiles []TileType
	size  int
}

func NewEnviroment(tiles []TileType, size int) Enviroment {
	return Enviroment{tiles, size}
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