package main

import (
	"fmt"
)

const POP_SIZE = 150
const ENV_SIZE = 15
const STEPS = ENV_SIZE * 2

func main() {
	var positionGenerator = NewPositionGenerator(ENV_SIZE)
	var neuralNetworkRandomFactory = NewNeuralNetworkRandomFactory()
	var populationRandomFactory = NewPopulationRandomFactory(positionGenerator, neuralNetworkRandomFactory)

	var firstPopulation = populationRandomFactory.Make(POP_SIZE)

	var mutator = NewMutator(10, 0.2)
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
	MainLoop(firstPopulation, executor, selector, evaluator, populationRandomFactory, populationReproductiveFactory)
}

func MainLoop(firstPopulation []*Microbe, executor Executor, selector PopulationSelector, evaluator Evaluator, populationRandomFactory PopulationRandomFactory, populationReproductiveFactory PopulationReproductiveFactory) {
	var population = firstPopulation
	var successfulness = 0.0

	for successfulness < 0.99 {
		executor.Execute(population)

		var selected = selector.SelectFrom(population, 0.1)
		var nSuccess = evaluator.GetNumberOfMicrobesAtSafeZone(population)
		successfulness = float64(nSuccess) / float64(POP_SIZE)

		if nSuccess == 0 {
			population = populationRandomFactory.Make(POP_SIZE)
		} else {
			population = populationReproductiveFactory.Make(selected, POP_SIZE)
		}

		fmt.Printf("%.2f%% (%d)\n", successfulness*100.0, nSuccess)
	}
}
