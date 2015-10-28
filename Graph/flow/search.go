package flow

import (
	"Graph/graph"
)

//获取增广路径流量，复杂度为O(VE)。
func search(shadow [][]graph.Path, matrix [][]uint, s *stack) uint {
	var size = len(matrix)
	//每一轮都至少会删除图的一条边
	for flow, stream := uint(0), uint(0); ; flow += stream {
		stream = ^uint(0)
		s.clear()
		for cur := 0; cur != size-1; {
			var sz = len(shadow[cur])
			if sz != 0 {
				s.push(cur, stream)
				var path = shadow[cur][sz-1]
				cur, stream = path.Next, min(stream, path.Weight)
			} else { //碰壁，退一步
				if s.isEmpty() { //退无可退
					return flow
				}
				cur, stream = s.pop()
				var last = len(shadow[cur]) - 1
				var path = shadow[cur][last]
				matrix[cur][path.Next] += path.Weight
				shadow[cur] = shadow[cur][:last]
			}
		}

		//该循环的每一轮复杂度为O(V)
		for !s.isEmpty() { //处理找到的增广路径
			var cur, _ = s.pop()
			var last = len(shadow[cur]) - 1
			var path = &shadow[cur][last]
			path.Weight -= stream
			matrix[path.Next][cur] += stream //逆流，防止贪心断路
			if path.Weight == 0 {
				shadow[cur] = shadow[cur][:last]
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

type stack struct {
	space []int
	spcx  []uint
	top   int
}

func (s *stack) bind(space []int, spcx []uint) {
	s.space, s.spcx = space, spcx
}
func (s *stack) clear() {
	s.top = 0
}
func (s *stack) isEmpty() bool {
	return s.top == 0
}
func (s *stack) push(id int, val uint) {
	s.space[s.top], s.spcx[s.top] = id, val
	s.top++
}
func (s *stack) pop() (id int, val uint) {
	s.top--
	return s.space[s.top], s.spcx[s.top]
}
