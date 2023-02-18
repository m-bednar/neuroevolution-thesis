package main

const POP_SIZE = 5

func main() {
	var positionGenerator = NewPositionGenerator(10, 10)
	var neuralNetworkRandomFactory = NewNeuralNetworkRandomFactory()
	var populationRandomFactory = NewPopulationRandomFactory(positionGenerator, neuralNetworkRandomFactory)

	var firstPopulation = populationRandomFactory.Make(POP_SIZE)

	var neuralNetworkReproductiveFactory = NewNeuralNetworkReproductionFactory()
	var populationReproductiveFactory = NewPopulationReproductiveFactory(positionGenerator, neuralNetworkReproductiveFactory)

	var reproduced = populationReproductiveFactory.Make(firstPopulation, POP_SIZE)[0]

	for i := 0; i < 5; i++ {
		reproduced.Process([]float64{0, float64(i) / 2.0, float64(i) / 4.0, 0})
	}
}
