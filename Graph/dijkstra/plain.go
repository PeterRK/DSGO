package dijkstra

//输入邻接矩阵(0指不通)，返回某点到各点的最短路径的长度(-1指不通)。
//本实现复杂度为O(V^2)。
func PlainDijkstra(matrix [][]uint, start int) []int {
	var size = len(matrix)
	if size == 0 || start < 0 || start >= size {
		return []int{}
	}
	var result = make([]int, size)

	var list = make([]vertex, size)
	for i := 0; i < size-1; i++ {
		list[i].index, list[i].dist = i, MaxDistance
	}
	list[start].index = size - 1
	list[size-1].index, list[size-1].dist = start, 0

	for last := size - 1; last > 0 && list[last].dist != MaxDistance; last-- {
		var best = 0
		for i := 0; i < last; i++ {
			var step = matrix[list[last].index][list[i].index]
			var distance uint
			if step != 0 { //通
				distance = list[last].dist + step
				if distance < list[i].dist {
					list[i].dist = distance
				} else {
					distance = list[i].dist
				}
			} else { //不通
				distance = list[i].dist
			}
			if distance < list[best].dist {
				best = i
			}
		}
		list[best], list[last-1] = list[last-1], list[best]
	}

	for i := 0; i < size; i++ {
		if list[i].dist == MaxDistance {
			result[list[i].index] = -1
		} else {
			result[list[i].index] = (int)(list[i].dist)
		}
	}
	return result
}

//输入邻接矩阵(0指不通)，返回两点间最短路径的长度(-1指不通)。
func PlainDijkstraX(matrix [][]uint, start int, end int) int {
	var size = len(matrix)
	if start < 0 || end < 0 || start >= size || end >= size {
		return -1
	}
	if start == end {
		return 0
	}

	var list = make([]vertex, size)
	for i := 0; i < size-1; i++ {
		list[i].index, list[i].dist = i, MaxDistance
	}
	list[start].index = size - 1
	list[size-1].index, list[size-1].dist = start, 0

	for last := size - 1; last > 0 && list[last].dist != MaxDistance; last-- {
		if list[last].index == end {
			return int(list[last].dist)
		}
		var best = 0
		for i := 0; i < last; i++ {
			var step = matrix[list[last].index][list[i].index]
			var distance uint
			if step != 0 { //通
				distance = list[last].dist + step
				if distance < list[i].dist {
					list[i].dist = distance
				} else {
					distance = list[i].dist
				}
			} else { //不通
				distance = list[i].dist
			}
			if distance < list[best].dist {
				best = i
			}
		}
		list[best], list[last-1] = list[last-1], list[best]
	}
	return -1
}
