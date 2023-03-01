package main

type StatsGatherer struct {
	enviroment *Enviroment
}

func NewStatsGatherer(enviroment *Enviroment) *StatsGatherer {
	return &StatsGatherer{enviroment}
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
