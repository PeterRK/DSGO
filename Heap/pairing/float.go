package pairingheap

//值挪移
func (heap *Heap) FloatUpX(target *Node, value int) *Node {
	if target == nil || target.key <= value {
		return target
	}
	for target != heap.root {
		var brother = target
		for brother.prev.child != brother {
			brother = brother.prev
		} //找到长兄节点和父节点
		var parent = brother.prev

		if parent.key <= value {
			break
		}
		target.key, target = parent.key, parent
	}
	target.key = value
	return target
}

//节点挪移
func (heap *Heap) FloatUp(target *Node, value int) {
	if target == nil || value >= target.key {
		return
	}
	target.key = value
	if target == heap.root {
		return
	}

	for {
		var brother = target
		for brother.prev.child != brother {
			brother = brother.prev
		} //找到长兄节点和父节点
		var parent = brother.prev
		if parent.key <= target.key {
			return
		}

		target.next, parent.next = parent.next, target.next
		if parent.next != nil {
			parent.next.prev = parent
		}
		if target.next != nil {
			target.next.prev = target
		}

		parent.child = target.child
		if parent.child != nil {
			parent.child.prev = parent
		}

		if brother != target {
			parent.prev, target.prev = target.prev, parent.prev
			parent.prev.next = parent
			target.child, brother.prev = brother, target
		} else { //target恰好是长兄
			target.prev = parent.prev
			target.child, parent.prev = parent, target
		}

		if target.prev == nil {
			heap.root = target
			break
		} else {
			var super = target.prev
			if super.next == parent {
				super.next = target
			} else {
				super.child = target
			}
		}
	}
}
