package main

type Executor struct {
	enviroment Enviroment
	steps      int
}

func NewExecutor(enviroment Enviroment, steps int) Executor {
	return Executor{enviroment, steps}
}

func (executor *Executor) Execute(population []*Microbe) {
	for i := 0; i < executor.steps; i++ {
		executor.ExecutePopulation(population)
	}
}

func (executor *Executor) ExecutePopulation(population []*Microbe) {
	for i := 0; i < len(population); i++ {
		var microbe = population[i]
		var inputs = executor.MakeNeuralNetworkInputs(microbe)
		var result = microbe.Process(inputs)
		if executor.enviroment.IsPassable(result) {
			microbe.MoveTo(result)
		}
	}
}

func (executor *Executor) MakeNeuralNetworkInputs(microbe *Microbe) []float64 {
	var borderDistWest = float64(microbe.position.x) / float64(executor.enviroment.size)
	var borderDistNorth = float64(microbe.position.y) / float64(executor.enviroment.size)
	return []float64{borderDistWest, borderDistNorth}
}
