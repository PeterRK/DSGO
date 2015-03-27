package binomial

//二项堆的Push和Top操作的复杂度为O(1)，其余核心操作复杂度为O(logN)。
type Heap struct {
	list *node
	top  *node
}
type node struct {
	key   int
	level uint
	peer  *node
	child *node
}

func (heap *Heap) IsEmpty() bool {
	return heap.list == nil
}

func (heap *Heap) Top() (key int, fail bool) {
	if heap.IsEmpty() {
		return 0, true
	}
	return heap.top.key, false
}
func (heap *Heap) Push(key int) {
	var unit = new(node)
	unit.key, unit.level = key, 0
	unit.peer, unit.child = nil, nil
	if heap.IsEmpty() {
		heap.list, heap.top = unit, unit
	} else {
		if key < heap.top.key {
			heap.top = unit
		}
		heap.merge(unit)
	}
}

//list是从少到多的，而child相反
func reverse(list *node) *node {
	var head *node = nil
	for list != nil {
		var current = list
		list = list.peer
		current.peer, head = head, current
	}
	return head
}
func (heap *Heap) Pop() (key int, fail bool) {
	if heap.IsEmpty() {
		return 0, true
	}
	key = heap.top.key

	var knot = fakeHead(&heap.list)
	for knot.peer != heap.top {
		knot = knot.peer
	}
	knot.peer = knot.peer.peer

	heap.merge(reverse(heap.top.child))
	heap.top = heap.list
	if heap.list != nil {
		for pt := heap.list.peer; pt != nil; pt = pt.peer {
			if pt.key < heap.top.key {
				heap.top = pt
			}
		}
	}
	return key, false
}
