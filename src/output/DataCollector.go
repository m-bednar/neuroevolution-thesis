package output

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
)

const (
	CAPTURE_MARGIN = 10
)

type GenerationSample struct {
	genomes        [][]int8
	steps          [][]Position
	averageFitness float64
	highestFitness float64
	captured       bool
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
	currentSteps    [][]Position
	captureModifier int
}

func CopyNestedPositions(source [][]Position) [][]Position {
	destination := make([][]Position, len(source))
	for i := range source {
		destination[i] = make([]Position, len(source[i]))
		copy(destination[i], source[i])
	}
	return destination
}

func NewDataCollector(gatherer *StatsGatherer, maxGenerations int, stepCount int, captureModifier int) *DataCollector {
	samples := make([]GenerationSample, maxGenerations+1)
	currentSteps := make([][]Position, stepCount)
	return &DataCollector{gatherer, samples, currentSteps, captureModifier}
}

func (collector *DataCollector) CollectStep(step int, population Population) {
	positions := population.CollectPositions()
	collector.currentSteps[step] = positions
}

func (collector *DataCollector) CollectGeneration(generation int, population Population) {
	sample := &collector.samples[generation]
	captured := collector.ShouldCapture(generation)
	if captured {
		sample.genomes = population.CollectNormalizedGenomes()
		sample.steps = CopyNestedPositions(collector.currentSteps)
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
