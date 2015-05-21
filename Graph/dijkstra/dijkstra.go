package dijkstra

import (
	"Graph/graph"
)

func Dijkstra(roads [][]graph.Path, start int) []int {
	var size = len(roads)
	if size == 0 || start < 0 || start >= size {
		return make([]int, 0)
	}

	var result = make([]int, size)
	if size == 1 {
		result[0] = 0
		return result
	}

	var root, list = InitHeap(size, start)
	for root != nil {
		var current = root
		root = Extract(root)
		for _, path := range roads[current.index] {
			var peer = &list[path.Next]
			if peer.index == path.Next {
				var distance = current.dist + path.Dist
				if distance < peer.dist {
					root = FloatUp(root, peer, distance)
				}
			}
		}
		current.index = -1
	}

	for i := 0; i < size; i++ {
		if list[i].dist == MaxDistance {
			result[i] = -1
		} else {
			result[i] = (int)(list[i].dist)
		}
	}
	return result
}

func DijkstraX(roads [][]graph.Path, start int, end int) int {
	var size = len(roads)
	if start < 0 || end < 0 || start >= size || end >= size {
		return -1
	}
	if start == end {
		return 0
	}

	var root, list = InitHeap(size, start)
	for root != nil && root.index != end {
		var current = root
		root = Extract(root)
		for _, path := range roads[current.index] {
			var peer = &list[path.Next]
			if peer.index == path.Next {
				var distance = current.dist + path.Dist
				if distance < peer.dist {
					root = FloatUp(root, peer, distance)
				}
			}
		}
		current.index = -1
	}
	if list[end].dist == MaxDistance {
		return -1
	}
	return (int)(list[end].dist)
}
