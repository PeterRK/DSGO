package flow

import (
	"DSGO/array"
	"DSGO/graph"
	"DSGO/utils"
	"math"
)

type contextM struct {
	matrix     [][]uint
	shadow     [][]graph.Path //分层残图
	queue      utils.Queue[int]
	stack      utils.Stack[memo]
	level      []uint32
	start, end int
}

type memo struct {
	idx    int
	stream uint
}

const fakeLevel = math.MaxUint32 - 1

func (ctx *contextM) init(matrix [][]uint, start, end int) {
	size := len(matrix)
	ctx.matrix = matrix
	ctx.shadow = make([][]graph.Path, size)
	ctx.level = make([]uint32, size)
	ctx.queue = array.NewQueue[int](size)
	ctx.stack = array.NewStack[memo](size)
	ctx.start, ctx.end = start, end
}

//输入邻接矩阵，返回最大流，复杂度为O(V^2 E)。
func DinicM(matrix [][]uint, start, end int) uint {
	size := len(matrix)
	if start < 0 || end < 0 ||
		start >= size || end >= size ||
		start == end {
		return 0
	}
	var ctx contextM
	ctx.init(matrix, start, end)
	return ctx.dinic()
}

func (ctx *contextM) dinic() uint {
	flow := uint(0)
	for ctx.separate() {
		//由于search一次会删除图的若干条边，所以循环次数为O(E/k)
		//循环内的separate和flushBack操作的复杂度都是O(V^2)，search的复杂度为O(Vk)
		flow += ctx.search()
		ctx.flushBack()
	}
	return flow
}

//获取增广路径流量，复杂度为O(VE)。
func (ctx *contextM) search() uint {
	//每一轮都至少会删除图的一条边
	for flow, stream := uint(0), uint(0); ; flow += stream {
		stream = math.MaxUint
		ctx.stack.Clear()
		for curr := ctx.start; curr != ctx.end; {
			sz := len(ctx.shadow[curr])
			if sz != 0 {
				ctx.stack.Push(memo{idx: curr, stream: stream})
				path := ctx.shadow[curr][sz-1]
				curr, stream = path.Next, utils.Min(stream, path.Weight)
			} else { //碰壁，退一步
				if ctx.stack.IsEmpty() { //退无可退
					return flow
				}
				tmp := ctx.stack.Pop()
				curr, stream = tmp.idx, tmp.stream
				last := len(ctx.shadow[curr]) - 1
				path := ctx.shadow[curr][last]
				ctx.matrix[curr][path.Next] += path.Weight
				ctx.shadow[curr] = ctx.shadow[curr][:last]
			}
		}

		//该循环的每一轮复杂度为O(V)
		for !ctx.stack.IsEmpty() { //处理找到的增广路径
			tmp := ctx.stack.Pop()
			curr := tmp.idx
			last := len(ctx.shadow[curr]) - 1
			path := &ctx.shadow[curr][last]
			path.Weight -= stream
			ctx.matrix[path.Next][curr] += stream //逆流，防止贪心断路
			if path.Weight == 0 {
				ctx.shadow[curr] = ctx.shadow[curr][:last]
			}
		}
	}
}

//宽度优先遍历
func (ctx *contextM) markLevel() bool {
	for i := 0; i < len(ctx.level); i++ {
		ctx.level[i] = fakeLevel
	}

	ctx.queue.Clear()
	ctx.level[ctx.start] = 0
	ctx.queue.Push(ctx.start)

	for !ctx.queue.IsEmpty() {
		curr := ctx.queue.Pop()
		if ctx.matrix[curr][ctx.end] != 0 {
			ctx.level[ctx.end] = ctx.level[curr] + 1
			endLevel := ctx.level[ctx.end]
			for !ctx.queue.IsEmpty() {
				curr = ctx.queue.Pop()
				if ctx.level[curr] == endLevel {
					ctx.level[curr] = fakeLevel
				}
			}
			return true
		}
		for i := 0; i < len(ctx.level); i++ {
			if ctx.level[i] == fakeLevel && ctx.matrix[curr][i] != 0 {
				ctx.level[i] = ctx.level[curr] + 1
				ctx.queue.Push(i)
			}
		}
	}
	return false
}

//筛分层次，生成分层残图，复杂度为O(V^2)。
func (ctx *contextM) separate() bool {
	if !ctx.markLevel() {
		return false
	}
	for curr := 0; curr < len(ctx.level); curr++ {
		if ctx.level[curr] == fakeLevel {
			continue
		}
		//ctx.shadow[curr] = ctx.shadow[curr][:0]
		for i := 0; i < len(ctx.level); i++ {
			if ctx.level[i] == ctx.level[curr]+1 && ctx.matrix[curr][i] != 0 {
				path := graph.Path{Next: i, Weight: ctx.matrix[curr][i]}
				ctx.shadow[curr] = append(ctx.shadow[curr], path)
				ctx.matrix[curr][i] = 0
			}
		}
	}
	return true
}

func (ctx *contextM) flushBack() {
	for i := 0; i < len(ctx.matrix); i++ {
		if len(ctx.shadow[i]) == 0 {
			continue
		}
		for _, path := range ctx.shadow[i] {
			ctx.matrix[i][path.Next] += path.Weight
		}
		ctx.shadow[i] = ctx.shadow[i][:0]
	}
}
