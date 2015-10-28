package flow

import (
	"Graph/graph"
)

//获取增广路径流量，复杂度为O(V^2 E)。
func (pk *data) search() uint {
	//每一轮都至少会删除图的一条边
	for flow, stream := uint(0), uint(0); ; flow += stream {
		stream = ^uint(0)
		pk.stack.clear()
		for cur := pk.start; cur != pk.end; {
			var sz = len(pk.shadow[cur])
			if sz != 0 {
				pk.stack.push(cur, stream)
				var path = pk.shadow[cur][sz-1]
				cur, stream = path.Next, min(stream, path.Weight)
			} else { //碰壁，退一步
				if pk.stack.isEmpty() { //退无可退
					return flow
				}
				cur, stream = pk.stack.pop()
				var last = len(pk.shadow[cur]) - 1
				pk.roads[cur] = patch(pk.roads[cur], pk.shadow[cur][last])
				pk.shadow[cur] = pk.shadow[cur][:last]
			}
		}

		//该循环的每一轮复杂度为O(V^2)
		for !pk.stack.isEmpty() { //处理找到的增广路径
			var cur, _ = pk.stack.pop()
			var last = len(pk.shadow[cur]) - 1
			var path = &pk.shadow[cur][last]
			path.Weight -= stream
			pk.roads[path.Next] = patch(pk.roads[path.Next],
				graph.Path{Next: cur, Weight: stream}) //逆流，防止贪心断路
			if path.Weight == 0 {
				pk.shadow[cur] = pk.shadow[cur][:last]
			}
		}
	}
	return 0
}

func patch(list []graph.Path, path graph.Path) []graph.Path {
	var spot = binarySearch(list, path.Next)
	if spot == len(list) || list[spot].Next != path.Next {
		list = append(list, path)
		for i := len(list) - 1; i > spot; i-- {
			list[i] = list[i-1]
		}
		list[spot] = path
	} else {
		list[spot].Weight += path.Weight
	}
	return list
}

func binarySearch(list []graph.Path, key int) int {
	var start, end = 0, len(list)
	for start < end {
		var mid = (start + end) / 2
		if key > list[mid].Next {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}
