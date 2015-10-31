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
	var con = new(deque)
	con.initialize()
	return con
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
	var con = new(stack)
	con.initialize()
	return con
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
	var con = new(queue)
	con.initialize()
	return con
}
func (q *queue) Push(key int) {
	q.PushBack(key)
}
func (q *queue) Pop() (int, error) {
	return q.PopFront()
}
