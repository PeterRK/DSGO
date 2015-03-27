package diysort

//比较少于SelectSort，挪移多于SelectSort
func InsertSort(list []int) {
	var size = len(list)
	for i := 1; i < size; i++ {
		var left, right = 0, i
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
}
