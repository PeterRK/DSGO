package sort

import (
	ll "DSGO/linkedlist"
	"golang.org/x/exp/constraints"
)

// 快速排序，平均复杂度为O(NlogN) & O(logN)，最坏情况是O(N^2) & O(N)，不具有稳定性。
// 这里采用递归实现，但实际上QuickSort不适合递归实现(有爆栈风险)。
func QuickSort[T constraints.Ordered](list *ll.Node[T]) *ll.Node[T] {
	if list != nil {
		list, _ = quickSort(list)
	}
	return list
}

func quickSort[T constraints.Ordered](list *ll.Node[T]) (head, tail *ll.Node[T]) {
	head, tail = sort2(list)
	if head == nil {
		lst1, pivot, lst2, _ := partition(list)
		head, tail = quickSort(lst1)
		tail.Next = pivot
		pivot.Next, tail = quickSort(lst2)
	}
	return head, tail
}

//list != nil
//done if head != nil
func sort2[T constraints.Ordered](list *ll.Node[T]) (head, tail *ll.Node[T]) {
	if list.Next == nil {
		return list, list
	}
	head, tail = list, list.Next
	if tail.Next == nil {
		if head.Val > tail.Val {
			tail.Next, head.Next = head, nil
			return tail, head
		}
		return head, tail
	}
	return nil, nil
}

func partition[T constraints.Ordered](list *ll.Node[T]) (lst1, pivot, lst2 *ll.Node[T], size int) {
	a := list
	b := a.Next
	c := b.Next
	tail := c.Next

	if a.Val > c.Val {
		a, c = c, a
	}
	switch {
	case b.Val < a.Val:
		a, b = b, a
	case b.Val > c.Val:
		b, c = c, b
	}

	lst1, lst2 = a, c

	size = 3
	for tail != nil {
		for a.Next = tail; tail != nil &&
			tail.Val <= b.Val; size++ {
			a, tail = tail, tail.Next
		}
		for c.Next = tail; tail != nil &&
			tail.Val > b.Val; size++ {
			c, tail = tail, tail.Next
		}
	}
	a.Next, c.Next = nil, nil
	return lst1, b, lst2, size
}
