package main

const (
	MIN_RAND_WEIGHT = -2
	MAX_RAND_WEIGHT = 2
)

type NeuralNetworkFactory interface {
	Make() NeuralNetwork
}
