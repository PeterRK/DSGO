package graph

import (
	"DSGO/linkedlist/deque"
	"errors"
)

//输入邻接表，返回拓扑序列

func TopologicalSort(roads [][]int) ([]int, error) {
	size := len(roads)
	list := make([]int, 0, size)
	book := make([]int, size)
	//	for i := 0; i < size; i++ {
	//		book[i] = 0
	//	}
	for i := 0; i < size; i++ {
		for _, next := range roads[i] {
			book[next]++ //标记存在几个上游
		}
	}

	free := deque.NewQueue[int]()
	for i := 0; i < size; i++ {
		if book[i] == 0 {
			free.Push(i) //最上游点
		}
	}

	for !free.IsEmpty() {
		curr := free.Pop()
		for _, next := range roads[curr] {
			book[next]--
			if book[next] == 0 {
				free.Push(next)
			}
		}
		list = append(list, curr)
	}

	if len(list) != size {
		return nil, errors.New("loops exist")
	}
	return list, nil
}
