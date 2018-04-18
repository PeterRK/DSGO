package sort

//增加判断条件，化不稳定排序为稳定排序
/*
func QuickSortS(list []Unit) {
	for i := 0; i < len(list); i++ {
		list[i].pad[0] = uint32(i)
	}
	doQuickSortS(list)
}

func doQuickSortS(list []Unit) {
	for len(list) > LOWER_BOUND_Y {
		fst, snd := triPartitionS(list)
		doQuickSortS(list[:fst])
		doQuickSortS(list[snd+1:])
		if list[fst] == list[snd] {
			return
		}
		list = list[fst+1 : snd]
	}
	SimpleSort(list)
}

func triPartitionS(list []Unit) (fst, snd int) {
	sz := len(list)
	m, s := sz/2, sz/8
	a, b := m-s, m+s
	if list[a].val > list[b].val ||
		(list[a].val == list[b].val && list[a].pad[0] > list[b].pad[0]) {
		a, b = b, a
	}
	pivot1, pivot2 := list[a], list[b]
	list[a], list[b] = list[0], list[sz-1]

	a, b = 1, sz-2
	for a <= b && (list[a].val < pivot1.val ||
		(list[a].val == pivot1.val && list[a].pad[0] < pivot1.pad[0])) {
		a++
	}
	for k := a; k <= b; k++ {
		if list[k].val > pivot2.val ||
			(list[k].val == pivot2.val && list[k].pad[0] > pivot2.pad[0]) {
			for k < b && (list[b].val > pivot2.val ||
				(list[b].val == pivot2.val && list[b].pad[0] > pivot2.pad[0])) {
				b--
			}
			list[k], list[b] = list[b], list[k]
			b--
		}
		if list[k].val < pivot1.val ||
			(list[k].val == pivot1.val && list[k].pad[0] < pivot1.pad[0]) {
			list[k], list[a] = list[a], list[k]
			a++
		}
	}

	list[0], list[a-1] = list[a-1], pivot1
	list[sz-1], list[b+1] = list[b+1], pivot2
	return a - 1, b + 1
}
*/
