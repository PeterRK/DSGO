package array

type Queue interface {
	Clear()
	IsFull() bool
	IsEmpty() bool
	Push(key int) (fail bool)
	Pop() (key int, fail bool)
	Front() (key int, fail bool)
}

func NewQueue(lv uint) Queue {
	var core = new(queue)
	core.initialize(lv)
	return core
}

type queue struct {
	space    []int
	mask     int
	rpt, wpt int
}

func (this *queue) initialize(lv uint) {
	if lv < 4 {
		lv = 4
	}
	this.space = make([]int, lv)
	this.mask = ^(int(-1) << lv)
	this.Clear()
}
func (this *queue) Clear() {
	this.rpt, this.wpt = 0, 0
}

func (this *queue) IsFull() bool {
	var next = (this.wpt + 1) & this.mask
	return next == this.rpt
}
func (this *queue) IsEmpty() bool {
	return this.rpt == this.wpt
}

func (this *queue) Push(key int) (fail bool) {
	var next = (this.wpt + 1) & this.mask
	if next == this.rpt {
		return true
	}
	this.space[this.wpt] = key
	this.wpt = next
	return false
}

func (this *queue) Front() (key int, fail bool) {
	return this.space[this.rpt], this.IsEmpty()
}
func (this *queue) Pop() (key int, fail bool) {
	if this.IsEmpty() {
		return 0, true
	}
	key, this.rpt = this.space[this.rpt], (this.rpt+1)&this.mask
	return key, fail
}
