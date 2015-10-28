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

type stack struct {
	space []int
	spcx  []uint
	top   int
}

func (s *stack) bind(space []int, spcx []uint) {
	s.space, s.spcx = space, spcx
}
func (s *stack) clear() {
	s.top = 0
}
func (s *stack) isEmpty() bool {
	return s.top == 0
}
func (s *stack) push(id int, val uint) {
	s.space[s.top], s.spcx[s.top] = id, val
	s.top++
}
func (s *stack) pop() (id int, val uint) {
	s.top--
	return s.space[s.top], s.spcx[s.top]
}

type queue struct {
	space    []int
	rpt, wpt int
}

func (q *queue) bind(space []int) {
	q.space = space
}
func (q *queue) clear() {
	q.rpt, q.wpt = 0, 0
}
func (q *queue) isEmpty() bool {
	return q.rpt == q.wpt
}
func (q *queue) push(key int) {
	q.space[q.wpt] = key
	q.wpt = (q.wpt + 1) % len(q.space)
}
func (q *queue) pop() int {
	var key = q.space[q.rpt]
	q.rpt = (q.rpt + 1) % len(q.space)
	return key
}
func (q *queue) traceBack() (int, error) {
	q.wpt--
	if q.wpt < 0 {
		return 0, errors.New("empty")
	}
	return q.space[q.wpt], nil
}
