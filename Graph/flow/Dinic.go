package flow

import (
	"Graph/graph"
)

type data struct {
	origin [][]graph.Path
	shadow [][]graph.Path
	reflux [][]graph.Path
	queue  arrayQueue
	stack  arrayStack
	memo   []uint
	size   int
	start  int
	end    int
}

//输入邻接表，返回最大流，复杂度为O(V^2 E)。
func Dinic(roads [][]graph.Path, start int, end int) uint {
	var size = len(roads)
	if start < 0 || end < 0 ||
		start >= size || end >= size ||
		start == end {
		return 0
	}

	var space1 = make([]int, size)  //临时空间
	var space2 = make([]uint, size) //临时空间

	var pack = data{
		shadow: make([][]graph.Path, size), //分层残图
		reflux: make([][]graph.Path, size), //逆流暂存
		origin: roads, size: size,
		start: start, end: end}
	pack.queue.bind(space1)
	pack.stack.bind(space1, space2)
	pack.memo = space2

	for i := 0; i < size; i++ {
		sort(roads[i]) //要求有序
	}

	var flow = uint(0)
	for pack.separate() {
		flow += pack.search()
		pack.flushBack()
	}
	return flow
}

type dataM struct {
	matrix [][]uint
	shadow [][]graph.Path
	queue  arrayQueue
	stack  arrayStack
	memo   []uint
	size   int
	start  int
	end    int
}

//输入邻接矩阵，返回最大流，复杂度为O(V^2 E)。
func DinicM(matrix [][]uint, start int, end int) uint {
	var size = len(matrix)
	if start < 0 || end < 0 ||
		start >= size || end >= size ||
		start == end {
		return 0
	}

	var space1 = make([]int, size)  //临时空间
	var space2 = make([]uint, size) //临时空间

	var pack = dataM{
		shadow: make([][]graph.Path, size), //分层残图
		matrix: matrix, size: size,
		start: start, end: end}
	pack.queue.bind(space1)
	pack.stack.bind(space1, space2)
	pack.memo = space2

	var flow = uint(0)
	for pack.separate() {
		//由于search一次会删除图的若干条边，所以循环次数为O(E/k)
		//循环内的separate和flushBack操作的复杂度都是O(V^2)，search的复杂度为O(Vk)
		flow += pack.search()
		pack.flushBack()
	}
	return flow
}
