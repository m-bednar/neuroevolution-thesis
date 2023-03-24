package utils

import (
	"sync"
)

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
	var value = mxtRand.rng.Float64()
	mxtRand.mutex.Unlock()
	return value
}
