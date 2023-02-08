package main

import (
	"fmt"
)

func main() {
	var neuralNetworkFactory = NewRandomNeuralNetworkFactory(5)
	var microbeFactory = NewMicrobeFactory(neuralNetworkFactory)
	var microbe = microbeFactory.Make(10, 20)
	fmt.Println(microbe.neuralNetwork.weights)
}
