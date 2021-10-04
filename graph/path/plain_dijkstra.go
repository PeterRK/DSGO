package path

import (
	"DSGO/array"
	"math"
)

//输入邻接矩阵(0指不通)，返回某点到各点的最短路径的长度(-1指不通)。
//本实现复杂度为O(V^2)。
func PlainDijkstra(matrix [][]uint, start int) []int {
	size := len(matrix)
	if size == 0 || start < 0 || start >= size {
		return nil
	}
	result := make([]int, size)

	memo := make([]vertex, size)
	for i := 0; i < size; i++ {
		memo[i].idx, memo[i].dist = i, math.MaxUint
	}
	memo[start].idx = size - 1
	memo[size-1].idx, memo[size-1].dist = start, 0

	for last := size - 1; last > 0 &&
		memo[last].dist != math.MaxUint; last-- {
		best := 0
		for i := 0; i < last; i++ {
			step := matrix[memo[last].idx][memo[i].idx]
			dist := memo[last].dist + step
			if step != 0 && dist < memo[i].dist {
				memo[i].dist = dist
			} else {
				dist = memo[i].dist
			}
			if dist < memo[best].dist {
				best = i
			}
		}
		memo[best], memo[last-1] = memo[last-1], memo[best]
	}

	for i := 0; i < size; i++ {
		if memo[i].dist == math.MaxUint {
			result[memo[i].idx] = -1
		} else {
			result[memo[i].idx] = (int)(memo[i].dist)
		}
	}
	return result
}

//输入邻接矩阵(0指不通)，返回两点间的最短路径及其长度(-1指不通)。
func PlainDijkstraPath(matrix [][]uint, start, end int) []int {
	size := len(matrix)
	if start < 0 || end < 0 || start >= size || end >= size {
		return nil
	}
	if start == end {
		return []int{start}
	}

	memo := make([]vertex, size)
	for i := 0; i < size-1; i++ {
		memo[i].idx, memo[i].dist = i, math.MaxUint
	}
	trace := func(yard []vertex) []int {
		var path []int
		for idx := 0; idx < len(yard)-1; {
			next := yard[idx].link
			for yard[idx].idx != next {
				idx++
			}
			path = append(path, next)
		}
		array.Reverse(path)
		path = append(path, end)
		return path
	}

	memo[start].idx = size - 1
	memo[size-1].idx, memo[size-1].dist = start, 0 //第一步
	for last := size - 1; last >= 0 &&
		memo[last].dist != math.MaxUint; last-- {
		if memo[last].idx == end {
			return trace(memo[last:])
		}
		best := 0
		for i := 0; i < last; i++ {
			step := matrix[memo[last].idx][memo[i].idx]
			dist := memo[last].dist + step
			if step != 0 && dist < memo[i].dist {
				memo[i].dist, memo[i].link = dist, memo[last].idx
			} else {
				dist = memo[i].dist
			}
			if dist < memo[best].dist {
				best = i
			}
		}
		memo[best], memo[last-1] = memo[last-1], memo[best]
	}
	return nil
}
