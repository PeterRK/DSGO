package linkedlist

import (
	"unsafe"
)

type Node struct {
	Val  int
	Next *Node
}

func FakeHead(this **Node) *Node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).Next)
	return (*Node)(unsafe.Pointer(base - off))
}

func Merge(list1 *Node, list2 *Node) *Node {
	var head *Node = nil
	var tail = FakeHead(&head)
	for ; list1 != nil && list2 != nil; tail = tail.Next {
		if list1.Val > list2.Val {
			tail.Next, list2 = list2, list2.Next
		} else {
			tail.Next, list1 = list1, list1.Next
		}
	}
	for ; list1 != nil; tail = tail.Next {
		tail.Next, list1 = list1, list1.Next
	}
	for ; list2 != nil; tail = tail.Next {
		tail.Next, list2 = list2, list2.Next
	}
	return head
}

func Reverse(list *Node) *Node {
	var head *Node = nil
	for list != nil {
		var node = list
		list = list.Next
		node.Next = head
		head = node
	}
	return head
}
