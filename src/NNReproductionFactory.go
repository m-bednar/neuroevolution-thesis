package main

type NNReproductionFactory struct{}

func NewNNReproductionFactory() *NNReproductionFactory {
	return &NNReproductionFactory{}
}

func Recombine(weights1 []float64, weights2 []float64) []float64 {
	var size = len(weights1)
	var result = make([]float64, size)
	for i := 0; i < size; i++ {
		result[i] = (weights1[i] + weights2[i]) / 2
	}
	return result
}

func (factory *NNReproductionFactory) Make(parent1 *Microbe, parent2 *Microbe) NeuralNetwork {
	var nn1 = parent1.neuralNetwork
	var nn2 = parent2.neuralNetwork
	var weights = Recombine(nn1.weights, nn2.weights)
	return NewNeuralNetwork(weights)
}
