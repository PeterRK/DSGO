package deque

type Deque interface {
	IsEmpty() bool
	Size() int
	PushFront(key int)
	PushBack(key int)
	PopFront() (key int, err bool)
	PopBack() (key int, err bool)
	Front() (key int, err bool)
	Back() (key int, err bool)
}

func NewDeque() Deque {
	var container = new(deque)
	container.initialize(DEQUE)
	return container
}

type stack struct {
	deque
}
type Stack interface {
	IsEmpty() bool
	Size() int
	Push(key int)
	Pop() (key int, err bool)
	Top() (key int, err bool)
}

func NewStack() Stack {
	var container = new(stack)
	container.initialize(STACK)
	return container
}
func (this *stack) Push(key int) {
	this.PushFront(key)
}
func (this *stack) Pop() (key int, err bool) {
	return this.PopFront()
}
func (this *stack) Top() (key int, err bool) {
	return this.Front()
}

type queue struct {
	deque
}
type Queue interface {
	IsEmpty() bool
	Size() int
	Push(key int)
	Pop() (key int, err bool)
	Front() (key int, err bool)
	Back() (key int, err bool)
}

func NewQueue() Queue {
	var container = new(queue)
	container.initialize(DEQUE)
	return container
}
func (this *queue) Push(key int) {
	this.PushBack(key)
}
func (this *queue) Pop() (key int, err bool) {
	return this.PopFront()
}
