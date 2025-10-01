package concurrency

import (
	"sync"
)

// ParallelMap applies a function f to each element of s in parallel and returns a new slice of results.
func ParallelMap[K any, T any](s []K, f func(K) T) []T {
	result := make([]T, len(s))
	var wg sync.WaitGroup
	for i, v := range s {
		wg.Add(1)
		go func(i int, v K) {
			defer wg.Done()
			result[i] = f(v)
		}(i, v)
	}
	wg.Wait()
	return result
}

// ParallelForEach applies a function f to each element of s in parallel (no result slice).
func ParallelForEach[K any](s []K, f func(K)) {
	var wg sync.WaitGroup
	for _, v := range s {
		wg.Add(1)
		go func(v K) {
			defer wg.Done()
			f(v)
		}(v)
	}
	wg.Wait()
}
