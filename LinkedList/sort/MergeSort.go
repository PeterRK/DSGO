package sort

import (
	"LinkedList/list"
)

func MergeSort(head *list.Node) *list.Node {
	if head != nil {
		head, _ = mergeSort(head)
	}
	return head
}
func mergeSort(head *list.Node) (first *list.Node, last *list.Node) { //head != nil
	first, last = head, list.FakeHead(&first)
	var size = 0
	for ; head != nil; size += 2 {
		if head.Next == nil {
			last = head
			size++
			break
		}
		var one, another = head, head.Next
		head = another.Next
		if one.Val > another.Val {
			last.Next, another.Next, one.Next = another, one, head
			last = one
		} else {
			last = another
		}
	}

	for step := 2; step < size; step *= 2 {
		head, last = first, list.FakeHead(&first)
		for head != nil {
			var left, right, node *list.Node
			left, head = head, cutPeice(head, step)
			right, head = head, cutPeice(head, step)

			last.Next, node = merge(left, right)
			last, node.Next = node, head
		}
	}
	return
}
func cutPeice(head *list.Node, sz int) *list.Node {
	for i := 1; i < sz && head != nil; i++ {
		head = head.Next
	}
	if head != nil {
		var last = head
		head, last.Next = head.Next, nil
	}
	return head
}

func merge(left *list.Node, right *list.Node) (first *list.Node, last *list.Node) {
	first, last = nil, list.FakeHead(&first)
	for ; left != nil && right != nil; last = last.Next {
		if left.Val > right.Val {
			last.Next, right = right, right.Next
		} else {
			last.Next, left = left, left.Next
		}
	}
	for ; left != nil; last = last.Next {
		last.Next, left = left, left.Next
	}
	for ; right != nil; last = last.Next {
		last.Next, right = right, right.Next
	}
	return
}
