package tree

import (
	"Graph/graph"
)

type memo struct {
	cnt   int
	group int
	next  *memo
}

//输入邻接表，返回最小生成树的权。
//复杂度为O(ElogE)，性能通常不如Prim
func Kruskal(roads []graph.Edge, size int) (sum uint, fail bool) {
	if size < 2 || len(roads) < size-1 {
		return 0, true
	}
	graph.Sort(roads)

	var list = make([]memo, size)
	for i := 0; i < size; i++ {
		list[i].cnt, list[i].group, list[i].next = 1, i, &list[i]
	}

	var active = &list[0]
	sum = uint(0)
	for _, path := range roads {
		if active.cnt == size {
			break
		}
		var grpA, grpB = list[path.A].group, list[path.B].group
		if grpA == grpB {
			continue
		}
		sum += path.Dist
		var another *memo
		if list[grpA].cnt > list[grpB].cnt {
			active, another = &list[grpA], &list[grpB]
		} else {
			active, another = &list[grpB], &list[grpA]
		}
		active.cnt += another.cnt
		var tail = active.next
		active.next, another.next = another.next, tail
		for another = active.next; another != tail; another = another.next {
			another.group = active.group
		}
	}
	return sum, active.cnt != size
}
