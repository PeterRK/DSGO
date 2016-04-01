package flow

import (
	"DSGO/Graph/graph"
)

//获取增广路径流量，复杂度为O(EVlogV)。
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
				fillBack(pk.origin[cur], pk.shadow[cur][last])
				pk.shadow[cur] = pk.shadow[cur][:last]
			}
		}

		//该循环的每一轮复杂度为O(V^2)
		for !pk.stack.isEmpty() { //处理找到的增广路径
			var cur, _ = pk.stack.pop()
			var last = len(pk.shadow[cur]) - 1
			var path = &pk.shadow[cur][last]
			path.Weight -= stream
			pk.reflux[path.Next] = append(pk.reflux[path.Next],
				graph.Path{Next: cur, Weight: stream}) //逆流，防止贪心断路
			if path.Weight == 0 {
				pk.shadow[cur] = pk.shadow[cur][:last]
			}
		}
	}
	return 0
}

type segment struct {
	space      []graph.Path
	start, end int
}

func (s *segment) fill(path graph.Path) int {
	var a, b = s.start, s.end
	for a < b {
		var m = (a + b) / 2
		switch {
		case path.Next > s.space[m].Next:
			a = m + 1
		case path.Next < s.space[m].Next:
			b = m
		default:
			s.space[m].Weight += path.Weight
			return m
		}
	}
	panic("no target") //目标必须存在
}
func fillBack(list []graph.Path, path graph.Path) {
	var seg = segment{
		space: list,
		start: 0, end: len(list)}
	seg.fill(path)
}
func fillBackVec(list []graph.Path, frag []graph.Path) {
	var seg = segment{
		space: list,
		start: 0, end: len(list)}
	var a, b = 0, len(frag) - 1
	for a < b {
		seg.start = seg.fill(frag[a])
		a++
		seg.end = seg.fill(frag[b])
		b--
	}
	if a == b {
		seg.fill(frag[a])
	}
}
