package prim

import (
	"Graph/graph"
)

//输入邻接表，返回一个以0号节点为根的树。
//对有向图不适用，多路同权时选择有问题（不能倒着用，可能选错）。
func Prim(roads [][]graph.Path) [][]int {
	var size = len(roads)
	var result = make([][]int, size)
	if size < 2 {
		return result
	}

	var list = make([]node, size)
	for i := 1; i < size; i++ {
		list[i].id = -1
	}
	list[0].id, list[0].dist, list[0].lnk = 0, 0, 0

	var root = insert(nil, &list[0])
	for root != nil {
		var current = root
		root = extract(root)
		result[current.lnk] = append(result[current.lnk], current.id)
		current.lnk = -1
		for _, path := range roads[current.id] {
			var peer = &list[path.Next]
			if peer.id == path.Next { //已经入围的点
				if peer.lnk != -1 && //还不是内点
					path.Dist < peer.dist {
					peer.lnk = current.id
					root = floatUp(root, peer, path.Dist)
				}
			} else {
				peer.id, peer.lnk, peer.dist = path.Next, current.id, path.Dist
				root = insert(root, peer)
			}
		}
	}
	result[0] = result[0][1:]
	return result
}

func PrimX(roads [][]graph.Path) (result uint) {
	var size = len(roads)
	result = 0
	if size < 2 {
		return
	}

	var list = make([]node, size)
	for i := 1; i < size; i++ {
		list[i].id = -1
	}
	list[0].id, list[0].dist, list[0].lnk = 0, 0, 0

	var root = insert(nil, &list[0])
	for root != nil {
		var current = root
		root = extract(root)
		result += current.dist
		current.lnk = -1
		for _, path := range roads[current.id] {
			var peer = &list[path.Next]
			if peer.id == path.Next { //已经入围的点
				if peer.lnk != -1 && //还不是内点
					path.Dist < peer.dist {
					peer.lnk = current.id
					root = floatUp(root, peer, path.Dist)
				}
			} else {
				peer.id, peer.lnk, peer.dist = path.Next, current.id, path.Dist
				root = insert(root, peer)
			}
		}
	}
	return result
}
