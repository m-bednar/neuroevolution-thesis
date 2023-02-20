package main

import "fmt"

const POP_SIZE = 200
const ENV_SIZE = 15
const STEPS = 20
const MAX_GEN = 1000

func main() {
	var positionGenerator = NewPositionGenerator(ENV_SIZE)
	var neuralNetworkRandomFactory = NewNeuralNetworkRandomFactory()
	var populationRandomFactory = NewPopulationRandomFactory(positionGenerator, neuralNetworkRandomFactory)

	var firstPopulation = populationRandomFactory.Make(POP_SIZE)

	var mutator = NewMutator(1, 0.1)
	var neuralNetworkReproductiveFactory = NewNeuralNetworkReproductionFactory()
	var populationReproductiveFactory = NewPopulationReproductiveFactory(positionGenerator, neuralNetworkReproductiveFactory, mutator)

	var tiles = []TileType{
		SafeZone, SafeZone, SafeZone, SafeZone, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		SafeZone, SafeZone, SafeZone, SafeZone, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		SafeZone, SafeZone, SafeZone, SafeZone, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		SafeZone, SafeZone, SafeZone, SafeZone, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
	}

	var enviroment = NewEnviroment(tiles, ENV_SIZE)
	var executor = NewExecutor(enviroment, STEPS)
	var evaluator = NewEvaluator(enviroment)
	var selector = NewPopulationSelector(evaluator)

	var population = firstPopulation

	for i := 0; i < MAX_GEN; i++ {
		executor.Execute(population)
		var selected = selector.SelectFrom(population)

		if len(selected) == 0 {
			population = populationRandomFactory.Make(POP_SIZE)
		} else {
			population = populationReproductiveFactory.Make(selected, POP_SIZE)
		}

		if i%5 == 0 {
			fmt.Printf("SUCCESS: %.2f%% (%d)\n", (float64(len(selected)) / float64(POP_SIZE) * 100.0), len(selected))
		}
	}

}
