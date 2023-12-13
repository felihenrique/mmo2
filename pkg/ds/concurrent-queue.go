package ds

import (
	"sync"
)

type ConcurrentQueue[T any] struct {
	list   []T
	locker sync.Mutex
}

func (iq *ConcurrentQueue[T]) Len() int {
	iq.locker.Lock()
	n := len(iq.list)
	iq.locker.Unlock()
	return n
}

func (iq *ConcurrentQueue[T]) Push(input T) {
	iq.locker.Lock()
	iq.list = append(iq.list, input)
	iq.locker.Unlock()
}

func (iq *ConcurrentQueue[T]) PopAll() []T {
	iq.locker.Lock()
	oldInputs := iq.list
	iq.list = []T{}
	iq.locker.Unlock()
	return oldInputs
}
