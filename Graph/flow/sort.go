package flow

import (
	"Graph/graph"
)

func sort(list []graph.Path) {
	var life = uint(12)
	for sz := len(list); sz != 0; sz /= 2 {
		life++
	}
	doIntroSort(list, life)
}
func doIntroSort(list []graph.Path, life uint) {
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

func partition(list []graph.Path) (start int, end int) {
	var size = len(list)
	var m1, m2 = len(list)/2 - 1, len(list) / 2
	if list[m1].Next > list[m2].Next {
		m1, m2 = m2, m1
	}
	var pivot1, pivot2 = list[m1], list[m2]
	list[m1], list[m2] = list[0], list[size-1]

	var left, right = 1, size - 2
	for k := left; k <= right; k++ {
		if list[k].Next > pivot2.Next {
			for k < right && list[right].Next > pivot2.Next {
				right--
			}
			list[k], list[right] = list[right], list[k]
			right--
		}
		if list[k].Next < pivot1.Next {
			list[k], list[left] = list[left], list[k]
			left++
		}
	}

	list[0], list[left-1] = list[left-1], pivot1
	list[size-1], list[right+1] = list[right+1], pivot2
	return left - 1, right + 2
}

func heapSort(list []graph.Path) {
	for idx := len(list)/2 - 1; idx >= 0; idx-- {
		down(list, idx)
	}
	for sz := len(list) - 1; sz > 0; sz-- {
		list[0], list[sz] = list[sz], list[0]
		down(list[:sz], 0)
	}
}
func down(list []graph.Path, spot int) {
	var key = list[spot]
	var left, right = spot*2 + 1, spot*2 + 2
	for right < len(list) {
		var kid int
		if list[left].Next > list[right].Next {
			kid = left
		} else {
			kid = right
		}
		if key.Next >= list[kid].Next {
			goto Label_OVER
		}
		list[spot] = list[kid]
		spot, left, right = kid, kid*2+1, kid*2+2
	}
	if right == len(list) && key.Next < list[left].Next {
		list[spot], list[left] = list[left], key
		return
	}
Label_OVER:
	list[spot] = key
}

func insertSort(list []graph.Path) {
	for i := 1; i < len(list); i++ {
		var key = list[i]
		var a, b = 0, i
		for a < b {
			var m = (a + b) / 2
			if key.Next < list[m].Next {
				b = m
			} else { //找第一个大于key的位置
				a = m + 1
			}
		} //不会越界
		for j := i; j > a; j-- {
			list[j] = list[j-1]
		}
		list[a] = key
	}
}
