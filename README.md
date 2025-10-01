
# fnkit

**fnkit** is a modern Go utility library inspired by the best of JavaScript (like Lodash, Array methods) and Rust (Result type, functional error handling). It brings expressive, type-safe, and composable utilities to Go, making your code more concise, robust, and fun to write.

`fnkit` provides a comprehensive set of generic, functional utilities for working with Go slices and error handling. All functions are generic and work with any type.

## Installation

```sh
go get github.com/kishankumarhs/fnkit
```

## Import

```go
import "github.com/kishankumarhs/fnkit"
```



## Usage & Examples


### Result[T] (Rust-like error handling)

`Result[T]` is a generic container for a value or an error, inspired by Rust. It allows for functional error handling without repetitive `if err != nil` checks.

#### Example: Database Query


```go
res := fnkit.Ok("user1")
if res.IsOk() {
    fmt.Println("User:", res.Value)
} else {
    fmt.Println("Error:", res.Err)
}

// Using ValueOr for fallback
name := res.ValueOr("default")
```

#### Example: Network Call


```go
call := fnkit.Err[int](errors.New("network error"))
code := call.ValueOr(404) // returns 404 if error
```

#### Example: Chaining


```go
db := mockDBQuery(1) // returns Result[string]
if db.IsOk() {
    fmt.Println("Found:", db.Value)
} else {
    fmt.Println("DB error:", db.Err)
}

net := mockNetworkCall("https://fail")
if !net.IsOk() {
    fmt.Println("Network error:", net.Err)
}
```


#### Edge Cases (Result[T])

- `Ok(zeroValue)` is valid and `IsOk()` is true.
- `Err[T](nil)` is treated as Ok (no error).
- Works with any type, including structs.


### Map

```go
nums := []int{1, 2, 3}
strs := fnkit.Map(nums, func(n int) string { return fmt.Sprintf("n=%d", n) })
// strs == []string{"n=1", "n=2", "n=3"}
```

### Reduce

```go
sum := fnkit.Reduce([]int{1, 2, 3}, 0, func(acc, v int) int { return acc + v })
// sum == 6
```

### At

```go
s := []string{"a", "b", "c"}
fnkit.At(s, 1) // "b"
fnkit.At(s, -1) // "c"
fnkit.At(s, 10) // "" (zero value)
fnkit.At(s, -10) // "" (zero value)
```

### Concat

```go
a := []int{1, 2}
b := []int{3, 4}
fnkit.Concat(a, b) // []int{1, 2, 3, 4}
```

### CopyWith

```go
orig := []int{1, 2, 3}
copy := fnkit.CopyWith(orig)
copy[0] = 99 // orig is unchanged
```

### Entries

```go
s := []string{"x", "y"}
entries := fnkit.Entries(s)
// entries == []fnkit.Entry[string]{{0, "x"}, {1, "y"}}
```

### Every

```go
allEven := fnkit.Every([]int{2, 4, 6}, func(i int) bool { return i%2 == 0 }) // true
allPositive := fnkit.Every([]int{1, 2, 3}, func(i int) bool { return i > 0 }) // true
empty := fnkit.Every([]int{}, func(i int) bool { return i > 0 }) // true (vacuous truth)
```

### Fill

```go
s := []int{1, 2, 3}
fnkit.Fill(s, 9) // s == []int{9, 9, 9}
```

### Find / FindIndex / FindLast / FindLastIndex

```go
s := []int{1, 2, 3, 2}
v, ok := fnkit.Find(s, func(i int) bool { return i == 2 }) // v==2, ok==true
idx := fnkit.FindIndex(s, func(i int) bool { return i == 2 }) // idx==1
v, ok = fnkit.FindLast(s, func(i int) bool { return i == 2 }) // v==2, ok==true
idx = fnkit.FindLastIndex(s, func(i int) bool { return i == 2 }) // idx==3
_, ok = fnkit.Find(s, func(i int) bool { return i == 100 }) // ok==false
```

### Flat / FlatMap

