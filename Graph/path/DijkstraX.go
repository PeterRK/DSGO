package path

import (
	"DSGO/Graph/graph"
	"DSGO/Graph/heap"
)

func DijkstraX(roads [][]graph.Path, start int) []int {
	var size = len(roads)
	if size == 0 || start < 0 || start >= size {
		return nil
	}

	var result = make([]int, size)
	if size == 1 {
		result[0] = 0
		return result
	}

	var list = heap.NewVectorBH(size)
	const FAKE = -1
	for i := 0; i < size; i++ {
		list[i].Link = FAKE
	}

	list[start].Index, list[start].Link, list[start].Dist = start, start, 0
	var hp = heap.NewBinaryHeap(size)
	hp.Push(&list[start])
	for !hp.IsEmpty() {
		var curr = hp.Pop()
		var index = curr.Index
		curr.Index = FAKE
		for _, path := range roads[index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE {
				peer.Index, peer.Link = path.Next, index
				peer.Dist = curr.Dist + path.Weight
				hp.Push(peer)
			} else if peer.Index != FAKE {
				var distance = curr.Dist + path.Weight
				if distance < peer.Dist {
					hp.FloatUp(peer, distance)
					//peer.Link = index
				}
			}
		}
	}

	for i := 0; i < size; i++ {
		if list[i].Dist == heap.MaxDistance {
			result[i] = -1
		} else {
			result[i] = (int)(list[i].Dist)
		}
	}
	return result
}

func DijkstraPathX(roads [][]graph.Path, start, end int) []int {
	var size = len(roads)
	if start < 0 || end < 0 || start >= size || end >= size {
		return nil
	}
	if start == end {
		return []int{start}
	}

	var list = heap.NewVectorBH(size)
	const FAKE = -1
	for i := 0; i < size; i++ {
		list[i].Link = FAKE
	}
	var trace = func() []int {
		var path []int
		for idx := end; idx != start; idx = list[idx].Link {
			path = append(path, idx)
		}
		path = append(path, start)
		reverse(path)
		return path
	}

	list[start].Index, list[start].Link, list[start].Dist = start, start, 0
	var hp = heap.NewBinaryHeap(size)
	hp.Push(&list[start])
	for !hp.IsEmpty() {
		var curr = hp.Pop()
		if curr.Index == end {
			return trace()
		}
		var index = curr.Index
		curr.Index = FAKE
		for _, path := range roads[index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE {
				peer.Index, peer.Link = path.Next, index
				peer.Dist = curr.Dist + path.Weight
				hp.Push(peer)
			} else if peer.Index != FAKE {
				var distance = curr.Dist + path.Weight
				if distance < peer.Dist {
					hp.FloatUp(peer, distance)
					peer.Link = index
				}
			}
		}
	}
	return nil
}
