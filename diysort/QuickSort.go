package diysort

//最坏情况是O(N)，不具有稳定性
func QuickSort(list []int) {
	var tasks stack
	tasks.push(0, len(list))
	for !tasks.isEmpty() {
		var start, end = tasks.pop()
		if end-start < 7 { //内建SelectSort
			for i := start; i < end-1; i++ {
				var pos = i
				for j := i + 1; j < end; j++ {
					if list[j] < list[pos] {
						pos = j
					}
				}
				list[pos], list[i] = list[i], list[pos]
			}
			continue
		}

		//三点取中法，最后保证seed落入中间，每轮至少解决此一处
		var seed = list[start]
		var mid = (start + end) / 2
		if list[start] > list[mid] {
			if list[mid] > list[end-1] {
				seed, list[mid] = list[mid], list[start]
			} else { //c >= b
				if list[start] > list[end-1] {
					seed, list[end-1] = list[end-1], list[start]
				}
			}
		} else { //b >= a
			if list[end-1] > list[start] {
				if list[mid] > list[end-1] {
					seed, list[end-1] = list[end-1], list[start]
				} else {
					seed, list[mid] = list[mid], list[start]
				}
			}
		}

		var left, right = start + 1, end - 1
		for { //注意对称性
			for list[left] < seed {
				left++
			}
			for list[right] > seed {
				right--
			}
			if left >= right {
				break
			}
			list[left], list[right] = list[right], list[left]
			left++
			right--
		}
		list[start], list[right] = list[right], seed
		//每轮保证至少解决一个，否则最坏情况可能是死循环
		tasks.push(right+1, end)
		tasks.push(start, right)
	}
}

type pair struct {
	start int
	end   int
}
type stack struct {
	core []pair
}

func (this *stack) clear() {
	this.core = this.core[:0]
}
func (this *stack) size() int {
	return len(this.core)
}
func (this *stack) isEmpty() bool {
	return len(this.core) == 0
}
func (this *stack) push(start int, end int) {
	this.core = append(this.core, pair{start, end})
}
func (this *stack) pop() (start int, end int) {
	var sz = len(this.core) - 1
	var unit = this.core[sz]
	this.core = this.core[:sz]
	return unit.start, unit.end
}
