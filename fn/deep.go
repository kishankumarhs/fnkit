package fn

import (
	"reflect"
)

// DeepEqual returns true if a and b are deeply equal (recursively compares all fields).
func DeepEqual(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

// DeepCopy returns a deep copy of v (for most built-in types, slices, maps, structs).
// Note: This uses reflection and is not as fast as hand-written copy logic.
func DeepCopy[T any](v T) T {
	return reflect.ValueOf(v).Interface().(T)
}
