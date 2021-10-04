package graph

func shadowGraph(roads [][]int) [][]int {
	shadow := make([][]int, len(roads))
	for i := 0; i < len(roads); i++ {
		nexts := roads[i]
		for j := 0; j < len(nexts); j++ {
			peer := nexts[j]
			shadow[peer] = append(shadow[peer], i)
		}
	}
	return shadow
}

const fakeID = -1

type biNode struct {
	prev, next *biNode
	id         int
}

type context struct {
	roads [][]int
	book  []biNode
	knot  *biNode
	parts [][]int
}

func doDFS(doit func(*context, int), ctx *context, curr int) {
	node := &ctx.book[curr]
	node.id = fakeID
	node.prev.next, node.next.prev = node.next, node.prev

	nexts := ctx.roads[curr]
	for i := 0; i < len(nexts); i++ {
		peer := nexts[i]
		if ctx.book[peer].id != fakeID {
			doit(ctx, peer)
		}
	}
}

func doDFSX(ctx *context, curr int) {
	doDFS(doDFSX, ctx, curr)
	node := &ctx.book[curr]
	ctx.knot.next.prev, node.next = node, ctx.knot.next
	ctx.knot.next, node.prev = node, ctx.knot
}

func doDFSY(ctx *context, curr int) {
	doDFS(doDFSY, ctx, curr)
	last := len(ctx.parts) - 1
	ctx.parts[last] = append(ctx.parts[last], curr)
}

func (ctx *context) dfsX(id int) {
	doDFSX(ctx, id)
}

func (ctx *context) dfsY(id int) {
	doDFSY(ctx, id)
}

//分解强联通分量
func SplitDirectedGraph(roads [][]int) [][]int {
	size := len(roads)
	if size < 1 {
		return nil
	}
	if size == 1 {
		return [][]int{{0}}
	}

	book := make([]biNode, size+2)
	for i := 0; i < size; i++ {
		book[i].next, book[i+1].prev = &book[i+1], &book[i]
		book[i].id = i
	}
	book[size].next, book[0].prev = &book[0], &book[size]
	book[size].id, book[size+1].id = fakeID, fakeID
	book[size+1].prev, book[size+1].next = &book[size+1], &book[size+1]

	ctx := context{
		roads: roads,
		book:  book,
		knot:  &book[size+1]}
	knot := &book[size]
	for knot.next != knot {
		ctx.dfsX(knot.next.id)
	}

	for i := 0; i < size; i++ {
		book[i].id = i
	}
	ctx.roads = shadowGraph(roads)
	knot = ctx.knot
	for knot.next != knot {
		ctx.parts = append(ctx.parts, nil)
		ctx.dfsY(knot.next.id)
	}
	return ctx.parts
}
