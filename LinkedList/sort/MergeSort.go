package sort

func mergeList(left *Node, right *Node) (head *Node, tail *Node) {
	head, tail = nil, FakeHead(&head)
	for ; left != nil && right != nil; tail = tail.next {
		if left.key > right.key {
			tail.next, right = right, right.next
		} else {
			tail.next, left = left, left.next
		}
	}
	for ; left != nil; tail = tail.next {
		tail.next, left = left, left.next
	}
	for ; right != nil; tail = tail.next {
		tail.next, right = right, right.next
	}
	return
}

func MergeSort(head *Node) *Node {
	var stop = false
	for step := 1; !stop; step *= 2 {
		tail, knot := head, FakeHead(&head)
		stop = true
		for {
			var left = tail
			for i := 1; i < step && tail != nil; i++ {
				tail = tail.next
			}
			if tail == nil || tail.next == nil {
				break
			}
			stop = false
			var last = tail
			tail = tail.next
			last.next = nil

			var right = tail
			for i := 1; i < step && tail != nil; i++ {
				tail = tail.next
			}
			if tail != nil {
				last = tail
				tail = tail.next
				last.next = nil
			}

			knot.next, last = mergeList(left, right)
			last.next = tail
			knot = last
		}
	}
	return head
}
