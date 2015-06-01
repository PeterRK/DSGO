package bptree

//成功返回true，冲突返回false。
func (tr *Tree) Insert(key int) bool {
	if tr.root == nil {
		var unit = newLeaf()
		unit.cnt, unit.data[0], unit.next = 1, key, nil
		tr.head, tr.root = unit, unit.asIndex()
		return true
	}

	tr.path.clear()
	var unit, place = (*leaf)(nil), 0

	var target = tr.root
	if key > tr.root.ceil() { //右界拓展
		for target.inner {
			var idx = target.cnt - 1
			target.data[idx] = key //之后难以修改，现在先改掉
			tr.path.push(target, idx)
			target = target.kids[idx]
		}
		unit, place = target.asLeaf(), target.asLeaf().cnt
	} else {
		for target.inner {
			var idx = target.locate(key)
			if key == target.data[idx] {
				return false
			}
			tr.path.push(target, idx)
			target = target.kids[idx]
		}
		unit, place = target.asLeaf(), target.asLeaf().locate(key)
		if key == unit.data[place] {
			return false
		}
	}

	var peer = unit.insert(place, key).asIndex()
	for peer != nil {
		if tr.path.isEmpty() {
			var unit = newIndex()
			unit.cnt = 2
			unit.data[0], unit.data[1] = target.ceil(), peer.ceil()
			unit.kids[0], unit.kids[1] = target, peer
			tr.root, peer = unit, nil
		} else {
			var parent, idx = tr.path.pop()
			parent.data[idx] = target.ceil()
			target, peer = parent, parent.insert(idx+1, peer)
		}
	}
	return true
}
