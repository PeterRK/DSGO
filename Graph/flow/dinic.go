package flow

type data struct {
	origin     [][]Path
	shadow     [][]Path
	reflux     [][]Path
	queue      arrayQueue
	stack      arrayStack
	memo       []uint
	start, end int
}

func newWorkObj(roads [][]Path, start, end int) *data {
	size := len(roads)

	pack := new(data)
	pack.shadow = make([][]Path, size) //分层残图
	pack.reflux = make([][]Path, size) //逆流暂存

	space1 := make([]int, size)  //临时空间
	space2 := make([]uint, size) //临时空间

	pack.origin = roads
	pack.start, pack.end = start, end
	pack.queue.bind(space1)
	pack.stack.bind(space1, space2)
	pack.memo = space2

	return pack
}

//输入邻接表，返回最大流，复杂度为O(V^2 E)。
func Dinic(roads [][]Path, start, end int) uint {
	size := len(roads)
	if start < 0 || end < 0 ||
		start >= size || end >= size ||
		start == end {
		return 0
	}

	pack := newWorkObj(roads, start, end)
	return pack.dinic()
}

func (pk *data) dinic() uint {
	size := len(pk.origin)
	for i := 0; i < size; i++ {
		sort(pk.origin[i]) //要求有序
	}

	flow := uint(0)
	for pk.separate() {
		flow += pk.search()
		pk.flushBack()
	}
	return flow
}

type dataM struct {
	matrix     [][]uint
	shadow     [][]Path
	queue      arrayQueue
	stack      arrayStack
	memo       []uint
	start, end int
}

func newWorkObjM(matrix [][]uint, start, end int) *dataM {
	size := len(matrix)

	pack := new(dataM)
	pack.shadow = make([][]Path, size) //分层残图

	space1 := make([]int, size)  //临时空间
	space2 := make([]uint, size) //临时空间

	pack.matrix = matrix
	pack.start, pack.end = start, end
	pack.queue.bind(space1)
	pack.stack.bind(space1, space2)
	pack.memo = space2

	return pack
}

//输入邻接矩阵，返回最大流，复杂度为O(V^2 E)。
func DinicM(matrix [][]uint, start, end int) uint {
	size := len(matrix)
	if start < 0 || end < 0 ||
		start >= size || end >= size ||
		start == end {
		return 0
	}

	pack := newWorkObjM(matrix, start, end)
	return pack.dinic()
}

func (pk *dataM) dinic() uint {
	flow := uint(0)
	for pk.separate() {
		//由于search一次会删除图的若干条边，所以循环次数为O(E/k)
		//循环内的separate和flushBack操作的复杂度都是O(V^2)，search的复杂度为O(Vk)
		flow += pk.search()
		pk.flushBack()
	}
	return flow
}
