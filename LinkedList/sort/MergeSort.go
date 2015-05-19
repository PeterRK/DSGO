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
func mergeSort(head *list.Node) (first *list.Node, last *list.Node) {
	var stop = false
	for step := 1; !stop; step *= 2 {
		tail, knot := head, list.FakeHead(&head)
		stop = true
		for {
			var left = tail
			for i := 1; i < step && tail != nil; i++ {
				last, tail = tail, tail.Next
			}
			if tail == nil {
				break
			} else if tail.Next == nil {
				last = tail
				break
			}
			stop = false
			last, tail = tail, tail.Next
			last.Next = nil

			var right = tail
			for i := 1; i < step && tail != nil; i++ {
				last, tail = tail, tail.Next
			}
			if tail != nil {
				last, tail = tail, tail.Next
				last.Next = nil
			}

			knot.Next, last = merge(left, right)
			last.Next = tail
			knot = last
		}
	}
	return head, last
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
