package main

const (
	LAYER_WIDTH     = 4
	MIN_RAND_WEIGHT = -2
	MAX_RAND_WEIGHT = 2
)

type NeuralNetworkFactory interface {
	Make() NeuralNetwork
}
