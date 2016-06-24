package graph

type memo struct {
	prev, next *memo
	id         int
}

const FakeID = -1

func shadowGraph(roads [][]int) [][]int {
	var shadow = make([][]int, len(roads))
	for i := 0; i < len(roads); i++ {
		var nexts = roads[i]
		for j := 0; j < len(nexts); j++ {
			var peer = nexts[j]
			shadow[peer] = append(shadow[peer], i)
		}
	}
	return shadow
}

type data struct {
	roads [][]int
	book  []memo
	knot  *memo
	parts [][]int
}

func doDFS(doit func(*data, int), pk *data, curr int) {
	var node = &pk.book[curr]
	node.id = FakeID
	node.prev.next, node.next.prev = node.next, node.prev

	var nexts = pk.roads[curr]
	for i := 0; i < len(nexts); i++ {
		var peer = nexts[i]
		if pk.book[peer].id != FakeID {
			doit(pk, peer)
		}
	}
}
func doDFSX(pk *data, curr int) {
	doDFS(doDFSX, pk, curr)
	var node = &pk.book[curr]
	pk.knot.next.prev, node.next = node, pk.knot.next
	pk.knot.next, node.prev = node, pk.knot
}
func doDFSY(pk *data, curr int) {
	doDFS(doDFSY, pk, curr)
	var last = len(pk.parts) - 1
	pk.parts[last] = append(pk.parts[last], curr)
}

func (pk *data) dfsX(id int) {
	doDFSX(pk, id)
}
func (pk *data) dfsY(id int) {
	doDFSY(pk, id)
}

//分解强联通分量
func SplitDirectedGraph(roads [][]int) [][]int {
	var size = len(roads)
	if size < 1 {
		return nil
	}
	if size == 1 {
		return [][]int{{0}}
	}

	var book = make([]memo, size+2)
	for i := 0; i < size; i++ {
		book[i].next, book[i+1].prev = &book[i+1], &book[i]
		book[i].id = i
	}
	book[size].next, book[0].prev = &book[0], &book[size]
	book[size].id, book[size+1].id = FakeID, FakeID
	book[size+1].prev, book[size+1].next = &book[size+1], &book[size+1]

	var pack = data{
		roads: roads,
		book:  book,
		knot:  &book[size+1]}
	var knot = &book[size]
	for knot.next != knot {
		pack.dfsX(knot.next.id)
	}

	for i := 0; i < size; i++ {
		book[i].id = i
	}
	pack.roads = shadowGraph(roads)
	knot = pack.knot
	for knot.next != knot {
		pack.parts = append(pack.parts, nil)
		pack.dfsY(knot.next.id)
	}
	return pack.parts
}
