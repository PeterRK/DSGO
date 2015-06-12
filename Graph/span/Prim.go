package tree

import (
	"Graph/graph"
)

//输入邻接表，返回最小生成树的权。
//复杂度为O(E+VlogV)，通常比Kruskal强。
//对有向图不适用，多路同权时选择有问题（不能倒着用，可能选错）。
func Prim(roads [][]graph.Path) (sum uint, fail bool) {
	var size = len(roads)
	sum = uint(0)
	if size < 2 {
		return 0, true
	}

	var list = graph.NewVector(size)
	for i := 1; i < size; i++ {
		list[i].Index = -1
	}
	list[0].Index, list[0].Dist, list[0].Link = 0, 0, 0

	var cnt int
	var root = graph.Insert(nil, &list[0])
	for cnt = 0; root != nil; cnt++ {
		var current = root
		root = graph.Extract(root)
		sum += current.Dist
		current.Link = -1
		for _, path := range roads[current.Index] {
			var peer = &list[path.Next]
			if peer.Index == path.Next { //已经入围的点
				if peer.Link != -1 && //还不是内点
					path.Dist < peer.Dist {
					peer.Link = current.Index
					root = graph.FloatUp(root, peer, path.Dist)
				}
			} else {
				peer.Index, peer.Link, peer.Dist = path.Next, current.Index, path.Dist
				root = graph.Insert(root, peer)
			}
		}
	}
	return sum, cnt != size
}

//输入邻接表，返回一个以0号节点为根的最小生成树。
func PrimTree(roads [][]graph.Path) (tree [][]int, fail bool) {
	var size = len(roads)
	if size < 2 {
		return [][]int{}, true
	}
	tree = make([][]int, size)

	var list = graph.NewVector(size)
	for i := 1; i < size; i++ {
		list[i].Index = -1
	}
	list[0].Index, list[0].Dist, list[0].Link = 0, 0, 0

	var cnt int
	var root = graph.Insert(nil, &list[0])
	for cnt = 0; root != nil; cnt++ {
		var current = root
		root = graph.Extract(root)
		tree[current.Link] = append(tree[current.Link], current.Index)
		current.Link = -1
		for _, path := range roads[current.Index] {
			var peer = &list[path.Next]
			if peer.Index == path.Next { //已经入围的点
				if peer.Link != -1 && //还不是内点
					path.Dist < peer.Dist {
					peer.Link = current.Index
					root = graph.FloatUp(root, peer, path.Dist)
				}
			} else {
				peer.Index, peer.Link, peer.Dist = path.Next, current.Index, path.Dist
				root = graph.Insert(root, peer)
			}
		}
	}
	tree[0] = tree[0][1:]
	return tree, cnt != size
}
