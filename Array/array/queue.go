package array

type Queue interface {
	Clear()
	IsFull() bool
	IsEmpty() bool
	Push(key int) (fail bool)
	Pop() (key int, fail bool)
	Front() (key int, fail bool)
}

func NewQueue(size int) Queue {
	var core = new(queue)
	core.initialize(size)
	return core
}

//实际容量为len(space)-1
type queue struct {
	space    []int
	mask     int
	rpt, wpt int
}

func (q *queue) initialize(size int) {
	if size > 0x10000 {
		size = 0x10000
	} else if size < 4 {
		size = 4
	}
	var sz = 4
	for sz < size {
		sz *= 2
	}
	q.space = make([]int, sz)
	q.mask = sz - 1
	q.Clear()
}
func (q *queue) Clear() {
	q.rpt, q.wpt = 0, 0
}

func (q *queue) IsFull() bool {
	var next = (q.wpt + 1) & q.mask
	return next == q.rpt
}
func (q *queue) IsEmpty() bool {
	return q.rpt == q.wpt
}

func (q *queue) Push(key int) (fail bool) {
	var next = (q.wpt + 1) & q.mask
	if next == q.rpt {
		return true
	}
	q.space[q.wpt] = key
	q.wpt = next
	return false
}

func (q *queue) Front() (key int, fail bool) {
	return q.space[q.rpt], q.IsEmpty()
}
func (q *queue) Pop() (key int, fail bool) {
	if q.IsEmpty() {
		return 0, true
	}
	key, q.rpt = q.space[q.rpt], (q.rpt+1)&q.mask
	return key, false
}
