package fn

// Pipeline enables chainable, functional operations on slices.
type Pipeline[T any] struct {
	data []T
}

func FromSlice[T any](s []T) Pipeline[T] {
	return Pipeline[T]{data: s}
}

func (p Pipeline[T]) Map(f func(T) T) Pipeline[T] {
	out := make([]T, len(p.data))
	for i, v := range p.data {
		out[i] = f(v)
	}
	return Pipeline[T]{data: out}
}

func (p Pipeline[T]) Filter(f func(T) bool) Pipeline[T] {
	out := make([]T, 0, len(p.data))
	for _, v := range p.data {
		if f(v) {
			out = append(out, v)
		}
	}
	return Pipeline[T]{data: out}
}

func (p Pipeline[T]) Reduce(init T, f func(T, T) T) T {
	acc := init
	for _, v := range p.data {
		acc = f(acc, v)
	}
	return acc
}

func (p Pipeline[T]) Slice() []T {
	return p.data
}
