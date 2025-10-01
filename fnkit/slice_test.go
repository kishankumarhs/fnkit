package fnkit_test

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/kishankumarhs/fnkit"
)

// Write a test for the ToFilter function
func TestToFilter(t *testing.T) {
	// Test case 1: Filter even numbers from a slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6}
	isEven := func(n int) bool {
		return n%2 == 0
	}
	expectedEven := []int{2, 4, 6}
	resultEven := fnkit.ToFilter(numbers, isEven)
	if len(resultEven) != len(expectedEven) {
		t.Errorf("Expected length %d, got %d", len(expectedEven), len(resultEven))
	}
	for i, v := range expectedEven {
		if resultEven[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, resultEven[i])
		}
	}

	// Test case 2: Filter strings longer than 3 characters
	words := []string{"Go", "is", "fun", "and", "powerful"}
	isLongerThan3 := func(s string) bool {
		return len(s) > 3
	}
	expectedLongWords := []string{"powerful"}
	resultLongWords := fnkit.ToFilter(words, isLongerThan3)
	if len(resultLongWords) != len(expectedLongWords) {
		t.Errorf("Expected length %d, got %d", len(expectedLongWords), len(resultLongWords))
	}
	for i, v := range expectedLongWords {
		if resultLongWords[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, resultLongWords[i])
		}
	}

	// Test case 3: Filter with no matches
	noMatch := func(n int) bool {
		return n > 10
	}
	expectedNoMatch := []int{}
	resultNoMatch := fnkit.ToFilter(numbers, noMatch)
	if len(resultNoMatch) != len(expectedNoMatch) {
		t.Errorf("Expected length %d, got %d", len(expectedNoMatch), len(resultNoMatch))
	}

	// Test case 4: Filter with all matches
	allMatch := func(n int) bool {
		return n > 0
	}
	expectedAllMatch := numbers
	resultAllMatch := fnkit.ToFilter(numbers, allMatch)
	if len(resultAllMatch) != len(expectedAllMatch) {
		t.Errorf("Expected length %d, got %d", len(expectedAllMatch), len(resultAllMatch))
	}
	for i, v := range expectedAllMatch {
		if resultAllMatch[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, resultAllMatch[i])
		}
	}

	// Test case 5: Filter with empty slice
	emptySlice := []int{}
	expectedEmpty := []int{}
	resultEmpty := fnkit.ToFilter(emptySlice, isEven)
	if len(resultEmpty) != len(expectedEmpty) {
		t.Errorf("Expected length %d, got %d", len(expectedEmpty), len(resultEmpty))
	}

	// Test case 6: Filter strings containing 'a'
	stringsWithA := func(s string) bool {
		matched, _ := regexp.MatchString("a", s)
		return matched
	}
	wordsWithA := []string{"apple", "banana", "cherry", "date"}
	expectedWithA := []string{"apple", "banana", "date"}
	resultWithA := fnkit.ToFilter(wordsWithA, stringsWithA)
	if len(resultWithA) != len(expectedWithA) {
		t.Errorf("Expected length %d, got %d", len(expectedWithA), len(resultWithA))
	}
	for i, v := range expectedWithA {
		if resultWithA[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, resultWithA[i])
		}
	}
}

func TestAt(t *testing.T) {
	// Test case 1: Valid index
	numbers := []int{10, 20, 30, 40, 50}
	if val := fnkit.At(numbers, 2); val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}

	// Test case 2: Negative index
	if val := fnkit.At(numbers, -1); val != 50 {
		t.Errorf("Expected 50, got %d", val)
	}

	// Test case 3: Out of bounds index
	if val := fnkit.At(numbers, 5); val != 0 {
		t.Errorf("Expected 0 (zero value), got %d", val)
	}

	// Test case 4: Negative out of bounds index
	if val := fnkit.At(numbers, -6); val != 0 {
		t.Errorf("Expected 0 (zero value), got %d", val)
	}

	// Test case 5: Empty slice
	empty := []int{}
	if val := fnkit.At(empty, 0); val != 0 {
		t.Errorf("Expected 0 (zero value), got %d", val)
	}
}

