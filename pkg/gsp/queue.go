package gsp

import (
	"sync"
)

type Queue[T any] struct {
	list   []T
	locker sync.Mutex
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		list: []T{},
	}
}

func (iq *Queue[T]) Len() int {
	iq.locker.Lock()
	n := len(iq.list)
	iq.locker.Unlock()
	return n
}

func (iq *Queue[T]) Push(input T) {
	iq.locker.Lock()
	iq.list = append(iq.list, input)
	iq.locker.Unlock()
}

func (iq *Queue[T]) PopAll() []T {
	iq.locker.Lock()
	oldInputs := iq.list
	iq.list = []T{}
	iq.locker.Unlock()
	return oldInputs
}
