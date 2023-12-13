package ds

type Queue[T any] struct {
	list []T
}

func (iq *Queue[T]) Len() int {
	n := len(iq.list)
	return n
}

func (iq *Queue[T]) Push(input T) {
	iq.list = append(iq.list, input)
}

func (iq *Queue[T]) PopAll() []T {
	oldInputs := iq.list
	iq.list = []T{}
	return oldInputs
}

func (iq *Queue[T]) Pop() T {
	item := iq.list[0]
	iq.list = iq.list[1:]
	return item
}
