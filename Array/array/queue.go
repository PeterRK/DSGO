package array

type Queue struct {
	space    []int
	mask     int
	rpt, wpt int
}

func (queue *Queue) Initialize(lv uint) {
	if lv < 4 {
		lv = 4
	}
	queue.space = make([]int, lv)
	queue.mask = ^(int(-1) << lv)
	queue.clear()
}
func (queue *Queue) clear() {
	queue.rpt, queue.wpt = 0, 0
}

func (queue *Queue) IsEmpty() bool {
	return queue.rpt == queue.wpt
}
func (queue *Queue) IsFull() bool {
	var next = (queue.wpt + 1) & queue.mask
	return next == queue.rpt
}

func (queue *Queue) Push(key int) (fail bool) {
	var next = (queue.wpt + 1) & queue.mask
	if next == queue.rpt {
		return true
	}
	queue.space[queue.wpt] = key
	queue.wpt = next
	return false
}

func (queue *Queue) Top() (key int, fail bool) {
	return queue.space[queue.rpt], queue.IsEmpty()
}
func (queue *Queue) Pop() (key int, fail bool) {
	if queue.IsEmpty() {
		return 0, true
	}
	key, queue.rpt = queue.space[queue.rpt], (queue.rpt+1)&queue.mask
	return key, fail
}
