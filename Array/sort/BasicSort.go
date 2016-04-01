package sort

// 冒泡排序，最原始的排序方法，具有稳定性。
// 比较操作是O(N^2)，挪移是O(N^2)，性能差。
func BubleSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := len(list) - 1; j > i; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
}

// 选择排序，不具有稳定性。
// 比较操作是O(N^2)，挪移是O(N)，综合性能不如InsertSort。
func SelectSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		var pos = i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[pos] {
				pos = j
			}
		}
		list[pos], list[i] = list[i], list[pos]
	}
}

// 插入排序，具有稳定性。
// 比较操作是O(NlogN)，挪移是O(N^2)，综合性能优于SelectSort。
func InsertSort(list []int) {
	for i := 1; i < len(list); i++ {
		var key = list[i]
		var a, b = 0, i
		for a < b {
			//var m = (a + b) / 2
			var m = a + (b-a)/2
			if key < list[m] {
				b = m
			} else { //找第一个大于key的位置
				a = m + 1
			}
		} //不会越界
		for j := i; j > a; j-- {
			list[j] = list[j-1]
		}
		list[a] = key
	}
}

func SimpleSortX(list []int) {
	if len(list) < 2 {
		return
	}
	var best = 0
	for i := 1; i < len(list); i++ {
		if list[i] < list[best] {
			best = i
		}
	}
	list[0], list[best] = list[best], list[0]
	for i := 1; i < len(list); i++ {
		var key, pos = list[i], i
		for list[pos-1] > key {
			list[pos] = list[pos-1]
			pos--
		}
		list[pos] = key
	}
}
func SimpleSort(list []int) {
	if len(list) < 2 {
		return
	}
	for i := 1; i < len(list); i++ {
		var key = list[i]
		if key < list[0] {
			for j := i; j > 0; j-- {
				list[j] = list[j-1]
			}
			list[0] = key
		} else {
			var pos = i
			for list[pos-1] > key {
				list[pos] = list[pos-1]
				pos--
			}
			list[pos] = key
		}
	}
}

const LOWER_BOUND = 16
const LOWER_BOUND_Y = 20
