package utils

import (
	"runtime"
	"sync"
)

type LimitFlag struct{}
type LoopFunc[Value any] func(index int, value Value)

func ConcurrentLoop[Value any](items []Value, loopFunc LoopFunc[Value]) {
	var wg = sync.WaitGroup{}
	var limit = runtime.NumCPU() * 2
	var limiter = make(chan LimitFlag, limit)

	wg.Add(len(items))
	for i, v := range items {
		limiter <- LimitFlag{}
		go func(index int, value Value) {
			loopFunc(index, value)
			<-limiter
			wg.Done()
		}(i, v)
	}
	wg.Wait()
}
