package main

import (
	"log"
	"math/rand"
)

type SpawnSelector struct {
	spawns []Position
	rng    *rand.Rand
}

func NewSpawnSelector(enviroment *Enviroment) *SpawnSelector {
	var rng = NewUnixTimeRng()
	var spawns = enviroment.GetAllTilesOfType(Spawn)
	if len(spawns) == 0 {
		log.Fatal("No spawn tiles in enviroment found.")
	}
	return &SpawnSelector{spawns, rng}
}

func (selector *SpawnSelector) GetRandomSpawnPosition() Position {
	var index = selector.rng.Intn(len(selector.spawns))
	return selector.spawns[index]
}
