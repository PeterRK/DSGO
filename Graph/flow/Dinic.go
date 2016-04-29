package flow

import (
	"DSGO/Graph/graph"
)

type data struct {
	origin     [][]graph.Path
	shadow     [][]graph.Path
	reflux     [][]graph.Path
	queue      arrayQueue
	stack      arrayStack
	memo       []uint
	start, end int
}

func newWorkObj(roads [][]graph.Path, start, end int) *data {
	var size = len(roads)

	var pack = new(data)
	pack.shadow = make([][]graph.Path, size) //分层残图
	pack.reflux = make([][]graph.Path, size) //逆流暂存

	var space1 = make([]int, size)  //临时空间
	var space2 = make([]uint, size) //临时空间

	pack.origin = roads
	pack.start, pack.end = start, end
	pack.queue.bind(space1)
	pack.stack.bind(space1, space2)
	pack.memo = space2

	return pack
}

//输入邻接表，返回最大流，复杂度为O(V^2 E)。
func Dinic(roads [][]graph.Path, start, end int) uint {
	var size = len(roads)
	if start < 0 || end < 0 ||
		start >= size || end >= size ||
		start == end {
		return 0
	}

	var pack = newWorkObj(roads, start, end)
	return pack.dinic()
}

func (pk *data) dinic() uint {
	var size = len(pk.origin)
	for i := 0; i < size; i++ {
		sort(pk.origin[i]) //要求有序
	}

	var flow = uint(0)
	for pk.separate() {
		flow += pk.search()
		pk.flushBack()
	}
	return flow
}

type dataM struct {
	matrix     [][]uint
	shadow     [][]graph.Path
	queue      arrayQueue
	stack      arrayStack
	memo       []uint
	start, end int
}

func newWorkObjM(matrix [][]uint, start, end int) *dataM {
	var size = len(matrix)

	var pack = new(dataM)
	pack.shadow = make([][]graph.Path, size) //分层残图

	var space1 = make([]int, size)  //临时空间
	var space2 = make([]uint, size) //临时空间

	pack.matrix = matrix
	pack.start, pack.end = start, end
	pack.queue.bind(space1)
	pack.stack.bind(space1, space2)
	pack.memo = space2

	return pack
}

//输入邻接矩阵，返回最大流，复杂度为O(V^2 E)。
func DinicM(matrix [][]uint, start, end int) uint {
	var size = len(matrix)
	if start < 0 || end < 0 ||
		start >= size || end >= size ||
		start == end {
		return 0
	}

	var pack = newWorkObjM(matrix, start, end)
	return pack.dinic()
}

func (pk *dataM) dinic() uint {
	var flow = uint(0)
	for pk.separate() {
		//由于search一次会删除图的若干条边，所以循环次数为O(E/k)
		//循环内的separate和flushBack操作的复杂度都是O(V^2)，search的复杂度为O(Vk)
		flow += pk.search()
		pk.flushBack()
	}
	return flow
}
