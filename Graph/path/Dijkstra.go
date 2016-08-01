package path

import (
	"DSGO/Graph/graph"
	"DSGO/Graph/heap"
)

//输入邻接表，返回某点到各点的最短路径的长度(-1指不通)。
//复杂度为O(E+VlogV)，已知最快的单源最短路径算法，对稀疏图尤甚。
//可以处理有向图，不能处理负权边。
func Dijkstra(roads [][]graph.Path, start int) []int {
	var size = len(roads)
	if size == 0 || start < 0 || start >= size {
		return nil
	}

	var result = make([]int, size)
	if size == 1 {
		result[0] = 0
		return result
	}

	const FAKE = -1
	var list = heap.NewVectorPH(size)
	for i := 0; i < size; i++ {
		list[i].Link = FAKE
	}
	list[start].Index, list[start].Link, list[start].Dist = start, start, 0
	var root = heap.Insert(nil, &list[start])

	for root != nil {
		var index, dist = root.Index, root.Dist
		root.Index, root = FAKE, heap.Extract(root) //入围
		for _, path := range roads[index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE { //未涉及点
				peer.Index, peer.Link = path.Next, index
				peer.Dist = dist + path.Weight
				root = heap.Insert(root, peer)
			} else if peer.Index != FAKE { //外围点
				var distance = dist + path.Weight
				if distance < peer.Dist {
					root = heap.FloatUp(root, peer, distance)
					//peer.Link = index
				}
			}
		}
	}

	for i := 0; i < size; i++ {
		if list[i].Dist == heap.MaxDistance {
			result[i] = -1
		} else {
			result[i] = (int)(list[i].Dist)
		}
	}
	return result
}

//输入邻接表，返回两点间的最短路径及其长度(-1指不通)。
func DijkstraPath(roads [][]graph.Path, start, end int) []int {
	var size = len(roads)
	if start < 0 || end < 0 || start >= size || end >= size {
		return nil
	}
	if start == end {
		return []int{start}
	}

	const FAKE = -1
	var list = heap.NewVectorPH(size)
	for i := 0; i < size; i++ {
		list[i].Link = FAKE
	}
	var trace = func() []int {
		var path []int
		for idx := end; idx != start; idx = list[idx].Link {
			path = append(path, idx)
		}
		path = append(path, start)
		reverse(path)
		return path
	}

	list[start].Index, list[start].Link, list[start].Dist = start, start, 0
	var root = heap.Insert(nil, &list[start]) //第一步
	for root != nil {
		var index, dist = root.Index, root.Dist
		if index == end {
			return trace()
		}
		root.Index, root = FAKE, heap.Extract(root) //入围
		for _, path := range roads[index] {
			var peer = &list[path.Next]
			if peer.Link == FAKE { //未涉及点
				peer.Index, peer.Link = path.Next, index
				peer.Dist = dist + path.Weight
				root = heap.Insert(root, peer)
			} else if peer.Index != FAKE { //外围点
				var distance = dist + path.Weight
				if distance < peer.Dist {
					root = heap.FloatUp(root, peer, distance)
					peer.Link = index
				}
			}
		}
	}
	return nil
}

func reverse(list []int) {
	for left, right := 0, len(list)-1; left < right; {
		list[left], list[right] = list[right], list[left]
		left++
		right--
	}
}
