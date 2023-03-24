package strategy

/*
Crossover strategy, that averages two weights vectors into single one.
*/
type ArithmeticCrossoverStrategy struct {
}

func NewArithmeticCrossoverStrategy() *ArithmeticCrossoverStrategy {
	return &ArithmeticCrossoverStrategy{}
}

func (strategy *ArithmeticCrossoverStrategy) Crossover(weights1 []float64, weights2 []float64) []float64 {
	var size = len(weights1)
	var result = make([]float64, size)
	for i := 0; i < size; i++ {
		result[i] = (weights1[i] + weights2[i]) / 2
	}
	return result
}
