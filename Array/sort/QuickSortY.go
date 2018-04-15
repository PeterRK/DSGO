package sort

// 三分快速排序，比二分版本略为复杂
func QuickSortY(list []Unit) {
	for len(list) > LOWER_BOUND_Y {
		fst, snd := triPartition(list)
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
func triPartition(list []Unit) (fst, snd int) {
	sz := len(list)
	m, s := sz/2, sz/8
	a, b := m-s, m+s
	if list[a].val > list[b].val {
		a, b = b, a
	}
	pivot1, pivot2 := list[a], list[b]
	list[a], list[b] = list[0], list[sz-1]

	a, b = 1, sz-2
	for a <= b && list[a].val < pivot1.val {
		a++
	}
	for k := a; k <= b; k++ {
		if list[k].val > pivot2.val {
			for k < b && list[b].val > pivot2.val {
				b--
			}
			list[k], list[b] = list[b], list[k]
			b--
		}
		if list[k].val < pivot1.val {
			list[k], list[a] = list[a], list[k]
			a++
		}
	}

	list[0], list[a-1] = list[a-1], pivot1
	list[sz-1], list[b+1] = list[b+1], pivot2
	return a - 1, b + 1
}

/*
func triPartition(list []Unit) (fst, snd int) {
	sz := len(list)
	m, s := sz/2, sz/8
	moveMedianTwo(list, m-s, 0, sz-1, m+s)
	pivot1, pivot2 := list[0], list[sz-1]

	a, b := 1, sz-2
	for list[a].val < pivot1.val {
		a++
	}
	for k := a; k <= b; k++ {
		if list[k].val > pivot2.val {
			for list[b].val > pivot2.val {
				b--
			}
			if k > b {
				break
			}
			list[k], list[b] = list[b], list[k]
			b--
		}
		if list[k].val < pivot1.val {
			list[k], list[a] = list[a], list[k]
			a++
		}
	}

	list[0], list[a-1] = list[a-1], pivot1
	list[sz-1], list[b+1] = list[b+1], pivot2
	return a - 1, b + 1
}

func moveMedianTwo(list []Unit, a, b, c, d int) {
	if list[a].val > list[b].val { //保证b非最小
		list[a], list[b] = list[b], list[a]
	}
	if list[c].val < list[d].val { //保证c非最大
		list[c], list[d] = list[d], list[c]
	}
	if list[b].val > list[c].val { //保证b<=c
		list[b], list[c] = list[c], list[b]
	}
}
*/

/*
func triPartition(list []Unit) (fst, snd int) {
	sz := len(list)

	pivot1, pivot2 := list[0], list[sz-1]
	if pivot1.val > pivot2.val {
		pivot1, pivot2 = pivot2, pivot1
		list[0], list[sz-1] = pivot1, pivot2
	}

	ax, bx := 1, sz-2
	a, b := ax, bx
	for {
		for ; a <= b; a++ {
			if list[a].val > pivot2.val {
				break
			} else if list[a].val <= pivot1.val {
				if a != ax {
					list[ax], list[a] = list[a], list[ax]
				}
				ax++
			}
		}
		for ; a <= b; b-- {
			if list[b].val < pivot1.val {
				break
			} else if list[b].val >= pivot2.val {
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
