package ds

import "math"

type SequentialID struct {
	current int16
}

func (s *SequentialID) Next() int16 {
	if s.current == math.MaxInt16 {
		s.current = 0
	}
	s.current += 1
	return s.current
}
