package main

type CrossoverStrategy interface {
	Crossover(weights1 []float64, weights2 []float64) []float64
}

type NNReproductionFactory struct {
	strategy CrossoverStrategy
}

func NewNNReproductionFactory(strategy CrossoverStrategy) *NNReproductionFactory {
	return &NNReproductionFactory{
		strategy: strategy,
	}
}

func (factory *NNReproductionFactory) Make(parent1 *Microbe, parent2 *Microbe) NeuralNetwork {
	var nn1 = parent1.neuralNetwork
	var nn2 = parent2.neuralNetwork
	var weights = factory.strategy.Crossover(nn1.weights, nn2.weights)
	return NewNeuralNetwork(weights)
}
