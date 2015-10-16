package pairing

import (
	"errors"
	"unsafe"
)

func fakeHead(spt **Node) *Node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).next)
	return (*Node)(unsafe.Pointer(base - off))
}

/*
func collect(head *Node) *Node {
	if head != nil {
		for head.next != nil {
			var list, knot = head, fakeHead(&head)
			for list != nil && list.next != nil { //两两配对
				var one, another = list, list.next
				list = another.next
				knot.next = merge(one, another)
				knot = knot.next
			}
			knot.next = list
		}
		head.prev = nil
	}
	return head
}
*/

func collect(head *Node) *Node {
	if head != nil && head.next != nil {
		var list, last = head, fakeHead(&head)
		for list != nil && list.next != nil { //两两配对
			var one, another = list, list.next
			list = another.next
			last.next = last.hook(merge(one, another))
			last = last.next
		}
		head.prev = nil
		if list == nil {
			head, list = last, last.prev
		} else {
			head, list = list, last
		}
		for list != nil {
			last, list = list, list.prev
			head = merge(head, last)
		}
		head.prev, head.next = nil, nil
	}
	return head
}

func (hp *Heap) PopNode() *Node {
	var unit = hp.root
	if unit != nil {
		hp.root = collect(unit.child)
	}
	return unit
}
func (hp *Heap) Pop() (int, error) {
	var node = hp.PopNode()
	if node == nil {
		return 0, errors.New("empty")
	}
	return node.key, nil
}

func (hp *Heap) Remove(target *Node) {
	if target != nil {
		if super := target.prev; super == nil { //根
			hp.root = collect(target.child)
		} else {
			if super.child == target { //super为父
				super.child = super.hook(target.next)
			} else { //super为兄
				super.next = super.hook(target.next)
			}
			if target = collect(target.child); target != nil {
				hp.root = merge(hp.root, target)
			}
		}
	}
}
func (hp *Heap) FloatUp(target *Node, value int) {
	if target != nil && value < target.key {
		target.key = value
		if super := target.prev; super != nil && super.key > value {
			target.prev = nil
			if super.next == target { //super为兄
				super.next, target.next = super.hook(target.next), nil
			} else { //super为父
				super.child, target.next = super.hook(target.next), nil
			}
			hp.root = merge(hp.root, target)
		}
	}
}
