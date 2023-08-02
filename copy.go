package stream

import "time"

const _hundred = 100

func Copy[T any](source *BaseStream[T]) chan T {
	output1 := make(chan T, _hundred)
	output2 := make(chan T, _hundred)

	go Distribute(source.C, output1, output2)

	time.Sleep(time.Microsecond)

	source.C = output1

	return output2
}
