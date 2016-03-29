package sort

// 内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
// 复杂度为O(NlogN) & O(logN)。
// 主要限制了QuickSort的最坏情况，适合递归实现(没有爆栈风险)。
func IntroSort(list []int) {
	var lv = uint(1)
	for sz := len(list); sz != 0; sz /= 2 {
		lv++
	}
	doIntroSort(list, lv*2)
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

// 三分内省排序
func IntroSortY(list []int) {
	var lv = uint(1)
	for sz := len(list); sz != 0; sz /= 2 {
		lv++
	}
	doIntroSortY(list, lv*3/2)
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
