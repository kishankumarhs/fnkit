
# fnkit

**fnkit** is a modern Go utility library inspired by the best of JavaScript (like Lodash, Array methods) and Rust (Result type, functional error handling). It brings expressive, type-safe, and composable utilities to Go, making your code more concise, robust, and fun to write.


## Installation

```sh
go get github.com/kishankumarhs/fnkit
```

## Import

```go
import "github.com/kishankumarhs/fnkit"
```





## String Utilities

fnkit provides a comprehensive set of string utility functions inspired by Python, JavaScript, and Rust. All are Unicode-safe and tested.

### Repeat

```go
fnkit.Repeat("ab", 3) // "ababab"
```

### ReplaceAll

```go
fnkit.ReplaceAll("foo bar foo", "foo", "baz") // "baz bar baz"
```

### HasPrefix / HasSuffix

```go
fnkit.HasPrefix("hello world", "hello") // true
fnkit.HasSuffix("hello world", "world") // true
```

### Contains / Count

```go
fnkit.Contains("banana", "nan") // true
fnkit.Count("banana", "na") // 2
```

### ReverseString

```go
fnkit.ReverseString("a😊b") // "b😊a"
```

### IsAlpha / IsNumeric

```go
fnkit.IsAlpha("abcXYZ") // true
fnkit.IsAlpha("abc123") // false
fnkit.IsNumeric("12345") // true
fnkit.IsNumeric("12a45") // false
```

### Capitalize

```go
fnkit.Capitalize("hELLO") // "Hello"
```

### StripLeft / StripRight

```go
fnkit.StripLeft("  hello  ") // "hello  "
fnkit.StripRight("  hello  ") // "  hello"
```

### Partition / Rpartition

```go
fnkit.Partition("foo-bar-baz", "-") // ("foo", "-", "bar-baz")
fnkit.Rpartition("foo-bar-baz", "-") // ("foo-bar", "-", "baz")
```

### Words

```go
fnkit.Words("Go is  awesome") // []string{"Go", "is", "awesome"}
```

### CamelCase / SnakeCase / KebabCase

```go
fnkit.CamelCase("hello world_test-case") // "helloWorldTestCase"
fnkit.SnakeCase("hello world_test-case") // "hello_world_test_case"
fnkit.KebabCase("hello world_test-case") // "hello-world-test-case"
```

### IsUpper / IsLower

```go
fnkit.IsUpper("ABC") // true
fnkit.IsUpper("AbC") // false
fnkit.IsLower("abc") // true
fnkit.IsLower("aBc") // false
```

### SwapCase

```go
fnkit.SwapCase("aBc") // "AbC"
```

### Remove / Keep

```go
fnkit.Remove("a1b2c3", unicode.IsDigit) // "abc"
fnkit.Keep("a1b2c3", unicode.IsDigit) // "123"
```

### PadCenter

```go
fnkit.PadCenter("hi", 6, '*') // "**hi**"
```

----


## Usage & Examples





## GroupBy

```go
nums := []int{1, 2, 3, 4, 5, 6}
grouped := fnkit.GroupBy(nums, func(n int) string {
    if n%2 == 0 {
        return "even"
    }
    return "odd"
})
// grouped["even"] == []int{2,4,6}
// grouped["odd"] == []int{1,3,5}
```

## Chunk

```go
nums := []int{1, 2, 3, 4, 5}
chunks := fnkit.Chunk(nums, 2) // [][]int{{1,2},{3,4},{5}}
```

## Unique

```go
nums := []int{1,2,2,3,1,4}
uniq := fnkit.Unique(nums) // []int{1,2,3,4}
```

## Flatten

```go
nested := [][]int{{1,2},{3,4}}
flat := fnkit.Flatten(nested) // []int{1,2,3,4}
```

## Result[T] (Rust-like error handling)

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


### Option[T] (Rust-like optional values)

`Option[T]` is a generic container for an optional value, inspired by Rust. It allows for safe, idiomatic handling of values that may or may not be present, without using nil pointers.

#### Example: Some and None

```go
opt := fnkit.Some(42)
if opt.IsSome() {
    fmt.Println("Value:", opt.Unwrap()) // prints 42
}

none := fnkit.None[int]()
if none.IsNone() {
    fmt.Println("No value")
}
```

#### Example: UnwrapOr

```go
opt := fnkit.Some("hello")
val := opt.UnwrapOr("default") // val == "hello"

none := fnkit.None[string]()
val2 := none.UnwrapOr("default") // val2 == "default"
```

#### Edge Cases (Option[T])

- `Some(zeroValue)` is valid and `IsSome()` is true.
- `None[T]()` is always `IsNone()`.
- Works with any type, including structs.


## GroupBy (slice)

```go
nums := []int{1, 2, 3, 4, 5, 6}
grouped := fnkit.GroupBy(nums, func(n int) string {
    if n%2 == 0 {
        return "even"
    }
    return "odd"
})
// grouped["even"] == []int{2,4,6}
// grouped["odd"] == []int{1,3,5}
```


## Chunk (slice)

```go
nums := []int{1, 2, 3, 4, 5}
chunks := fnkit.Chunk(nums, 2) // [][]int{{1,2},{3,4},{5}}
```


## Unique (slice)

```go
nums := []int{1,2,2,3,1,4}
uniq := fnkit.Unique(nums) // []int{1,2,3,4}
```


## Flatten (slice)

```go
nested := [][]int{{1,2},{3,4}}
flat := fnkit.Flatten(nested) // []int{1,2,3,4}
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

### Any

```go
s := []int{1, 2, 3}
fnkit.Any(s, func(i int) bool { return i == 2 }) // true
fnkit.Any(s, func(i int) bool { return i == 100 }) // false
fnkit.Any([]int{}, func(i int) bool { return true }) // false
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

----

## Edge Cases

- All functions handle empty slices gracefully.
- Functions like `At`, `Pop`, `Shift` return the zero value and false if out of bounds or empty.
- Negative indices are supported in `At`, `Slice`, and `Splice` (like Python).
- `Filter` modifies the slice in-place; `ToFilter` returns a new filtered slice.
- `Every` on an empty slice returns true (vacuous truth); `Any` on empty returns false.
- All functions are safe for any type (thanks to Go generics).

----

## License

MIT
