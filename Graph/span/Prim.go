package span

import (
	"Graph/graph"
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
	var list = graph.NewVector(size)
	for i := 1; i < size; i++ {
		list[i].Link = FAKE
	}
	list[0].Index, list[0].Link, list[0].Dist = 0, 0, 0
	var root = graph.Insert(nil, &list[0])

	var cnt int
	for cnt = 0; root != nil; cnt++ {
		var current = root
		root = graph.Extract(root)
		sum += current.Dist
		for _, path := range roads[current.Index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE { //未涉及点
				peer.Index, peer.Link, peer.Dist = path.Next, current.Index, path.Dist
				root = graph.Insert(root, peer)
			} else if peer.Index != FAKE && //外围点
				path.Dist < peer.Dist { //可更新
				peer.Link = current.Index
				root = graph.FloatUp(root, peer, path.Dist)
			}
		}
		current.Index = FAKE //入围
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
		return []Edge{}, errors.New("illegal input")
	}
	var edges = make([]Edge, 0, size-1)

	const FAKE = -1
	var list = graph.NewVector(size)
	for i := 1; i < size; i++ {
		list[i].Link = FAKE
	}
	list[0].Index, list[0].Link, list[0].Dist = 0, 0, 0
	var root = graph.Insert(nil, &list[0])

	for {
		var current = root
		root = graph.Extract(root)
		for _, path := range roads[current.Index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE { //未涉及点
				peer.Index, peer.Link, peer.Dist = path.Next, current.Index, path.Dist
				root = graph.Insert(root, peer)
			} else if peer.Index != FAKE && //外围点
				path.Dist < peer.Dist { //可更新
				peer.Link = current.Index
				root = graph.FloatUp(root, peer, path.Dist)
			}
		}
		current.Index = FAKE //入围
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
