/**
 * @project neuroevolution/enviroment
 * @file    SpawnPositionGenerator.go
 * @author  Martin Bednář (xbedna77@fit.vut.cz)
 */

package enviroment

import (
	"log"

	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
)

type SpawnSelector struct {
	spawns []Position
	rng    *Rng
}

func NewSpawnSelector(enviroment *Enviroment) *SpawnSelector {
	rng := NewTimeSeedRng()
	spawns := enviroment.GetAllTilesOfType(Spawn)
	if len(spawns) == 0 {
		log.Fatal("No spawn tiles in enviroment found.")
	}
	return &SpawnSelector{spawns, rng}
}

func (selector *SpawnSelector) GetPosition() Position {
	index := selector.rng.Intn(len(selector.spawns))
	return selector.spawns[index]
}
