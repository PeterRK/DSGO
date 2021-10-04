package bplus

import (
	"constraints"
	"unsafe"
)

type bNode[T constraints.Ordered] struct {
	inner bool
	view  []T
}

func (node *bNode[T]) ceil() T {
	return node.view[len(node.view)-1]
}

const iHalfSize = 15
const iFullSize = iHalfSize * 2
const iQuarterSize = iHalfSize / 2

type iNode[T constraints.Ordered] struct {
	bNode[T] //inner==true
	data     [iFullSize]T
	kids     [iFullSize]*iNode[T]
}

func newIndex[T constraints.Ordered]() *iNode[T] {
	node := new(iNode[T])
	node.inner = true
	node.view = node.data[:0]
	return node
}

const lHalfSize = 30
const lFullSize = lHalfSize * 2
const lQuarterSize = lHalfSize / 2

type lNode[T constraints.Ordered] struct {
	bNode[T] //inner==false
	data     [lFullSize]T
	next     *lNode[T]
}

func newLeaf[T constraints.Ordered]() *lNode[T] {
	node := new(lNode[T])
	//node.inner = false
	node.view = node.data[:0]
	//node.next = nil
	return node
}

func (node *iNode[T]) asLeaf() *lNode[T] {
	return (*lNode[T])(unsafe.Pointer(node))
}

func (node *lNode[T]) asIndex() *iNode[T] {
	return (*iNode[T])(unsafe.Pointer(node))
}

func (node *iNode[T]) remove(pos int) {
	for i := pos; i < len(node.view)-1; i++ {
		node.data[i] = node.data[i+1]
		node.kids[i] = node.kids[i+1]
	}
	node.view = node.data[:len(node.view)-1]
}

func (node *lNode[T]) remove(pos int) {
	for i := pos; i < len(node.view)-1; i++ {
		node.data[i] = node.data[i+1]
	}
	node.view = node.data[:len(node.view)-1]
}

//peer为分裂项，peer为nil时表示不分裂
func (node *iNode[T]) insert(pos int, kid *iNode[T]) (peer *iNode[T]) {
	if len(node.view) < iFullSize {
		for i := len(node.view); i > pos; i-- {
			node.data[i] = node.data[i-1]
			node.kids[i] = node.kids[i-1]
		}
		node.view = node.data[:len(node.view)+1]
		node.data[pos] = kid.ceil()
		node.kids[pos] = kid
		return nil
	}
	//iFullSize+1 => iHalfSize+1 & iHalfSize
	peer = newIndex[T]()
	if pos <= iHalfSize {
		for i := 0; i < iHalfSize; i++ {
			peer.data[i] = node.data[i+iHalfSize]
			peer.kids[i] = node.kids[i+iHalfSize]
		}
		for i := iHalfSize; i > pos; i-- {
			node.data[i] = node.data[i-1]
			node.kids[i] = node.kids[i-1]
		}
		node.data[pos] = kid.ceil()
		node.kids[pos] = kid
	} else {
		pos -= iHalfSize + 1
		for i := 0; i < pos; i++ {
			peer.data[i] = node.data[i+(iHalfSize+1)]
			peer.kids[i] = node.kids[i+(iHalfSize+1)]
		}
		peer.data[pos] = kid.ceil()
		peer.kids[pos] = kid
		for i := pos + 1; i < iHalfSize; i++ {
			peer.data[i] = node.data[i+iHalfSize]
			peer.kids[i] = node.kids[i+iHalfSize]
		}
	}
	node.view = node.data[:iHalfSize+1]
	peer.view = peer.data[:iHalfSize]
	return peer
}

