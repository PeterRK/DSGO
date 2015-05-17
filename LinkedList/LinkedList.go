package linkedlist

import (
	"unsafe"
)

type Node struct {
	val  int
	next *Node
}

func FakeHead(this **Node) *Node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).next)
	return (*Node)(unsafe.Pointer(base - off))
}

func Merge(list1 *Node, list2 *Node) *Node {
	var head *Node = nil
	var tail = FakeHead(&head)
	for ; list1 != nil && list2 != nil; tail = tail.next {
		if list1.val > list2.val {
			tail.next, list2 = list2, list2.next
		} else {
			tail.next, list1 = list1, list1.next
		}
	}
	for ; list1 != nil; tail = tail.next {
		tail.next, list1 = list1, list1.next
	}
	for ; list2 != nil; tail = tail.next {
		tail.next, list2 = list2, list2.next
	}
	return head
}

func Reverse(list *Node) *Node {
	var head *Node = nil
	for list != nil {
		var node = list
		list = list.next
		node.next = head
		head = node
	}
	return head
}
