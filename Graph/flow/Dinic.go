package flow

import (
	"errors"
)

type edge struct {
	next int
	val  uint
}

//输入邻接矩阵，返回头节点到尾顶点见的最大流。
//复杂度为O(V^2 E)。
func Dinic(matrix [][]uint) uint {
	var size = len(matrix)
	if size == 0 {
		return 0
	}

	var space1 = make([]int, size)  //临时空间
	var space2 = make([]uint, size) //临时空间
	var q queue
	q.bind(space1)
	var s stack
	s.bind(space1, space2)

	var shadow = make([][]edge, size-1) //分层残图

	var flow = uint(0)
	for separate(shadow, matrix, &q, space2) {
		//由于search一次会删除图的若干条边，所以循环次数为O(E/k)
		//循环内的separate和flushBack操作的复杂度都是O(V^2)，search的复杂度为O(Vk)
		flow += search(shadow, matrix, &s)
		flushBack(shadow, matrix)
	}
	return flow
}

////////////////////////////////////////////////////////////////////////////////
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

////////////////////////////////////////////////////////////////////////////////
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
