package flow

type edge struct {
	next int
	val  uint
}

//输入邻接矩阵，返回头节点到尾顶点见的最大流。
//复杂度为O(V^2 E)。
func Dinic(matrix [][]uint) uint {
	var size = len(matrix)
	if size == 0 {
		return 0
	}

	var space1 = make([]int, size)  //临时空间
	var space2 = make([]uint, size) //临时空间
	var q queue
	q.bind(space1)
	var s stack
	s.bind(space1, space2)

	var shadow = make([][]edge, size-1) //分层残图

	var flow = uint(0)
	for separate(shadow, matrix, &q, space2) {
		//由于search一次会删除图的若干条边，所以循环次数为O(E/k)
		//循环内的separate和flushBack操作的复杂度都是O(V^2)，search的复杂度为O(Vk)
		flow += search(shadow, matrix, &s)
		flushBack(shadow, matrix)
	}
	return flow
}
