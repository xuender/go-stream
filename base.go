package stream

type BaseStream[T any] struct {
	C chan T
}

func NewBase[T any](input chan T) *BaseStream[T] {
	return &BaseStream[T]{input}
}

func (p *BaseStream[T]) Count() int {
	return Count(p.C)
}

func (p *BaseStream[T]) FindFirst() T {
	ret := <-p.C

	go Count(p.C)

	return ret
}
