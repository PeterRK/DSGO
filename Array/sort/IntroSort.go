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
		var line = partition(list)
		doIntroSort(list[:line], life-1)
		doIntroSort(list[line:], life-1)
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
	if len(list) < LOWER_BOUND_Y {
		SimpleSort(list)
	} else if life == 0 {
		HeapSort(list)
	} else {
		var fst, snd = triPartition(list)
		if list[fst] != list[snd] {
			doIntroSortY(list[fst+1:snd], life-1)
		}
		doIntroSortY(list[:fst], life-1)
		doIntroSortY(list[snd+1:], life-1)
	}
}
