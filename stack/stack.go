package stack

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) { // need to use pointer to modify the Stack, else it will be a copy
	s.data = append(s.data, v) // append to end
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) <= 0 {
		var zeroValue T
		return zeroValue, false
	}

	lastIndex := len(s.data) - 1
	value := s.data[lastIndex]
	s.data = s.data[0:lastIndex] // remove the last element

	return value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) <= 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}
