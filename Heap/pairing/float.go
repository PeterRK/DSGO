package pairing

//值挪移
func (hp *Heap) FloatUpX(target *Node, value int) *Node {
	if target == nil || target.key <= value {
		return target
	}
	for target != hp.root {
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
func (hp *Heap) FloatUp(target *Node, value int) {
	if target == nil || value >= target.key {
		return
	}
	target.key = value
	if target == hp.root {
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

		target.next, parent.next = target.hook(parent.next), parent.hook(target.next)
		parent.child = parent.hook(target.child)

		if brother != target {
			parent.prev, target.prev = target.prev, parent.prev
			parent.prev.next = parent
			target.child, brother.prev = brother, target
		} else { //target恰好是长兄
			target.prev = parent.prev
			target.child, parent.prev = parent, target
		}

		if target.prev == nil {
			hp.root = target
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
