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

	var space = make([]int, size)       //临时空间
	var book = make([]uint, size)       //临时空间
	var shadow = make([][]edge, size-1) //分层残图

	var flow = uint(0)
	for separate(shadow, matrix, book, space) {
		//由于search一次会删除图的若干条边，所以循环次数为O(E/k)
		//循环内的separate和flushBack操作的复杂度都是O(V^2)，search的复杂度为O(Vk)
		flow += search(shadow, matrix, book, space)
		flushBack(shadow, matrix)
	}
	return flow
}

//筛分层次，生成分层残图，复杂度为(V^2)。
func separate(shadow [][]edge, matrix [][]uint, book []uint, space []int) (ok bool) {
	var size = len(matrix)
	const FAKE_LEVEL = ^uint(0)
	for i := 1; i < size; i++ {
		book[i] = FAKE_LEVEL
	}
	book[0] = 0

	var rpt, wpt = 0, 1
	for space[0] = 0; rpt != wpt; rpt++ {
		var current = space[rpt]
		if matrix[current][size-1] != 0 {
			book[size-1] = book[current] + 1
			goto ReachGoal //到终点层
		}
		for i := 1; i < size-1; i++ {
			if book[i] == FAKE_LEVEL && matrix[current][i] != 0 {
				book[i] = book[current] + 1
				space[wpt] = i
				wpt++
			}
		}
	}
	return false

ReachGoal:
	for wpt--; wpt >= 0; wpt-- {
		var current = space[wpt]
		//shadow[current] = shadow[current][:0]
		for i := 1; i < size; i++ {
			if book[i] == book[current]+1 && matrix[current][i] != 0 {
				shadow[current] = append(shadow[current], edge{next: i, val: matrix[current][i]})
				matrix[current][i] = 0
			}
		}
		if len(shadow[current]) == 0 {
			book[current] = FAKE_LEVEL
		}
	}
	return true
}

func flushBack(shadow [][]edge, matrix [][]uint) {
	var size = len(shadow) //len(shadow) == len(matrix)-1
	for i := 1; i < size; i++ {
		if len(shadow[i]) != 0 {
			for _, path := range shadow[i] {
				matrix[i][path.next] += path.val
			}
			shadow[i] = shadow[i][:0]
		}
	}
}
