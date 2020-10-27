package sort

func indexSimpleSort(list []Unit, index []uint32) {
	if len(index) < 2 {
		return
	}
	for i := 1; i < len(index); i++ {
		key := index[i]
		if list[key].val < list[index[0]].val ||
			(list[key].val == list[index[0]].val && key < index[0]) {
			for j := i; j > 0; j-- {
				index[j] = index[j-1]
			}
			index[0] = key
		} else {
			pos := i
			for list[index[pos-1]].val > list[key].val ||
				(list[index[pos-1]].val == list[key].val && index[pos-1] > key) {
				index[pos] = index[pos-1]
				pos--
			}
			index[pos] = key
		}
	}
}

func indexHeapSort(list []Unit, index []uint32) {
	for idx := len(index)/2 - 1; idx >= 0; idx-- {
		indexDown(list, index, idx)
	}
	for sz := len(index) - 1; sz > 0; sz-- {
		index[0], index[sz] = index[sz], index[0]
		indexDown(list, index[:sz], 0)
	}
}

func indexDown(list []Unit, index []uint32, pos int) {
	key := index[pos]
	kid, last := pos*2+1, len(index)-1
	for kid < last {
		if list[index[kid+1]].val > list[index[kid]].val ||
			(list[index[kid+1]].val == list[index[kid]].val && index[kid+1] > index[kid]) {
			kid++
		}
		if list[key].val > list[index[kid]].val ||
			(list[key].val == list[index[kid]].val && key > index[kid]) {
			break
		}
		index[pos] = index[kid]
		pos, kid = kid, kid*2+1
	}
	if kid == last && (list[key].val < list[index[kid]].val ||
		(list[key].val == list[index[kid]].val && key < index[kid])) {
		index[pos], pos = index[kid], kid
	}
	index[pos] = key
}

func indexTriPartition(list []Unit, index []uint32) (fst, snd int) {
	sz := len(index)
	m, s := sz/2, sz/8
	a, b := m-s, m+s
	if list[index[a]].val > list[index[b]].val ||
		(list[index[a]].val == list[index[b]].val && index[a] > index[b]) {
		a, b = b, a
	}
	pivot1, pivot2 := index[a], index[b]
	index[a], index[b] = index[0], index[sz-1]

	a, b = 1, sz-2
	for a <= b && (list[index[a]].val < list[pivot1].val ||
		(list[index[a]].val == list[pivot1].val && index[a] < pivot1)) {
		a++
	}
	for k := a; k <= b; k++ {
		if list[index[k]].val > list[pivot2].val ||
			(list[index[k]].val == list[pivot2].val && index[k] > pivot2) {
			for k < b && (list[index[b]].val > list[pivot2].val ||
				(list[index[b]].val == list[pivot2].val && index[b] > pivot2)) {
				b--
			}
			index[k], index[b] = index[b], index[k]
			b--
		}
		if list[index[k]].val < list[pivot1].val ||
			(list[index[k]].val == list[pivot1].val && index[k] < pivot1) {
			index[k], index[a] = index[a], index[k]
			a++
		}
	}

	index[0], index[a-1] = index[a-1], pivot1
	index[sz-1], index[b+1] = index[b+1], pivot2
	return a - 1, b + 1
}

func reorder(list []Unit, index []uint32) {
	for i, size := uint32(0), uint32(len(index)); i < size; i++ {
		if index[i] == i {
			continue
		}
		tmp := list[i]
		j, k := i, index[i]
		list[j], index[j] = list[k], j
		for index[k] != i {
			j, k = k, index[k]
			list[j], index[j] = list[k], j
		}
		list[k], index[k] = tmp, k
	}
}

func IndexIntroSort(list []Unit) {
	life := log2ceil(uint(len(list))) * 3 / 2
	index := make([]uint32, len(list))
	for i := 0; i < len(list); i++ {
		index[i] = uint32(i)
	}
	indexIntroSort(list, index, life)

	//根据index对list进行最终处理
	reorder(list, index)
}
func indexIntroSort(list []Unit, index []uint32, life uint) {
	for len(index) > LOWER_BOUND_Y {
		if life == 0 {
			indexHeapSort(list, index)
			return
		}
		life--
		fst, snd := indexTriPartition(list, index)
		indexIntroSort(list, index[:fst], life)
		indexIntroSort(list, index[snd+1:], life)
		index = index[fst+1 : snd]
	}
	indexSimpleSort(list, index)
}
