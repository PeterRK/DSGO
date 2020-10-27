package sort

// 快速排序，平均复杂度为O(NlogN) & O(logN)，最坏情况是O(N^2) & O(N)，不具有稳定性。
// 这里采用递归实现，但实际上QuickSort不适合递归实现(有爆栈风险)。
func QuickSort(head *Node) *Node {
	if head != nil {
		head, _ = doQuickSort(head)
	}
	return head
}
func doQuickSort(head *Node) (first, last *Node) {
	first, last = sortOnlyTwo(head)
	if first == nil {
		left, center, right, _ := partition(head)
		var knot *Node
		first, knot = doQuickSort(left)
		knot.Next = center
		center.Next, last = doQuickSort(right)
	}
	return first, last
}

//head != nil
//done if firt != nil
func sortOnlyTwo(head *Node) (first, last *Node) {
	if head.Next == nil {
		return head, head
	}
	node := head.Next
	if node.Next == nil {
		if head.Val > node.Val {
			node.Next, head.Next = head, nil
			return node, head
		}
		return head, node
	}
	return nil, nil
}

func partition(nd0 *Node) (left, center, right *Node, size int) {
	nd1 := nd0.Next
	nd2 := nd1.Next
	tail := nd2.Next

	if nd1.Val > nd2.Val {
		nd1, nd2 = nd2, nd1
	}
	switch {
	case nd0.Val < nd1.Val:
		left, center, right = nd0, nd1, nd2
	case nd0.Val > nd2.Val:
		left, center, right = nd1, nd2, nd0
	default:
		left, center, right = nd1, nd0, nd2
	}

	size = 3
	nd1, nd2 = left, right
	for tail != nil {
		for nd1.Next = tail; tail != nil &&
			tail.Val <= center.Val; size++ {
			nd1, tail = tail, tail.Next
		}
		for nd2.Next = tail; tail != nil &&
			tail.Val > center.Val; size++ {
			nd2, tail = tail, tail.Next
		}
	}
	nd1.Next, nd2.Next = nil, nil
	return left, center, right, size
}
