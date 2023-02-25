package main

type PopulationSelector struct {
	enviroment Enviroment
}

func NewPopulationSelector(enviroment Enviroment) PopulationSelector {
	return PopulationSelector{enviroment}
}

func (selector *PopulationSelector) IsMicrobeInsideSafeZone(microbe *Microbe) bool {
	return selector.enviroment.GetTile(microbe.position).IsSafeZone()
}

func (selector *PopulationSelector) SelectFrom(population []*Microbe) []*Microbe {
	var selected = make([]*Microbe, 0, len(population))
	for _, microbe := range population {
		if selector.IsMicrobeInsideSafeZone(microbe) {
			selected = append(selected, microbe)
		}
	}
	return selected
}
