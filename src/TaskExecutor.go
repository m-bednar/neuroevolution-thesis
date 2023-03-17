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
		executor.collector.CollectStep(i, population)
	}
	for _, microbe := range population {
		microbe.fitness += executor.evaluator.GetEvaluation(microbe)
	}
	executor.collector.CollectGeneration(generation, population)
}

func (executor *TaskExecutor) ExecuteStep(population Population) {
	for _, microbe := range population {
		var inputs = executor.inputsMaker.MakeInputsFor(microbe)
		var output = microbe.Process(inputs)
		var action = executor.selector.SelectMoveAction(output)
		var result = microbe.position.AddToDirection(action)
		if executor.enviroment.IsPassable(result) {
			microbe.MoveTo(result)
		}
	}
}
