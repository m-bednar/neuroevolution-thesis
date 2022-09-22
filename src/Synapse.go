package main

type Synapse struct {
	weight float64
	source NeuronProcessor
}

func (synapse *Synapse) GetValue(microbe Microbe) float64 {
	return synapse.source.ProcessNeuron(microbe) * synapse.weight
}
