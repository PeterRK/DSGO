package dijkstra

import (
	"Graph/graph"
)

//输入邻接表，返回某点到各点的最短路径的长度(-1指不通)。
//本实现复杂度为O(ElogV)，理论上最佳为O(E+VlogV)
//已知最快的单源最短路径算法，对稀疏图尤甚。
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

	var root, list = graph.NewHeap(size, start)
	for root != nil && root.Dist != graph.MaxDistance {
		var current = root
		root = graph.Extract(root)
		for _, path := range roads[current.Index] {
			var peer = &list[path.Next]
			if peer.Index == path.Next { //针对未处理的点
				var distance = current.Dist + path.Dist
				if distance < peer.Dist {
					root = graph.FloatUp(root, peer, distance)
				}
			}
		}
		current.Index = -1 //标记为已经处理
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
func DijkstraPath(roads [][]graph.Path, start int, end int) (Dist int, marks []int) {
	var size = len(roads)
	if start < 0 || end < 0 || start >= size || end >= size {
		return -1, []int{}
	}
	if start == end {
		return 0, []int{start}
	}

	var root, list = graph.NewHeap(size, start)
	for root != nil && root.Dist != graph.MaxDistance {
		if root.Index == end {
			for idx := end; idx != start; idx = list[idx].Link {
				marks = append(marks, idx)
			}
			marks = append(marks, start)
			for left, right := 0, len(marks)-1; left < right; {
				marks[left], marks[right] = marks[right], marks[left]
				left++
				right--
			}
			return (int)(root.Dist), marks
		}
		var current = root
		root = graph.Extract(root)
		for _, path := range roads[current.Index] {
			var peer = &list[path.Next]
			if peer.Index == path.Next { //针对未处理的点
				var distance = current.Dist + path.Dist
				if distance < peer.Dist {
					peer.Link = current.Index
					root = graph.FloatUp(root, peer, distance)
				}
			}
		}
		current.Index = -1 //标记为已经处理
	}
	return -1, []int{}
}
