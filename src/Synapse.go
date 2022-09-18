package main

type Synapse struct {
	weight float32
	source NeuronProcessor
}

func (synapse *Synapse) GetValue(microbe Microbe) float32 {
	return synapse.source.ProcessNeuron(microbe) * synapse.weight
}