```go
nested := [][]int{{1, 2}, {3, 4}}
flat := fnkit.Flat(nested) // []int{1, 2, 3, 4}

nums := []int{1, 2}
flatMapped := fnkit.FlatMap(nums, func(i int) []int { return []int{i, i} }) // [1, 1, 2, 2]
```

### ForEach

```go
sum := 0
fnkit.ForEach([]int{1, 2, 3}, func(i int) { sum += i })
// sum == 6
```

### Includes / IndexOf / LastIndexOf

```go
s := []int{1, 2, 3, 2}
fnkit.Includes(s, 2) // true
fnkit.IndexOf(s, 2) // 1
fnkit.LastIndexOf(s, 2) // 3
fnkit.Includes(s, 99) // false
```

### Join

```go
s := []int{1, 2, 3}
str := fnkit.Join(s, "-") // "1-2-3"
empty := fnkit.Join([]int{}, ",") // ""
```

### Keys / Values

```go
s := []string{"a", "b", "c"}
fnkit.Keys(s) // []int{0, 1, 2}
fnkit.Values(s) // []string{"a", "b", "c"} (copy)
```

### Pop / Push / Shift / Unshift

```go
s := []int{1, 2, 3}
v, ok := fnkit.Pop(&s) // v==3, ok==true, s==[1,2]
fnkit.Push(&s, 4) // s==[1,2,4]
v, ok = fnkit.Shift(&s) // v==1, ok==true, s==[2,4]
fnkit.Unshift(&s, 0) // s==[0,2,4]
empty := []int{}
_, ok = fnkit.Pop(&empty) // ok==false
_, ok = fnkit.Shift(&empty) // ok==false
```

### ReduceRight

```go
s := []int{1, 2, 3}
sum := fnkit.ReduceRight(s, 0, func(acc, v int) int { return acc + v }) // 6
```

### Reverse

```go
s := []int{1, 2, 3}
fnkit.Reverse(s) // s == [3,2,1]
```

### Slice

```go
s := []int{1, 2, 3, 4}
fnkit.Slice(s, 1, 3) // [2,3]
fnkit.Slice(s, -1, 10) // [4]
fnkit.Slice(s, 10, 20) // []
```

### Some

```go
s := []int{1, 2, 3}
fnkit.Some(s, func(i int) bool { return i == 2 }) // true
fnkit.Some(s, func(i int) bool { return i == 100 }) // false
fnkit.Some([]int{}, func(i int) bool { return true }) // false
```

### Splice

```go
s := []int{1, 2, 3, 4}
removed := fnkit.Splice(&s, 1, 2, []int{9, 9}) // removed==[2,3], s==[1,9,9,4]
removed = fnkit.Splice(&s, 10, 2, []int{5}) // removed==[], s==[1,9,9,4,5]
removed = fnkit.Splice(&s, -1, 1, nil) // removed==[5], s==[1,9,9,4]
```

### ToLocaleString

```go
s := []int{1, 2, 3}
str := fnkit.ToLocaleString(s) // "1,2,3"
```

### Without

```go
s := []int{1, 2, 3, 2}
fnkit.Without(s, 2) // [1,3]
```

### Filter (in-place) / ToFilter (returns new slice)

```go
s := []int{1, 2, 3, 4}
fnkit.Filter(&s, func(i int) bool { return i%2 == 0 }) // s==[2,4]
s2 := []int{1, 2, 3, 4}
evens := fnkit.ToFilter(s2, func(i int) bool { return i%2 == 0 }) // evens==[2,4], s2 unchanged
```

---

## Edge Cases

- All functions handle empty slices gracefully.
- Functions like `At`, `Pop`, `Shift` return the zero value and false if out of bounds or empty.
- Negative indices are supported in `At`, `Slice`, and `Splice` (like Python).
- `Filter` modifies the slice in-place; `ToFilter` returns a new filtered slice.
- `Every` on an empty slice returns true (vacuous truth); `Some` on empty returns false.
- All functions are safe for any type (thanks to Go generics).

---

## License

MIT