func TestFilter(t *testing.T) {
	// Test case 1: Filter even numbers from a slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6}
	isEven := func(n int) bool {
		return n%2 == 0
	}
	expectedEven := []int{2, 4, 6}
	fnkit.Filter(&numbers, isEven)
	if len(numbers) != len(expectedEven) {
		t.Errorf("Expected length %d, got %d", len(expectedEven), len(numbers))
	}
	for i, v := range expectedEven {
		if numbers[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, numbers[i])
		}
	}

	// Test case 2: Filter strings longer than 3 characters
	words := []string{"Go", "is", "fun", "and", "powerful"}
	isLongerThan3 := func(s string) bool {
		return len(s) > 3
	}
	expectedLongWords := []string{"powerful"}
	fnkit.Filter(&words, isLongerThan3)
	if len(words) != len(expectedLongWords) {
		t.Errorf("Expected length %d, got %d", len(expectedLongWords), len(words))
	}
	for i, v := range expectedLongWords {
		if words[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, words[i])
		}
	}

	// Test case 3: Filter with no matches
	numbersNoMatch := []int{1, 2, 3}
	noMatch := func(n int) bool {
		return n > 10
	}
	expectedNoMatch := []int{}
	fnkit.Filter(&numbersNoMatch, noMatch)
	if len(numbersNoMatch) != len(expectedNoMatch) {
		t.Errorf("Expected length %d, got %d", len(expectedNoMatch), len(numbersNoMatch))
	}

	// Test case 4: Filter with all matches
	numbersAllMatch := []int{1, 2, 3}
	allMatch := func(n int) bool {
		return n > 0
	}
	expectedAllMatch := []int{1, 2, 3}
	fnkit.Filter(&numbersAllMatch, allMatch)
	if len(numbersAllMatch) != len(expectedAllMatch) {
		t.Errorf("Expected length %d, got %d", len(expectedAllMatch), len(numbersAllMatch))
	}
	for i, v := range expectedAllMatch {
		if numbersAllMatch[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, numbersAllMatch[i])
		}
	}

	// Test case 5: Filter with empty slice
	emptySlice := []int{}
	expectedEmpty := []int{}
	fnkit.Filter(&emptySlice, isEven)
	if len(emptySlice) != len(expectedEmpty) {
		t.Errorf("Expected length %d, got %d", len(expectedEmpty), len(emptySlice))
	}
}

// write a test for all the remaining functions in slice_utils.go

// Test Map
func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []string{"1", "2", "3"}
	got := fnkit.Map(input, func(i int) string { return fmt.Sprint(i) })
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Map() = %v, want %v", got, expected)
	}
}

func TestReduce(t *testing.T) {
	input := []int{1, 2, 3}
	expected := 6
	got := fnkit.Reduce(input, 0, func(acc, v int) int { return acc + v })
	if got != expected {
		t.Errorf("Reduce() = %v, want %v", got, expected)
	}
}

func TestConcat(t *testing.T) {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	expected := []int{1, 2, 3, 4}
	got := fnkit.Concat(s1, s2)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Concat() = %v, want %v", got, expected)
	}
}

func TestCopyWith(t *testing.T) {
	s := []int{1, 2, 3}
	got := fnkit.CopyWith(s)
	if !reflect.DeepEqual(got, s) {
		t.Errorf("CopyWith() = %v, want %v", got, s)
	}
	got[0] = 100
	if s[0] == 100 {
		t.Errorf("CopyWith() did not make a copy")
	}
}

func TestEntries(t *testing.T) {
	s := []string{"a", "b"}
	expected := []fnkit.Entry[string]{{0, "a"}, {1, "b"}}
	got := fnkit.Entries(s)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Entries() = %v, want %v", got, expected)
	}
}

