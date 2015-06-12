package pairing

import (
	"unsafe"
)

func fakeHead(spt **Node) *Node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).next)
	return (*Node)(unsafe.Pointer(base - off))
}

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

func (hp *Heap) PopNode() *Node {
	var unit = hp.root
	if unit != nil {
		hp.root = collect(unit.child)
	}
	return unit
}
func (hp *Heap) Pop() (key int, fail bool) {
	var node = hp.PopNode()
	if node == nil {
		return 0, true
	}
	return node.key, false
}

func (hp *Heap) Remove(target *Node) {
	if target != nil {
		if target == hp.root { //根
			hp.root = collect(target.child)
		} else {
			var super = target.prev
			if super.child == target { //super为父
				super.child = super.hook(target.next)
			} else { //super为兄
				super.next = super.hook(target.next)
			}
		}
	}
}
func (hp *Heap) FloatUp(target *Node, value int) {
	if target != nil && value < target.key {
		target.key = value
		var super = target.prev
		if super != nil { //非根
			if super.child == target { //super为父
				if super.key > value { //但被超越
					super.child, target.next = super.hook(target.next), nil
					hp.root = merge(hp.root, target)
				}
			} else { //super为兄
				super.next, target.next = super.hook(target.next), nil
				hp.root = merge(hp.root, target)
			}
		}
	}
}
