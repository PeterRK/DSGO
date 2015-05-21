package sort

import (
	"LinkedList/list"
)

//归并排序，复杂度为O(2NlogN) & O(1)，具有稳定性。
func MergeSort(head *list.Node) *list.Node {
	if head != nil {
		head, _ = doMergeSort(head)
	}
	return head
}
func doMergeSort(head *list.Node) (first *list.Node, last *list.Node) { //head != nil
	first, last = head, list.FakeHead(&first)
	var size = 0
	for ; head != nil; size += 2 {
		if head.Next == nil {
			last = head
			size++
			break
		}
		var node0, node1 = head, head.Next
		head = node1.Next
		if node0.Val > node1.Val {
			last.Next, node1.Next, node0.Next = node1, node0, head
			last = node0
		} else {
			last = node1
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
	return first, last
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
	return first, last
}
