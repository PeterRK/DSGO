package sort

//内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
//复杂度为O(NlogN) & O(logN)。
//主要限制了QuickSort的最坏情况，适合递归实现(没有爆栈风险)。
func IntroSort(list []int) {
	var size = len(list)
	var life uint
	for life = 12; size != 0; life++ {
		size /= 2
	}
	doIntroSort(list, life)
}
func doIntroSort(list []int, life uint) {
	var size = len(list)
	if size < sz_limit {
		InsertSort(list)
	} else if life == 0 {
		HeapSort(list)
	} else {
		var knot = partition(list)
		doIntroSort(list[0:knot], life-1)
		doIntroSort(list[knot+1:size], life-1)
	}
}
