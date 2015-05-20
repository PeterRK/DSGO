package sort

//内省排序是QuickSort的变种，混合排序
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
			var knot = part(list[start:end]) + start
			tasks.push(knot+1, end)
			tasks.push(start, knot)
		} //每轮保证至少解决一个，否则最坏情况可能是死循环
	}
}
