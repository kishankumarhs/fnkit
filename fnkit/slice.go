// Package fnkit provides functional utilities for Go slices.
package fnkit

import "fmt"

// Map applies a transformation function 'f' to each element of the input slice 's'
// and returns a new slice containing the results.
//
// The function uses Go generics:
// - K: Represents the type of elements in the input slice (e.g., int, string).
// - T: Represents the type of elements in the output slice after transformation.
// - S: The input slice, of type []K.
// - f: The transformation function, which takes an element of type K and returns an element of type T.
func Map[K any, T any](s []K, f func(K) T) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Reduce applies a binary function 'f' to each element of the input slice 's'
// and an accumulator 'initial', reducing the slice to a single value.
func Reduce[K any, T any](s []K, initial T, f func(T, K) T) T {
	accumulator := initial
	for _, v := range s {
		accumulator = f(accumulator, v)
	}
	return accumulator
}

// At returns the element at the specified index from the slice.
// If the index is out of bounds, it returns the zero value of the element type.
func At[K any](s []K, index int) K {
	if index >= len(s) {
		var zero K
		return zero
	}
	if index < 0 && index+len(s) < 0 {
		var zero K
		return zero
	}
	if index < 0 {
		index = len(s) + index
	}
	return s[index]
}

// Concat concatenates two slices of the same type and returns a new slice.
func Concat[K any](s1, s2 []K) []K {
	result := make([]K, len(s1)+len(s2))
	copy(result, s1)
	copy(result[len(s1):], s2)
	return result
}

// CopyWith makes a shallow copy of the input slice 's' and returns it.
func CopyWith[K any](s []K) []K {
	result := make([]K, len(s))
	copy(result, s)
	return result
}

// Entries returns a slice of key-value pairs from the input slice 's'.
// Each key-value pair is represented as a struct with 'Index' and 'Value' fields.
type Entry[K any] struct {
	Index int
	Value K
}

func Entries[K any](s []K) []Entry[K] {
	result := make([]Entry[K], len(s))
	for i, v := range s {
		result[i] = Entry[K]{Index: i, Value: v}
	}
	return result
}

// Every checks if all elements in the slice 's' satisfy the predicate function 'f'.
// It returns true if all elements satisfy the condition, otherwise false.
func Every[K any](s []K, f func(K) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

// Fill fills the slice 's' with the specified value 'value'.
func Fill[K any](s []K, value K) {
	for i := range s {
		s[i] = value
	}
}

// Find returns the first element in the slice 's' that satisfies the predicate function 'f'.
// If no element satisfies the condition, it returns the zero value of the element type
// and a boolean value 'false'. If an element is found, it returns the element and 'true'.
func Find[K any](s []K, f func(K) bool) (K, bool) {
	for _, v := range s {
		if f(v) {
			return v, true
		}
	}
	var zero K
	return zero, false
}

// FindIndex returns the index of the first element in the slice 's' that satisfies
// the predicate function 'f'. If no element satisfies the condition, it returns -1.
func FindIndex[K any](s []K, f func(K) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}

// FindLast returns the last element in the slice 's' that satisfies the predicate function 'f'.
// If no element satisfies the condition, it returns the zero value of the element type
// and a boolean value 'false'. If an element is found, it returns the element and 'true'.
func FindLast[K any](s []K, f func(K) bool) (K, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return s[i], true
		}
	}
	var zero K
	return zero, false
}

// FindLastIndex returns the index of the last element in the slice 's' that satisfies
// the predicate function 'f'. If no element satisfies the condition, it returns -1.
func FindLastIndex[K any](s []K, f func(K) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

// Flat flattens a slice of slices 's' into a single slice.
func Flat[K any](s [][]K) []K {
	totalLength := 0
	for _, subSlice := range s {
		totalLength += len(subSlice)
	}
	result := make([]K, 0, totalLength)
	for _, subSlice := range s {
		result = append(result, subSlice...)
	}
	return result
}

// FlatMap applies a transformation function 'f' to each element of the input slice 's'
// and flattens the resulting slices into a single slice.
func FlatMap[K any, T any](s []K, f func(K) []T) []T {
	var result []T
	for _, v := range s {
		result = append(result, f(v)...)
	}
	return result
}

// ForEach applies a function 'f' to each element of the input slice 's'.
func ForEach[K any](s []K, f func(K)) {
	for _, v := range s {
		f(v)
	}
}

// Includes checks if the slice 's' contains the specified value 'value'.
// It returns true if the value is found, otherwise false.
func Includes[K comparable](s []K, value K) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of the specified value 'value'
// in the slice 's'. If the value is not found, it returns -1.
func IndexOf[K comparable](s []K, value K) int {
	for i, v := range s {
		if v == value {
			return i
		}
	}
	return -1
}

// Join concatenates the elements of the slice 's' into a single string,
// with the specified separator 'sep' between elements.
func Join[K any](s []K, sep string) string {
	if len(s) == 0 {
		return ""
	}
	var result string
	for i, v := range s {
		if i > 0 {
			result += sep
		}
		result += fmt.Sprint(v)
	}
	return result
}

// LastIndexOf returns the index of the last occurrence of the specified value 'value'
// in the slice 's'. If the value is not found, it returns -1.
func LastIndexOf[K comparable](s []K, value K) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == value {
			return i
		}
	}
	return -1
}

