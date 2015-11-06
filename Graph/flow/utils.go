package flow

import (
	"errors"
)

func min(a uint, b uint) uint {
	if a > b {
		return b
	}
	return a
}

type arrayStack struct {
	space []int
	spcx  []uint
	top   int
}

func (s *arrayStack) bind(space []int, spcx []uint) {
	s.space, s.spcx = space, spcx
}
func (s *arrayStack) clear() {
	s.top = 0
}
func (s *arrayStack) isEmpty() bool {
	return s.top == 0
}
func (s *arrayStack) push(id int, val uint) {
	s.space[s.top], s.spcx[s.top] = id, val
	s.top++
}
func (s *arrayStack) pop() (id int, val uint) {
	s.top--
	return s.space[s.top], s.spcx[s.top]
}

type arrayQueue struct {
	space    []int
	rpt, wpt int
}

func (q *arrayQueue) bind(space []int) {
	q.space = space
}
func (q *arrayQueue) clear() {
	q.rpt, q.wpt = 0, 0
}
func (q *arrayQueue) isEmpty() bool {
	return q.rpt == q.wpt
}
func (q *arrayQueue) push(key int) {
	q.space[q.wpt] = key
	q.wpt = (q.wpt + 1) % len(q.space)
}
func (q *arrayQueue) pop() int {
	var key = q.space[q.rpt]
	q.rpt = (q.rpt + 1) % len(q.space)
	return key
}
func (q *arrayQueue) traceBack() (int, error) {
	q.wpt--
	if q.wpt < 0 {
		return 0, errors.New("empty")
	}
	return q.space[q.wpt], nil
}
