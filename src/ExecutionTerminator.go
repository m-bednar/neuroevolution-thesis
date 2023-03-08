package main

type ExecutionTerminator struct {
	stats     *StatsGatherer
	arguments *ProgramArguments
}

func NewExecutionTerminator(stats *StatsGatherer, arguments *ProgramArguments) *ExecutionTerminator {
	return &ExecutionTerminator{stats, arguments}
}

func (terminator *ExecutionTerminator) ShouldTerminate(generation int, population []*Microbe) bool {
	var successRate = terminator.stats.GetSuccessRate(population)
	if successRate >= terminator.arguments.minSuccessRate {
		return true
	}
	if generation >= terminator.arguments.maxGenerations {
		return true
	}
	return false
}
