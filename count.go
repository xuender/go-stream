package stream

func (p *Stream[T]) Count() int {
	return len(p.elems)
}
