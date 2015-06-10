package pairing

//配对堆的理论复杂度较好，Pop和FloatUp为O(log N)，其余核心操作为O(1)。
//虽然Fibonacci堆中FloatUp操作的理论复杂度更好，但配对堆实际上更为实用。
type Heap struct {
	root *Node
}
type Node struct {
	key   int
	child *Node
	prev  *Node //父兄节点
	next  *Node //弟节点
}

func (hp *Heap) IsEmpty() bool {
	return hp.root == nil
}
func (hp *Heap) Top() (key int, fail bool) {
	if hp.IsEmpty() {
		return 0, true
	}
	return hp.root.key, false
}

func merge(one *Node, another *Node) *Node {
	if one.key > another.key {
		one, another = another, one
	}
	another.next = one.child
	if one.child != nil {
		one.child.prev = another
	}
	one.child, another.prev = another, one
	return one
}
func (hp *Heap) Merge(peer *Heap) {
	if hp == peer || peer.root == nil {
		return
	}
	if hp.root == nil {
		hp.root = peer.root
	} else {
		hp.root = merge(hp.root, peer.root)
	}
	peer.root = nil
}

//这货Push时不怎么管整理，到Pop时再做
func (hp *Heap) PushNode(unit *Node) {
	if unit != nil {
		unit.prev, unit.next, unit.child = nil, nil, nil
		if hp.root == nil {
			hp.root = unit
		} else {
			hp.root = merge(hp.root, unit)
		}
	}
}
func (hp *Heap) Push(key int) *Node {
	var unit = new(Node)
	unit.key = key
	hp.PushNode(unit)
	return unit
}
