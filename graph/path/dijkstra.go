package path

import (
	bheap "DSGO/Heap/binary"
	pheap "DSGO/Heap/pairing"
	"DSGO/array"
	"DSGO/graph"
	"math"
)

type vertex struct {
	idx  int  //本顶点编号
	link int  //关联顶点编号
	dist uint //与关联顶点间的距离
}

func nearer(a, b *vertex) bool {
	return a.dist < b.dist
}

//输入邻接表，返回某点到各点的最短路径的长度(-1指不通)。
//复杂度为O(E+VlogV)，已知最快的单源最短路径算法，对稀疏图尤甚。
//可以处理有向图，不能处理负权边。
func Dijkstra(roads [][]graph.Path, start int) []int {
	size := len(roads)
	if size == 0 || start < 0 || start >= size {
		return nil
	}

	result := make([]int, size)
	if size == 1 {
		result[0] = 0
		return result
	}

	const fakeIdx = -1
	nodes := make([]pheap.NodeG[vertex], size)
	for i := 0; i < size; i++ {
		nodes[i].Val = vertex{idx: i, link: fakeIdx, dist: math.MaxUint}
	}

	nodes[start].Val = vertex{idx: start, link: start, dist: 0}
	hp := pheap.New(nearer)
	hp.Push(&nodes[start])

	for !hp.IsEmpty() {
		curr := hp.Pop()
		idx := curr.Val.idx
		curr.Val.idx = fakeIdx //入围
		for _, path := range roads[idx] {
			peer := &nodes[path.Next]
			if peer.Val.link == fakeIdx { //未涉及点
				peer.Val.link = idx
				peer.Val.dist = curr.Val.dist + path.Weight
				hp.Push(peer)
			} else if peer.Val.idx != fakeIdx { //外围点
				dist := curr.Val.dist + path.Weight
				if dist < peer.Val.dist {
					peer.Val.link = idx
					peer.Val.dist = dist
					hp.FloatUp(peer)
				}
			}
		}
	}

	for i := 0; i < size; i++ {
		if nodes[i].Val.dist == math.MaxUint {
			result[i] = -1
		} else {
			result[i] = (int)(nodes[i].Val.dist)
		}
	}
	return result
}

//输入邻接表，返回两点间的最短路径及其长度(-1指不通)。
func DijkstraPath(roads [][]graph.Path, start, end int) []int {
	size := len(roads)
	if start < 0 || end < 0 || start >= size || end >= size {
		return nil
	}
	if start == end {
		return []int{start}
	}

	const fakeIdx = -1
	nodes := make([]bheap.NodeG[vertex], size)
	for i := 0; i < size; i++ {
		nodes[i].Val = vertex{idx: i, link: fakeIdx, dist: math.MaxUint}
	}

	nodes[start].Val = vertex{idx: start, link: start, dist: 0}
	hp := bheap.New(nearer)
	hp.Push(&nodes[start])

	for !hp.IsEmpty() {
		curr := hp.Top()
		idx := curr.Val.idx
		if idx == end {
			var path []int
			for idx := end; idx != start; idx = nodes[idx].Val.link {
				path = append(path, idx)
			}
			path = append(path, start)
			array.Reverse(path)
			return path
		}
		curr.Val.idx = fakeIdx
		hp.Pop()
		for _, path := range roads[idx] {
			peer := &nodes[path.Next]
			if peer.Val.link == fakeIdx { //未涉及点
				peer.Val.link = idx
				peer.Val.dist = curr.Val.dist + path.Weight
				hp.Push(peer)
			} else if peer.Val.idx != fakeIdx { //外围点
				dist := curr.Val.dist + path.Weight
				if dist < peer.Val.dist {
					peer.Val.link = idx
					peer.Val.dist = dist
					hp.FloatUp(peer)
				}
			}
		}
	}
	return nil
}
