package main

type StepSample struct {
	positions []Position
}

type GenerationSample struct {
	steps []StepSample
}

type StatsData struct {
	survivability []float64
}

type DataCollector struct {
	gatherer *StatsGatherer
	samples  []GenerationSample
	stats    StatsData
}

func NewDataCollector(gatherer *StatsGatherer, maxGenerations int) *DataCollector {
	var samples = make([]GenerationSample, maxGenerations+1) // FIXME: Can be bigger, than needed after minSurvivability achieved
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
	var survivability = collector.gatherer.GetSuccessRate(population)
	collector.stats.survivability = append(collector.stats.survivability, survivability)
}
