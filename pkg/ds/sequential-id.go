package ds

import (
	"math"
	"sync"
)

type SequentialID struct {
	current int16
	lock    sync.Mutex
}

func (s *SequentialID) Next() int16 {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.current == math.MaxInt16 {
		s.current = 0
	}
	s.current += 1
	return s.current
}
