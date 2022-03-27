package queue

type Queue struct {
	arr []interface{}
}

func (s *Queue) Push(value interface{}) {
	s.arr = append(s.arr, value)
}

func (s *Queue) Pop() interface{} {
	value := s.arr[0]
	s.arr = s.arr[1:]
	return value
}

func (s *Queue) Empty() bool {
	return len(s.arr) == 0
}
