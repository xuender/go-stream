package stream

type ParallelStream[T any] struct {
	BaseStream[T]
	Size int
}

func NewParallel[T any](input chan T, size int) *ParallelStream[T] {
	return &ParallelStream[T]{BaseStream[T]{input}, size}
}

func (p *BaseStream[T]) Parallel(size int) *ParallelStream[T] {
	return NewParallel[T](p.C, size)
}
