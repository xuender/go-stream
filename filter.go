package stream

import "sync"

func (p *BaseStream[T]) Filter(predicate func(T) bool) *BaseStream[T] {
	p.C = Filter(p.C, predicate)

	return p
}

func (p *ParallelStream[T]) Filter(predicate func(T) bool) *ParallelStream[T] {
	p.C = FilterParallel(p.C, p.Size, predicate)

	return p
}

func Filter[T any](input <-chan T, predicate func(T) bool) chan T {
	output := make(chan T)

	go func() {
		for i := range input {
			if predicate(i) {
				output <- i
			}
		}

		close(output)
	}()

	return output
}

func FilterParallel[T any](input <-chan T, size int, predicate func(T) bool) chan T {
	output := make(chan T)
	group := sync.WaitGroup{}

	group.Add(size)

	for i := 0; i < size; i++ {
		go func() {
			for i := range input {
				if predicate(i) {
					output <- i
				}
			}

			group.Done()
		}()
	}

	go func() {
		group.Wait()
		close(output)
	}()

	return output
}
