package span

import (
	"DSGO/Graph/graph"
	"DSGO/Graph/heap"
	"errors"
)

func PrimX(roads [][]graph.Path) (uint, error) {
	var size = len(roads)
	var sum = uint(0)
	if size < 2 {
		return 0, errors.New("illegal input")
	}

	const FAKE = -1
	var list = heap.NewVectorBH(size)
	for i := 1; i < size; i++ {
		list[i].Link = FAKE
	}
	list[0].Index, list[0].Link, list[0].Dist = 0, 0, 0

	var cnt = 0
	var hp = heap.NewBinaryHeap(size)
	hp.Push(&list[0])
	for !hp.IsEmpty() {
		var curr = hp.Pop()
		sum += curr.Dist
		var index = curr.Index
		curr.Index = FAKE
		for _, path := range roads[index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE {
				peer.Index, peer.Link, peer.Dist = path.Next, index, path.Weight
				hp.Push(peer)
			} else if peer.Index != FAKE &&
				path.Weight < peer.Dist {
				hp.FloatUp(peer, path.Weight)
				peer.Link = index
			}
		}
		cnt++
	}
	if cnt != size {
		return sum, errors.New("isolated part exist")
	}
	return sum, nil
}

func PrimTreeX(roads [][]graph.Path) ([]Edge, error) {
	var size = len(roads)
	if size < 2 {
		return nil, errors.New("illegal input")
	}
	var edges = make([]Edge, 0, size-1)

	const FAKE = -1
	var list = heap.NewVectorBH(size)
	for i := 1; i < size; i++ {
		list[i].Link = FAKE
	}

	var hp = heap.NewBinaryHeap(size)
	var curr = &list[0]
	for {
		var index = curr.Index
		curr.Index = FAKE
		for _, path := range roads[index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE {
				peer.Index, peer.Link, peer.Dist = path.Next, index, path.Weight
				hp.Push(peer)
			} else if peer.Index != FAKE &&
				path.Weight < peer.Dist {
				peer.Link = index
				hp.FloatUp(peer, path.Weight)
			}
		}
		if hp.IsEmpty() {
			break
		}
		curr = hp.Pop()
		edges = append(edges, Edge{curr.Link, curr.Index})
	}
	if len(edges) != size-1 {
		return edges, errors.New("isolated part exist")
	}
	return edges, nil
}
