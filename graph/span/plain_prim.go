package span

import (
	"DSGO/graph"
	"errors"
	"math"
)

//输入邻接矩阵(0指不通)，返回最小生成树的权。
//本实现复杂度为O(V^2)。
func PlainPrim(matrix [][]uint) (uint, error) {
	size := len(matrix)
	sum := uint(0)
	if size < 2 {
		return 0, errors.New("illegal input")
	}

	memo := make([]vertex, size)
	for i := 0; i < size; i++ {
		memo[i].idx, memo[i].dist = i, math.MaxUint
	}
	memo[size-1].dist = 0

	for last := size - 1; last > 0; last-- {
		best := 0
		for i := 0; i < last; i++ {
			dist := matrix[memo[last].idx][memo[i].idx]
			if dist != 0 && dist < memo[i].dist {
				memo[i].dist = dist
			} else {
				dist = memo[i].dist
			}
			if dist < memo[best].dist {
				best = i
			}
		}
		if memo[best].dist == math.MaxUint {
			return 0, errors.New("isolated part exist")
		}
		sum += memo[best].dist
		memo[best], memo[last-1] = memo[last-1], memo[best]
	}
	return sum, nil
}

func PlainPrimTree(matrix [][]uint) ([]graph.SimpleEdge, error) {
	size := len(matrix)
	if size < 2 {
		return nil, errors.New("illegal input")
	}
	edges := make([]graph.SimpleEdge, 0, size-1)

	memo := make([]vertex, size)
	for i := 0; i < size-1; i++ {
		memo[i].idx, memo[i].dist = i+1, math.MaxUint
	}
	memo[size-1] = vertex{idx: 0, link: 0, dist: 0}

	for last := size - 1; last > 0; last-- {
		best := 0
		for i := 0; i < last; i++ {
			dist := matrix[memo[last].idx][memo[i].idx]
			if dist != 0 && dist < memo[i].dist {
				memo[i].dist, memo[i].link = dist, memo[last].idx
			} else {
				dist = memo[i].dist
			}
			if dist < memo[best].dist {
				best = i
			}
		}
		if memo[best].dist == math.MaxUint {
			return nil, errors.New("isolated part exist")
		}
		edges = append(edges, graph.SimpleEdge{memo[best].link, memo[best].idx})
		memo[best], memo[last-1] = memo[last-1], memo[best]
	}
	return edges, nil
}
