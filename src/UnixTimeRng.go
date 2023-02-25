package main

import (
	"math/rand"
	"time"
)

func NewUnixTimeRng() *rand.Rand {
	var seed = time.Now().UnixNano()
	var source = rand.NewSource(seed)
	return rand.New(source)
}
