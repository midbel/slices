package slices

type Stack[T any] struct {
	list []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(t T) {
	s.list = append(s.list, t)
}

func (s *Stack[T]) Pop() T {
	var (
		t = Lst(s.list)
		n = len(s.list)
	)
	if n > 0 {
		s.list = s.list[:n-1]
	}
	return t
}

func (s *Stack[T]) Len() int {
	return len(s.list)
}
