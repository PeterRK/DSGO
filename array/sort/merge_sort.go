package sort

import (
	"DSGO/array"
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
		half := len(a) / 2
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

func SymMergeSort[T constraints.Ordered](list []T) {
	size := len(list)

	step := lowerBound
	a, b := 0, step
	for b <= size {
		SimpleSort(list[a:b])
		a = b
		b += step
	}
	SimpleSort(list[a:])

	for step < size {
		a, b = 0, step*2
		for b <= size {
			symmerge(list[a:b], step)
			a = b
			b += step * 2
		}
		if a+step < size {
			symmerge(list[a:], step)
		}
		step *= 2
	}
}

func symmerge[T constraints.Ordered](list []T, border int) {
	size := len(list)

	if border == 1 {
		curr := list[0]
		a, b := 1, size
		for a < b {
			m := int(uint(a+b)/2)
			if list[m] < curr {
				a = m + 1
			} else {
				b = m
			}
		}
		for i := 1; i < a; i++ {
			list[i-1] = list[i]
		}
		list[a-1] = curr
		return
	}

	if border == size-1 {
		curr := list[border]
		a, b := 0, border
		for a < b {
			m := int(uint(a+b)/2)
			if curr < list[m] {
				b = m
			} else {
				a = m + 1
			}
		}
		for i := border; i > a; i-- {
			list[i] = list[i-1]
		}
		list[a] = curr
		return
	}

	half := size / 2
	n := border + half
	a, b := 0, border
	if border > half {
		a, b = n-size, half
	}
	p := n - 1
	for a < b {
		m := int(uint(a+b)/2)
		if list[p-m] < list[m] {
			b = m
		} else {
			a = m + 1
		}
	}
	b = n - a
	if a < border && border < b {
		rotate(list[a:b], border-a)
	}
	if 0 < a && a < half {
		symmerge(list[:half], a)
	}
	if half < b && b < size {
		symmerge(list[half:], b-half)
	}
}

func rotate[T any](list []T, border int) {
	array.Reverse(list[:border])
	array.Reverse(list[border:])
	array.Reverse(list)
}
