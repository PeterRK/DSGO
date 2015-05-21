package array

//在由小到大的序列中寻找目标，找打返回索引，没有则返回-1
func BinarySearch(list []int, key int) int {
	var start, end = 0, len(list)
	for start < end {
		var mid = (start + end) / 2
		if key > list[mid] {
			start = mid + 1
		} else if key < list[mid] {
			end = mid
		} else {
			return mid
		}
	}
	return -1
}

//在由小到大的序列中寻找目标，找打返回索引范围，没有则返回(-1,-1)
func SearchRange(list []int, key int) (first int, last int) {
	if len(list) == 0 {
		return -1, -1
	}
	var left, right = 0, len(list) - 1
	for left < right {
		var mid = (left + right) / 2 //偏前
		if key > list[mid] {
			left = mid + 1
		} else { //等于归前
			right = mid
		}
	}
	if list[left] != key {
		return -1, -1
	}
	first = left

	left, right = 0, len(list)-1
	for left < right {
		var mid = (left + right + 1) / 2 //偏后
		if key < list[mid] {
			right = mid - 1
		} else { //等于归后
			left = mid
		}
	}
	return first, right
}

//在由小到大的序列中寻找第一个大于key的位置
func SearchInsertPoint(list []int, key int) int {
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

//保持顺序
func Delete(list []int, idx int) []int {
	var last = len(list) - 1
	if idx < 0 || idx > last {
		return list
	}
	for i := idx; i < last; i++ {
		list[i] = list[i+1] //前移
	}
	return list[:last]
}

//不保持顺序
func Erease(list []int, idx int) []int {
	var last = len(list) - 1
	if idx < 0 || idx > last {
		return list
	}
	list[idx] = list[last]
	return list[:last]
}
