package sort

func fakeTriPatition(list []int) (m0, m1 int) {
	var m, last = len(list) / 2, len(list) - 1

	if list[0] < list[m] {
		list[0], list[m] = list[m], list[0]
	}
	if list[last] < list[0] {
		list[last], list[0] = list[0], list[last]
		if list[0] < list[m] {
			list[0], list[m] = list[m], list[0]
		}
	}
	var pivot = list[0]

	var a, b = 1, last - 1
	for a <= b && list[a] < pivot {
		a++
	}
	for {
		for a <= b && list[a] <= pivot {
			a++
		}
		for a <= b && list[b] > pivot {
			b--
		}
		if a > b {
			break
		}
		list[a], list[b] = list[b], list[a]
		a++
		b--
	}

	list[0], list[a-1] = list[a-1], list[0]
	return a - 1, b + 1
}

func QuickSortZ(list []int) {
	if len(list) < LOWER_BOUND {
		SimpleSort(list)
	} else {
		m0, m1 := fakeTriPatition(list)
		QuickSort(list[:m0])
		QuickSort(list[m1:])
	}
}
