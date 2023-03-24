package utils

import (
	"math/rand"
	"time"
)

func NewTimeSeedRng() *rand.Rand {
	var seed = time.Now().UnixNano()
	var source = rand.NewSource(seed)
	return rand.New(source)
}
