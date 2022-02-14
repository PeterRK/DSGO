package linkedlist

import (
	"golang.org/x/exp/constraints"
	"unsafe"
)

type Node[T any] struct {
	Next *Node[T]
	Val  T
}

func FakeHead[T constraints.Ordered](spt **Node[T]) *Node[T] {
	//base := uintptr(unsafe.Pointer(spt))
	//off := unsafe.Offsetof((*spt).Next)
	//return (*Node[T])(unsafe.Pointer(base - off))
	return (*Node[T])(unsafe.Pointer(spt))
}

func Merge[T constraints.Ordered](lst1, lst2 *Node[T]) (list *Node[T]) {
	last := FakeHead(&list)
	for {
		last.Next = lst1
		if lst2 == nil {
			return list
		}
		for lst1 != nil && lst1.Val <= lst2.Val {
			last, lst1 = lst1, lst1.Next
		}

		last.Next = lst2
		if lst1 == nil {
			return list
		}
		for lst2 != nil && lst1.Val > lst2.Val {
			last, lst2 = lst2, lst2.Next
		}
	}
}

func IsSorted[T constraints.Ordered](head *Node[T]) bool {
	if head == nil {
		return true
	}
	for ; head.Next != nil; head = head.Next {
		if head.Val > head.Next.Val {
			return false
		}
	}
	return true
}

func Reverse[T any](list *Node[T]) *Node[T] {
	head := (*Node[T])(nil)
	for list != nil {
		node := list
		list = list.Next
		node.Next, head = head, node
	}
	return head
}
