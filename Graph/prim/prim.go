package prim

import (
	"graph"
)

//Prim对有向图不适用，多路同权时选择有问题（不能倒着用，可能选错）
//返回一个以0号节点为根的树
func Prim(roads [][]graph.Path) [][]int {
	var size = len(roads)
	var result = make([][]int, size)
	if size < 2 {
		return result
	}

	var list = make([]Node, size)
	for i := 1; i < size; i++ {
		list[i].id = -1
	}
	list[0].id, list[0].dist, list[0].next = 0, 0, 0

	var root = Insert(nil, &list[0])
	for root != nil {
		var current = root
		root = Extract(root)
		result[current.next] = append(result[current.next], current.id)
		current.next = -1
		for _, path := range roads[current.id] {
			var peer = &list[path.Next]
			if peer.id == path.Next { //已经入围的点
				if peer.next != -1 && //还不是内点
					path.Dist < peer.dist {
					peer.next = current.id
					root = FloatUp(root, peer, path.Dist)
				}
			} else {
				peer.id, peer.next, peer.dist = path.Next, current.id, path.Dist
				root = Insert(root, peer)
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

	var list = make([]Node, size)
	for i := 1; i < size; i++ {
		list[i].id = -1
	}
	list[0].id, list[0].dist, list[0].next = 0, 0, 0

	var root = Insert(nil, &list[0])
	for root != nil {
		var current = root
		root = Extract(root)
		result += current.dist
		current.next = -1
		for _, path := range roads[current.id] {
			var peer = &list[path.Next]
			if peer.id == path.Next { //已经入围的点
				if peer.next != -1 && //还不是内点
					path.Dist < peer.dist {
					peer.next = current.id
					root = FloatUp(root, peer, path.Dist)
				}
			} else {
				peer.id, peer.next, peer.dist = path.Next, current.id, path.Dist
				root = Insert(root, peer)
			}
		}
	}
	return result
}
