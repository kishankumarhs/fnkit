package fnkit

// Result is a generic container that holds either a successful value of type T, or an error.
// It allows for functional chaining without constant 'if err != nil' checks.
// The T must be comparable (like basic types and structs without non-comparable fields).
type Result[T any] struct {
	Value T
	Err   error
}

// Ok is a constructor for a successful Result containing the given value.
func Ok[T any](val T) Result[T] {
	return Result[T]{Value: val, Err: nil}
}

// Err is a constructor for a failed Result containing the given error.
func Err[T any](err error) Result[T] {
	// Value will be the zero value of type T
	var zeroValue T
	return Result[T]{Value: zeroValue, Err: err}
}

// IsOk returns true if the Result contains a value and no error.
func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

// ValueOr returns the contained value if successful, or the provided fallback value if an error exists.
func (r Result[T]) ValueOr(fallback T) T {
	if r.IsOk() {
		return r.Value
	}
	return fallback
}
