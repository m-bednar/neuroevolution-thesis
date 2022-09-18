package main

import "fmt"

type Microbe struct{}

type NeuronProcessor interface {
	ProcessNeuron(microbe Microbe) float32
}

type Neuron struct {
	inputs []Synapse
}

type Synapse struct {
	weight float32
	source NeuronProcessor
}

func (synapse *Synapse) GetValue(microbe Microbe) float32 {
	return synapse.source.ProcessNeuron(microbe) * synapse.weight
}

func (neuron *Neuron) GetSynapticSum(microbe Microbe) float32 {
	var value float32 = 0
	for _, synapse := range neuron.inputs {
		value += synapse.GetValue(microbe)
	}
	return value
}

type MoveXNeuron struct {
	Neuron
}

func (neuron *MoveXNeuron) ProcessNeuron(microbe Microbe) float32 {
	neuron.GetSynapticSum(microbe)
	// move
	return 0
}

/*
func Signum(input float32) float32 {
	if input > 0.0 {
		return 1.0
	}
	if input < 0.0 {
		return -1.0
	}
	return 0.0
}

const BRAIN_SIZE = 8

type InputNeuronId int8
type OutputNeuronId int8

type NeuronProcess func() float32

type Neuron struct {
	name    string
	process NeuronProcess
}

type Point struct {
	x uint8
	y uint8
}

type Synapse struct {
	weight float32
	input  *Neuron
	output *Neuron
}

type NeuralNet struct {
	connections [BRAIN_SIZE]Synapse
}

type Microbe struct {
	genome   [BRAIN_SIZE]byte
	location Point
	brain    NeuralNet
}

func main() {
	var microbe = Microbe{}
	fmt.Println(500 * unsafe.Sizeof(microbe))
}
*/

func main() {
	var microbe = Microbe{}
	var neuron = MoveXNeuron{
		Neuron{
			inputs: []Synapse{},
		},
	}
	fmt.Println(neuron.ProcessNeuron(microbe))
}
