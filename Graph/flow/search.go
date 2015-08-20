package flow

//获取增广路径流量，复杂度为O(VE)。
func search(shadow [][]edge, matrix [][]uint, s *stack) uint {
	var size = len(matrix)
	//每一轮都至少会删除图的一条边
	for flow, stream := uint(0), uint(0); ; flow += stream {
		stream = ^uint(0)
		s.clear()
		for current := 0; current != size-1; {
			var sz = len(shadow[current])
			if sz != 0 {
				s.push(current, stream)
				var path = shadow[current][sz-1]
				current, stream = path.next, min(stream, path.val)
			} else { //碰壁，退一步
				if s.isEmpty() { //退无可退
					return flow
				}
				current, stream = s.pop()
				var last = len(shadow[current]) - 1
				var path = shadow[current][last]
				matrix[current][path.next] += path.val
				shadow[current] = shadow[current][:last]
			}
		}

		//该循环的每一轮复杂度为O(V)
		for !s.isEmpty() { //处理找到的增广路径
			var current, _ = s.pop()
			var last = len(shadow[current]) - 1
			var path = &shadow[current][last]
			path.val -= stream
			matrix[path.next][current] += stream //逆流，防止贪心断路
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
