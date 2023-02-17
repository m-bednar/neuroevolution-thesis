package main

import (
	"fmt"
)

func main() {
	var neuralNetworkFactory = NewNeuralNetworkRandomizedFactory(3)
	var microbeFactory = NewMicrobeFactory(&neuralNetworkFactory)
	var microbe = microbeFactory.Make(10, 20)
	fmt.Println(microbe.neuralNetwork.weights)
}
