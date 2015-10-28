package span

import (
	"Graph/graph"
)

func sort(list []graph.Edge) {
	var life = uint(12)
	for sz := len(list); sz != 0; sz /= 2 {
		life++
	}
	doIntroSort(list, life)
}
func doIntroSort(list []graph.Edge, life uint) {
	if life == 0 {
		heapSort(list)
	} else if len(list) < 8 {
		insertSort(list)
	} else {
		var start, end = partition(list)
		if list[start] != list[end-1] {
			doIntroSort(list[start+1:end-1], life-1)
		}
		doIntroSort(list[:start], life-1)
		doIntroSort(list[end:], life-1)
	}
}

func partition(list []graph.Edge) (start int, end int) {
	var size = len(list)
	var m1, m2 = len(list)/2 - 1, len(list) / 2
	if list[m1].Weight > list[m2].Weight {
		m1, m2 = m2, m1
	}
	var pivot1, pivot2 = list[m1], list[m2]
	list[m1], list[m2] = list[0], list[size-1]

	var left, right = 1, size - 2
	for k := left; k <= right; k++ {
		if list[k].Weight > pivot2.Weight {
			for k < right && list[right].Weight > pivot2.Weight {
				right--
			}
			list[k], list[right] = list[right], list[k]
			right--
		}
		if list[k].Weight < pivot1.Weight {
			list[k], list[left] = list[left], list[k]
			left++
		}
	}

	list[0], list[left-1] = list[left-1], pivot1
	list[size-1], list[right+1] = list[right+1], pivot2
	return left - 1, right + 2
}

func heapSort(list []graph.Edge) {
	for idx := len(list)/2 - 1; idx >= 0; idx-- {
		down(list, idx)
	}
	for sz := len(list) - 1; sz > 0; sz-- {
		list[0], list[sz] = list[sz], list[0]
		down(list[:sz], 0)
	}
}
func down(list []graph.Edge, spot int) {
	var key = list[spot]
	var left, right = spot*2 + 1, spot*2 + 2
	for right < len(list) {
		var kid int
		if list[left].Weight > list[right].Weight {
			kid = left
		} else {
			kid = right
		}
		if key.Weight >= list[kid].Weight {
			goto Label_OVER
		}
		list[spot] = list[kid]
		spot, left, right = kid, kid*2+1, kid*2+2
	}
	if right == len(list) && key.Weight < list[left].Weight {
		list[spot], list[left] = list[left], key
		return
	}
Label_OVER:
	list[spot] = key
}

func insertSort(list []graph.Edge) {
	for i := 1; i < len(list); i++ {
		var key = list[i]
		var start, end = 0, i
		for start < end {
			var mid = (start + end) / 2
			if key.Weight < list[mid].Weight {
				end = mid
			} else { //找第一个大于key的位置
				start = mid + 1
			}
		} //不会越界
		for j := i; j > start; j-- {
			list[j] = list[j-1]
		}
		list[start] = key
	}
}
