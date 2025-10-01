package concurrency_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/kishankumarhs/fnkit/concurrency"
)

func TestParallelMap(t *testing.T) {
	s := []int{1, 2, 3, 4}
	doubled := concurrency.ParallelMap(s, func(x int) int { return x * 2 })
	want := []int{2, 4, 6, 8}
	for i, v := range want {
		if doubled[i] != v {
			t.Errorf("ParallelMap: got %v, want %v", doubled, want)
		}
	}
}

func TestParallelForEach(t *testing.T) {
	s := []int{1, 2, 3, 4}
	var sum int64
	concurrency.ParallelForEach(s, func(x int) { atomic.AddInt64(&sum, int64(x)) })
	if sum != 10 {
		t.Errorf("ParallelForEach: got %d, want 10", sum)
	}
}

func TestDebounce(t *testing.T) {
	var count int32
	debounced := concurrency.Debounce(func() { atomic.AddInt32(&count, 1) }, 50*time.Millisecond)
	for i := 0; i < 5; i++ {
		debounced()
	}
	time.Sleep(100 * time.Millisecond)
	if count != 1 {
		t.Errorf("Debounce: got %d, want 1", count)
	}
}

func TestThrottle(t *testing.T) {
	var count int32
	throttled := concurrency.Throttle(func() { atomic.AddInt32(&count, 1) }, 50*time.Millisecond)
	for i := 0; i < 5; i++ {
		throttled()
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(100 * time.Millisecond)
	if count < 2 || count > 3 {
		t.Errorf("Throttle: got %d, want 2 or 3", count)
	}
}
