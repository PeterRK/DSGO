package sort

type Unit struct {
	pad [0]uint32
	val int
}

// 冒泡排序，最原始的排序方法，具有稳定性。
// 比较操作是O(N^2)，挪移是O(N^2)，性能差。
func BubleSort(list []Unit) {
	for i := 0; i < len(list)-1; i++ {
		for j := len(list) - 1; j > i; j-- {
			if list[j].val < list[j-1].val {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
}

// 选择排序，不具有稳定性。
// 比较操作是O(N^2)，挪移是O(N)，综合性能不如InsertSort。
func SelectSort(list []Unit) {
	for i := 0; i < len(list)-1; i++ {
		pos := i
		for j := i + 1; j < len(list); j++ {
			if list[j].val < list[pos].val {
				pos = j
			}
		}
		list[pos], list[i] = list[i], list[pos]
	}
}

// 插入排序，具有稳定性。
// 比较操作是O(NlogN)，挪移是O(N^2)，综合性能优于SelectSort。
func InsertSort(list []Unit) {
	for i := 1; i < len(list); i++ {
		key := list[i]
		a, b := 0, i
		for a < b {
			m := a + (b-a)/2 //(a+b)/2
			if key.val < list[m].val {
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

func SimpleSortX(list []Unit) {
	if len(list) < 2 {
		return
	}
	best := 0
	for i := 1; i < len(list); i++ {
		if list[i].val < list[best].val {
			best = i
		}
	}
	list[0], list[best] = list[best], list[0]
	for i := 1; i < len(list); i++ {
		key, pos := list[i], i
		for list[pos-1].val > key.val {
			list[pos] = list[pos-1]
			pos--
		}
		list[pos] = key
	}
}
func SimpleSort(list []Unit) {
	if len(list) < 2 {
		return
	}
	for i := 1; i < len(list); i++ {
		key := list[i]
		if key.val < list[0].val {
			for j := i; j > 0; j-- {
				list[j] = list[j-1]
			}
			list[0] = key
		} else {
			pos := i
			for list[pos-1].val > key.val {
				list[pos] = list[pos-1]
				pos--
			}
			list[pos] = key
		}
	}
}

const LOWER_BOUND = 16
const LOWER_BOUND_Y = 20
