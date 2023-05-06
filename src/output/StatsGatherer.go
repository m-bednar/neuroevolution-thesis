/**
 * @project neuroevolution/output
 * @file    StatsGatherer.go
 * @author  Martin Bednář (xbedna77@fit.vut.cz)
 */

package output

import (
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/microbe"
)

type StatsGatherer struct {
	enviroment *Enviroment
}

func NewStatsGatherer(enviroment *Enviroment) *StatsGatherer {
	return &StatsGatherer{enviroment}
}

func (gatherer *StatsGatherer) IsMicrobeInsideSafeZone(microbe *Microbe) bool {
	return gatherer.enviroment.GetTile(microbe.GetPosition()).IsSafe()
}

func (gatherer *StatsGatherer) GetAverageFitness(population Population) float64 {
	sum := 0.0
	for _, microbe := range population {
		sum += microbe.GetFitness()
	}
	return sum / float64(len(population))
}

func (gatherer *StatsGatherer) GetHighestFitness(population Population) float64 {
	microbe := population.SelectOneWithHighestFitness()
	return microbe.GetFitness()
}
