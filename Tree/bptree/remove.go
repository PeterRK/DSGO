package bptree

//B+树的删除比较复杂，本实现采用积极合并策略。
//成功返回true，没有返回false。
func (tree *Tree) Remove(key int) bool {
	if tree.root == nil ||
		key > tree.root.ceil() {
		return false
	}
	tree.path.clear()

	var target = tree.root
	for target.inner {
		var idx = target.locate(key)
		tree.path.push(target, idx)
		target = target.kids[idx]
	}
	var unit = target.asLeaf()
	var place = unit.locate(key)
	if key != unit.data[place] {
		return false
	}

	unit.remove(place)
	if tree.path.isEmpty() {
		if unit.cnt == 0 {
			tree.root, tree.head = nil, nil
		}
		return true
	} //除了到根节点，unit.cnt >= 2
	var shrink, new_ceil = (place == unit.cnt), unit.ceil()

	parent, place := tree.path.pop()
	if unit.cnt <= LEAF_QUARTER {
		var peer = unit
		if place == parent.cnt-1 {
			unit = parent.kids[place-1].asLeaf()
		} else {
			place++
			peer, shrink = parent.kids[place].asLeaf(), false
		}
		var combined = unit.combine(peer)
		parent.data[place-1] = unit.ceil()

		for combined {
			var unit = parent //此后代码与之前类似，但unit的类型已经不同
			unit.remove(place)
			if tree.path.isEmpty() {
				if unit.cnt == 1 {
					tree.root = unit.kids[0]
				}
				return true
			}

			parent, place = tree.path.pop()
			if unit.cnt <= INDEX_QUARTER {
				var peer = unit
				if place == parent.cnt-1 {
					unit = parent.kids[place-1]
				} else {
					place++
					peer, shrink = parent.kids[place], false
				}
				combined = unit.combine(peer)
				parent.data[place-1] = unit.ceil()
				continue
			}
			break
		}
	}
	if shrink { //
		parent.data[place] = new_ceil
		for place == parent.cnt-1 && //级联
			!tree.path.isEmpty() {
			parent, place = tree.path.pop()
			parent.data[place] = new_ceil
		}
	}
	return true
}
