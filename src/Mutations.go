package main

import (
	"math"
	"math/rand"
)

func clamp(value int) int8 {
	if value < math.MinInt8 {
		return math.MinInt8 
	}
	if value > math.MaxInt8 {
		return math.MaxInt8 
	}
	return int8(value)
}

func sign(x float64) int {
	if x <= 0 {
		return 1
	}
	return -1
}

func getRandomChange() int {
	const RND_INTERVAL_RADIUS = 1000
	const RND_INTERVAL_DIAMETER = RND_INTERVAL_RADIUS * 2
	var rnd = rand.Float64() * RND_INTERVAL_DIAMETER - RND_INTERVAL_RADIUS
	var base = math.Pow(1.05, math.Abs(rnd) / 10.0) 
	return int(base) * sign(rnd)
}

func MutateWeight(weight int8) int8 {
	var change = getRandomChange()
	var new = int(weight) + change
	return clamp(new)
}
