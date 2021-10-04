package sort

import (
	ll "DSGO/linkedlist"
	"DSGO/utils"
	"constraints"
)

// 内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
// 主要限制了QuickSort的最坏情况，适合递归实现(没有爆栈风险)。
func IntroSort[T constraints.Ordered](list *ll.Node[T]) *ll.Node[T] {
	if list == nil {
		return nil
	} else if head, _ := sort2(list); head != nil {
		return head
	}

	lst1, pivot, lst2, size := partition(list)
	life := utils.Log2Ceil(uint(size)) * 2

	list, tail := introSort(lst1, life)
	tail.Next = pivot
	pivot.Next, _ = introSort(lst2, life)
	return list
}

func introSort[T constraints.Ordered](list *ll.Node[T], life uint) (head, tail *ll.Node[T]) {
	head, tail = sort2(list)
	if head == nil {
		if life == 0 {
			head, tail = mergeSort(list)
		} else {
			lst1, pivot, lst2, _ := partition(list)
			head, tail = introSort(lst1, life-1)
			tail.Next = pivot
			pivot.Next, tail = introSort(lst2, life-1)
		}
	}
	return head, tail
}
