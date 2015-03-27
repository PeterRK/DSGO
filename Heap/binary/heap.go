package binary

//二叉堆，底层采用数组。
//Build的复杂度为O(N)，Top的复杂度为O(1)，其余核心操作复杂度为O(logN)。
type Heap struct {
	core []int
}

func (heap *Heap) Size() int {
	return len(heap.core)
}
func (heap *Heap) IsEmpty() bool {
	return len(heap.core) == 0
}

func (heap *Heap) Top() (key int, fail bool) {
	if heap.IsEmpty() {
		return 0, true
	}
	return heap.core[0], false
}

func (heap *Heap) Build(list []int) {
	var size = len(list)
	heap.core = list
	for idx := size/2 - 1; idx >= 0; idx-- {
		heap.adjustDown(idx)
	}
}
func (heap *Heap) Push(key int) {
	var place = len(heap.core)
	heap.core = append(heap.core, key)
	heap.adjustUp(place)
}
func (heap *Heap) Pop() (key int, fail bool) {
	var size = heap.Size()
	if size == 0 {
		return 0, true
	}

	key = heap.core[0]
	if size == 1 {
		heap.core = heap.core[:0]
	} else {
		heap.core[0] = heap.core[size-1]
		heap.core = heap.core[:size-1]
		heap.adjustDown(0)
	}
	return key, false
}
