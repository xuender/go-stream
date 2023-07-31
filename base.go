package stream

type BaseStream[T any] struct {
	C chan T
}

func NewBase[T any](input chan T) *BaseStream[T] {
	return &BaseStream[T]{input}
}

func (p *BaseStream[T]) Count() int {
	count := 0

	for range p.C {
		count++
	}

	return count
}
