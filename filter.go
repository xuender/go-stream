package stream

import "sync"

type FilterAction[T any] func(T) bool

func (p *BaseStream[T]) Filter(action FilterAction[T]) *BaseStream[T] {
	p.C = Filter(p.C, action)

	return p
}

func Filter[T any](input <-chan T, action FilterAction[T]) chan T {
	output := make(chan T)

	go filter(input, output, action)

	return output
}

func filter[T any](input <-chan T, output chan<- T, action FilterAction[T]) {
	for i := range input {
		if action(i) {
			output <- i
		}
	}

	close(output)
}

func (p *ParallelStream[T]) Filter(action FilterAction[T]) *ParallelStream[T] {
	p.C = FilterParallel(p.C, p.Size, action)

	return p
}

func FilterParallel[T any](input <-chan T, size int, action FilterAction[T]) chan T {
	output := make(chan T)
	group := &sync.WaitGroup{}

	group.Add(size)

	for i := 0; i < size; i++ {
		go filterParallel(input, output, group, action)
	}

	go filterClose(output, group)

	return output
}

func filterClose[T any](output chan<- T, group *sync.WaitGroup) {
	group.Wait()
	close(output)
}

func filterParallel[T any](input <-chan T, output chan<- T, group *sync.WaitGroup, action FilterAction[T]) {
	for i := range input {
		if action(i) {
			output <- i
		}
	}

	group.Done()
}