//peer为分裂项，peer为nil时表示不分裂
func (node *lNode[T]) insert(pos int, key T) (peer *lNode[T]) {
	if len(node.view) < lFullSize {
		for i := len(node.view); i > pos; i-- {
			node.data[i] = node.data[i-1]
		}
		node.view = node.data[:len(node.view)+1]
		node.data[pos] = key
		return nil
	}
	//lFullSize+1 => lHalfSize+1 & lHalfSize
	peer = newLeaf[T]()
	if pos <= lHalfSize {
		for i := 0; i < lHalfSize; i++ {
			peer.data[i] = node.data[i+lHalfSize]
		}
		for i := lHalfSize; i > pos; i-- {
			node.data[i] = node.data[i-1]
		}
		node.data[pos] = key
	} else {
		pos -= lHalfSize + 1
		for i := 0; i < pos; i++ {
			peer.data[i] = node.data[i+(lHalfSize+1)]
		}
		peer.data[pos] = key
		for i := pos + 1; i < lHalfSize; i++ {
			peer.data[i] = node.data[i+lHalfSize]
		}
	}
	node.view = node.data[:lHalfSize+1]
	peer.view = peer.data[:lHalfSize]
	peer.next, node.next = node.next, peer
	return peer
}

//要求peer为node后的节点，发生合并返回true
func (node *iNode[T]) combine(peer *iNode[T]) bool {
	total := len(node.view) + len(peer.view)
	if total < iHalfSize+iQuarterSize {
		for i := len(node.view); i < total; i++ {
			node.data[i] = peer.data[i-len(node.view)]
			node.kids[i] = peer.kids[i-len(node.view)]
		}
		node.view = node.data[:total]
		return true
	}
	//分流而不合并
	nodeCnt := total / 2
	if len(node.view) == nodeCnt {
		return false
	}
	peerCnt := total - nodeCnt
	if len(peer.view) > peerCnt { //后填前
		diff := len(peer.view) - peerCnt
		for i := len(node.view); i < nodeCnt; i++ {
			node.data[i] = peer.data[i-len(node.view)]
			node.kids[i] = peer.kids[i-len(node.view)]
		}
		for i := 0; i < peerCnt; i++ {
			peer.data[i] = peer.data[i+diff]
			peer.kids[i] = peer.kids[i+diff]
		}
	} else { //前填后
		diff := peerCnt - len(peer.view)
		for i := len(peer.view) - 1; i >= 0; i-- {
			peer.data[i+diff] = peer.data[i]
			peer.kids[i+diff] = peer.kids[i]
		}
		for i := len(node.view) - 1; i >= nodeCnt; i-- {
			peer.data[i-nodeCnt] = node.data[i]
			peer.kids[i-nodeCnt] = node.kids[i]
		}
	}
	node.view = node.data[:nodeCnt]
	peer.view = peer.data[:peerCnt]
	return false
}

//要求peer为unit后的节点，发生合并返回 true
func (node *lNode[T]) combine(peer *lNode[T]) bool {
	total := len(node.view) + len(peer.view)
	if total <= lHalfSize+lQuarterSize {
		for i := len(node.view); i < total; i++ {
			node.data[i] = peer.data[i-len(node.view)]
		}
		node.view = node.data[:total]
		return true
	}
	//分流而不合并
	nodeCnt := total / 2
	if len(node.view) == nodeCnt {
		return false
	}
	peerCnt := total - nodeCnt
	if len(peer.view) > peerCnt { //后填前
		diff := len(peer.view) - peerCnt
		for i := len(node.view); i < nodeCnt; i++ {
			node.data[i] = peer.data[i-len(node.view)]
		}
		for i := 0; i < peerCnt; i++ {
			peer.data[i] = peer.data[i+diff]
		}
	} else { //前填后
		diff := peerCnt - len(peer.view)
		for i := len(peer.view) - 1; i >= 0; i-- {
			peer.data[i+diff] = peer.data[i]
		}
		for i := len(node.view) - 1; i >= nodeCnt; i-- {
			peer.data[i-nodeCnt] = node.data[i]
		}
	}
	node.view = node.data[:nodeCnt]
	peer.view = peer.data[:peerCnt]
	return false
}
