package array

import (
	"errors"
)

type Queue interface {
	Clear()
	IsFull() bool
	IsEmpty() bool
	Push(key int) error
	Pop() (int, error)
	Front() (int, error)
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
		return errors.New("illefal queue size")
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

func (q *queue) Push(key int) error {
	var next = (q.wpt + 1) & q.mask
	if next == q.rpt {
		return errors.New("full")
	}
	q.space[q.wpt] = key
	//memory barrier
	q.wpt = next
	return nil
}

func (q *queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty")
	}
	return q.space[q.rpt], nil
}
func (q *queue) Pop() (int, error) {
	var key, err = q.Front()
	if err == nil {
		//memory barrier
		q.rpt = (q.rpt + 1) & q.mask
	}
	return key, nil
}
