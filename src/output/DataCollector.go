package output

import (
	"math"

	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
	. "github.com/m-bednar/neuroevolution-thesis/src/neuralnet"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
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
		destination[i] = Clone(source[i])
	}
	return destination
}

// Maps <NN_WEIGHT_LIMIT;NN_WEIGHT_LIMIT> to <0;1>
func NormalizeWeigth(weight float64) float64 {
	shifted := weight + NN_WEIGHT_LIMIT
	normalized := shifted / (NN_WEIGHT_LIMIT * 2)
	return normalized
}

func WeightToUint8(weight float64) uint8 {
	normalized := NormalizeWeigth(weight)
	return uint8(normalized * math.MaxUint8)
}

func GetAverageGenomeValueRepresentation(genome []float64) uint8 {
	sum := 0.0
	for _, weight := range genome {
		sum += weight
	}
	average := sum / float64(len(genome))
	return WeightToUint8(average)
}

func GetGenomeRepresentation(genome []float64) GenomeRepresentation {
	part1 := len(genome) / 3
	part2 := part1 * 2
	r := GetAverageGenomeValueRepresentation(genome[:part1])
	g := GetAverageGenomeValueRepresentation(genome[part1:part2])
	b := GetAverageGenomeValueRepresentation(genome[part2:])
	return GenomeRepresentation{r, g, b}
}

func GetGenomeRepresentations(genomes [][]float64) []GenomeRepresentation {
	representations := make([]GenomeRepresentation, len(genomes))
	for i, genome := range genomes {
		representations[i] = GetGenomeRepresentation(genome)
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
		sample.representations = GetGenomeRepresentations(population.CollectGenomes())
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
