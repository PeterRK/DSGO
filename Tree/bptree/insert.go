package bptree

import (
	"unsafe"
)

func (node *leaf) insert(place int, key int) (peer *leaf) {
	if node.cnt < LEAF_FULL {
		for i := node.cnt; i > place; i-- {
			node.data[i] = node.data[i-1]
		}
		node.data[place] = key
		node.cnt++
		return nil
	}

	peer = new(leaf)
	node.cnt, peer.cnt = LEAF_HALF, LEAF_HALF
	peer.next, node.next = node.next, peer
	if place < LEAF_HALF { //落在前半
		for i := 0; i < LEAF_HALF; i++ {
			peer.data[i] = node.data[i+(LEAF_HALF-1)]
		}
		for i := LEAF_HALF - 1; i > place; i-- {
			node.data[i] = node.data[i-1]
		}
		node.data[place] = key
	} else { //落在后半
		for i := LEAF_FULL; i > place; i-- {
			peer.data[i-LEAF_HALF] = node.data[i-1]
		}
		peer.data[place-LEAF_HALF] = key
		for i := LEAF_HALF; i < place; i++ {
			peer.data[i-LEAF_HALF] = node.data[i]
		}
	}
	return peer
}

func (node *index) insert(place int, kid *index) (peer *index) {
	if (-node.cnt) < INDEX_FULL {
		for i := (-node.cnt); i > place; i-- {
			node.data[i] = node.data[i-1]
			node.kids[i] = node.kids[i-1]
		}
		node.data[place] = kid.ceil()
		node.kids[place] = kid
		node.cnt--
		return nil
	}

	node.cnt = -INDEX_HALF
	peer = new(index)
	peer.cnt = -INDEX_HALF
	if place < INDEX_HALF { //落在前半
		for i := 0; i < INDEX_HALF; i++ {
			peer.data[i] = node.data[i+(INDEX_HALF-1)]
			peer.kids[i] = node.kids[i+(INDEX_HALF-1)]
		}
		for i := INDEX_HALF - 1; i > place; i-- {
			node.data[i] = node.data[i-1]
			node.kids[i] = node.kids[i-1]
		}
		node.data[place] = kid.ceil()
		node.kids[place] = kid
	} else { //落在后半
		for i := INDEX_FULL; i > place; i-- {
			peer.data[i-INDEX_HALF] = node.data[i-1]
			peer.kids[i-INDEX_HALF] = node.kids[i-1]
		}
		peer.data[place-INDEX_HALF] = kid.ceil()
		peer.kids[place-INDEX_HALF] = kid
		for i := INDEX_HALF; i < place; i++ {
			peer.data[i-INDEX_HALF] = node.data[i]
			peer.kids[i-INDEX_HALF] = node.kids[i]
		}
	}
	return peer
}

func (tree *Tree) Insert(key int) bool {
	if tree.root == nil {
		tree.head = new(leaf)
		tree.head.cnt = 1
		tree.head.next = nil
		tree.head.data[0] = key
		tree.root = (*index)(unsafe.Pointer(tree.head))
		return true
	}

	tree.path.clear()
	var node *leaf
	var place int

	var target = tree.root
	if key > tree.root.ceil() { //右界拓展
		for target.cnt < 0 { //index节点
			var idx = (-target.cnt) - 1
			target.data[idx] = key //之后难以修改，现在先改掉
			tree.path.push(target, idx)
			target = target.kids[idx]
		}
		node = (*leaf)(unsafe.Pointer(target))
		place = node.cnt
	} else {
		for target.cnt < 0 { //index节点
			var idx = target.locate(key)
			if key == target.data[idx] {
				return false
			}
			tree.path.push(target, idx)
			target = target.kids[idx]
		}
		node = (*leaf)(unsafe.Pointer(target)) //叶节点
		place = node.locate(key)
		if key == node.data[place] {
			return false
		}
	}

	var another = (*index)(unsafe.Pointer(node.insert(place, key)))
	for another != nil {
		if tree.path.isEmpty() {
			tree.root = new(index)
			tree.root.cnt = -2
			tree.root.data[0], tree.root.data[1] = target.ceil(), another.ceil()
			tree.root.kids[0], tree.root.kids[1] = target, another
			break
		}
		parent, idx := tree.path.pop()
		parent.data[idx] = target.ceil()
		target = parent
		another = target.insert(idx+1, another)
	}
	return true
}
