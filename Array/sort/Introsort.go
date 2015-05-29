package sort

//内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
//复杂度为O(NlogN) & O(logN)。
func IntroSort(list []int) {
	var size = len(list)
	var tasks stack
	tasks.push(0, size)

	var level int
	for level = 3; size != 0; level++ {
		size /= 2
	}
	for !tasks.isEmpty() {
		var start, end = tasks.pop()
		if end-start < sz_limit {
			InsertSort(list[start:end])
		} else if tasks.size() == level {
			HeapSort(list[start:end])
		} else {
			var knot = partition(list[start:end]) + start
			tasks.push(knot+1, end)
			tasks.push(start, knot)
		} //每轮保证至少解决一个，否则最坏情况可能是死循环
	}
}
