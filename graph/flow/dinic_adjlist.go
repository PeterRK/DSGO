package flow

import (
	"DSGO/array"
	"DSGO/array/sort"
	"DSGO/graph"
	"DSGO/utils"
	"math"
)

type contextL struct {
	origin     [][]graph.Path
	shadow     [][]graph.Path //分层残图
	reflux     [][]graph.Path //逆流暂存
	level      []uint32
	queue      utils.Queue[int]
	stack      utils.Stack[memo]
	ranker     sort.Ranker[graph.Path]
	start, end int
}

func nextOrder(a, b *graph.Path) bool {
	return a.Next < b.Next
}

func (ctx *contextL) init(roads [][]graph.Path, start, end int) {
	ctx.ranker.Less = nextOrder
	size := len(roads)
	for i := 0; i < size; i++ {
		ctx.ranker.Sort(roads[i]) //要求有序
	}
	ctx.origin = roads
	ctx.shadow = make([][]graph.Path, size)
	ctx.reflux = make([][]graph.Path, size)
	ctx.level = make([]uint32, size)
	ctx.queue = array.NewQueue[int](size)
	ctx.stack = array.NewStack[memo](size)
	ctx.start, ctx.end = start, end
}

//输入邻接表，返回最大流，复杂度为O(V^2 E)。
func DinicL(roads [][]graph.Path, start, end int) uint {
	size := len(roads)
	if start < 0 || end < 0 ||
		start >= size || end >= size ||
		start == end {
		return 0
	}
	var ctx contextL
	ctx.init(roads, start, end)
	return ctx.dinic()
}

func (ctx *contextL) dinic() uint {
	flow := uint(0)
	for ctx.separate() {
		flow += ctx.search()
		ctx.flushBack()
	}
	return flow
}

//获取增广路径流量，复杂度为O(EVlogV)。
func (ctx *contextL) search() uint {
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
				fillBack(ctx.origin[curr], ctx.shadow[curr][last])
				ctx.shadow[curr] = ctx.shadow[curr][:last]
			}
		}

		//该循环的每一轮复杂度为O(V^2)
		for !ctx.stack.IsEmpty() { //处理找到的增广路径
			tmp := ctx.stack.Pop()
			curr := tmp.idx
			last := len(ctx.shadow[curr]) - 1
			path := &ctx.shadow[curr][last]
			path.Weight -= stream
			ctx.reflux[path.Next] = append(ctx.reflux[path.Next],
				graph.Path{Next: curr, Weight: stream}) //逆流，防止贪心断路
			if path.Weight == 0 {
				ctx.shadow[curr] = ctx.shadow[curr][:last]
			}
		}
	}
}

func (ctx *contextL) markLevel() bool {
	for i := 0; i < len(ctx.level); i++ {
		ctx.level[i] = fakeLevel
	}

	ctx.queue.Clear()
	ctx.level[ctx.start] = 0
	ctx.queue.Push(ctx.start)

	for !ctx.queue.IsEmpty() {
		curr := ctx.queue.Pop()
		for _, path := range ctx.origin[curr] {
			if ctx.level[path.Next] != fakeLevel {
				continue
			}
			ctx.level[path.Next] = ctx.level[curr] + 1
			if path.Next == ctx.end {
				endLevel := ctx.level[path.Next]
				for !ctx.queue.IsEmpty() {
					curr = ctx.queue.Pop()
					if ctx.level[curr] == endLevel {
						ctx.level[curr] = fakeLevel
					}
				}
				return true
			}
			ctx.queue.Push(path.Next)
		}
	}
	return false
}

//筛分层次，生成分层残图，复杂度为O(E)。
func (ctx *contextL) separate() bool {
	if !ctx.markLevel() {
		return false
	}
	for curr := 0; curr < len(ctx.level); curr++ {
		if ctx.level[curr] == fakeLevel {
			continue
		}
		paths := ctx.origin[curr]
		for i := 0; i < len(paths); i++ {
			next := paths[i].Next
			if ctx.level[next] == ctx.level[curr]+1 {
				ctx.shadow[curr] = append(ctx.shadow[curr], paths[i])
				paths[i].Weight = 0
			}
		}
	}
	return true
}

func (ctx *contextL) flushBack() {
	for i := 0; i < len(ctx.origin); i++ {
		origin, shadow, reflux := ctx.origin[i], ctx.shadow[i], ctx.reflux[i]
		if len(shadow) != 0 {
			fillBackVec(origin, shadow)
			ctx.shadow[i] = shadow[:0]
		}
		if len(reflux) != 0 {
			ctx.ranker.Sort(reflux)
			origin = merge(origin, reflux)
		}
		ctx.origin[i] = compact(origin) //去重去零
	}
}

func merge(base, part []graph.Path) []graph.Path {
	a, b := len(base)-1, len(part)-1
	base = append(base, part...)
	for c := len(base) - 1; b >= 0; c-- {
		if a < 0 {
			for ; b >= 0; b-- {
				base[b] = part[b]
			}
			break
		}
		if base[a].Next > part[b].Next {
			base[c] = base[a]
			a--
		} else {
			base[c] = part[b]
			b--
		}
	}
	return base
}

func compact(list []graph.Path) []graph.Path {
	size := len(list)
	if size == 0 {
		return list
	}
	last := 0
	for i := 1; i < size; i++ {
		if list[i].Next == list[last].Next {
			list[last].Weight += list[i].Weight
		} else {
			if list[last].Weight != 0 {
				last++
			}
			list[last] = list[i]
		}
	}
	if list[last].Weight != 0 {
		last++
	}
	return list[:last]
}

type segment struct {
	space      []graph.Path
	start, end int
}

func (s *segment) fill(path graph.Path) int {
	a, b := s.start, s.end
	for a < b {
		m := (a + b) / 2
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
	seg := segment{
		space: list,
		start: 0, end: len(list)}
	seg.fill(path)
}

func fillBackVec(list []graph.Path, frag []graph.Path) {
	seg := segment{
		space: list,
		start: 0, end: len(list)}
	a, b := 0, len(frag)-1
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
