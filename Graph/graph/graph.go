package graph

//邻接矩阵
//棋盘

//邻接表
type Path struct {
	Next int
	Dist uint
}

type PathX struct {
	A, B int
	Dist uint
}

const MaxDistance = ^uint(0)

type Vertex struct {
	Index int  //本顶点编号
	Link  int  //关联顶点编号
	Dist  uint //与关联顶点间的距离
}
