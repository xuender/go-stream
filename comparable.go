package stream

type ComparableStream[C comparable] struct {
	BaseStream[C]
}

func NewComparable[C comparable](input chan C) *ComparableStream[C] {
	return &ComparableStream[C]{*NewBase(input)}
}

func (p *ComparableStream[C]) Distinct() *ComparableStream[C] {
	p.C = Distinct(p.C)

	return p
}

func Distinct[C comparable](input <-chan C) chan C {
	output := make(chan C)

	go distinct(input, output)

	return output
}

func distinct[C comparable](input <-chan C, output chan<- C) {
	old := map[C]struct{}{}
	none := struct{}{}

	for elem := range input {
		if _, has := old[elem]; has {
			continue
		}

		old[elem] = none
		output <- elem
	}

	close(output)
}
