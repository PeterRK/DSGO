package deque

type Deque interface {
	Clear()
	Size() int
	IsEmpty() bool
	PushFront(key int)
	PushBack(key int)
	PopFront() (int, error)
	PopBack() (int, error)
	Front() (int, error)
	Back() (int, error)
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
	Pop() (int, error)
	Top() (int, error)
}

func NewStack() Stack {
	var container = new(stack)
	container.initialize()
	return container
}
func (s *stack) Push(key int) {
	s.PushFront(key)
}
func (s *stack) Pop() (int, error) {
	return s.PopFront()
}
func (s *stack) Top() (int, error) {
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
	Pop() (int, error)
	Front() (int, error)
	Back() (int, error)
}

func NewQueue() Queue {
	var container = new(queue)
	container.initialize()
	return container
}
func (q *queue) Push(key int) {
	q.PushBack(key)
}
func (q *queue) Pop() (int, error) {
	return q.PopFront()
}
