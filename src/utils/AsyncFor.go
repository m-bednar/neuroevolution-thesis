package utils

import (
	"runtime"
	"sync"
)

type ForFunc[T any] func(index int, value T)

/*
Runs the *fn* function on every single element of *items* concurrently, while
limiting the number of used concurrent goroutines to number of available CPU cores.
The thread running this function is being blocked until all the goroutines finish.
*/
func AsyncFor[T any](items []T, fn ForFunc[T]) {
	wg := sync.WaitGroup{}
	limit := runtime.NumCPU()
	limiter := make(chan struct{}, limit)

	wg.Add(len(items))
	for i, v := range items {
		limiter <- struct{}{}
		go func(index int, value T) {
			fn(index, value)
			<-limiter
			wg.Done()
		}(i, v)
	}
	wg.Wait()
}
