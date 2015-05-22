package path

import (
	"Graph/graph"
)

//输入邻接矩阵(0指不通)，返回某点到各点的最短路径的长度(-1指不通)。
//本实现复杂度为O(V^2)。
func PlainDijkstra(matrix [][]uint, start int) []int {
	var size = len(matrix)
	if size == 0 || start < 0 || start >= size {
		return []int{}
	}
	var result = make([]int, size)

	var list = make([]graph.Vertex, size)
	for i := 0; i < size-1; i++ {
		list[i].Index, list[i].Dist = i, graph.MaxDistance
	}
	list[start].Index = size - 1
	list[size-1].Index, list[size-1].Dist = start, 0

	for last := size - 1; last > 0 && list[last].Dist != graph.MaxDistance; last-- {
		var best = 0
		for i := 0; i < last; i++ {
			var step = matrix[list[last].Index][list[i].Index]
			var distance = list[last].Dist + step
			if step != 0 && distance < list[i].Dist {
				list[i].Dist = distance
			} else {
				distance = list[i].Dist
			}
			if distance < list[best].Dist {
				best = i
			}
		}
		list[best], list[last-1] = list[last-1], list[best]
	}

	for i := 0; i < size; i++ {
		if list[i].Dist == graph.MaxDistance {
			result[list[i].Index] = -1
		} else {
			result[list[i].Index] = (int)(list[i].Dist)
		}
	}
	return result
}

//输入邻接矩阵(0指不通)，返回两点间的最短路径及其长度(-1指不通)。
func PlainDijkstraPath(matrix [][]uint, start int, end int) (Dist int, marks []int) {
	var size = len(matrix)
	if start < 0 || end < 0 || start >= size || end >= size {
		return -1, []int{}
	}
	if start == end {
		return 0, []int{start}
	}

	var list = make([]graph.Vertex, size)
	for i := 0; i < size-1; i++ {
		list[i].Index, list[i].Dist = i, graph.MaxDistance
	}
	list[start].Index = size - 1
	list[size-1].Index, list[size-1].Dist = start, 0

	for last := size - 1; last >= 0 && list[last].Dist != graph.MaxDistance; last-- {
		if list[last].Index == end {
			for idx := last; idx < size-1; {
				var next = list[idx].Link
				for list[idx].Index != next {
					idx++
				}
				marks = append(marks, next)
			}
			for left, right := 0, len(marks)-1; left < right; {
				marks[left], marks[right] = marks[right], marks[left]
				left++
				right--
			}
			marks = append(marks, end)
			return (int)(list[last].Dist), marks
		}
		var best = 0
		for i := 0; i < last; i++ {
			var step = matrix[list[last].Index][list[i].Index]
			var distance = list[last].Dist + step
			if step != 0 && distance < list[i].Dist {
				list[i].Dist, list[i].Link = distance, list[last].Index
			} else {
				distance = list[i].Dist
			}
			if distance < list[best].Dist {
				best = i
			}
		}
		list[best], list[last-1] = list[last-1], list[best]
	}
	return -1, []int{}
}
