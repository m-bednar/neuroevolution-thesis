package main

import "fmt"

const POP_SIZE = 10
const ENV_SIZE = 8
const STEPS = 10

func main() {
	var positionGenerator = NewPositionGenerator(ENV_SIZE)
	var neuralNetworkRandomFactory = NewNeuralNetworkRandomFactory()
	var populationRandomFactory = NewPopulationRandomFactory(positionGenerator, neuralNetworkRandomFactory)

	var firstPopulation = populationRandomFactory.Make(POP_SIZE)

	var neuralNetworkReproductiveFactory = NewNeuralNetworkReproductionFactory()
	var populationReproductiveFactory = NewPopulationReproductiveFactory(positionGenerator, neuralNetworkReproductiveFactory)

	var reproduced = populationReproductiveFactory.Make(firstPopulation, POP_SIZE)

	var tiles = []TileType{
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, SafeZone, SafeZone, SafeZone, Empty, Empty, Empty, Empty,
		Empty, SafeZone, SafeZone, SafeZone, Empty, Empty, Empty, Empty,
		Empty, SafeZone, SafeZone, SafeZone, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
	}

	var enviroment = NewEnviroment(tiles, ENV_SIZE)

	var executor = NewExecutor(enviroment, STEPS)
	executor.Execute(reproduced)

	var evaluator = NewEvaluator(enviroment)
	var selector = NewPopulationSelector(evaluator)
	var selected = selector.SelectFrom(reproduced)

	fmt.Println(len(selected), selected)
}
