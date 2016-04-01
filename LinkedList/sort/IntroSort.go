package sort

import (
	"DSGO/LinkedList/list"
)

func log2ceil(num uint) uint {
	var ceil uint
	for ceil = 1; num != 0; ceil++ {
		num /= 2
	}
	return ceil
}

// 内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
// 主要限制了QuickSort的最坏情况，适合递归实现(没有爆栈风险)。
func IntroSort(head *list.Node) *list.Node {
	if head == nil {
		return nil
	}
	if first, _ := sortOnlyTwo(head); first != nil {
		return first
	}

	var left, center, right, size = partition(head)
	var life = log2ceil(uint(size)) * 2

	var knot *list.Node
	head, knot = doIntroSort(left, life)
	knot.Next = center
	center.Next, _ = doIntroSort(right, life)
	return head
}

func doIntroSort(head *list.Node, life uint) (first, last *list.Node) {
	first, last = sortOnlyTwo(head)
	if first == nil {
		if life == 0 {
			first, last = doMergeSort(head)
		} else {
			var left, center, right, _ = partition(head)
			var knot *list.Node
			first, knot = doIntroSort(left, life-1)
			knot.Next = center
			center.Next, last = doIntroSort(right, life-1)
		}
	}
	return first, last
}
