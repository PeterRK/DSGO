package span

import (
	"DSGO/array/sort"
	"DSGO/graph"
	"errors"
)

func edgeLess(a, b *graph.Edge) bool {
	return a.Weight < b.Weight
}

//输入边集，返回最小生成树的权。
//复杂度为O(ElogE)，性能通常不如Prim。
func Kruskal(roads []graph.Edge, size int) (uint, error) {
	if size < 2 || len(roads) < size-1 {
		return 0, errors.New("illegal input")
	}
	ranker := sort.Ranker[graph.Edge]{Less: edgeLess}
	ranker.Sort(roads)

	nodes := make([]int, size)
	for i := 0; i < size; i++ {
		nodes[i] = -1 //正数表示归属，负数表示个数（仅首领项）
	}
	trace := func(id int) int {
		sid := id
		for nodes[sid] >= 0 {
			sid = nodes[sid]
		}
		if sid != id {
			nodes[id] = sid
		}
		return sid
	}

	sum := uint(0)
	for _, path := range roads {
		active, another := trace(path.A), trace(path.B)
		if active == another {
			continue
		}
		sum += path.Weight

		if -nodes[active] < -nodes[another] {
			active, another = another, active
		}
		nodes[active] += nodes[another]
		nodes[another] = active

		if -nodes[active] == size {
			return sum, nil
		}
	}
	return sum, errors.New("isolated part exist")
}

func Kruskal_v2(roads []graph.Edge, size int) (uint, error) {
	if size < 2 || len(roads) < size-1 {
		return 0, errors.New("illegal input")
	}
	ranker := sort.Ranker[graph.Edge]{Less: edgeLess}
	ranker.Sort(roads)

	type node struct {
		mark int //正数表示归属
		next *node
	}
	nodes := make([]node, size)
	for i := 0; i < size; i++ {
		nodes[i].mark, nodes[i].next = -1, &nodes[i]
	}
	trace := func(id int) int {
		if nodes[id].mark < 0 {
			return id
		} else {
			return nodes[id].mark
		}
	}

	sum := uint(0)
	for _, path := range roads {
		active, another := trace(path.A), trace(path.B)
		if active == another {
			continue
		}
		sum += path.Weight

		if -nodes[active].mark < -nodes[another].mark {
			active, another = another, active
		}
		nodes[active].mark += nodes[another].mark

		tail := nodes[active].next
		nodes[active].next, nodes[another].next = nodes[another].next, tail
		for u := nodes[active].next; u != tail; u = u.next {
			u.mark = active
		}

		if -nodes[active].mark == size {
			return sum, nil
		}
	}
	return sum, errors.New("isolated part exist")
}
