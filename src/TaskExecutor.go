package main

type TaskExecutor struct {
	enviroment  *Enviroment
	collector   *DataCollector
	evaluator   *FitnessEvaluator
	selector    *ActionSelector
	inputsMaker *NeuralInputsMaker
	steps       int
}

func NewTaskExecutor(enviroment *Enviroment, collector *DataCollector, evaluator *FitnessEvaluator, steps int) *TaskExecutor {
	var selector = NewActionSelector()
	var inputsMaker = NewNeuralInputsMaker(enviroment)
	return &TaskExecutor{enviroment, collector, evaluator, selector, inputsMaker, steps}
}

func (executor *TaskExecutor) ExecuteTask(generation int, population Population) {
	for i := 0; i < executor.steps; i++ {
		executor.ExecuteStep(population)
		executor.collector.CollectPositions(generation, population)
	}
	for _, microbe := range population {
		microbe.fitness += executor.evaluator.GetFinalEvaluation(microbe)
	}
}

func (executor *TaskExecutor) ExecuteStep(population Population) {
	for _, microbe := range population {
		var inputs = executor.inputsMaker.MakeInputsFor(microbe)
		var result = microbe.Process(executor.selector, inputs)
		microbe.fitness += executor.evaluator.EvaluateMove(microbe.position, result)
		if executor.enviroment.IsPassable(result) {
			microbe.MoveTo(result)
		}
	}
}
