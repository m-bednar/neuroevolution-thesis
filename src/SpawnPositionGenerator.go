package main

import "math/rand"

type SpawnSelector struct {
	spawns []Position
	rng    *rand.Rand
}

func NewSpawnSelector(enviroment *Enviroment) *SpawnSelector {
	var rng = NewUnixTimeRng()
	var spawns = enviroment.GetAllTilesOfType(Spawn)
	return &SpawnSelector{spawns, rng}
}

func (selector *SpawnSelector) GetRandomSpawnPosition() Position {
	var index = selector.rng.Intn(len(selector.spawns))
	return selector.spawns[index]
}
