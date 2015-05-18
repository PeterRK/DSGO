package pairingheap

//配对堆的理论复杂度较好，Pop和floatUp为O(log N)，其余核心操作为O(1)
//虽然Fibonacci堆中floatUp操作的理论复杂度更好，但配对堆实际上更为实用

type Node struct {
	key     int
	child   *Node
	brother *Node
	prev    *Node
}
type Heap struct {
	root *Node
}

func (heap *Heap) IsEmpty() bool {
	return heap.root == nil
}
func (heap *Heap) Top() int {
	if heap.IsEmpty() {
		return 0
	}
	return heap.root.key
}

func merge(one *Node, another *Node) *Node {
	if one.key > another.key {
		one, another = another, one
	}
	another.brother = one.child
	if one.child != nil {
		one.child.prev = another
	}
	one.child, another.prev = another, one
	return one
}
func (heap *Heap) Merge(peer *Heap) {
	if heap == peer || peer.root == nil {
		return
	}
	if heap.root == nil {
		heap.root = peer.root
	} else {
		heap.root = merge(heap.root, peer.root)
	}
	peer.root = nil
}

//这货Push时不怎么管整理，到Pop时再做
func (heap *Heap) PushNode(unit *Node) {
	if unit == nil {
		return
	}
	unit.child, unit.brother, unit.prev = nil, nil, nil
	if heap.root == nil {
		heap.root = unit
	} else {
		heap.root = merge(heap.root, unit)
	}
}
func (heap *Heap) Push(key int) *Node {
	var unit = new(Node)
	unit.key = key
	heap.PushNode(unit)
	return unit
}
