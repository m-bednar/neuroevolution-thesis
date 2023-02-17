package main

import (
	"fmt"
)

type NeuralNetworkParams struct {
	internalWidth uint
}

func main() {
	var positionGenerator = NewPositionGenerator(10, 10)
	var neuralNetworkFactory = NewNeuralNetworkRandomizedFactory()
	var microbeFactory = NewMicrobeFactory(positionGenerator, &neuralNetworkFactory)
	var microbe = microbeFactory.Make()
	fmt.Println(microbe.neuralNetwork.Process([]float64{1, 2, 0, 1}))
}
