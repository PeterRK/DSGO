package prim

import (
	"Graph/graph"
)

//输入邻接表，返回最小生成树的权。
//与本实现复杂度为O(ElogV)，不比Kruskal差。
//对有向图不适用，多路同权时选择有问题（不能倒着用，可能选错）。
func Prim(roads [][]graph.Path) (sum uint, fail bool) {
	var size = len(roads)
	sum = uint(0)
	if size < 2 {
		return sum, true
	}

	var list = make([]node, size)
	for i := 1; i < size; i++ {
		list[i].index = -1
	}
	list[0].index, list[0].dist, list[0].link = 0, 0, 0

	var cnt int
	var root = insert(nil, &list[0])
	for cnt = 0; root != nil; cnt++ {
		var current = root
		root = extract(root)
		sum += current.dist
		current.link = -1
		for _, path := range roads[current.index] {
			var peer = &list[path.Next]
			if peer.index == path.Next { //已经入围的点
				if peer.link != -1 && //还不是内点
					path.Dist < peer.dist {
					peer.link = current.index
					root = floatUp(root, peer, path.Dist)
				}
			} else {
				peer.index, peer.link, peer.dist = path.Next, current.index, path.Dist
				root = insert(root, peer)
			}
		}
	}
	return sum, cnt != size
}

//输入邻接表，返回一个以0号节点为根的最小生成树。
func PrimTree(roads [][]graph.Path) (tree [][]int, fail bool) {
	var size = len(roads)
	if size < 2 {
		return tree, true
	}
	tree = make([][]int, size)

	var list = make([]node, size)
	for i := 1; i < size; i++ {
		list[i].index = -1
	}
	list[0].index, list[0].dist, list[0].link = 0, 0, 0

	var cnt int
	var root = insert(nil, &list[0])
	for cnt = 0; root != nil; cnt++ {
		var current = root
		root = extract(root)
		tree[current.link] = append(tree[current.link], current.index)
		current.link = -1
		for _, path := range roads[current.index] {
			var peer = &list[path.Next]
			if peer.index == path.Next { //已经入围的点
				if peer.link != -1 && //还不是内点
					path.Dist < peer.dist {
					peer.link = current.index
					root = floatUp(root, peer, path.Dist)
				}
			} else {
				peer.index, peer.link, peer.dist = path.Next, current.index, path.Dist
				root = insert(root, peer)
			}
		}
	}
	tree[0] = tree[0][1:]
	return tree, cnt != size
}
