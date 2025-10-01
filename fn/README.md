# fnkit/fn

Advanced Functional Utilities for Go

This module provides:
- **Generic, thread-safe Map and Set types**
- **DeepEqual** and **DeepCopy** for complex/nested structures
- **Functional pipelines** for chainable slice operations

## Usage

### Map/Set
```go
import "github.com/kishankumarhs/fnkit/fn"

m := fn.NewMap[string, int]()
m.Set("foo", 42)
v, ok := m.Get("foo")

s := fn.NewSet[int]()
s.Add(1)
if s.Has(1) { /* ... */ }
```

### DeepEqual / DeepCopy
```go
import "github.com/kishankumarhs/fnkit/fn"

a := map[string]int{"x": 1}
b := map[string]int{"x": 1}
fn.DeepEqual(a, b) // true

c := fn.DeepCopy(a)
```

### Functional Pipelines
```go
import "github.com/kishankumarhs/fnkit/fn"

result := fn.FromSlice([]int{1,2,3,4,5}).
    Filter(func(x int) bool { return x%2==0 }).
    Map(func(x int) int { return x*10 }).
    Slice() // [20, 40]
```

## License
MIT
