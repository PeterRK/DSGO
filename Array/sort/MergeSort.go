package sort

// 归并排序，改进的插入排序，具有稳定性。
// 复杂度为O(NlogN) & O(N)。
// 其中比较操作是O(NlogN)，但常数小于HeapSort；挪移操作是O(NlogN)，常数与HeapSort相当。
func MergeSort(list []int) {
	if len(list) < LOWER_BOUND {
		InsertSort(list)
	} else {
		var shadow = make([]int, len(list))
		for i := 0; i < len(list); i++ {
			shadow[i] = list[i]
		}
		doMergeSort(shadow, list)
	}
}
func doMergeSort(in []int, out []int) {
	if len(in) < LOWER_BOUND {
		if len(in) == 0 {
			return
		}
		out[0] = in[0]
		for i := 1; i < len(in); i++ {
			var key = in[i]
			if key < out[0] {
				for j := i; j > 0; j-- {
					out[j] = out[j-1]
				}
				out[0] = key
			} else {
				var pos = i
				for out[pos-1] > key {
					out[pos] = out[pos-1]
					pos--
				}
				out[pos] = key
			}
		}
	} else {
		var half = len(in) / 2
		doMergeSort(out[:half], in[:half])
		doMergeSort(out[half:], in[half:])

		var pos, i, j, size = 0, 0, half, len(in)
		for ; i < half && j < size; pos++ {
			if in[i] > in[j] {
				out[pos] = in[j]
				j++
			} else {
				out[pos] = in[i]
				i++
			}
		}
		for ; i < half; pos++ {
			out[pos] = in[i]
			i++
		}
		for ; j < size; pos++ {
			out[pos] = in[j]
			j++
		}
	}
}

/*
func merge(in1 []int, in2 []int, out []int) {
	var i, j, k = 0, 0, 0
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
*/
