package span

import (
	"DSGO/Graph/heap"
	"errors"
)

func PrimX(roads [][]Path) (uint, error) {
	size := len(roads)
	sum := uint(0)
	if size < 2 {
		return 0, errors.New("illegal input")
	}

	const FAKE = -1
	list := heap.NewVectorBH(size)
	for i := 1; i < size; i++ {
		list[i].Link = FAKE
	}
	list[0].Index, list[0].Link, list[0].Dist = 0, 0, 0

	cnt := 0
	hp := heap.NewBinaryHeap(size)
	hp.Push(&list[0])
	for !hp.IsEmpty() {
		curr := hp.Pop()
		sum += curr.Dist
		index := curr.Index
		curr.Index = FAKE
		for _, path := range roads[index] {
			peer := &list[path.Next]
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

func PrimTreeX(roads [][]Path) ([]Edge, error) {
	size := len(roads)
	if size < 2 {
		return nil, errors.New("illegal input")
	}
	edges := make([]Edge, 0, size-1)

	const FAKE = -1
	list := heap.NewVectorBH(size)
	for i := 1; i < size; i++ {
		list[i].Link = FAKE
	}

	hp := heap.NewBinaryHeap(size)
	curr := &list[0]
	for {
		index := curr.Index
		curr.Index = FAKE
		for _, path := range roads[index] {
			peer := &list[path.Next]
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
