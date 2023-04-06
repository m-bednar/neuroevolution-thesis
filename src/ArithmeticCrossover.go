package main

/*
Crossover strategy, that averages two weights vectors into single one.
*/
type ArithmeticCrossover struct {
}

func NewArithmeticCrossover() *ArithmeticCrossover {
	return &ArithmeticCrossover{}
}

func (crossover *ArithmeticCrossover) Crossover(weights1 []float64, weights2 []float64) []float64 {
	size := len(weights1)
	result := make([]float64, size)
	for i := 0; i < size; i++ {
		result[i] = (weights1[i] + weights2[i]) / 2
	}
	return result
}
