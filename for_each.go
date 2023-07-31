package stream

import "sync"

type Action[T any] func(T)

func (p *BaseStream[T]) ForEach(action Action[T]) {
	for elem := range p.C {
		action(elem)
	}
}

func (p *ParallelStream[T]) ForEach(action Action[T]) {
	group := sync.WaitGroup{}
	group.Add(p.Size)

	for i := 0; i < p.Size; i++ {
		go func() {
			p.BaseStream.ForEach(action)
			group.Done()
		}()
	}

	group.Wait()
}
