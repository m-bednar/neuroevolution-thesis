package main

import (
	"fmt"
	"math"
)

type MoveXNeuron struct {
	Neuron
}

func (neuron *MoveXNeuron) ProcessNeuron(microbe Microbe) float64 {
	neuron.GetSynapticSum(microbe)
	// move
	return 0
}

/*
const BRAIN_SIZE = 8

type InputNeuronId int8
type OutputNeuronId int8

type NeuronProcess func() float64

type Neuron struct {
	name    string
	process NeuronProcess
}

type Point struct {
	x uint8
	y uint8
}

type Synapse struct {
	weight float64
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

/*
Conversion from signed 8-bit to float
char a = 0b10000000;
char b = 0b01111111;
(a + 0.5f) / 32.0f  =  -3.984375
(b + 0.5f) / 32.0f  =   3.984375
*/

func ToSynapticWeight(encoded int8) float64 {
	return (float64(encoded) + 0.5) / 16.0
}

func main() {
	var microbe = Microbe{}
	var neuron = MoveXNeuron{
		Neuron{
			inputs: []Synapse{},
		},
	}
	fmt.Println(neuron.ProcessNeuron(microbe))

	var a int8 = 127
	var b int8 = -128

	var x = math.Tanh(1.0 * ToSynapticWeight(a))
	var y = math.Tanh(1.0 * ToSynapticWeight(b))

	fmt.Println(x)
	fmt.Println(y)
}
