package sort

import (
	"constraints"
)

// 归并排序，改进的插入排序，具有稳定性。
// 复杂度为O(NlogN) & O(N)。
// 其中比较操作是O(NlogN)，但常数小于HeapSort；挪移操作是O(NlogN)，常数与HeapSort相当。
func MergeSort[T constraints.Ordered](list []T) {
	if len(list) < lowerBound {
		SimpleSort(list)
	} else {
		temp := make([]T, len(list))
		copy(temp, list)
		mergeSort(temp, list)
	}
}

func mergeSort[T constraints.Ordered](a, b []T) {
	if size := len(a); size < lowerBound {
		if size < 2 {
			return
		}
		b[0] = a[0]
		for i := 1; i < len(a); i++ {
			if key := a[i]; key < b[0] {
				for j := i; j > 0; j-- {
					b[j] = b[j-1]
				}
				b[0] = key
			} else {
				pos := i
				for ; b[pos-1] > key; pos-- {
					b[pos] = b[pos-1]
				}
				b[pos] = key
			}
		}
	} else {
		half := len(a)/2
		mergeSort(b[:half], a[:half])
		mergeSort(b[half:], a[half:])

		pos, i, j := 0, 0, half
		for ; i < half && j < size; pos++ {
			if a[i] > a[j] {
				b[pos] = a[j]
				j++
			} else {
				b[pos] = a[i]
				i++
			}
		}
		for ; i < half; pos++ {
			b[pos] = a[i]
			i++
		}
		for ; j < size; pos++ {
			b[pos] = a[j]
			j++
		}
	}
}

func merge[T constraints.Ordered](in1, in2, out []T) {
	i, j, k := 0, 0, 0
	for ; i < len(in1) && j < len(in2); k++ {
		if in1[i] <= in2[j] {
			out[k] = in1[i]
			i++
		} else {
			out[k] = in2[j]
			j++
		}
	}
	for ; i < len(in1); k++ {
		out[k] = in1[i]
		i++
	}
	for ; j < len(in2); k++ {
		out[k] = in2[j]
		j++
	}
}
