package flow

import (
	"DSGO/Graph/graph"
)

func (pk *dataM) flushBack() {
	for i := 0; i < len(pk.matrix); i++ {
		if len(pk.shadow[i]) == 0 {
			continue
		}
		for _, path := range pk.shadow[i] {
			pk.matrix[i][path.Next] += path.Weight
		}
		pk.shadow[i] = pk.shadow[i][:0]
	}
}

//宽度优先遍历
func (pk *dataM) markLevel() bool {
	for i := 0; i < len(pk.memo); i++ {
		pk.memo[i] = fakeLevel
	}

	pk.memo[pk.start] = 0
	pk.queue.push(pk.start)

	for !pk.queue.isEmpty() {
		var cur = pk.queue.pop()
		if pk.matrix[cur][pk.end] != 0 {
			pk.memo[pk.end] = pk.memo[cur] + 1
			return true
		}
		for i := 0; i < len(pk.memo); i++ {
			if pk.memo[i] == fakeLevel && pk.matrix[cur][i] != 0 {
				pk.memo[i] = pk.memo[cur] + 1
				pk.queue.push(i)
			}
		}
	}
	return false
}

//筛分层次，生成分层残图，复杂度为O(V^2)。
func (pk *dataM) separate() bool {
	pk.queue.clear()
	if !pk.markLevel() {
		return false
	}
	for { //队列pop出的点并没有实际删除，可回溯遍历所有访问过的点
		var cur, err = pk.queue.traceBack()
		if err != nil {
			break
		}
		//pk.shadow[cur] = pk.shadow[cur][:0]
		for i := 0; i < len(pk.memo); i++ {
			if pk.memo[i] == pk.memo[cur]+1 && pk.matrix[cur][i] != 0 {
				var path = graph.Path{Next: i, Weight: pk.matrix[cur][i]}
				pk.shadow[cur] = append(pk.shadow[cur], path)
				pk.matrix[cur][i] = 0
			}
		}
		if len(pk.shadow[cur]) == 0 {
			pk.memo[cur] = fakeLevel
		}
	}
	return true
}

//获取增广路径流量，复杂度为O(VE)。
func (pk *dataM) search() uint {
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
				var path = pk.shadow[cur][last]
				pk.matrix[cur][path.Next] += path.Weight
				pk.shadow[cur] = pk.shadow[cur][:last]
			}
		}

		//该循环的每一轮复杂度为O(V)
		for !pk.stack.isEmpty() { //处理找到的增广路径
			var cur, _ = pk.stack.pop()
			var last = len(pk.shadow[cur]) - 1
			var path = &pk.shadow[cur][last]
			path.Weight -= stream
			pk.matrix[path.Next][cur] += stream //逆流，防止贪心断路
			if path.Weight == 0 {
				pk.shadow[cur] = pk.shadow[cur][:last]
			}
		}
	}
	return 0
}
