package sort

import (
	"LinkedList/list"
)

//快速排序，平均复杂度为O(NlogN) & O(logN)，最坏情况是O(N^2) & O(N)，不具有稳定性。
//这里采用递归实现，但实际上QuickSort不适合递归实现(有爆栈风险)。
func QuickSort(head *list.Node) *list.Node {
	if head != nil {
		head, _ = doQuickSort(head)
	}
	return head
}
func doQuickSort(head *list.Node) (first *list.Node, last *list.Node) {
	first, last = sortOnlyTwo(head)
	if first == nil {
		var left, center, right, _ = partition(head)
		var knot *list.Node
		first, knot = doQuickSort(left)
		knot.Next = center
		center.Next, last = doQuickSort(right)
	}
	return first, last
}

//head != nil
//done if firt != nil
func sortOnlyTwo(head *list.Node) (first *list.Node, last *list.Node) {
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
	return nil, nil
}

func partition(node0 *list.Node) (left *list.Node, center *list.Node, right *list.Node, size int) {
	var node1 = node0.Next
	var node2 = node1.Next
	var tail = node2.Next

	if node1.Val > node2.Val {
		node1, node2 = node2, node1
	}
	switch {
	case node0.Val < node1.Val:
		left, center, right = node0, node1, node2
	case node0.Val > node2.Val:
		left, center, right = node1, node2, node0
	default:
		left, center, right = node1, node0, node2
	}

	size = 3
	node1, node2 = left, right
	for tail != nil {
		for node1.Next = tail; tail != nil && tail.Val <= center.Val; size++ {
			node1, tail = tail, tail.Next
		}
		for node2.Next = tail; tail != nil && tail.Val > center.Val; size++ {
			node2, tail = tail, tail.Next
		}
	}
	node1.Next, node2.Next = nil, nil
	return left, center, right, size
}
