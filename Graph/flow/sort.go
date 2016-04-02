package flow

import (
	"DSGO/Graph/graph"
)

func sort(list []graph.Path) {
	var lv = uint(1)
	for sz := len(list); sz != 0; sz /= 2 {
		lv++
	}
	doIntroSort(list, lv*2)
}
func doIntroSort(list []graph.Path, life uint) {
	if len(list) < 16 {
		simpleSort(list)
	} else if life == 0 {
		heapSort(list)
	} else {
		var m = partition(list)
		doIntroSort(list[:m], life-1)
		doIntroSort(list[m:], life-1)
	}
}

func partition(list []graph.Path) int {
	var pivot = list[len(list)/2].Next
	var a, b = 0, len(list) - 1
	for {
		for list[a].Next < pivot {
			a++
		}
		for list[b].Next > pivot {
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

func heapSort(list []graph.Path) {
	for idx := len(list)/2 - 1; idx >= 0; idx-- {
		down(list, idx)
	}
	for sz := len(list) - 1; sz > 0; sz-- {
		list[0], list[sz] = list[sz], list[0]
		down(list[:sz], 0)
	}
}
func down(list []graph.Path, root int) {
	var key = list[root]
	var kid, last = root*2 + 1, len(list) - 1
	for kid < last {
		if list[kid+1].Next > list[kid].Next {
			kid++
		}
		if key.Next >= list[kid].Next {
			break
		}
		list[root] = list[kid]
		root, kid = kid, kid*2+1
	}
	if kid == last && key.Next < list[kid].Next {
		list[root], root = list[kid], kid
	}
	list[root] = key
}

func simpleSort(list []graph.Path) {
	if len(list) < 2 {
		return
	}
	var best = 0
	for i := 1; i < len(list); i++ {
		if list[i].Next < list[best].Next {
			best = i
		}
	}
	list[0], list[best] = list[best], list[0]
	for i := 1; i < len(list); i++ {
		var key, pos = list[i], i
		for list[pos-1].Next > key.Next {
			list[pos] = list[pos-1]
			pos--
		}
		list[pos] = key
	}
}
