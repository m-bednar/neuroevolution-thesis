package main

type ExecutionTerminator struct {
	stats     *StatsGatherer
	arguments *ProgramArguments
}

func NewExecutionTerminator(stats *StatsGatherer, arguments *ProgramArguments) *ExecutionTerminator {
	return &ExecutionTerminator{stats, arguments}
}

func (terminator *ExecutionTerminator) ShouldTerminate(generation int, population Population) bool {
	return generation >= terminator.arguments.maxGenerations
}
