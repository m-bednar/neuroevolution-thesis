package main

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
	var hasWallInDirection, distanceToWall = maker.enviroment.GetDistanceToWallTileInDirection(origin, direction)
	if hasWallInDirection && distanceToWall <= WALL_SENSORY_RANGE {
		return (WALL_SENSORY_RANGE - (distanceToWall - 1)) / WALL_SENSORY_RANGE
	}
	return 0.0
}

func (maker *NeuralInputsMaker) MakeInputsFor(microbe *Microbe) []float64 {
	var position = microbe.position
	var enviroment = maker.enviroment
	var enviromentSize = float64(enviroment.size)
	var borderDistWest = float64(position.x) / enviromentSize
	var borderDistNorth = float64(position.y) / enviromentSize
	var wallNorth = maker.GetSignalForWallTileInDirection(position, North)
	var wallSouth = maker.GetSignalForWallTileInDirection(position, South)
	var wallWest = maker.GetSignalForWallTileInDirection(position, West)
	var wallEast = maker.GetSignalForWallTileInDirection(position, East)

	return []float64{borderDistWest, borderDistNorth, wallNorth, wallSouth, wallWest, wallEast}
}
