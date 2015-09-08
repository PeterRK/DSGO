package array

//在由小到大的序列中寻找第一个大于key的位置
func SearchAfter(list []int, key int) int {
	var start, end = 0, len(list)
	for start < end {
		var mid = (start + end) / 2
		if key < list[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

//寻找key的位置(未必是第一个)，没有返回-1
func Search(list []int, key int) int {
	for a, b := 0, len(list); a < b; {
		var m = (a + b) / 2
		switch {
		case key > list[m]:
			a = m + 1
		case key < list[m]:
			b = m
		default:
			return m
		}
	}
	return -1
}

//在由小到大的序列中寻找第一个大于或等于key的位置
func SearchFirst(list []int, key int) int {
	var start, end = 0, len(list)
	for start < end {
		var mid = (start + end) / 2
		if key > list[mid] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

//在由小到大的序列中寻找最后一个小于或等于key的位置
func SearchLast(list []int, key int) int {
	var start, end = len(list) - 1, -1
	for start > end {
		//"(start + end + 2) / 2"也可以，但"(start+end)/2 + 1"不行
		var mid = (start + end + 1) / 2
		if key < list[mid] {
			start = mid - 1
		} else {
			end = mid
		}
	}
	return start
}

//在由小到大的序列中寻找目标，找打返回索引范围，没有则返回false
func SearchRange(list []int, key int) (first int, last int, ok bool) {
	last = SearchLast(list, key)
	if last == -1 || list[last] != key {
		return -1, -1, false
	}
	first = SearchFirst(list, key)
	return first, last, true
}

func Insert(list []int, key int) []int {
	var spot = SearchAfter(list, key)
	list = append(list, 0)
	for i := len(list) - 1; i > spot; i-- {
		list[i] = list[i-1] //后移
	}
	list[spot] = key
	return list
}
