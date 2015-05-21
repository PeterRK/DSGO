package binomialheap

type node struct {
	key     int
	height  uint
	brother *node
	child   *node
}
type Heap struct {
	list *node
} //二项堆的Top和Push可以到O(1)，但本实现的Top为O(log N)

func (heap *Heap) IsEmpty() bool {
	return heap.list == nil
}

func (heap *Heap) Top() (key int, fail bool) {
	if heap.IsEmpty() {
		return 0, true
	}
	var best = heap.list.key
	for pt := heap.list.brother; pt != nil; pt = pt.brother {
		if pt.key < best {
			best = pt.key
		}
	}
	return best, false
}
func (heap *Heap) Push(key int) {
	var peer = new(node)
	peer.key = key
	peer.height = 0
	peer.brother, peer.child = nil, nil
	heap.merge(peer)
}

func reverse(head *node) *node {
	if head == nil {
		return nil
	}
	var tail = head.brother
	head.brother = nil
	for tail != nil {
		var current = tail
		tail = tail.brother
		current.brother, head = head, current
	}
	return head
}
func (heap *Heap) Pop() (key int, fail bool) {
	if heap.IsEmpty() {
		return 0, true
	}

	var knot = fakeHead(&heap.list)
	for pt := heap.list; pt.brother != nil; pt = pt.brother {
		if pt.brother.key < knot.brother.key {
			knot = pt
		}
	}
	key = knot.brother.key
	var list = reverse(knot.brother.child)
	knot.brother = knot.brother.brother
	heap.merge(list)
	return key, false
}
