package flow

import (
	"Graph/graph"
	"errors"
)

func flushBack(shadow [][]graph.Path, matrix [][]uint) {
	var size = len(shadow) //len(shadow) == len(matrix)-1
	for i := 1; i < size; i++ {
		if len(shadow[i]) != 0 {
			for _, path := range shadow[i] {
				matrix[i][path.Next] += path.Weight
			}
			shadow[i] = shadow[i][:0]
		}
	}
}

const fake_level = ^uint(0)

//宽度优先遍历
func markLevel(matrix [][]uint, q *queue, memo []uint) bool {
	var last = len(matrix) - 1
	for i := 1; i <= last; i++ {
		memo[i] = fake_level
	}
	memo[0] = 0

	q.push(0)
	for !q.isEmpty() {
		var cur = q.pop()
		if matrix[cur][last] != 0 {
			memo[last] = memo[cur] + 1
			return true
		}
		for i := 1; i < last; i++ {
			if memo[i] == fake_level && matrix[cur][i] != 0 {
				memo[i] = memo[cur] + 1
				q.push(i)
			}
		}
	}
	return false
}

//筛分层次，生成分层残图，复杂度为(V^2)。
func separate(shadow [][]graph.Path, matrix [][]uint, q *queue, memo []uint) bool {
	q.clear()
	if !markLevel(matrix, q, memo) {
		return false
	}
	for { //队列pop出的点并没有实际删除，可回溯遍历所有访问过的点
		var cur, err = q.traceBack()
		if err != nil {
			break
		}
		//shadow[cur] = shadow[cur][:0]
		for i := 1; i < len(matrix); i++ {
			if memo[i] == memo[cur]+1 && matrix[cur][i] != 0 {
				var path = graph.Path{Next: i, Weight: matrix[cur][i]}
				shadow[cur] = append(shadow[cur], path)
				matrix[cur][i] = 0
			}
		}
		if len(shadow[cur]) == 0 {
			memo[cur] = fake_level
		}
	}
	return true
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
