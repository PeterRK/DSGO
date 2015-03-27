package flow

//获取增广路径流量，复杂度为O(VE)。
func search(shadow [][]edge, matrix [][]uint, book []uint, space []int) uint {
	var size = len(matrix)
	//该循环的每一轮复杂度为O(V)，每一轮都至少会删除图的一条边
	for flow, tips := uint(0), uint(0); ; flow += tips {
		tips = ^uint(0)
		var top = 0
		for current := 0; current != size-1; {
			var sz = len(shadow[current])
			if sz != 0 {
				space[top], book[top] = current, tips
				top++
				var path = shadow[current][sz-1]
				current, tips = path.next, min(tips, path.val)
			} else { //碰壁，退一步
				if top == 0 { //退无可退
					return flow
				}
				top--
				current, tips = space[top], book[top]
				var last = len(shadow[current]) - 1
				var path = shadow[current][last]
				matrix[current][path.next] += path.val
				shadow[current] = shadow[current][:last]
			}
		}
		//找到增广路径
		for top--; top >= 0; top-- {
			var current = space[top]
			var last = len(shadow[current]) - 1
			var path = &shadow[current][last]
			path.val -= tips
			matrix[path.next][current] += tips //逆流，防止贪心断路
			if path.val == 0 {
				shadow[current] = shadow[current][:last]
			}
		}
	}
	return 0
}

func min(a uint, b uint) uint {
	if a > b {
		return b
	}
	return a
}
