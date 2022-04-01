package span

import (
	bheap "DSGO/heap/binary"
	pheap "DSGO/heap/pairing"
	"DSGO/graph"
	"errors"
)

type vertex struct {
	idx  int  //本顶点编号
	link int  //关联顶点编号
	dist uint //与关联顶点间的距离
}

func nearer(a, b *vertex) bool {
	return a.dist < b.dist
}

//输入邻接表，返回最小生成树的权。
//复杂度为O(E+VlogV)，通常比Kruskal强。
//对有向图不适用，多路同权时选择有问题（不能倒着用，可能选错）。
func Prim(roads [][]graph.Path) (uint, error) {
	size := len(roads)
	sum := uint(0)
	if size < 2 {
		return 0, errors.New("illegal input")
	}

	const fakeIdx = -1
	nodes := make([]pheap.NodeG[vertex], size)
	for i := 1; i < size; i++ {
		nodes[i].Val = vertex{idx: i, link: fakeIdx, dist: 0}
	}
	nodes[0].Val = vertex{idx: 0, link: 0, dist: 0}
	hp := pheap.New(nearer)
	hp.Push(&nodes[0])

	var cnt int
	for cnt = 0; !hp.IsEmpty(); cnt++ {
		curr := hp.Pop()
		sum += curr.Val.dist
		idx := curr.Val.idx
		curr.Val.idx = fakeIdx //入围
		for _, path := range roads[idx] {
			peer := &nodes[path.Next]
			if peer.Val.link == fakeIdx { //未涉及点
				peer.Val.link = idx
				peer.Val.dist = path.Weight
				hp.Push(peer)
			} else if peer.Val.idx != fakeIdx && //外围点
				path.Weight < peer.Val.dist { //可更新
				peer.Val.link = idx
				peer.Val.dist = path.Weight
				hp.FloatUp(peer)
			}
		}
	}
	if cnt != size {
		return sum, errors.New("isolated part exist")
	}
	return sum, nil
}

//输入邻接表，返回一个以0号节点为根的最小生成树。
func PrimTree(roads [][]graph.Path) ([]graph.SimpleEdge, error) {
	size := len(roads)
	if size < 2 {
		return nil, errors.New("illegal input")
	}
	edges := make([]graph.SimpleEdge, 0, size-1)

	const fakeIdx = -1
	nodes := make([]bheap.NodeG[vertex], size)
	for i := 1; i < size; i++ {
		nodes[i].Val = vertex{idx: i, link: fakeIdx, dist: 0}
	}
	nodes[0].Val = vertex{idx: 0, link: 0, dist: 0}
	hp := bheap.New(nearer)
	hp.Push(&nodes[0])

	for {
		curr := hp.Pop()
		idx := curr.Val.idx
		curr.Val.idx = fakeIdx //入围
		for _, path := range roads[idx] {
			peer := &nodes[path.Next]
			if peer.Val.link == fakeIdx { //未涉及点
				peer.Val.link = idx
				peer.Val.dist = path.Weight
				hp.Push(peer)
			} else if peer.Val.idx != fakeIdx && //外围点
				path.Weight < peer.Val.dist { //可更新
				peer.Val.link = idx
				peer.Val.dist = path.Weight
				hp.FloatUp(peer)
			}
		}
		if hp.IsEmpty() {
			break
		}
		curr = hp.Top()
		edges = append(edges, graph.SimpleEdge{A: curr.Val.link, B: curr.Val.idx})
	}
	if len(edges) != size-1 {
		return edges, errors.New("isolated part exist")
	}
	return edges, nil
}
