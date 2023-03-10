package main

type StatsCollector struct {
	gatherer      *StatsGatherer
	survivability []float64
}

func NewStatsCollector(gatherer *StatsGatherer) *StatsCollector {
	return &StatsCollector{
		gatherer: gatherer,
	}
}

func (collector *StatsCollector) Collect(population []*Microbe) {
	var survivability = collector.gatherer.GetSuccessRate(population)
	collector.survivability = append(collector.survivability, survivability)
}
