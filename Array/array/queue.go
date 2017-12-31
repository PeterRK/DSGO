package array

import (
	"errors"
)

// 环状队列
type Queue interface {
	Clear()
	IsFull() bool
	IsEmpty() bool
	Push(key int) error
	Pop() (int, error)
	Front() (int, error)
}

// 申请环状队列
func NewQueue(size int) (Queue, error) {
	core := new(queue)
	if err := core.initialize(size); err != nil {
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
	sz := 4
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
	return (q.wpt+1)&q.mask == q.rpt
}
func (q *queue) IsEmpty() bool {
	return q.rpt == q.wpt
}

func (q *queue) Push(key int) error {
	next := (q.wpt + 1) & q.mask
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
	key, err := q.Front()
	if err == nil {
		//memory barrier
		q.rpt = (q.rpt + 1) & q.mask
	}
	return key, nil
}
