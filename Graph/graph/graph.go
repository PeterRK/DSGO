package graph

//图通常不用节点对象表示
//对稠密图可以用邻接矩阵，对稀疏图可以用邻接表
//N维数组也是一种隐式图

type Path struct {
	Next   int
	Weight uint
}
type Edge struct {
	A, B   int
	Weight uint
}

const (
	MarkO = iota
	MarkX
	MarkV
)

//以下函数从某起点深度优先遍历二维数组形式的图（MarkX表示不通，MarkO表示通而未达，MarkV表示已达）
func DFS(table [][]int, x int, y int) {
	//defer func() { recover() }()
	const (
		LEFT = iota
		RIGHT
		UP
		DOWN
	)
	var length, width = len(table), len(table[0])
	var path []int
	for {
		if x >= 0 && x < length &&
			y >= 0 && y < width &&
			table[x][y] == MarkO {
			table[x][y] = MarkV
			y--
			path = append(path, LEFT)
		} else {
		Label_LOOP:
			var last = len(path) - 1
			if last < 0 {
				return
			}
			switch path[last] {
			case LEFT:
				y += 2
				path[last] = RIGHT
			case RIGHT:
				y--
				x--
				path[last] = UP
			case UP:
				x += 2
				path[last] = DOWN
			case DOWN:
				x--
				path = path[:last]
				goto Label_LOOP
			}
		}
	}
}
