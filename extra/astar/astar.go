package astar

type Path struct {
	Next int
	Dist uint
}

const MaxDistance = ^uint(0)

//这是个启发式算法，有猜测成分，不一定能得到最优解
func AStar(roads [][]Path, start int, end int,
	evaluate func(int) uint) []int {
	var size = len(roads)
	if start < 0 || end < 0 || start >= size || end >= size {
		return []int{}
	}
	if start == end {
		return []int{start}
	}

	const FAKE = -1
	var list = make([]node, size)
	for i := 0; i < size; i++ {
		list[i].link = FAKE
	}
	var trace = func() []int {
		var path []int
		for idx := end; idx != start; idx = list[idx].link {
			path = append(path, idx)
		}
		path = append(path, start)
		reverse(path)
		return path
	}

	list[start].index, list[start].link = start, start
	list[start].dist, list[start].weight = 0, evaluate(start)
	var root = insert(nil, &list[start])
	for root != nil && root.dist != MaxDistance {
		var index, dist = root.index, root.dist
		if index == end {
			return trace()
		}
		root.index, root = FAKE, extract(root) //入围

		for _, path := range roads[index] {
			var peer = &list[path.Next]
			if peer.link == FAKE { //未涉及点
				peer.index, peer.link = path.Next, index
				peer.dist = dist + path.Dist
				peer.weight = peer.dist + evaluate(peer.index)
				root = insert(root, peer)
			} else if peer.index != FAKE { //外围点
				var distance = dist + path.Dist
				if distance < peer.dist {
					root = floatUp(root, peer, peer.dist-distance)
					peer.dist = distance
					peer.link = index
				}
			}
		}
	}
	return []int{}
}
func reverse(list []int) {
	for left, right := 0, len(list)-1; left < right; {
		list[left], list[right] = list[right], list[left]
		left++
		right--
	}
}
