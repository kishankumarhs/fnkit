package fn_test

import (
	"fmt"
	"github.com/kishankumarhs/fnkit/fn"
)

func ExampleMapSet() {
	m := fn.NewMap[string, int]()
	m.Set("foo", 42)
	m.Set("bar", 7)
	fmt.Println(m.Get("foo")) // 42 true
	fmt.Println(m.Keys())      // [foo bar] (order not guaranteed)

	s := fn.NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(1)
	fmt.Println(s.Has(2)) // true
	fmt.Println(s.Len())  // 2
	// Output:
	// 42 true
	// [foo bar]
	// true
	// 2
}

func ExampleDeepEqualCopy() {

a := map[string]int{"x": 1}
b := map[string]int{"x": 1}
fmt.Println(fn.DeepEqual(a, b)) // true

c := fn.DeepCopy(a)
c["x"] = 99
fmt.Println(a["x"]) // 1
fmt.Println(c["x"]) // 99
// Output:
// true
// 1
// 99
}

func ExamplePipeline() {
	result := fn.FromSlice([]int{1, 2, 3, 4, 5}).
		Filter(func(x int) bool { return x%2 == 0 }).
		Map(func(x int) int { return x * 10 }).
		Slice()
	fmt.Println(result) // [20 40]
	// Output:
	// [20 40]
}
