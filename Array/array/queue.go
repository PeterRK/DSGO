package array

import (
	"errors"
)

type Queue interface {
	Clear()
	IsFull() bool
	IsEmpty() bool
	Push(key int) (fail bool)
	Pop() (key int, fail bool)
	Front() (key int, fail bool)
}

func NewQueue(size int) (Queue, error) {
	var core = new(queue)
	var err = core.initialize(size)
	if err != nil {
		return nil, err
	}
	return core, nil
}

type queue struct {
	space    []int
	mask     int
	rpt, wpt int
}

func (q *queue) initialize(size int) error {
	if size < 1 || size > 0xffff {
		return errors.New("Illefal queue size")
	}
	var sz = 4
	for sz <= size {
		sz *= 2
	} //实际容量为sz-1
	q.space = make([]int, sz)
	q.mask = sz - 1
	q.Clear()
	return nil
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
	//memory barrier
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
	key = q.space[q.rpt]
	//memory barrier
	q.rpt = (q.rpt + 1) & q.mask
	return key, false
}
