package stream

import "sync"

func (p *BaseStream[T]) ForEach(consum func(elem T)) {
	for elem := range p.C {
		consum(elem)
	}
}

func (p *ParallelStream[T]) ForEach(consum func(elem T)) {
	group := sync.WaitGroup{}
	group.Add(p.Size)

	for i := 0; i < p.Size; i++ {
		go func() {
			p.BaseStream.ForEach(consum)
			group.Done()
		}()
	}

	group.Wait()
}
