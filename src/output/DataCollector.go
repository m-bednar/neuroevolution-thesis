package output

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
	. "github.com/m-bednar/neuroevolution-thesis/src/neuralnet"
)

const (
	CAPTURE_MARGIN   = 10
	NORM_GENOME_SIZE = 2
)

type GenomeRepresentation struct {
	red   uint8
	green uint8
	blue  uint8
}

type MicrobePath []Position

type GenerationSample struct {
	representations []GenomeRepresentation
	paths           []MicrobePath
	averageFitness  float64
	highestFitness  float64
	captured        bool
}

type CapturedGenerationSample struct {
	generation int
	GenerationSample
}

/*
Collects statistical data and data required for rendering.
*/
type DataCollector struct {
	gatherer        *StatsGatherer
	samples         []GenerationSample
	currentPaths    []MicrobePath
	captureModifier int
}

func CopyPaths(source []MicrobePath) []MicrobePath {
	destination := make([]MicrobePath, len(source))
	for i := range source {
		destination[i] = make(MicrobePath, len(source[i]))
		copy(destination[i], source[i])
	}
	return destination
}

func GetNormalizedGenomeColor(genome []int8) GenomeRepresentation {
	div := len(genome) / 3
	rem := len(genome) % 3
	values := []uint8{0, 0, 0}
	lengths := []int{div, div, div + rem}
	start := 0

	for i, length := range lengths {
		sum := 0
		for j := start; j < start+length; j++ {
			sum += int(genome[j])
		}
		avg := float64(sum) / float64(length)
		norm := avg / NN_WEIGHT_LIMIT
		values[i] = uint8((norm * 20) + 100)
		start += length
	}

	return GenomeRepresentation{values[0], values[1], values[2]}
}

func GetGenomeRepresentations(genomes [][]int8) []GenomeRepresentation {
	representations := make([]GenomeRepresentation, len(genomes))
	for i, genome := range genomes {
		representations[i] = GetNormalizedGenomeColor(genome)
	}
	return representations
}

func NewDataCollector(gatherer *StatsGatherer, maxGenerations int, stepCount int, captureModifier int) *DataCollector {
	samples := make([]GenerationSample, maxGenerations+1)
	currentPaths := make([]MicrobePath, stepCount)
	return &DataCollector{gatherer, samples, currentPaths, captureModifier}
}

func (collector *DataCollector) CollectStep(step int, population Population) {
	positions := population.CollectPositions()
	collector.currentPaths[step] = positions
}

func (collector *DataCollector) CollectGeneration(generation int, population Population) {
	sample := &collector.samples[generation]
	captured := collector.ShouldCapture(generation)
	if captured {
		sample.representations = GetGenomeRepresentations(population.CollectNormalizedGenomes())
		sample.paths = CopyPaths(collector.currentPaths)
	}
	sample.averageFitness = collector.gatherer.GetAverageFitness(population)
	sample.highestFitness = collector.gatherer.GetHighestFitness(population)
	sample.captured = captured
}

func (collector *DataCollector) ShouldCapture(generation int) bool {
	return generation <= CAPTURE_MARGIN || (generation%collector.captureModifier) == 0
}

func (collector *DataCollector) GetCapturedGenerationSamples() []CapturedGenerationSample {
	captured := make([]CapturedGenerationSample, 0, len(collector.samples))
	for i, sample := range collector.samples {
		if sample.captured {
			captured = append(captured, CapturedGenerationSample{i, sample})
		}
	}
	return captured
}

func (collector *DataCollector) GetAverageFitnesses() []float64 {
	fitnesses := make([]float64, len(collector.samples))
	for i, sample := range collector.samples {
		fitnesses[i] = sample.averageFitness
	}
	return fitnesses
}

func (collector *DataCollector) GetHighestFitnesses() []float64 {
	fitnesses := make([]float64, len(collector.samples))
	for i, sample := range collector.samples {
		fitnesses[i] = sample.highestFitness
	}
	return fitnesses
}
