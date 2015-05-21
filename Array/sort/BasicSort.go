package sort

//冒泡排序，最原始的排序方法，具有稳定性。
//比较操作是O(N^2)，挪移是O(N^2)，性能差。
func BubleSort(list []int) {
	var size = len(list)
	for i := 0; i < size-1; i++ {
		for j := size - 1; j > i; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
}

//选择排序，不具有稳定性。
//比较操作是O(N^2)，挪移是O(N)，综合性能不如InsertSort。
func SelectSort(list []int) {
	var size = len(list)
	for i := 0; i < size-1; i++ {
		var pos = i
		for j := i + 1; j < size; j++ {
			if list[j] < list[pos] {
				pos = j
			}
		}
		list[pos], list[i] = list[i], list[pos]
	}
}

//插入排序，具有稳定性。
//比较操作是O(NlogN)，挪移是O(N^2)，综合性能优于SelectSort。
func InsertSort(list []int) {
	var size = len(list)
	for i := 1; i < size; i++ {
		var left, right = 0, i
		var key = list[i]
		for left < right {
			var mid = (left + right) / 2
			if key < list[mid] {
				right = mid
			} else { //找第一个大于key的位置
				left = mid + 1
			}
		} //不会越界
		for j := i; j > left; j-- {
			list[j] = list[j-1]
		}
		list[left] = key
	}
}
