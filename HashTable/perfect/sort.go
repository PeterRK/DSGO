package perfect

func sort(list []memo) {
	var lv = uint(1)
	for sz := len(list); sz != 0; sz /= 2 {
		lv++
	}
	doIntroSort(list, lv*2)
}
func doIntroSort(list []memo, life uint) {
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

func partition(list []memo) int {
	var pivot = len(list[len(list)/2].lst)
	var a, b = 0, len(list) - 1
	for {
		for len(list[a].lst) < pivot {
			a++
		}
		for len(list[b].lst) > pivot {
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

func heapSort(list []memo) {
	for idx := len(list)/2 - 1; idx >= 0; idx-- {
		down(list, idx)
	}
	for sz := len(list) - 1; sz > 0; sz-- {
		list[0], list[sz] = list[sz], list[0]
		down(list[:sz], 0)
	}
}
func down(list []memo, root int) {
	var key = list[root]
	var kid, last = root*2 + 1, len(list) - 1
	for kid < last {
		if len(list[kid+1].lst) > len(list[kid].lst) {
			kid++
		}
		if len(key.lst) >= len(list[kid].lst) {
			break
		}
		list[root] = list[kid]
		root, kid = kid, kid*2+1
	}
	if kid == last && len(key.lst) < len(list[kid].lst) {
		list[root], root = list[kid], kid
	}
	list[root] = key
}

func simpleSort(list []memo) {
	if len(list) < 2 {
		return
	}
	var best = 0
	for i := 1; i < len(list); i++ {
		if len(list[i].lst) < len(list[best].lst) {
			best = i
		}
	}
	list[0], list[best] = list[best], list[0]
	for i := 1; i < len(list); i++ {
		var key, pos = list[i], i
		for len(list[pos-1].lst) > len(key.lst) {
			list[pos] = list[pos-1]
			pos--
		}
		list[pos] = key
	}
}
