package main

type TaskExecutor struct {
	enviroment *Enviroment
	capturer   *VideoCapturer
	evaluator  *FitnessEvaluator
	selector   *ActionSelector
	steps      int
}

func NewTaskExecutor(enviroment *Enviroment, capturer *VideoCapturer, evaluator *FitnessEvaluator, selector *ActionSelector, steps int) *TaskExecutor {
	return &TaskExecutor{enviroment, capturer, evaluator, selector, steps}
}

func (executor *TaskExecutor) ExecuteTask(generation int, population []*Microbe) {
	for i := 0; i < executor.steps; i++ {
		executor.ExecuteStep(population)
		executor.capturer.CaptureScene(generation, population)
	}
	for _, microbe := range population {
		microbe.fitness += executor.evaluator.GetFinalEvaluation(microbe)
	}
}

func (executor *TaskExecutor) ExecuteStep(population []*Microbe) {
	for _, microbe := range population {
		var inputs = executor.MakeNeuralNetworkInputs(microbe)
		var result = microbe.Process(executor.selector, inputs)
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
