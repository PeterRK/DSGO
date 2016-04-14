package sort

// 三分快速排序，比二分版本略为复杂
func QuickSortY(list []int) {
	for len(list) > LOWER_BOUND_Y {
		var fst, snd = triPartition(list)
		QuickSortY(list[:fst])
		QuickSortY(list[snd+1:])
		if list[fst] == list[snd] {
			return
		}
		list = list[fst+1 : snd]
	}
	SimpleSort(list)
}

// 返回两个分界元素的位置
func triPartition(list []int) (fst, snd int) {
	var sz = len(list)
	var m, s = sz / 2, sz / 8
	var a, b = m - s, m + s
	if list[a] > list[b] {
		a, b = b, a
	}
	var pivot1, pivot2 = list[a], list[b]
	list[a], list[b] = list[0], list[sz-1]

	a, b = 1, sz-2
	for a <= b && list[a] < pivot1 {
		a++
	}
	for k := a; k <= b; k++ {
		if list[k] > pivot2 {
			for k < b && list[b] > pivot2 {
				b--
			}
			list[k], list[b] = list[b], list[k]
			b--
		}
		if list[k] < pivot1 {
			list[k], list[a] = list[a], list[k]
			a++
		}
	}

	list[0], list[a-1] = list[a-1], pivot1
	list[sz-1], list[b+1] = list[b+1], pivot2
	return a - 1, b + 1
}

/*
func triPartition(list []int) (fst, snd int) {
	var sz = len(list)
	var m, s = sz / 2, sz / 8
	moveMedianTwo(list, m-s, 0, sz-1, m+s)
	var pivot1, pivot2 = list[0], list[sz-1]

	var a, b = 1, sz - 2
	for list[a] < pivot1 {
		a++
	}
	for k := a; k <= b; k++ {
		if list[k] > pivot2 {
			for list[b] > pivot2 {
				b--
			}
			if k > b {
				break
			}
			list[k], list[b] = list[b], list[k]
			b--
		}
		if list[k] < pivot1 {
			list[k], list[a] = list[a], list[k]
			a++
		}
	}

	list[0], list[a-1] = list[a-1], pivot1
	list[sz-1], list[b+1] = list[b+1], pivot2
	return a - 1, b + 1
}

func moveMedianTwo(list []int, a, b, c, d int) {
	if list[a] > list[b] { //保证b非最小
		list[a], list[b] = list[b], list[a]
	}
	if list[c] < list[d] { //保证c非最大
		list[c], list[d] = list[d], list[c]
	}
	if list[b] > list[c] { //保证b<=c
		list[b], list[c] = list[c], list[b]
	}
}
*/

/*
func triPartition(list []int) (fst, snd int) {
	var sz = len(list)

	var pivot1, pivot2 = list[0], list[sz-1]
	if pivot1 > pivot2 {
		pivot1, pivot2 = pivot2, pivot1
		list[0], list[sz-1] = pivot1, pivot2
	}

	var ax, bx = 1, sz - 2
	var a, b = ax, bx
	for {
		for ; a <= b; a++ {
			if list[a] > pivot2 {
				break
			} else if list[a] <= pivot1 {
				if a != ax {
					list[ax], list[a] = list[a], list[ax]
				}
				ax++
			}
		}
		for ; a <= b; b-- {
			if list[b] < pivot1 {
				break
			} else if list[b] >= pivot2 {
				if b != bx {
					list[bx], list[b] = list[b], list[bx]
				}
				bx--
			}
		}
		if a >= b {
			break
		}
		if a == ax || b == bx {
			list[a], list[b] = list[b], list[a]
		} else {
			list[a], list[bx] = list[bx], list[a]
			list[b], list[ax] = list[ax], list[b]
			ax++
			bx--
			a++
			b--
		}
	}
	list[0], list[ax-1] = list[ax-1], pivot1
	list[sz-1], list[bx+1] = list[bx+1], pivot2
	return ax - 1, bx + 1
}
*/
