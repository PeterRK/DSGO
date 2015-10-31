package sort

//内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
//复杂度为O(NlogN) & O(logN)。
//主要限制了QuickSort的最坏情况，适合递归实现(没有爆栈风险)。
func IntroSort(list []int) {
	var life = uint(24)
	for sz := len(list); sz != 0; sz /= 2 {
		life++
	}
	doIntroSort(list, life)
}
func doIntroSort(list []int, life uint) {
	if life == 0 {
		HeapSort(list)
	} else if len(list) < LOWER_BOUND {
		InsertSort(list)
	} else {
		var line = partition(list)
		doIntroSort(list[:line], life-1)
		doIntroSort(list[line:], life-1)
	}
}

func IntroSortY(list []int) {
	var life = uint(16)
	for sz := len(list); sz != 0; sz /= 2 {
		life++
	}
	doIntroSortY(list, life)
}
func doIntroSortY(list []int, life uint) {
	if life == 0 {
		HeapSort(list)
	} else if len(list) < LOWER_BOUND {
		InsertSort(list)
	} else {
		var start, end = triPartition(list)
		if list[start] != list[end-1] {
			doIntroSortY(list[start+1:end-1], life-1)
		}
		doIntroSortY(list[:start], life-1)
		doIntroSortY(list[end:], life-1)
	}
}
