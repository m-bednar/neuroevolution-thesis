package main

import (
	"encoding/binary"
	"math"
	"os"
)

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
