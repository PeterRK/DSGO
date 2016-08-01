package span

import (
	"DSGO/Graph/graph"
	"DSGO/Graph/heap"
	"errors"
)

//输入邻接表，返回最小生成树的权。
//复杂度为O(E+VlogV)，通常比Kruskal强。
//对有向图不适用，多路同权时选择有问题（不能倒着用，可能选错）。
func Prim(roads [][]graph.Path) (uint, error) {
	var size = len(roads)
	var sum = uint(0)
	if size < 2 {
		return 0, errors.New("illegal input")
	}

	const FAKE = -1
	var list = heap.NewVectorPH(size)
	for i := 1; i < size; i++ {
		list[i].Link = FAKE
	}
	list[0].Index, list[0].Link, list[0].Dist = 0, 0, 0
	var root = heap.Insert(nil, &list[0])

	var cnt int
	for cnt = 0; root != nil; cnt++ {
		sum += root.Dist
		var index = root.Index
		root.Index, root = FAKE, heap.Extract(root) //入围
		for _, path := range roads[index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE { //未涉及点
				peer.Index, peer.Link, peer.Dist = path.Next, index, path.Weight
				root = heap.Insert(root, peer)
			} else if peer.Index != FAKE && //外围点
				path.Weight < peer.Dist { //可更新
				root = heap.FloatUp(root, peer, path.Weight)
				peer.Link = index
			}
		}
	}
	if cnt != size {
		return sum, errors.New("isolated part exist")
	}
	return sum, nil
}

type Edge struct {
	A, B int
}

//输入邻接表，返回一个以0号节点为根的最小生成树。
func PrimTree(roads [][]graph.Path) ([]Edge, error) {
	var size = len(roads)
	if size < 2 {
		return nil, errors.New("illegal input")
	}
	var edges = make([]Edge, 0, size-1)

	const FAKE = -1
	var list = heap.NewVectorPH(size)
	for i := 1; i < size; i++ {
		list[i].Link = FAKE
	}
	list[0].Index, list[0].Link, list[0].Dist = 0, 0, 0
	var root = heap.Insert(nil, &list[0])

	for {
		var index = root.Index
		root.Index, root = FAKE, heap.Extract(root) //入围
		for _, path := range roads[index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE { //未涉及点
				peer.Index, peer.Link, peer.Dist = path.Next, index, path.Weight
				root = heap.Insert(root, peer)
			} else if peer.Index != FAKE && //外围点
				path.Weight < peer.Dist { //可更新
				peer.Link = index
				root = heap.FloatUp(root, peer, path.Weight)
			}
		}
		if root == nil {
			break
		}
		edges = append(edges, Edge{root.Link, root.Index})
	}
	if len(edges) != size-1 {
		return edges, errors.New("isolated part exist")
	}
	return edges, nil
}
