/**
 * @project neuroevolution/utils
 * @file    MutexedRand.go
 * @author  Martin Bednář (xbedna77@fit.vut.cz)
 */

package utils

import (
	"sync"
)

/*
Thread-safe random number generator.
Uses mutex to prevent race conditions.
*/
type MutexedRand struct {
	rng   *Rng
	mutex sync.Mutex
}

func NewMutexedRand() *MutexedRand {
	return &MutexedRand{
		rng: NewTimeSeedRng(),
	}
}

func (mxtRand *MutexedRand) Float64() float64 {
	mxtRand.mutex.Lock()
	value := mxtRand.rng.Float64()
	mxtRand.mutex.Unlock()
	return value
}
