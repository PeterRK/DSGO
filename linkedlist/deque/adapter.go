package deque

import (
	"DSGO/utils"
)

type stack[T any] struct {
	core deque[T]
}

func NewStack[T any]() utils.Stack[T] {
	obj := new(stack[T])
	obj.core.init()
	return obj
}

func (s *stack[T]) Size() int {
	return s.core.Size()
}

func (s *stack[T]) IsEmpty() bool {
	return s.core.IsEmpty()
}

func (s *stack[T]) Clear() {
	s.core.Clear()
}

func (s *stack[T]) Push(unit T) {
	s.core.PushFront(unit)
}

func (s *stack[T]) Pop() T {
	return s.core.PopFront()
}

func (s *stack[T]) Top() T {
	return s.core.Front()
}

type queue[T any] struct {
	core deque[T]
}

func NewQueue[T any]() utils.Queue[T] {
	obj := new(queue[T])
	obj.core.init()
	return obj
}

func (q *queue[T]) Size() int {
	return q.core.Size()
}

func (q *queue[T]) IsEmpty() bool {
	return q.core.IsEmpty()
}

func (q *queue[T]) Clear() {
	q.core.Clear()
}

func (q *queue[T]) Push(unit T) {
	q.core.PushBack(unit)
}

func (q *queue[T]) Pop() T {
	return q.core.PopFront()
}

func (q *queue[T]) Front() T {
	return q.core.Front()
}

func (q *queue[T]) Back() T {
	return q.core.Back()
}
