package diysort

//内省排序是QuickSort的变种，混合排序
func Introsort(list []int) {
	var size = len(list)
	if size < 2 {
		return
	}
	var tasks stack
	tasks.push(0, size)

	var level int
	for level = 0; size != 0; level++ {
		size /= 2
	}

	for !tasks.isEmpty() {
		var start, end = tasks.pop()
		if end-start < 7 { //内建SelectSort
			for i := start; i < end-1; i++ {
				var pos = i
				for j := i + 1; j < end; j++ {
					if list[j] < list[pos] {
						pos = j
					}
				}
				list[pos], list[i] = list[i], list[pos]
			}
			continue
		}
		if tasks.size() == level { //约束最坏情况
			HeapSort(list[start:end])
			continue
		}

		//三点取中法，最后保证seed落入中间，每轮至少解决此一处
		var seed = list[start]
		var mid = (start + end) / 2
		if list[start] > list[mid] {
			if list[mid] > list[end-1] {
				seed, list[mid] = list[mid], list[start]
			} else { //c >= b
				if list[start] > list[end-1] {
					seed, list[end-1] = list[end-1], list[start]
				}
			}
		} else { //b >= a
			if list[end-1] > list[start] {
				if list[mid] > list[end-1] {
					seed, list[end-1] = list[end-1], list[start]
				} else {
					seed, list[mid] = list[mid], list[start]
				}
			}
		}

		var left, right = start + 1, end - 1
		for left < right {
			if list[right] <= seed {
				if list[left] > seed {
					list[left], list[right] = list[right], list[left]
					right--
				}
				left++
			} else {
				right--
			}
		}
		if list[right] > seed {
			right--
		} //这里不能用left，left可能越界
		list[start], list[right] = list[right], seed
		//每轮保证至少解决一个，否则最坏情况可能是死循环
		tasks.push(right+1, end)
		tasks.push(start, right)
	}
}