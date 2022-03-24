package stack

type Stack struct {
	arr []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.arr = append(s.arr, value)
}

func (s *Stack) Pop() interface{} {
	l := len(s.arr)
	value := s.arr[l - 1]
	s.arr = s.arr[:l - 1]
	return value
}

func (s *Stack) Top() interface{} {
	l := len(s.arr)
	value := s.arr[l - 1]
	return value
}

func (s *Stack) Empty() bool {
	return len(s.arr) == 0
}