package bptree

//B+树的删除比较复杂，本实现采用积极合并策略。
//成功返回true，没有返回false。
func (tr *Tree) Remove(key int) bool {
	if tr.root == nil ||
		key > tr.root.ceil() {
		return false
	}
	tr.path.clear()

	target := tr.root
	for target.inner {
		idx := target.locate(key)
		tr.path.push(target, idx)
		target = target.kids[idx]
	}
	unit := target.asLeaf()
	place := unit.locate(key)
	if key != unit.data[place] {
		return false
	}

	unit.remove(place)
	if tr.path.isEmpty() {
		if unit.cnt == 0 {
			tr.root, tr.head = nil, nil
		}
		return true
	} //除了到根节点，unit.cnt >= 2
	shrink, new_ceil := (place == unit.cnt), unit.ceil()

	parent, place := tr.path.pop()
	if unit.cnt <= LEAF_QUARTER {
		peer := unit
		if place == parent.cnt-1 {
			unit = parent.kids[place-1].asLeaf()
		} else {
			place++
			peer, shrink = parent.kids[place].asLeaf(), false
		}
		combined := unit.combine(peer)
		parent.data[place-1] = unit.ceil()

		for combined {
			unit := parent //此后代码与之前类似，但unit的类型已经不同
			unit.remove(place)
			if tr.path.isEmpty() {
				if unit.cnt == 1 {
					tr.root = unit.kids[0]
				}
				return true
			}

			parent, place = tr.path.pop()
			if unit.cnt <= INDEX_QUARTER {
				peer := unit
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
			!tr.path.isEmpty() {
			parent, place = tr.path.pop()
			parent.data[place] = new_ceil
		}
	}
	return true
}
