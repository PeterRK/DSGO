package path

import (
	"DSGO/array"
	"DSGO/graph"
	"errors"
	"math"
)

//输入邻接表，返回某点到各点的最短路径的长度(MAX_DIST指不通)。
//实为改良的Bellman-Ford算法，复杂度为O(EV)，逊于Dijkstra，但可以处理负权边。
func SPFA(roads [][]graph.PathS, start int) ([]int, error) {
	size := len(roads)
	if size == 0 || start < 0 || start >= size {
		return nil, errors.New("illegal input")
	}

	q := array.NewQueue[int](size)
	dist := make([]int, size)
	age := make([]int, size) //绝对值记录入队次数，负值表示在队
	for i := 0; i < size; i++ {
		dist[i], age[i] = math.MaxInt, 0
	}

	q.Push(start)
	dist[start], age[start] = 0, -1
	for !q.IsEmpty() {
		curr := q.Pop()
		age[curr] = -age[curr]
		for _, path := range roads[curr] {
			distance := dist[curr] + path.Weight
			peer := path.Next
			if distance < dist[peer] {
				dist[peer] = distance
				if age[peer] >= 0 { //未入队
					q.Push(peer)
					age[peer]++
					if age[peer] > size { //负回路
						return nil, errors.New("bad loops exist")
					}
					age[peer] = -age[peer] //入队
				}
			}
		}
	}
	return dist, nil
}
