package sort

import (
	ll "DSGO/linkedlist"
	"golang.org/x/exp/constraints"
)

// 归并排序，复杂度为O(2NlogN) & O(1)，具有稳定性。
func MergeSort[T constraints.Ordered](list *ll.Node[T]) *ll.Node[T] {
	list, _ = mergeSort(list)
	return list
}

func mergeSort[T constraints.Ordered](list *ll.Node[T]) (head, tail *ll.Node[T]) {
	head, tail = list, ll.FakeHead(&head)
	size := 0
	for ; list != nil; size += 2 {
		if list.Next == nil {
			tail = list
			size++
			break
		}
		a, b := list, list.Next
		list = b.Next
		if a.Val > b.Val {
			tail.Next, b.Next, a.Next = b, a, list
			tail = a
		} else {
			tail = b
		}
	}
	for step := 2; step < size; step *= 2 {
		list, tail = head, ll.FakeHead(&head)
		for list != nil {
			lst1 := list
			lst2 := cut(lst1, step)
			list = cut(lst2, step)
			var node *ll.Node[T]
			tail.Next, node = merge(lst1, lst2)
			tail, node.Next = node, list
		}
	}
	return head, tail
}

func cut[T any](list *ll.Node[T], size int) *ll.Node[T] {
	for ; list != nil && size > 1; size-- {
		list = list.Next
	}
	if list != nil {
		list, list.Next = list.Next, nil
	}
	return list
}

func merge[T constraints.Ordered](lst1, lst2 *ll.Node[T]) (head, tail *ll.Node[T]) {
	head, tail = nil, ll.FakeHead(&head)
	for {
		tail.Next = lst1
		if lst2 == nil {
			break
		}
		for lst1 != nil && lst1.Val <= lst2.Val {
			tail, lst1 = lst1, lst1.Next
		}
		tail.Next = lst2
		if lst1 == nil {
			break
		}
		for lst2 != nil && lst1.Val > lst2.Val {
			tail, lst2 = lst2, lst2.Next
		}
	}
	for tail.Next != nil {
		tail = tail.Next
	}
	return head, tail
}