// Keys returns a slice containing the indices of the input slice 's'.
func Keys[K any](s []K) []int {
	result := make([]int, len(s))
	for i := range s {
		result[i] = i
	}
	return result
}

// Values returns a copy of the input slice 's'.
func Values[K any](s []K) []K {
	result := make([]K, len(s))
	copy(result, s)
	return result
}

// Pop removes and returns the last element of the slice 's'.
// If the slice is empty, it returns the zero value of the element type
// and a boolean value 'false'. If an element is removed, it returns the element and 'true'.
func Pop[K any](s *[]K) (K, bool) {
	if len(*s) == 0 {
		var zero K
		return zero, false
	}
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last, true
}

// Push appends the specified value 'value' to the end of the slice 's'.
func Push[K any](s *[]K, value K) {
	*s = append(*s, value)
}

// ReduceRight applies a binary function 'f' to each element of the input slice 's'
// and an accumulator 'initial', reducing the slice to a single value, starting from the end.
func ReduceRight[K any, T any](s []K, initial T, f func(T, K) T) T {
	accumulator := initial
	for i := len(s) - 1; i >= 0; i-- {
		accumulator = f(accumulator, s[i])
	}
	return accumulator
}

// Reverse reverses the elements of the slice 's' in place.
func Reverse[K any](s []K) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Shift removes and returns the first element of the slice 's'.
// If the slice is empty, it returns the zero value of the element type
// and a boolean value 'false'. If an element is removed, it returns the element and 'true'.
func Shift[K any](s *[]K) (K, bool) {
	if len(*s) == 0 {
		var zero K
		return zero, false
	}
	first := (*s)[0]
	*s = (*s)[1:]
	return first, true
}

// Slice returns a new slice that is a sub-slice of the input slice 's',
// starting from index 'start' up to, but not including, index 'end'.
// If 'end' is greater than the length of the slice, it is set to the length of the slice.
// If 'start' is less than 0, it is set to 0. If 'start' is greater than or equal to 'end',
// an empty slice is returned.
func Slice[K any](s []K, start, end int) []K {
	if start < 0 {
		start = 0
	}
	if end > len(s) {
		end = len(s)
	}
	if start >= end {
		return []K{}
	}
	return s[start:end]
}

// Some returns true if at least one element in the slice 's' satisfies the predicate function 'f'.
// It returns false if no elements satisfy the condition.
func Some[K any](s []K, f func(K) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// Splice removes 'deleteCount' elements from the slice 's' starting at index 'start',
// and inserts the elements from the 'items' slice at that position.
// It returns a new slice containing the removed elements.
//
// If 'start' is negative, it is treated as 'len(s) + start'.
// If 'start' is greater than the length of the slice, it is set to the length of the slice.
// If 'deleteCount' is greater than the number of elements from 'start' to the end of the slice,
// it is set to that number.
func Splice[K any](s *[]K, start, deleteCount int, items []K) []K {
	if start < 0 {
		start = len(*s) + start
	}
	if start > len(*s) {
		start = len(*s)
	}
	if deleteCount < 0 {
		deleteCount = 0
	}
	if deleteCount > len(*s)-start {
		deleteCount = len(*s) - start
	}
	removed := make([]K, deleteCount)
	copy(removed, (*s)[start:start+deleteCount])
	result := make([]K, 0, len(*s)-deleteCount+len(items))
	result = append(result, (*s)[:start]...)
	result = append(result, items...)
	result = append(result, (*s)[start+deleteCount:]...)
	*s = result
	return removed
}

// ToLocaleString converts each element of the slice 's' to its string representation
// and joins them into a single string, separated by commas.
func ToLocaleString[K any](s []K) string {
	return Join(s, ",")
}

// Unshift adds the specified value 'value' to the beginning of the slice 's'.
func Unshift[K any](s *[]K, value K) {
	*s = append([]K{value}, *s...)
}

// Without returns a new slice that excludes all occurrences of the specified value 'value'
// from the input slice 's'.
func Without[K comparable](s []K, value K) []K {
	var result []K
	for _, v := range s {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

// Filter returns a new slice containing only the elements of the input slice 's'
// for which the predicate function 'f' returns true.
//
// The function uses Go generics:
// - K: Represents the type of elements in the input slice (e.g., int, string).
// - S: The input slice, of type []K.
// - f: The predicate function, which takes an element of type K and returns a bool.
// Filter filters the slice s in-place, keeping only the elements for which
// the function f returns true. It modifies the original slice directly.
func Filter[K any](s *[]K, f func(K) bool) {
	w := 0
	for _, v := range *s {
		if f(v) {
			(*s)[w] = v
			w++
		}
	}
	*s = (*s)[:w]
}

// ToFilter returns a new slice containing only the elements of the input slice 's'
// for which the predicate function 'f' returns true.
//
// The function uses Go generics:
// - K: Represents the type of elements in the input slice (e.g., int, string).
// - S: The input slice, of type []K.
// - f: The predicate function, which takes an element of type K and returns a bool.
func ToFilter[K any](s []K, f func(K) bool) []K {
	var result []K
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
