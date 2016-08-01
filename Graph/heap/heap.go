package heap

//为图算法定制的辅助堆

const MaxDistance = ^uint(0)

type Vertex struct {
	Index int  //本顶点编号
	Link  int  //关联顶点编号
	Dist  uint //与关联顶点间的距离
}
