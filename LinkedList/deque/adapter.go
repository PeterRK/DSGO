package deque

type Deque interface {
	Clear()
	Size() int
	IsEmpty() bool
	PushFront(key int)
	PushBack(key int)
	PopFront() (key int, fail bool)
	PopBack() (key int, fail bool)
	Front() (key int, fail bool)
	Back() (key int, fail bool)
}

func NewDeque() Deque {
	var container = new(deque)
	container.initialize()
	return container
}

type stack struct {
	deque
}
type Stack interface {
	Clear()
	Size() int
	IsEmpty() bool
	Push(key int)
	Pop() (key int, fail bool)
	Top() (key int, fail bool)
}

func NewStack() Stack {
	var container = new(stack)
	container.initialize()
	return container
}
func (s *stack) Push(key int) {
	s.PushFront(key)
}
func (s *stack) Pop() (key int, fail bool) {
	return s.PopFront()
}
func (s *stack) Top() (key int, fail bool) {
	return s.Front()
}

type queue struct {
	deque
}
type Queue interface {
	Clear()
	Size() int
	IsEmpty() bool
	Push(key int)
	Pop() (key int, fail bool)
	Front() (key int, fail bool)
	Back() (key int, fail bool)
}

func NewQueue() Queue {
	var container = new(queue)
	container.initialize()
	return container
}
func (q *queue) Push(key int) {
	q.PushBack(key)
}
func (q *queue) Pop() (key int, fail bool) {
	return q.PopFront()
}
