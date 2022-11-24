package main

type NeuralNetwork struct {
	in      uint8
	out     uint8
	width   uint8
	depth   uint8
	weights []int8
}

func NewNeuralNetwork(in uint8, out uint8, width uint8, depth uint8, genome Genome) NeuralNetwork {
	var n = in*width + out*width + (depth-1)*(width*width)
	var weights = make([]int8, n)
	return NeuralNetwork{in, out, width, depth, weights}
}
