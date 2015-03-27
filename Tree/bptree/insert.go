package bptree

//成功返回true，冲突返回false。
func (tree *Tree) Insert(key int) bool {
	if tree.root == nil {
		var unit = newLeaf()
		unit.cnt, unit.data[0], unit.next = 1, key, nil
		tree.head, tree.root = unit, unit.asIndex()
		return true
	}

	tree.path.clear()
	var unit, place = (*leaf)(nil), 0

	var target = tree.root
	if key > tree.root.ceil() { //右界拓展
		for target.inner {
			var idx = target.cnt - 1
			target.data[idx] = key //之后难以修改，现在先改掉
			tree.path.push(target, idx)
			target = target.kids[idx]
		}
		unit, place = target.asLeaf(), target.asLeaf().cnt
	} else {
		for target.inner {
			var idx = target.locate(key)
			if key == target.data[idx] {
				return false
			}
			tree.path.push(target, idx)
			target = target.kids[idx]
		}
		unit, place = target.asLeaf(), target.asLeaf().locate(key)
		if key == unit.data[place] {
			return false
		}
	}

	var peer = unit.insert(place, key).asIndex()
	for peer != nil {
		if tree.path.isEmpty() {
			var unit = newIndex()
			unit.cnt = 2
			unit.data[0], unit.data[1] = target.ceil(), peer.ceil()
			unit.kids[0], unit.kids[1] = target, peer
			tree.root, peer = unit, nil
		} else {
			var parent, idx = tree.path.pop()
			parent.data[idx] = target.ceil()
			target, peer = parent, parent.insert(idx+1, peer)
		}
	}
	return true
}
