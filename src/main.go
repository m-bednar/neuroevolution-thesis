package main

import "fmt"

func main() {
	var positionGenerator = NewPositionGenerator(10, 10)
	var neuralNetworkRandomFactory = NewNeuralNetworkRandomFactory()
	var populationRandomFactory = NewPopulationRandomFactory(positionGenerator, neuralNetworkRandomFactory)

	var firstPopulation = populationRandomFactory.Make(5)

	var neuralNetworkReproductiveFactory = NewNeuralNetworkReproductiveFactory()
	var populationReproductiveFactory = NewPopulationReproductiveFactory(positionGenerator, neuralNetworkReproductiveFactory)

	var reproduced = populationReproductiveFactory.Make(firstPopulation)

	fmt.Println(reproduced)
}
