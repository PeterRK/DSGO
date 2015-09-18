package path

import (
	"Graph/graph"
)

//输入邻接表，返回某点到各点的最短路径的长度(-1指不通)。
//复杂度为O(E+VlogV)，已知最快的单源最短路径算法，对稀疏图尤甚。
//可以处理有向图，不能处理负权边。
func Dijkstra(roads [][]graph.Path, start int) []int {
	var size = len(roads)
	if size == 0 || start < 0 || start >= size {
		return []int{}
	}

	var result = make([]int, size)
	if size == 1 {
		result[0] = 0
		return result
	}

	const FAKE = -1
	var list = graph.NewVector(size)
	for i := 0; i < size; i++ {
		list[i].Link = FAKE
	}
	list[start].Index, list[start].Link, list[start].Dist = start, start, 0
	var root = graph.Insert(nil, &list[start])

	for root != nil && root.Dist != graph.MaxDistance {
		var current = root
		root = graph.Extract(root)
		for _, path := range roads[current.Index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE { //未涉及点
				peer.Index, peer.Link = path.Next, current.Index
				peer.Dist = current.Dist + path.Dist
				root = graph.Insert(root, peer)
			} else if peer.Index != FAKE { //外围点
				var distance = current.Dist + path.Dist
				if distance < peer.Dist {
					//peer.Link = current.Index
					root = graph.FloatUp(root, peer, distance)
				}
			}
		}
		current.Index = FAKE //入围
	}

	for i := 0; i < size; i++ {
		if list[i].Dist == graph.MaxDistance {
			result[i] = -1
		} else {
			result[i] = (int)(list[i].Dist)
		}
	}
	return result
}

//输入邻接表，返回两点间的最短路径及其长度(-1指不通)。
func DijkstraPath(roads [][]graph.Path, start int, end int) (
	Dist int, marks []int) {
	var size = len(roads)
	if start < 0 || end < 0 || start >= size || end >= size {
		return -1, []int{}
	}
	if start == end {
		return 0, []int{start}
	}

	const FAKE = -1
	var list = graph.NewVector(size)
	for i := 0; i < size; i++ {
		list[i].Link = FAKE
	}
	list[start].Index, list[start].Link, list[start].Dist = start, start, 0
	var root = graph.Insert(nil, &list[start])

	for root != nil && root.Dist != graph.MaxDistance {
		var current = root
		if current.Index == end {
			for idx := end; idx != start; idx = list[idx].Link {
				marks = append(marks, idx)
			}
			marks = append(marks, start)
			for left, right := 0, len(marks)-1; left < right; {
				marks[left], marks[right] = marks[right], marks[left]
				left++
				right--
			}
			return (int)(current.Dist), marks
		}
		root = graph.Extract(root)
		for _, path := range roads[current.Index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE { //未涉及点
				peer.Index, peer.Link = path.Next, current.Index
				peer.Dist = current.Dist + path.Dist
				root = graph.Insert(root, peer)
			} else if peer.Index != FAKE { //外围点
				var distance = current.Dist + path.Dist
				if distance < peer.Dist {
					peer.Link = current.Index
					root = graph.FloatUp(root, peer, distance)
				}
			}
		}
		current.Index = FAKE //入围
	}
	return -1, []int{}
}
