package span

import (
	"Graph/graph"
)

//输入邻接矩阵(0指不通)，返回最小生成树的权。
//本实现复杂度为O(V^2)。
func PlainPrim(matrix [][]uint) (sum uint, fail bool) {
	var size = len(matrix)
	sum = uint(0)
	if size < 2 {
		return 0, true
	}

	var list = make([]graph.Vertex, size)
	for i := 0; i < size; i++ {
		list[i].Index, list[i].Dist = i, graph.MaxDistance
	}
	list[size-1].Dist = 0

	for last := size - 1; last > 0; last-- {
		var best = 0
		for i := 0; i < last; i++ {
			var distance = matrix[list[last].Index][list[i].Index]
			if distance != 0 && distance < list[i].Dist {
				list[i].Dist = distance
			} else {
				distance = list[i].Dist
			}
			if distance < list[best].Dist {
				best = i
			}
		}
		if list[best].Dist == graph.MaxDistance {
			return 0, true
		}
		sum += list[best].Dist
		list[best], list[last-1] = list[last-1], list[best]
	}
	return sum, false
}

func PlainPrimTree(matrix [][]uint) (edges []Edge, fail bool) {
	var size = len(matrix)
	if size < 2 {
		return []Edge{}, true
	}
	edges = make([]Edge, 0, size-1)

	var list = make([]graph.Vertex, size)
	for i := 0; i < size-1; i++ {
		list[i].Index, list[i].Dist = i+1, graph.MaxDistance
	}
	list[size-1].Index, list[size-1].Dist, list[size-1].Link = 0, 0, 0

	for last := size - 1; last > 0; last-- {
		var best = 0
		for i := 0; i < last; i++ {
			var distance = matrix[list[last].Index][list[i].Index]
			if distance != 0 && distance < list[i].Dist {
				list[i].Dist, list[i].Link = distance, list[last].Index
			} else {
				distance = list[i].Dist
			}
			if distance < list[best].Dist {
				best = i
			}
		}
		if list[best].Dist == graph.MaxDistance {
			return []Edge{}, true
		}
		edges = append(edges, Edge{list[best].Link, list[best].Index})
		list[best], list[last-1] = list[last-1], list[best]
	}
	return edges, false
}
