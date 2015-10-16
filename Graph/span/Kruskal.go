package span

import (
	"Graph/graph"
	"errors"
)

//输入边集，返回最小生成树的权。
//复杂度为O(ElogE)，性能通常不如Prim。
func Kruskal(roads []graph.Edge, size int) (uint, error) {
	if size < 2 || len(roads) < size-1 {
		return 0, errors.New("illegal input")
	}
	graph.Sort(roads)

	type memo struct {
		cnt   int
		group int
		next  *memo
	}
	var list = make([]memo, size)
	for i := 0; i < size; i++ {
		list[i].cnt, list[i].group, list[i].next = 1, i, &list[i]
	}

	var sum = uint(0)
	for _, path := range roads {
		var grpA, grpB = list[path.A].group, list[path.B].group
		if grpA == grpB {
			continue
		}
		sum += path.Dist

		var active, another *memo
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

		if active.cnt == size {
			return sum, nil
		}
	}
	return sum, errors.New("isolated part exist")
}

func KruskalS(roads []graph.Edge, size int) (uint, error) {
	if size < 2 || len(roads) < size-1 {
		return 0, errors.New("illegal input")
	}
	graph.Sort(roads)

	type memo struct {
		cnt   int
		super int
	}
	var list = make([]memo, size)
	for i := 0; i < size; i++ {
		list[i].cnt, list[i].super = 1, i
	}
	var trace = func(space []memo, id int) (super int, depth uint) {
		super = space[id].super
		for depth = uint(0); super != id; depth++ {
			id, super = super, space[super].super
		}
		return super, depth
	}

	var sum = uint(0)
	for _, path := range roads {
		var grpA, dpA = trace(list, path.A)
		var grpB, dpB = trace(list, path.B)
		if grpA == grpB {
			continue
		}
		sum += path.Dist

		var active, another int
		if dpA > dpB {
			active, another = grpA, grpB
		} else {
			active, another = grpB, grpA
		}
		list[active].cnt += list[another].cnt
		list[another].super = active

		if list[active].cnt == size {
			return sum, nil
		}
	}
	return sum, errors.New("isolated part exist")
}
