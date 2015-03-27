package sort

import (
	"LinkedList/list"
)

//内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
func IntroSort(head *list.Node) *list.Node {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}
	var node = head.Next
	if node.Next == nil {
		if head.Val > node.Val {
			node.Next, head.Next = head, nil
			return node
		}
		return head
	}

	var left, center, right, size = part(head, node, node.Next)

	var life uint
	for life = 3; size != 0; life++ {
		size /= 2
	}

	head, node = doIntroSort(left, life-1)
	node.Next = center
	center.Next, _ = doIntroSort(right, life-1)
	return head
}
func doIntroSort(head *list.Node, life uint) (first *list.Node, last *list.Node) { //head != nil
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
		return doMergeSort(head)
	}

	var left, center, right, _ = part(head, node, node.Next)

	first, node = doIntroSort(left, life-1)
	node.Next = center
	center.Next, last = doIntroSort(right, life-1)
	return first, last
}
