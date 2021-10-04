package array

import (
	"DSGO/utils"
)

type Queue[T any] interface {
	utils.Queue[T]
	IsFull() bool
}

func NewQueue[T any](size int) Queue[T] {
	q := new(queue[T])
	q.init(size)
	return q
}

type queue[T any] struct {
	r, w  int
	space []T
}

func (q *queue[T]) init(size int) {
	if size < 7 {
		size = 7
	}
	q.space = make([]T, size+1)
	q.r, q.w = 0, 0
}

func (q *queue[T]) Clear() {
	q.r, q.w = 0, 0
}

func (q *queue[T]) IsEmpty() bool {
	return q.r == q.w
}

func (q *queue[T]) IsFull() bool {
	return (q.w+1)%len(q.space) == q.r
}

func (q *queue[T]) Size() int {
	size := q.r - q.w
	if size < 0 {
		size += len(q.space)
	}
	return size
}

func (q *queue[T]) Push(unit T) {
	w := (q.w + 1) % len(q.space)
	if w == q.r {
		panic("full queue")
	}
	q.space[q.w] = unit
	q.w = w
}

func (q *queue[T]) Pop() T {
	if q.IsEmpty() {
		panic("empty queue")
	}
	unit := q.space[q.r]
	q.r = (q.r + 1) % len(q.space)
	return unit
}

func (q *queue[T]) Front() T {
	if q.IsEmpty() {
		panic("empty queue")
	}
	return q.space[q.r]
}

func (q *queue[T]) Back() T {
	if q.IsEmpty() {
		panic("empty queue")
	}
	if q.w == 0 {
		return q.space[len(q.space)-1]
	}
	return q.space[q.w-1]
}
