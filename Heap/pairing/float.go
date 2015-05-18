package pairingheap

func (heap *Heap) FloatUp(target *Node, value int) {
	if target == nil || value >= target.key {
		return
	}
	target.key = value
	if target == heap.root {
		return
	}

	for {
		var big_bro = target
		for big_bro.prev.child != big_bro {
			big_bro = big_bro.prev
		}
		var parent = big_bro.prev
		if parent.key <= target.key {
			return
		}

		parent.brother, target.brother = target.brother, parent.brother
		if parent.brother != nil {
			parent.brother.prev = parent
		}
		if target.brother != nil {
			target.brother.prev = target
		}

		parent.child = target.child
		if parent.child != nil {
			parent.child.prev = parent
		}

		if big_bro != target {
			parent.prev, target.prev = target.prev, parent.prev
			parent.prev.brother = parent
			target.child, big_bro.prev = big_bro, target
		} else { //target恰好是左子
			target.prev = parent.prev
			target.child, parent.prev = parent, target
		}

		if target.prev == nil {
			heap.root = target
			break
		} else {
			var super = target.prev
			if super.brother == parent {
				super.brother = target
			} else {
				super.child = target
			}
		}
	}
}
