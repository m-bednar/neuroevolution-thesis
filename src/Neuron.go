package main

type NeuronProcessor interface {
	ProcessNeuron(microbe Microbe) float64
}

type Neuron struct {
	inputs []Synapse
}

func (neuron *Neuron) GetSynapticSum(microbe Microbe) float64 {
	var value float64 = 0
	for _, synapse := range neuron.inputs {
		value += synapse.GetValue(microbe)
	}
	return value
}
