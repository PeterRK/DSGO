package list

import (
	"unsafe"
)

type Node struct {
	Val  int
	Next *Node
}

func FakeHead(spt **Node) *Node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).Next)
	return (*Node)(unsafe.Pointer(base - off))
}

func Merge(lst1 *Node, lst2 *Node) (list *Node) {
	var last = FakeHead(&list)
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

func Reverse(list *Node) *Node {
	var head = (*Node)(nil)
	for list != nil {
		var node = list
		list = list.Next
		node.Next, head = head, node
	}
	return head
}
