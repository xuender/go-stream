package stream

func (p *BaseStream[T]) AnyMatch(action FilterAction[T]) bool {
	for i := range p.C {
		if action(i) {
			go Consume(p.C)

			return true
		}
	}

	return false
}

func (p *BaseStream[T]) AllMatch(action FilterAction[T]) bool {
	for i := range p.C {
		if !action(i) {
			go Consume(p.C)

			return false
		}
	}

	return true
}

func (p *BaseStream[T]) NoneMatch(action FilterAction[T]) bool {
	return !p.AllMatch(action)
}
