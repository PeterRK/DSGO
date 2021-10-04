package graph

//图通常不用节点对象表示
//对稠密图可以用邻接矩阵，对稀疏图可以用邻接表
//N维数组也是一种隐式图

type Path struct {
	Next   int
	Weight uint
}

type PathS struct {
	Next   int
	Weight int
}

type SimpleEdge struct {
	A, B int
}

type Edge struct {
	A, B   int
	Weight uint
}
