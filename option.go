package fnkit

// Option is a generic container that holds either a value (Some) or nothing (None).
// Inspired by Rust's Option<T> type, it allows for safe handling of optional values.
type Option[T any] struct {
	value   T
	present bool
}

// Some returns an Option containing a value.
func Some[T any](val T) Option[T] {
	return Option[T]{value: val, present: true}
}

// None returns an Option with no value.
func None[T any]() Option[T] {
	var zero T
	return Option[T]{value: zero, present: false}
}

// IsSome returns true if the Option contains a value.
func (o Option[T]) IsSome() bool {
	return o.present
}

// IsNone returns true if the Option contains no value.
func (o Option[T]) IsNone() bool {
	return !o.present
}

// Unwrap returns the value if present, otherwise panics.
func (o Option[T]) Unwrap() T {
	if !o.present {
		panic("called Unwrap on None Option")
	}
	return o.value
}

// UnwrapOr returns the value if present, or the provided fallback otherwise.
func (o Option[T]) UnwrapOr(fallback T) T {
	if o.present {
		return o.value
	}
	return fallback
}
