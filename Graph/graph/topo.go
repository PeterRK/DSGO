package graph

import (
	"errors"
)

//输入邻接表，返回拓扑序列
func TopologicalSort(roads [][]int) ([]int, error) {
	var list []int
	var size = len(roads)
	var book = make([]int, size)
	//	for i := 0; i < size; i++ {
	//		book[i] = 0
	//	}
	for i := 0; i < size; i++ {
		for _, next := range roads[i] {
			book[next]++ //标记存在几个上游
		}
	}

	var space []int
	for i := 0; i < size; i++ {
		if book[i] == 0 {
			space = append(space, i) //最上游点
		}
	}

	for len(space) != 0 {
		var last = len(space) - 1
		var current = space[last]
		space = space[:last]

		for _, next := range roads[current] {
			book[next]--
			if book[next] == 0 {
				space = append(space, next)
			}
		}
		list = append(list, current)
	}

	if len(list) != size {
		return []int{}, errors.New("loops exist")
	}
	return list, nil
}
