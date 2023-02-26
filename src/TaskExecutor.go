package main

type TaskExecutor struct {
	enviroment *Enviroment
	evaluator  *FitnessEvaluator
	steps      int
}

func NewTaskExecutor(enviroment *Enviroment, evaluator *FitnessEvaluator, steps int) *TaskExecutor {
	return &TaskExecutor{enviroment, evaluator, steps}
}

func (executor *TaskExecutor) ExecuteTask(population []*Microbe) {
	for i := 0; i < executor.steps; i++ {
		executor.ExecuteStep(population)
	}
}

func (executor *TaskExecutor) ExecuteStep(population []*Microbe) {
	for _, microbe := range population {
		var inputs = executor.MakeNeuralNetworkInputs(microbe)
		var result = microbe.Process(inputs)
		if executor.enviroment.IsPassable(result) {
			microbe.fitness += executor.evaluator.EvaluateMove(microbe.position, result)
			microbe.MoveTo(result)
		}
	}
}

func (executor *TaskExecutor) MakeNeuralNetworkInputs(microbe *Microbe) []float64 {
	var borderDistWest = float64(microbe.position.x) / float64(executor.enviroment.size)
	var borderDistNorth = float64(microbe.position.y) / float64(executor.enviroment.size)
	return []float64{borderDistWest, borderDistNorth}
}
