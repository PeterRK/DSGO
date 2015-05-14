package sort

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
		if end-start < sz_limit { //内建InsertSort
			for i := start + 1; i < end; i++ {
				var left, right = start, i
				var key = list[i]
				for left < right {
					var mid = (left + right) / 2
					if key < list[mid] {
						right = mid
					} else {
						left = mid + 1
					}
				}
				for j := i; j > left; j-- {
					list[j] = list[j-1]
				}
				list[left] = key
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
		for { //注意对称性
			for list[left] < seed {
				left++
			}
			for list[right] > seed {
				right--
			}
			if left >= right {
				break
			}
			list[left], list[right] = list[right], list[left]
			left++
			right--
		}
		list[start], list[right] = list[right], seed
		//每轮保证至少解决一个，否则最坏情况可能是死循环
		tasks.push(right+1, end)
		tasks.push(start, right)
	}
}
