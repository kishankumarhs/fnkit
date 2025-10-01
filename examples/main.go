package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/kishankumarhs/fnkit"
	"github.com/kishankumarhs/fnkit/concurrency"
	"github.com/kishankumarhs/fnkit/validations"
)

func main() {
	// --- Real-world: Email validation and normalization ---
	emails := []string{"foo@bar.com", "bademail", "test@example.org"}
	valid := fnkit.ToFilter(emails, validations.IsEmail)
	fmt.Println("Valid emails:", valid)

	// --- Real-world: Parallel processing of data ---
	nums := []int{1, 2, 3, 4, 5}
	doubled := concurrency.ParallelMap(nums, func(x int) int { return x * 2 })
	fmt.Println("Doubled in parallel:", doubled)

	// --- Real-world: Debounce user input (simulate rapid calls) ---
	debounced := concurrency.Debounce(func() { fmt.Println("Debounced action!") }, 100*time.Millisecond)
	for i := 0; i < 5; i++ {
		debounced()
	}
	time.Sleep(200 * time.Millisecond)

	// --- Real-world: String utilities for slug generation ---
	titles := []string{"Go is Awesome!", "Hello, World!"}
	slugs := fnkit.Map(titles, func(s string) string {
		s = strings.ToLower(s)
		s = strings.ReplaceAll(s, " ", "-")
		s = strings.ReplaceAll(s, "!", "")
		return s
	})
	fmt.Println("Slugs:", slugs)

	// --- Real-world: GroupBy for analytics ---
	ages := []int{21, 34, 21, 40, 34, 21}
	grouped := fnkit.GroupBy(ages, func(age int) string {
		if age < 30 {
			return "young"
		}
		return "adult"
	})
	fmt.Println("Grouped ages:", grouped)

	// --- Real-world: Validate and convert user input ---
	inputs := []any{"42", 17, "notanumber", 3.14}
	for _, v := range inputs {
		if n, ok := validations.ToInt(v); ok {
			fmt.Printf("%v as int: %d\n", v, n)
		} else {
			fmt.Printf("%v is not an int\n", v)
		}
	}

	// --- Real-world: ParallelForEach for concurrent side effects ---
	words := []string{"go", "is", "fast"}
	concurrency.ParallelForEach(words, func(w string) {
		fmt.Println("Word:", strings.ToUpper(w))
	})

	// --- Real-world: Throttle for rate-limited logging ---
	throttled := concurrency.Throttle(func() { fmt.Println("Throttled log!") }, 100*time.Millisecond)
	for i := 0; i < 5; i++ {
		throttled()
		time.Sleep(30 * time.Millisecond)
	}
	// Wait to see throttled effect
	time.Sleep(200 * time.Millisecond)
}
