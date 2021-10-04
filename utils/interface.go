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