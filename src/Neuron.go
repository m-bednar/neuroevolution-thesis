package main

type NeuronProcessor interface {
	ProcessNeuron(microbe Microbe) float32
}

type Neuron struct {
	inputs []Synapse
}

func (neuron *Neuron) GetSynapticSum(microbe Microbe) float32 {
	var value float32 = 0
	for _, synapse := range neuron.inputs {
		value += synapse.GetValue(microbe)
	}
	return value
}
