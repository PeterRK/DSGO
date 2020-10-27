package path

import (
	"errors"
)

type PathS struct {
	Next int
	Dist int
}

const MAX_DIST = int((^uint(0)) >> 1)

//输入邻接表，返回某点到各点的最短路径的长度(MAX_DIST指不通)。
//实为改良的Bellman-Ford算法，复杂度为O(EV)，逊于Dijkstra，但可以处理负权边。
func SPFA(roads [][]PathS, start int) ([]int, error) {
	size := len(roads)
	if size == 0 || start < 0 || start >= size {
		return nil, errors.New("illegal input")
	}

	q := newQueue(size)
	dists := make([]int, size)
	cnts := make([]int, size) //绝对值记录入队次数，负值表示在队
	for i := 0; i < size; i++ {
		dists[i], cnts[i] = MAX_DIST, 0
	}

	q.push(start)
	dists[start], cnts[start] = 0, -1
	for !q.isEmpty() {
		cur := q.pop()
		cnts[cur] = -cnts[cur]
		for _, path := range roads[cur] {
			distance := dists[cur] + path.Dist
			peer := path.Next
			if distance < dists[peer] {
				dists[peer] = distance
				if cnts[peer] >= 0 { //未入队
					q.push(peer)
					cnts[peer]++
					if cnts[peer] > size { //负回路
						return nil, errors.New("bad loops exist")
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
	q := new(queue)
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
	key := q.space[q.rpt]
	q.rpt = (q.rpt + 1) % len(q.space)
	return key
}
