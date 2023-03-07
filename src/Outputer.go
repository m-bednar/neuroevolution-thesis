package main

import (
	"encoding/binary"
	"math"
	"os"
)

/*
population status:
 8 8 8  64 64   8 8 8  64 64 ...
aabbcc (10, 2) 00aaff (8, 4) ... x pop size

simulation step (x steps):
  64 64   64 64 ...
(10, 2) (10, 2) ... x pop size

generation stats:
  64   f64   f64
 155  4.15  8.12
*/

const (
	MICROBE_STATUS_OUTPUT_SIZE   = 3 + 4 + 4
	MICROBE_STEP_OUTPUT_SIZE     = 4 + 4
	GENERATION_STATS_OUTPUT_SIZE = 4 + 8 + 8
)

func OutputPopulationStatus(population []*Microbe) {
	var buffer = make([]byte, 0, MICROBE_STATUS_OUTPUT_SIZE*len(population))
	for _, microbe := range population {
		var r, g, b = microbe.GetRGBHexCode()
		buffer = append(buffer, r, g, b)
		buffer = binary.BigEndian.AppendUint32(buffer, uint32(microbe.position.x))
		buffer = binary.BigEndian.AppendUint32(buffer, uint32(microbe.position.y))
	}
	os.Stdout.Write(buffer)
}

func OutputSimulationStep(population []*Microbe) {
	var buffer = make([]byte, 0, MICROBE_STEP_OUTPUT_SIZE*len(population))
	for _, microbe := range population {
		buffer = binary.BigEndian.AppendUint32(buffer, uint32(microbe.position.x))
		buffer = binary.BigEndian.AppendUint32(buffer, uint32(microbe.position.y))
	}
	os.Stdout.Write(buffer)
}

func OutputGenerationStats(population []*Microbe, stats *StatsGatherer) {
	var safe = stats.CountMicrobesInSafeZone(population)
	var averageFitness = stats.GetAverageFitness(population)
	var highestFitness = stats.GetHighestFitness(population)
	var buffer = make([]byte, 0, GENERATION_STATS_OUTPUT_SIZE)
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(safe))
	buffer = binary.BigEndian.AppendUint64(buffer, math.Float64bits(averageFitness))
	buffer = binary.BigEndian.AppendUint64(buffer, math.Float64bits(highestFitness))
	os.Stdout.Write(buffer)
}
