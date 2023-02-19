package main

type Evaluator struct {
	executor Executor
}

func NewEvaluator(executor Executor, enviroment Enviroment) Evaluator {
	return Evaluator{
		executor: executor,
	}
}

func (evaluator *Evaluator) Evaluate() {
	// TODO
}
