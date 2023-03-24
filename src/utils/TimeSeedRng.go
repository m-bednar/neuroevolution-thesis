package utils

import (
	"math/rand"
	"time"
)

type Rng struct {
	*rand.Rand
}

func NewTimeSeedRng() *Rng {
	var seed = time.Now().UnixNano()
	var source = rand.NewSource(seed)
	var rnd = rand.New(source)
	return &Rng{Rand: rnd}
}

func (rng *Rng) NormFloat64(deviation float64) float64 {
	return rng.Rand.NormFloat64() * deviation
}
