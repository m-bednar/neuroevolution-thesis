package main

type StatsGatherer struct {
	enviroment *Enviroment
	selector   *ParentSelector
}

func NewStatsGatherer(enviroment *Enviroment, selector *ParentSelector) *StatsGatherer {
	return &StatsGatherer{enviroment, selector}
}

func (gatherer *StatsGatherer) IsMicrobeInsideSafeZone(microbe *Microbe) bool {
	return gatherer.enviroment.GetTile(microbe.position).IsSafe()
}

func (gatherer *StatsGatherer) CountMicrobesInSafeZone(population []*Microbe) int {
	var count = 0
	for _, microbe := range population {
		if gatherer.IsMicrobeInsideSafeZone(microbe) {
			count++
		}
	}
	return count
}

func (gatherer *StatsGatherer) GetAverageFitness(population []*Microbe) float64 {
	var sum = 0.0
	for _, microbe := range population {
		sum += microbe.fitness
	}
	return sum / float64(len(population))
}

func (gatherer *StatsGatherer) GetHighestFitness(population []*Microbe) float64 {
	var microbe = gatherer.selector.SelectOneWithHighestFitness(population)
	return microbe.fitness
}

func (gatherer *StatsGatherer) GetSuccessRate(population []*Microbe) float64 {
	return float64(gatherer.CountMicrobesInSafeZone(population)) / float64(len(population))
}
