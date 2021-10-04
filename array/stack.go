package array

import (
	"DSGO/utils"
)

type Stack[T any] struct {
	vec []T
}

func NewStack[T any](size int) utils.Stack[T] {
	s := new(Stack[T])
	s.vec = make([]T, 0, size)
	return s
}

func (s *Stack[T]) Size() int {
	return len(s.vec)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.vec) == 0
}

func (s *Stack[T]) Clear() {
	s.vec = s.vec[:0]
}

func (s *Stack[T]) Push(elem T) {
	s.vec = append(s.vec, elem)
}

func (s *Stack[T]) Pop() T {
	if len(s.vec) == 0 {
		panic("empty stack")
	}
	last := len(s.vec) - 1
	elem := s.vec[last]
	s.vec = s.vec[:last]
	return elem
}

func (s *Stack[T]) Top() T {
	if len(s.vec) == 0 {
		panic("empty stack")
	}
	return s.vec[len(s.vec)-1]
}
