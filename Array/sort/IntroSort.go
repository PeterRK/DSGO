package sort

// 内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
// 复杂度为O(NlogN) & O(logN)。
// 主要限制了QuickSort的最坏情况，适合递归实现(没有爆栈风险)。
func IntroSort(list []int) {
	var life = log2ceil(uint(len(list))) * 2
	doIntroSort(list, life)
}
func doIntroSort(list []int, life uint) {
	if len(list) < LOWER_BOUND {
		SimpleSort(list)
	} else if life == 0 {
		HeapSort(list)
	} else {
		var m = partition(list)
		doIntroSort(list[:m], life-1)
		doIntroSort(list[m:], life-1)
	}
}

func log2ceil(num uint) uint {
	var ceil uint
	for ceil = 1; num != 0; ceil++ {
		num /= 2
	}
	return ceil
}

// 三分内省排序
func IntroSortY(list []int) {
	var life = log2ceil(uint(len(list))) * 3 / 2
	doIntroSortY(list, life)
}
func doIntroSortY(list []int, life uint) {
	for len(list) > LOWER_BOUND_Y {
		if life == 0 {
			HeapSort(list)
			return
		}
		life--
		var fst, snd = triPartition(list)
		doIntroSortY(list[:fst], life)
		doIntroSortY(list[snd+1:], life)
		if list[fst] == list[snd] {
			return
		}
		list = list[fst+1 : snd]
	}
	SimpleSort(list)
}
