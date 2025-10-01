# fnkit

A collection of functional programming utilities for Golang.

## Features

- Higher-order functions
- Map, Filter, Reduce utilities
- Function composition

## Installation

```bash
go get github.com/kishankumarhs/fnkit
```

## Usage

### Map

Apply a function to each element in a slice.

```go
import "github.com/kishankumarhs/fnkit"

nums := []int{1, 2, 3}
doubled := fnkit.Map(nums, func(x int) int { return x * 2 })
// doubled: [2, 4, 6]
```

### Filter

Filter elements in a slice based on a predicate.

```go
evens := fnkit.Filter(nums, func(x int) bool { return x%2 == 0 })
// evens: [2]
```

### Reduce

Reduce a slice to a single value.

```go
sum := fnkit.Reduce(nums, 0, func(acc, x int) int { return acc + x })
// sum: 6
```

### Compose

Compose multiple functions into one.

```go
double := func(x int) int { return x * 2 }
increment := func(x int) int { return x + 1 }
composed := fnkit.Compose(double, increment)
result := composed(3) // (3 + 1) * 2 = 8
```

## License

MIT