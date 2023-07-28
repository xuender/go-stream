package stream

type Stream[T any] struct {
	elems []T
}

func New[T any](elems ...T) *Stream[T] {
	return &Stream[T]{elems: elems}
}
