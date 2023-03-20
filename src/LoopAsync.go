package main

import (
	"runtime"
	"sync"
)

type LoopFunc[Value any] func(index int, value Value)

func LoopAsync[Value any](items []Value, loopFunc LoopFunc[Value]) {
	var wg = sync.WaitGroup{}
	var limit = runtime.NumCPU() + 1
	var limiter = make(chan struct{}, limit)
	wg.Add(len(items))
	for i, v := range items {
		limiter <- struct{}{}
		go func(index int, value Value) {
			loopFunc(index, value)
			<-limiter
			wg.Done()
		}(i, v)
	}
	wg.Wait()
}
