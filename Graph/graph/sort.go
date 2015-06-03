package graph

import (
	"time"
)

func Sort(list []Edge) {
	var size = len(list)
	var life uint
	for life = 12; size != 0; life++ {
		size /= 2
	}
	magic = uint(time.Now().Unix())
	doIntroSort(list, life)
}
func doIntroSort(list []Edge, life uint) {
	var size = len(list)
	if size < 7 {
		insertSort(list)
	} else if life == 0 {
		heapSort(list)
	} else {
		var knot = partition(list)
		doIntroSort(list[0:knot], life-1)
		doIntroSort(list[knot+1:size], life-1)
	}
}

var magic = ^uint(0)

func partition(list []Edge) int {
	var size = len(list)

	var x, y = int(magic % uint(size-1)), int(magic % uint(size-2))
	magic = magic*1103515245 + 12345
	var a, b = 1 + x, 1 + (1+x+y)%(size-1) //a != b
	var barrier = list[0]
	if list[0].Dist > list[a].Dist {
		if list[a].Dist > list[b].Dist {
			barrier, list[a] = list[a], list[0]
		} else { //c >= b
			if list[0].Dist > list[b].Dist {
				barrier, list[b] = list[b], list[0]
			}
		}
	} else { //b >= a
		if list[b].Dist > list[0].Dist {
			if list[a].Dist > list[b].Dist {
				barrier, list[b] = list[b], list[0]
			} else {
				barrier, list[a] = list[a], list[0]
			}
		}
	}

	a, b = 1, size-1
	for { //注意对称性
		for list[a].Dist < barrier.Dist {
			a++
		}
		for list[b].Dist > barrier.Dist {
			b--
		}
		if a >= b {
			break
		}
		list[a], list[b] = list[b], list[a]
		a++
		b--
	}
	list[0], list[b] = list[b], barrier
	return b
}

func heapSort(list []Edge) {
	var size = len(list)
	for idx := size/2 - 1; idx >= 0; idx-- {
		down(list, idx)
	}
	for sz := size - 1; sz > 0; sz-- {
		list[0], list[sz] = list[sz], list[0]
		down(list[:sz], 0)
	}
}
func down(list []Edge, spot int) {
	var size = len(list)
	var key = list[spot]
	var left, right = spot*2 + 1, spot*2 + 2
	for right < size {
		var kid int
		if list[left].Dist > list[right].Dist {
			kid = left
		} else {
			kid = right
		}
		if key.Dist >= list[kid].Dist {
			goto LabelOver
		}
		list[spot] = list[kid]
		spot, left, right = kid, kid*2+1, kid*2+2
	}
	if right == size && key.Dist < list[left].Dist {
		list[spot], list[left] = list[left], key
		return
	}
LabelOver:
	list[spot] = key
}

func insertSort(list []Edge) {
	var size = len(list)
	for i := 1; i < size; i++ {
		var left, right = 0, i
		var key = list[i]
		for left < right {
			var mid = (left + right) / 2
			if key.Dist < list[mid].Dist {
				right = mid
			} else { //找第一个大于key的位置
				left = mid + 1
			}
		} //不会越界
		for j := i; j > left; j-- {
			list[j] = list[j-1]
		}
		list[left] = key
	}
}