func TestEvery(t *testing.T) {
	s := []int{2, 4, 6}
	if !fnkit.Every(s, func(i int) bool { return i%2 == 0 }) {
		t.Errorf("Every() = false, want true")
	}
	if fnkit.Every(s, func(i int) bool { return i > 2 }) {
		t.Errorf("Every() = true, want false")
	}
}

func TestFill(t *testing.T) {
	s := []int{1, 2, 3}
	fnkit.Fill(s, 9)
	expected := []int{9, 9, 9}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Fill() = %v, want %v", s, expected)
	}
}

func TestFind(t *testing.T) {
	s := []int{1, 2, 3}
	got, ok := fnkit.Find(s, func(i int) bool { return i == 2 })
	if !ok || got != 2 {
		t.Errorf("Find() = %v, %v, want 2, true", got, ok)
	}
	_, ok = fnkit.Find(s, func(i int) bool { return i == 100 })
	if ok {
		t.Errorf("Find() found non-existent element")
	}
}

func TestFindIndex(t *testing.T) {
	s := []int{1, 2, 3}
	if got := fnkit.FindIndex(s, func(i int) bool { return i == 2 }); got != 1 {
		t.Errorf("FindIndex() = %v, want 1", got)
	}
	if got := fnkit.FindIndex(s, func(i int) bool { return i == 100 }); got != -1 {
		t.Errorf("FindIndex() = %v, want -1", got)
	}
}

func TestFindLast(t *testing.T) {
	s := []int{1, 2, 3, 2}
	got, ok := fnkit.FindLast(s, func(i int) bool { return i == 2 })
	if !ok || got != 2 {
		t.Errorf("FindLast() = %v, %v, want 2, true", got, ok)
	}
}

func TestFindLastIndex(t *testing.T) {
	s := []int{1, 2, 3, 2}
	if got := fnkit.FindLastIndex(s, func(i int) bool { return i == 2 }); got != 3 {
		t.Errorf("FindLastIndex() = %v, want 3", got)
	}
}

func TestFlat(t *testing.T) {
	s := [][]int{{1, 2}, {3, 4}}
	expected := []int{1, 2, 3, 4}
	got := fnkit.Flat(s)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Flat() = %v, want %v", got, expected)
	}
}

func TestFlatMap(t *testing.T) {
	s := []int{1, 2, 3}
	expected := []int{1, 1, 2, 2, 3, 3}
	got := fnkit.FlatMap(s, func(i int) []int { return []int{i, i} })
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("FlatMap() = %v, want %v", got, expected)
	}
}

func TestForEach(t *testing.T) {
	s := []int{1, 2, 3}
	sum := 0
	fnkit.ForEach(s, func(i int) { sum += i })
	if sum != 6 {
		t.Errorf("ForEach() sum = %v, want 6", sum)
	}
}

func TestIncludes(t *testing.T) {
	s := []int{1, 2, 3}
	if !fnkit.Includes(s, 2) {
		t.Errorf("Includes() = false, want true")
	}
	if fnkit.Includes(s, 100) {
		t.Errorf("Includes() = true, want false")
	}
}

func TestIndexOf(t *testing.T) {
	s := []int{1, 2, 3, 2}
	if got := fnkit.IndexOf(s, 2); got != 1 {
		t.Errorf("IndexOf() = %v, want 1", got)
	}
	if got := fnkit.IndexOf(s, 100); got != -1 {
		t.Errorf("IndexOf() = %v, want -1", got)
	}
}

func TestJoin(t *testing.T) {
	s := []int{1, 2, 3}
	expected := "1-2-3"
	got := fnkit.Join(s, "-")
	if got != expected {
		t.Errorf("Join() = %v, want %v", got, expected)
	}
}

