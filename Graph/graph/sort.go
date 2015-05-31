package graph

func Sort(list []Edge) {
	var size = len(list)
	var tasks stack
	tasks.push(0, size)

	var level int
	for level = 3; size != 0; level++ {
		size /= 2
	}
	for !tasks.isEmpty() {
		var start, end = tasks.pop()
		if end-start < 7 {
			insertSort(list[start:end])
		} else if tasks.size() == level {
			heapSort(list[start:end])
		} else {
			var knot = part(list[start:end]) + start
			tasks.push(knot+1, end)
			tasks.push(start, knot)
		} //每轮保证至少解决一个，否则最坏情况可能是死循环
	}
}

func part(list []Edge) int {
	var size = len(list)

	var seed = list[0]
	var mid, last = size / 2, size - 1
	if list[0].Dist > list[mid].Dist {
		if list[mid].Dist > list[last].Dist {
			seed, list[mid] = list[mid], list[0]
		} else { //c >= b
			if list[0].Dist > list[last].Dist {
				seed, list[last] = list[last], list[0]
			}
		}
	} else { //b >= a
		if list[last].Dist > list[0].Dist {
			if list[mid].Dist > list[last].Dist {
				seed, list[last] = list[last], list[0]
			} else {
				seed, list[mid] = list[mid], list[0]
			}
		}
	}

	var left, right = 1, last
	for { //注意对称性
		for list[left].Dist < seed.Dist {
			left++
		}
		for list[right].Dist > seed.Dist {
			right--
		}
		if left >= right {
			break
		}
		list[left], list[right] = list[right], list[left]
		left++
		right--
	}
	list[0], list[right] = list[right], seed

	return right
}

func heapSort(list []Edge) {
	var size = len(list)
	for idx := size/2 - 1; idx >= 0; idx-- {
		down(list, idx)
	}
	for sz := size - 1; sz > 0; sz-- {
		list[0], list[sz] = list[sz], list[0]
		down(list[:sz], 0)
	}
}

func down(list []Edge, spot int) {
	var size = len(list)
	var key = list[spot]
	var left, right = spot*2 + 1, spot*2 + 2
	for right < size {
		var kid int
		if list[left].Dist > list[right].Dist {
			kid = left
		} else {
			kid = right
		}
		if key.Dist >= list[kid].Dist {
			goto LabelOver
		}
		list[spot] = list[kid]
		spot, left, right = kid, kid*2+1, kid*2+2
	}
	if right == size && key.Dist < list[left].Dist {
		list[spot], list[left] = list[left], key
		return
	}
LabelOver:
	list[spot] = key
}

func insertSort(list []Edge) {
	var size = len(list)
	for i := 1; i < size; i++ {
		var left, right = 0, i
		var key = list[i]
		for left < right {
			var mid = (left + right) / 2
			if key.Dist < list[mid].Dist {
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

type pair struct {
	start int
	end   int
}
type stack struct {
	core []pair
}

func (s *stack) clear() {
	s.core = s.core[:0]
}
func (s *stack) size() int {
	return len(s.core)
}
func (s *stack) isEmpty() bool {
	return len(s.core) == 0
}
func (s *stack) push(start int, end int) {
	s.core = append(s.core, pair{start, end})
}
func (s *stack) pop() (start int, end int) {
	var sz = len(s.core) - 1
	var unit = s.core[sz]
	s.core = s.core[:sz]
	return unit.start, unit.end
}
