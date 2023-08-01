package stream

import "golang.org/x/exp/constraints"

type OrderedStream[O constraints.Ordered] struct {
	ComparableStream[O]
}

func NewOrdered[O constraints.Ordered](input chan O) *OrderedStream[O] {
	return &OrderedStream[O]{*NewComparable(input)}
}

func (p *OrderedStream[O]) Max() O {
	var max O

	one := true

	for num := range p.C {
		if num > max || one {
			max = num
			one = false
		}
	}

	return max
}

func (p *OrderedStream[O]) Min() O {
	var min O

	one := true

	for num := range p.C {
		if num < min || one {
			min = num
			one = false
		}
	}

	return min
}
