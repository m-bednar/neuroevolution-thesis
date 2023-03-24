package main

import (
	"log"

	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
)

type SpawnSelector struct {
	spawns []Position
	rng    *Rng
}

func NewSpawnSelector(enviroment *Enviroment) *SpawnSelector {
	var rng = NewTimeSeedRng()
	var spawns = enviroment.GetAllTilesOfType(Spawn)
	if len(spawns) == 0 {
		log.Fatal("No spawn tiles in enviroment found.")
	}
	return &SpawnSelector{spawns, rng}
}

func (selector *SpawnSelector) GetPosition() Position {
	var index = selector.rng.Intn(len(selector.spawns))
	return selector.spawns[index]
}
