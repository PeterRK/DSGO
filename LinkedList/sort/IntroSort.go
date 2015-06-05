package sort

import (
	"LinkedList/list"
)

//内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
//主要限制了QuickSort的最坏情况，适合递归实现(没有爆栈风险)。
func IntroSort(head *list.Node) *list.Node {
	if head == nil || head.Next == nil {
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

	var left, center, right, size = partition(head)

	var life uint
	for life = 12; size != 0; life++ {
		size /= 2
	}

	head, node = doIntroSort(left, life)
	node.Next = center
	center.Next, _ = doIntroSort(right, life)
	return head
}
func doIntroSort(head *list.Node, life uint) (first *list.Node, last *list.Node) {
	if head.Next == nil { //head != nil
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
		first, last = doMergeSort(head)
	} else {
		var left, center, right, _ = partition(head)
		first, node = doIntroSort(left, life-1)
		node.Next = center
		center.Next, last = doIntroSort(right, life-1)
	}
	return first, last
}
