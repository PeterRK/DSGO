package sort

func QuickSort(head *Node) *Node {
	if head != nil {
		head, _ = quickSort(head)
	}
	return head
}

func quickSort(list *Node) (head *Node, tail *Node) {
	if list.next == nil {
		return list, list
	}
	var first = list.next
	if first.next == nil {
		if list.key > first.key {
			first.next, list.next = list, nil
			return first, list
		}
		return list, first
	}
	var second = first.next
	tail = second.next

	if list.key > first.key { //a > b
		if first.key > second.key { //b > c		//b a c = c b a
			first, list, second = second, first, list
		} else { //c >= b
			if list.key > second.key { //a > c		//b a c = b c a
				list, second = second, list
			} //else b a c = b a c
		}
	} else { //b >= a
		if second.key > list.key { //c > a
			if first.key > second.key { //b > c		//b a c = a c b
				first, list, second = list, second, first
			} else { //b a c = a b c
				first, list = list, first
			}
		} else { //b a c = c a b
			first, second = second, first
		}
	}

	var tail1, tail2 = first, second
	for ; tail != nil; tail = tail.next {
		if tail.key < list.key {
			tail1.next = tail
			tail1 = tail
		} else {
			tail2.next = tail
			tail2 = tail
		}
	}
	tail1.next, tail2.next = nil, nil

	head, tail1 = quickSort(first)
	tail1.next = list
	list.next, tail = quickSort(second)
	return
}
