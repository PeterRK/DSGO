package span

import (
	"Graph/graph"
	"errors"
)

//输入邻接矩阵(0指不通)，返回最小生成树的权。
//本实现复杂度为O(V^2)。
func PlainPrim(matrix [][]uint) (uint, error) {
	var size = len(matrix)
	var sum = uint(0)
	if size < 2 {
		return 0, errors.New("illegal input")
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
			return 0, errors.New("isolated part exist")
		}
		sum += list[best].Dist
		list[best], list[last-1] = list[last-1], list[best]
	}
	return sum, nil
}

func PlainPrimTree(matrix [][]uint) ([]Edge, error) {
	var size = len(matrix)
	if size < 2 {
		return []Edge{}, errors.New("illegal input")
	}
	var edges = make([]Edge, 0, size-1)

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
			return []Edge{}, errors.New("isolated part exist")
		}
		edges = append(edges, Edge{list[best].Link, list[best].Index})
		list[best], list[last-1] = list[last-1], list[best]
	}
	return edges, nil
}
