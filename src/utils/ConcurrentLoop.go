package utils

import (
	"runtime"
	"sync"
)

type LoopFunc[T any] func(index int, value T)

func ConcurrentLoop[T any](items []T, loopFunc LoopFunc[T]) {
	wg := sync.WaitGroup{}
	limit := runtime.NumCPU()
	limiter := make(chan struct{}, limit)

	wg.Add(len(items))
	for i, v := range items {
		limiter <- struct{}{}
		go func(index int, value T) {
			loopFunc(index, value)
			<-limiter
			wg.Done()
		}(i, v)
	}
	wg.Wait()
}
