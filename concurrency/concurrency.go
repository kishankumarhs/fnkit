package concurrency

import (
	"sync"
	"time"
)

// Debounce returns a function that will only invoke fn if no calls have been made for the given duration.
func Debounce(fn func(), d time.Duration) func() {
	var mu sync.Mutex
	var timer *time.Timer
	return func() {
		mu.Lock()
		defer mu.Unlock()
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(d, fn)
	}
}

// Throttle returns a function that will invoke fn at most once per duration d.
func Throttle(fn func(), d time.Duration) func() {
	var mu sync.Mutex
	var last time.Time
	return func() {
		mu.Lock()
		defer mu.Unlock()
		now := time.Now()
		if now.Sub(last) >= d {
			last = now
			go fn()
		}
	}
}
