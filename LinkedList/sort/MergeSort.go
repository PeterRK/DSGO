package sort

import (
	"DSGO/LinkedList/list"
)

// 归并排序，复杂度为O(2NlogN) & O(1)，具有稳定性。
func MergeSort(head *list.Node) *list.Node {
	head, _ = doMergeSort(head)
	return head
}
func doMergeSort(head *list.Node) (first, last *list.Node) { //head != nil
	first, last = head, list.FakeHead(&first)
	var size = 0
	for ; head != nil; size += 2 {
		if head.Next == nil {
			last = head
			size++
			break
		}
		var nd0, nd1 = head, head.Next
		head = nd1.Next
		if nd0.Val > nd1.Val {
			last.Next, nd1.Next, nd0.Next = nd1, nd0, head
			last = nd0
		} else {
			last = nd1
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
func merge(left, right *list.Node) (first, last *list.Node) {
	first, last = nil, list.FakeHead(&first)
	for {
		last.Next = left
		if right == nil {
			break
		}
		for left != nil && left.Val <= right.Val {
			last, left = left, left.Next
		}
		last.Next = right
		if left == nil {
			break
		}
		for right != nil && left.Val > right.Val {
			last, right = right, right.Next
		}
	}
	for last.Next != nil {
		last = last.Next
	}
	return first, last
}
