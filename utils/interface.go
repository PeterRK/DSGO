package utils

type Stack[T any] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Push(T)
	Pop() T
	Top() T
}

type Queue[T any] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Push(T)
	Pop() T
	Front() T
	Back() T
}

type Cache[K comparable, V any] interface {
	Clear()
	Size() int
	Capacity() int
	Put(K, V)
	Get(K) (V, bool)
	Discard(K)
}

type Random interface {
	Next() uint32
}

type StrSet interface {
	Size() int
	IsEmpty() bool
	Clear()
	Search(string) bool
	Insert(string) bool
	Remove(string) bool
}