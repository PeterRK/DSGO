package sort

import (
	"LinkedList/list"
)

func IntroSort(head *list.Node) *list.Node {
	if head != nil {
		var size = 0
		for node := head; node != nil; node = node.Next {
			size++
		}
		var life uint
		for life = 2; size != 0; life++ {
			size /= 2
		}
		head, _ = introSort(head, life)
	}
	return head
}
func introSort(head *list.Node, life uint) (first *list.Node, last *list.Node) {
	if head.Next == nil {
		return head, head
	}
	var node = head.Next
	if node.Next == nil {
		if head.Val > node.Val {
			node.Next, head.Next = head, nil
			return node, head
		}
		return head, node
	}
	if life == 0 {
		return mergeSort(head)
	}

	var left, center, right = part(head, node, node.Next)

	first, node = introSort(left, life-1)
	node.Next = center
	center.Next, last = introSort(right, life-1)
	return
}
