package span

import (
	"Graph/graph"
)

func sort(list []graph.Edge) {
	var lv = uint(1)
	for sz := len(list); sz != 0; sz /= 2 {
		lv++
	}
	doIntroSort(list, lv*2)
}
func doIntroSort(list []graph.Edge, life uint) {
	if len(list) < 16 {
		simpleSort(list)
	} else if life == 0 {
		heapSort(list)
	} else {
		var line = partition(list)
		doIntroSort(list[:line], life-1)
		doIntroSort(list[line:], life-1)
	}
}

func partition(list []graph.Edge) int {
	var pivot = list[len(list)/2].Weight
	var a, b = 0, len(list) - 1
	for {
		for list[a].Weight < pivot {
			a++
		}
		for list[b].Weight > pivot {
			b--
		}
		if a >= b {
			break
		}
		list[a], list[b] = list[b], list[a]
		a++
		b--
	}
	return a
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

func simpleSort(list []graph.Edge) {
	if len(list) < 2 {
		return
	}
	var best = 0
	for i := 1; i < len(list); i++ {
		if list[i].Weight < list[best].Weight {
			best = i
		}
	}
	list[0], list[best] = list[best], list[0]
	for i := 1; i < len(list); i++ {
		var key, pos = list[i], i
		for list[pos-1].Weight > key.Weight {
			list[pos] = list[pos-1]
			pos--
		}
		list[pos] = key
	}
}
