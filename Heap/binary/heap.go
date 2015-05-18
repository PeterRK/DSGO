package binaryheap

//二叉堆，底层采用数组，缺点是合并代价较高
type Heap struct {
	core []int
}

func (heap *Heap) Size() int {
	return len(heap.core)
}
func (heap *Heap) IsEmpty() bool {
	return len(heap.core) == 0
}

func (heap *Heap) Top() int {
	if heap.IsEmpty() {
		return 0
	}
	return heap.core[0]
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
func (heap *Heap) Pop() (key int) {
	var size = len(heap.core)
	if size == 0 {
		return 0
	}
	key = heap.core[0]
	if size == 1 {
		heap.core = heap.core[:0]
	} else {
		heap.core[0] = heap.core[size-1]
		heap.core = heap.core[:size-1]
		heap.adjustDown(0)
	}
	return
}
