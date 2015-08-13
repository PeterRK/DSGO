package path

import (
	"errors"
)

type Path struct {
	Next int
	Dist int
}

const MAX_DIST = int((^uint(0)) >> 1)

//输入邻接表，返回某点到各点的最短路径的长度(MAX_DIST指不通)。
//实为改良的Bellman-Ford算法，复杂度为O(EV)，逊于Dijkstra，但可以处理负权边。
func SPFA(roads [][]Path, start int) ([]int, error) {
	var size = len(roads)
	if size == 0 || start < 0 || start >= size {
		return []int{}, errors.New("illegal input")
	}

	var q = newQueue(size)
	var dists = make([]int, size)
	var cnts = make([]int, size) //绝对值记录入队次数，负值表示在队
	for i := 0; i < size; i++ {
		dists[i], cnts[i] = MAX_DIST, 0
	}

	q.push(start)
	dists[start], cnts[start] = 0, -1
	for !q.isEmpty() {
		var current = q.pop()
		cnts[current] = -cnts[current]
		for _, path := range roads[current] {
			var distance = dists[current] + path.Dist
			var peer = path.Next
			if distance < dists[peer] {
				dists[peer] = distance
				if cnts[peer] >= 0 { //未入队
					q.push(peer)
					cnts[peer]++
					if cnts[peer] > size { //负回路
						return []int{}, errors.New("bad loops exist")
					}
					cnts[peer] = -cnts[peer] //入队
				}
			}
		}
	}
	return dists, nil
}

type queue struct {
	space    []int
	rpt, wpt int
}

func newQueue(size int) *queue {
	var q = new(queue)
	q.space = make([]int, size)
	q.rpt, q.wpt = 0, 0
	return q
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
