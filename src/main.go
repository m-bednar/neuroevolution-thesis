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

func ComputeNeuronOutputValue(sum float64, weight float64) float64 {
	return math.Tanh(sum * weight * 0.5)
}

func main() {
	/*
	var microbe = Microbe{}
	var neuron = MoveXNeuron{}
	fmt.Println(neuron.ProcessNeuron(microbe))

	for i := int8(-128); i < int8(127); i++ {
		var w = ToSynapticWeight(i)
		var x = ComputeNeuronOutputValue(1.0, w)
		fmt.Print(i, x)
	}

	fmt.Println()
	Render()
	*/

	dict := make(map[int]int)

	for i := 0; i < 1000; i++ {
		x := int(MutateWeight(0))
		dict[x]++
	}

	fmt.Println(dict)
}
