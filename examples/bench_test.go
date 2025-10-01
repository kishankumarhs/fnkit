package main

import (
	"strings"
	"testing"

	"github.com/kishankumarhs/fnkit"
	"github.com/kishankumarhs/fnkit/concurrency"
	"github.com/kishankumarhs/fnkit/validations"
)

func BenchmarkParallelMap(b *testing.B) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = concurrency.ParallelMap(data, func(x int) int { return x * 2 })
	}
}

func BenchmarkStdlibMap(b *testing.B) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := make([]int, len(data))
		for j, v := range data {
			out[j] = v * 2
		}
	}
}

func BenchmarkToFilter(b *testing.B) {
	emails := []string{"foo@bar.com", "bademail", "test@example.org", "hello@world.com"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fnkit.ToFilter(emails, validations.IsEmail)
	}
}

func BenchmarkStdlibFilter(b *testing.B) {
	emails := []string{"foo@bar.com", "bademail", "test@example.org", "hello@world.com"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var out []string
		for _, e := range emails {
			if strings.Contains(e, "@") && strings.Contains(e, ".") {
				out = append(out, e)
			}
		}
	}
}

func BenchmarkParallelForEach(b *testing.B) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrency.ParallelForEach(data, func(x int) { _ = x * 2 })
	}
}

func BenchmarkStdlibForEach(b *testing.B) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range data {
			_ = v * 2
		}
	}
}

func BenchmarkGroupBy(b *testing.B) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i % 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fnkit.GroupBy(data, func(x int) int { return x })
	}
}

func BenchmarkUnique(b *testing.B) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i % 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fnkit.Unique(data)
	}
}

func BenchmarkIsEmail(b *testing.B) {
	emails := []string{"foo@bar.com", "bademail", "test@example.org", "hello@world.com"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range emails {
			_ = validations.IsEmail(e)
		}
	}
}
