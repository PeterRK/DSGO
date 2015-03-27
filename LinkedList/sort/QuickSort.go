package sort

import (
	"LinkedList/list"
)

//快速排序，平均复杂度为O(NlogN) & O(logN)，最坏情况是O(N^2) & O(N)，不具有稳定性。
func QuickSort(head *list.Node) *list.Node {
	if head != nil {
		head, _ = doQuickSort(head)
	}
	return head
}
func doQuickSort(head *list.Node) (first *list.Node, last *list.Node) {
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
	var left, center, right, _ = part(head, node, node.Next)

	first, node = doQuickSort(left)
	node.Next = center
	center.Next, last = doQuickSort(right)
	return first, last
}

func part(node0 *list.Node, node1 *list.Node, node2 *list.Node) (left *list.Node, center *list.Node, right *list.Node, size int) {
	var tail = node2.Next

	if node0.Val > node1.Val { //a > b
		if node1.Val > node2.Val { //b > c		//a b c = b c a
			center, left, right = node1, node2, node0
		} else { //c >= b
			if node0.Val > node2.Val { //a > c	//a b c = c b a
				center, left, right = node2, node1, node0
			} else { //a b c = a b c
				center, left, right = node0, node1, node2
			}
		}
	} else { //b >= a
		if node2.Val > node0.Val { //c > a
			if node1.Val > node2.Val { //b > c	//a b c = c a b
				center, left, right = node2, node0, node1
			} else { //a b c = b a c
				center, left, right = node1, node0, node2
			}
		} else { //a b c = a c b
			center, left, right = node0, node2, node1
		}
	}

	size = 3
	node1, node2 = left, right
	for ; tail != nil; tail = tail.Next {
		if tail.Val < center.Val {
			node1.Next = tail
			node1 = tail
		} else {
			node2.Next = tail
			node2 = tail
		}
		size++
	}
	node1.Next, node2.Next = nil, nil
	return left, center, right, size
}
