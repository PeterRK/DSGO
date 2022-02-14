package heap

import (
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Top() T
	Push(T)
	Pop() T
}

type NodeHeap[Node any] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Top() *Node
	Push(*Node)
	Pop() *Node
	FloatUp(*Node)
}
