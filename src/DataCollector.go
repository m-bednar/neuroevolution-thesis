package main

type StepSample struct {
	positions []Position
}

type GenerationSample struct {
	genomes [][]int8
	steps   []StepSample
}

type StatsData struct {
	averageFitness []float64
	highestFitness []float64
}

type DataCollector struct {
	gatherer *StatsGatherer
	samples  []GenerationSample
	stats    StatsData
}

func NewDataCollector(gatherer *StatsGatherer, maxGenerations int) *DataCollector {
	var samples = make([]GenerationSample, maxGenerations+1)
	return &DataCollector{gatherer, samples, StatsData{}}
}

func NewStepSample(population Population) StepSample {
	var positions = population.CollectPositions()
	return StepSample{positions}
}

func (collector *DataCollector) CollectPositions(generation int, population Population) {
	var stepSample = NewStepSample(population)
	collector.samples[generation].steps = append(collector.samples[generation].steps, stepSample)
}

func (collector *DataCollector) CollectStats(generation int, population Population) {
	var averageFitness = collector.gatherer.GetAverageFitness(population)
	var highestFitness = collector.gatherer.GetHighestFitness(population)
	collector.stats.averageFitness = append(collector.stats.averageFitness, averageFitness)
	collector.stats.highestFitness = append(collector.stats.highestFitness, highestFitness)
	collector.samples[generation].genomes = population.CollectNormalizedGenomes()
}
