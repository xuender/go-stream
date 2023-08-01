package stream

func (p *ParallelStream[T]) Sequential() *BaseStream[T] {
	return &p.BaseStream
}
