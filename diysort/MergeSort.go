package diysort

//复杂度为O(NlogN) & O(N)，具有稳定性
//比较操作是O(NlogN)，常数小于HeapSort
//挪移操作是O(NlogN)，常数与HeapSort相当
func MergeSort(list []int) {
	var size = len(list)
	if size < sz_limit {
		InsertSort(list)
	} else {
		var shadow = make([]int, size)
		for i := 0; i < size; i++ {
			shadow[i] = list[i]
		}
		doMergeSort(shadow, list)
	}
}

func doMergeSort(in []int, out []int) {
	var size = len(in)
	if size < sz_limit { //内建InsertSort
		for i := 1; i < size; i++ {
			var left, right = 0, i
			var key = in[i]
			for left < right {
				var mid = (left + right) / 2
				if key < out[mid] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			for j := i; j > left; j-- {
				out[j] = out[j-1]
			}
			out[left] = key
		}
		return
	}
	var half = size / 2
	doMergeSort(out[:half], in[:half])
	doMergeSort(out[half:], in[half:])
	var pos, i, j = 0, 0, half
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
