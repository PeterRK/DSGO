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

	var root, list = newHeap(size, start)
	for root != nil && root.dist != MaxDistance {
		var current = root
		root = extract(root)
		for _, path := range roads[current.index] {
			var peer = &list[path.Next]
			if peer.index == path.Next { //针对未处理的点
				var distance = current.dist + path.Dist
				if distance < peer.dist {
					root = floatUp(root, peer, distance)
				}
			}
		}
		current.index = -1 //标记为已经处理
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

//输入邻接表，返回两点间的最短路径及其长度(-1指不通)。
func DijkstraPath(roads [][]graph.Path, start int, end int) (dist int, marks []int) {
	var size = len(roads)
	if start < 0 || end < 0 || start >= size || end >= size {
		return -1, marks
	}
	if start == end {
		return 0, marks
	}

	var root, list = newHeap(size, start)
	for root != nil && root.dist != MaxDistance {
		if root.index == end {
			for idx := end; idx != start; idx = list[idx].link {
				marks = append(marks, idx)
			}
			marks = append(marks, start)
			for left, right := 0, len(marks)-1; left < right; {
				marks[left], marks[right] = marks[right], marks[left]
				left++
				right--
			}
			return (int)(root.dist), marks
		}
		var current = root
		root = extract(root)
		for _, path := range roads[current.index] {
			var peer = &list[path.Next]
			if peer.index == path.Next { //针对未处理的点
				var distance = current.dist + path.Dist
				if distance < peer.dist {
					peer.link = current.index
					root = floatUp(root, peer, distance)
				}
			}
		}
		current.index = -1 //标记为已经处理
	}
	return -1, marks
}
