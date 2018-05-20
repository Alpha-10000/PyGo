package main

type Stack struct {
	data []interface{}
	size int
}

func (s *Stack) Push(elt interface{}) {
	s.data = append(s.data, elt)
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size < 1 {		
		return nil
	}
	var val = s.data[0]
	s.data = append(s.data[:0], s.data[1:]...)
	s.size--
	return val
}

