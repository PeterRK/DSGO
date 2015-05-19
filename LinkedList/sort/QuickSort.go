package sort

import (
	"linkedlist"
)

func QuickSort(head *linkedlist.Node) *linkedlist.Node {
	if head != nil {
		head, _ = quickSort(head)
	}
	return head
}

func quickSort(list *linkedlist.Node) (head *linkedlist.Node, tail *linkedlist.Node) {
	if list.Next == nil {
		return list, list
	}
	var first = list.Next
	if first.Next == nil {
		if list.Val > first.Val {
			first.Next, list.Next = list, nil
			return first, list
		}
		return list, first
	}
	var second = first.Next
	tail = second.Next

	if list.Val > first.Val { //a > b
		if first.Val > second.Val { //b > c		//b a c = c b a
			first, list, second = second, first, list
		} else { //c >= b
			if list.Val > second.Val { //a > c		//b a c = b c a
				list, second = second, list
			} //else b a c = b a c
		}
	} else { //b >= a
		if second.Val > list.Val { //c > a
			if first.Val > second.Val { //b > c		//b a c = a c b
				first, list, second = list, second, first
			} else { //b a c = a b c
				first, list = list, first
			}
		} else { //b a c = c a b
			first, second = second, first
		}
	}

	var tail1, tail2 = first, second
	for ; tail != nil; tail = tail.Next {
		if tail.Val < list.Val {
			tail1.Next = tail
			tail1 = tail
		} else {
			tail2.Next = tail
			tail2 = tail
		}
	}
	tail1.Next, tail2.Next = nil, nil

	head, tail1 = quickSort(first)
	tail1.Next = list
	list.Next, tail = quickSort(second)
	return
}
