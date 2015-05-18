package binaryheap

func (heap *Heap) adjustDown(spot int) {
	var size = len(heap.core)
	var key = heap.core[spot]
	var left, right = spot*2 + 1, spot*2 + 2
	for right < size {
		var kid = 0
		if heap.core[left] < heap.core[right] {
			kid = left
		} else {
			kid = right
		}
		if key <= heap.core[kid] {
			goto LabelOver
		}
		heap.core[spot] = heap.core[kid]
		spot, left, right = kid, kid*2+1, kid*2+2
	}
	if right == size && key > heap.core[left] {
		heap.core[spot], heap.core[left] = heap.core[left], key
		return
	}
LabelOver:
	heap.core[spot] = key
}

func (heap *Heap) adjustUp(spot int) {
	var key = heap.core[spot]
	for spot > 0 {
		var parent = (spot - 1) / 2
		if heap.core[parent] <= key {
			break
		}
		heap.core[spot], spot = heap.core[parent], parent
	}
	heap.core[spot] = key
}
