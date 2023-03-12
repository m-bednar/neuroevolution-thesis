package main

type TaskExecutor struct {
	enviroment *Enviroment
	collector  *DataCollector
	evaluator  *FitnessEvaluator
	selector   *ActionSelector
	steps      int
}

func GetNormalizeDistanceToImpassableTile(enviroment *Enviroment, origin Position, direction Direction) float64 {
	var enviromentSize = float64(enviroment.size - 1)
	var distance = enviroment.GetDistanceToImpassableTileInDirection(origin, direction)
	return (distance - 1) / enviromentSize
}

func NewTaskExecutor(enviroment *Enviroment, collector *DataCollector, evaluator *FitnessEvaluator, steps int) *TaskExecutor {
	var selector = NewActionSelector()
	return &TaskExecutor{enviroment, collector, evaluator, selector, steps}
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
		var inputs = executor.MakeNeuralNetworkInputs(microbe)
		var result = microbe.Process(executor.selector, inputs)
		if executor.enviroment.IsPassable(result) {
			microbe.fitness += executor.evaluator.EvaluateMove(microbe.position, result)
			microbe.MoveTo(result)
		}
	}
}

func (executor *TaskExecutor) MakeNeuralNetworkInputs(microbe *Microbe) []float64 {
	var position = microbe.position
	var enviroment = executor.enviroment
	var enviromentSize = float64(enviroment.size)

	var borderDistWest = float64(position.x) / enviromentSize
	var borderDistNorth = float64(position.y) / enviromentSize
	var impDistNorth = GetNormalizeDistanceToImpassableTile(enviroment, position, North)
	var impDistSouth = GetNormalizeDistanceToImpassableTile(enviroment, position, South)
	var impDistWest = GetNormalizeDistanceToImpassableTile(enviroment, position, West)
	var impDistEast = GetNormalizeDistanceToImpassableTile(enviroment, position, East)

	return []float64{borderDistWest, borderDistNorth, impDistNorth, impDistSouth, impDistWest, impDistEast}
}
