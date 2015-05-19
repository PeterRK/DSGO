package sort

import (
	"linkedlist"
)

func merge(left *linkedlist.Node, right *linkedlist.Node) (head *linkedlist.Node, tail *linkedlist.Node) {
	head, tail = nil, linkedlist.FakeHead(&head)
	for ; left != nil && right != nil; tail = tail.Next {
		if left.Val > right.Val {
			tail.Next, right = right, right.Next
		} else {
			tail.Next, left = left, left.Next
		}
	}
	for ; left != nil; tail = tail.Next {
		tail.Next, left = left, left.Next
	}
	for ; right != nil; tail = tail.Next {
		tail.Next, right = right, right.Next
	}
	return
}

func MergeSort(head *linkedlist.Node) *linkedlist.Node {
	var stop = false
	for step := 1; !stop; step *= 2 {
		tail, knot := head, linkedlist.FakeHead(&head)
		stop = true
		for {
			var left = tail
			for i := 1; i < step && tail != nil; i++ {
				tail = tail.Next
			}
			if tail == nil || tail.Next == nil {
				break
			}
			stop = false
			var last = tail
			tail = tail.Next
			last.Next = nil

			var right = tail
			for i := 1; i < step && tail != nil; i++ {
				tail = tail.Next
			}
			if tail != nil {
				last = tail
				tail = tail.Next
				last.Next = nil
			}

			knot.Next, last = merge(left, right)
			last.Next = tail
			knot = last
		}
	}
	return head
}
