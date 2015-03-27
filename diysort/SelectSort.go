package diysort

//比较多于InsertSort，挪移少于InsertSort
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
