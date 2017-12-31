package span

import (
	"errors"
)

//输入边集，返回最小生成树的权。
//复杂度为O(ElogE)，性能通常不如Prim。
func Kruskal(roads []EdgeW, size int) (uint, error) {
	if size < 2 || len(roads) < size-1 {
		return 0, errors.New("illegal input")
	}
	sort(roads)

	type memo struct {
		mark int //正数表示归属，负数表示个数（仅首领项）
		next *memo
	}
	list := make([]memo, size)
	for i := 0; i < size; i++ {
		list[i].mark, list[i].next = -1, &list[i]
	}
	trace := func(id int) int {
		if list[id].mark < 0 {
			return id
		} else {
			return list[id].mark
		}
	}

	sum := uint(0)
	for _, path := range roads {
		active, another := trace(path.A), trace(path.B)
		if active == another {
			continue
		}
		sum += path.Weight

		if -list[active].mark < -list[another].mark {
			active, another = another, active
		}
		list[active].mark += list[another].mark

		tail := list[active].next
		list[active].next, list[another].next = list[another].next, tail
		for u := list[active].next; u != tail; u = u.next {
			u.mark = active
		}

		if -list[active].mark == size {
			return sum, nil
		}
	}
	return sum, errors.New("isolated part exist")
}

func KruskalS(roads []EdgeW, size int) (uint, error) {
	if size < 2 || len(roads) < size-1 {
		return 0, errors.New("illegal input")
	}
	sort(roads)

	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = -1 //正数表示归属，负数表示个数（仅首领项）
	}
	trace := func(id int) int {
		for list[id] >= 0 {
			id = list[id]
		}
		return id
	}

	sum := uint(0)
	for _, path := range roads {
		active, another := trace(path.A), trace(path.B)
		if active == another {
			continue
		}
		sum += path.Weight

		if -list[active] < -list[another] {
			active, another = another, active
		}
		list[active] += list[another]
		list[another] = active

		if -list[active] == size {
			return sum, nil
		}
	}
	return sum, errors.New("isolated part exist")
}
