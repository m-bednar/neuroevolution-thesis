package main

import "fmt"

const POP_SIZE = 1
const ENV_SIZE = 15
const STEPS = 1
const MAX_GEN = 1

func main() {
	var positionGenerator = NewPositionGenerator(ENV_SIZE)
	var neuralNetworkRandomFactory = NewNeuralNetworkRandomFactory()
	var populationRandomFactory = NewPopulationRandomFactory(positionGenerator, neuralNetworkRandomFactory)

	var firstPopulation = populationRandomFactory.Make(POP_SIZE)

	var mutator = NewMutator(0, 0.1)
	var neuralNetworkReproductiveFactory = NewNeuralNetworkReproductionFactory()
	var populationReproductiveFactory = NewPopulationReproductiveFactory(positionGenerator, neuralNetworkReproductiveFactory, mutator)

	var tiles = []TileType{
		SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone,
		SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone,
		SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone,
		SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone, SafeZone,
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
			var random = populationRandomFactory.Make(POP_SIZE / 4)
			var reproduced = populationReproductiveFactory.Make(selected, (POP_SIZE/4)*3)
			population = append(random, reproduced...)
		}

		if i%10000 == 0 {
			fmt.Println("SUCCESS:", float64(len(selected))/float64(POP_SIZE))
		}
	}

}