func TestLastIndexOf(t *testing.T) {
	s := []int{1, 2, 3, 2}
	if got := fnkit.LastIndexOf(s, 2); got != 3 {
		t.Errorf("LastIndexOf() = %v, want 3", got)
	}
	if got := fnkit.LastIndexOf(s, 100); got != -1 {
		t.Errorf("LastIndexOf() = %v, want -1", got)
	}
}

func TestKeys(t *testing.T) {
	s := []string{"a", "b", "c"}
	expected := []int{0, 1, 2}
	got := fnkit.Keys(s)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Keys() = %v, want %v", got, expected)
	}
}

func TestValues(t *testing.T) {
	s := []int{1, 2, 3}
	got := fnkit.Values(s)
	if !reflect.DeepEqual(got, s) {
		t.Errorf("Values() = %v, want %v", got, s)
	}
	got[0] = 100
	if s[0] == 100 {
		t.Errorf("Values() did not make a copy")
	}
}

func TestPop(t *testing.T) {
	s := []int{1, 2, 3}
	got, ok := fnkit.Pop(&s)
	if !ok || got != 3 {
		t.Errorf("Pop() = %v, %v, want 3, true", got, ok)
	}
	if len(s) != 2 {
		t.Errorf("Pop() did not remove element")
	}
}

func TestPush(t *testing.T) {
	s := []int{1, 2}
	fnkit.Push(&s, 3)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Push() = %v, want %v", s, expected)
	}
}

func TestReduceRight(t *testing.T) {
	s := []int{1, 2, 3}
	expected := 6
	got := fnkit.ReduceRight(s, 0, func(acc, v int) int { return acc + v })
	if got != expected {
		t.Errorf("ReduceRight() = %v, want %v", got, expected)
	}
}

func TestReverse(t *testing.T) {
	s := []int{1, 2, 3}
	fnkit.Reverse(s)
	expected := []int{3, 2, 1}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Reverse() = %v, want %v", s, expected)
	}
}

func TestShift(t *testing.T) {
	s := []int{1, 2, 3}
	got, ok := fnkit.Shift(&s)
	if !ok || got != 1 {
		t.Errorf("Shift() = %v, %v, want 1, true", got, ok)
	}
	if len(s) != 2 {
		t.Errorf("Shift() did not remove element")
	}
}

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 4}
	expected := []int{2, 3}
	got := fnkit.Slice(s, 1, 3)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Slice() = %v, want %v", got, expected)
	}
}

func TestSome(t *testing.T) {
	s := []int{1, 2, 3}
	if !fnkit.Some(s, func(i int) bool { return i == 2 }) {
		t.Errorf("Some() = false, want true")
	}
	if fnkit.Some(s, func(i int) bool { return i == 100 }) {
		t.Errorf("Some() = true, want false")
	}
}

func TestSplice(t *testing.T) {
	s := []int{1, 2, 3, 4}
	removed := fnkit.Splice(&s, 1, 2, []int{9, 9})
	expected := []int{2, 3}
	if !reflect.DeepEqual(removed, expected) {
		t.Errorf("Splice() removed = %v, want %v", removed, expected)
	}
	expectedS := []int{1, 9, 9, 4}
	if !reflect.DeepEqual(s, expectedS) {
		t.Errorf("Splice() s = %v, want %v", s, expectedS)
	}
}

func TestToLocaleString(t *testing.T) {
	s := []int{1, 2, 3}
	expected := "1,2,3"
	got := fnkit.ToLocaleString(s)
	if got != expected {
		t.Errorf("ToLocaleString() = %v, want %v", got, expected)
	}
}

func TestUnshift(t *testing.T) {
	s := []int{2, 3}
	fnkit.Unshift(&s, 1)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Unshift() = %v, want %v", s, expected)
	}
}

func TestWithout(t *testing.T) {
	s := []int{1, 2, 3, 2}
	expected := []int{1, 3}
	got := fnkit.Without(s, 2)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Without() = %v, want %v", got, expected)
	}
}
