package prim

const MaxDistance = ^uint(0)

//输入邻接矩阵(0指不通)，返回最小生成树的权。
//本实现复杂度为O(V^2)。
func PlainPrim(matrix [][]uint) (sum uint, fail bool) {
	var size = len(matrix)
	sum = uint(0)
	if size < 2 {
		return sum, true
	}

	var list = make([]vertex, size)
	for i := 0; i < size; i++ {
		list[i].index, list[i].dist = i, MaxDistance
	}
	list[size-1].dist = 0

	for last := size - 1; last > 0; last-- {
		var best = 0
		for i := 0; i < last; i++ {
			var distance = matrix[list[last].index][list[i].index]
			if distance != 0 && distance < list[i].dist {
				list[i].dist = distance
			} else {
				distance = list[i].dist
			}
			if distance < list[best].dist {
				best = i
			}
		}
		if list[best].dist == MaxDistance {
			return sum, true
		}
		sum += list[best].dist
		list[best], list[last-1] = list[last-1], list[best]
	}
	return sum, false
}

func PlainPrimTree(matrix [][]uint) (tree [][]int, fail bool) {
	var size = len(matrix)
	if size < 2 {
		return tree, true
	}
	tree = make([][]int, size)

	var list = make([]vertex, size)
	for i := 0; i < size-1; i++ {
		list[i].index, list[i].dist = i+1, MaxDistance
	}
	list[size-1].index, list[size-1].dist, list[size-1].link = 0, 0, 0

	for last := size - 1; last > 0; last-- {
		var best = 0
		for i := 0; i < last; i++ {
			var distance = matrix[list[last].index][list[i].index]
			if distance != 0 && distance < list[i].dist {
				list[i].dist, list[i].link = distance, list[last].index
			} else {
				distance = list[i].dist
			}
			if distance < list[best].dist {
				best = i
			}
		}
		if list[best].dist == MaxDistance {
			return tree, true
		}
		tree[list[best].link] = append(tree[list[best].link], list[best].index)
		list[best], list[last-1] = list[last-1], list[best]
	}
	return tree, false
}
