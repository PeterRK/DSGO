package tree

const piece_sz = 30

type piece struct {
	fw, bw *piece
	space  [piece_sz]*Node
}
type index struct {
	pt  *piece
	idx int
}
type queue struct {
	front, back index
	cnt         int
}

func newQ() *queue {
	var q = new(queue)
	q.cnt = 0
	var block = new(piece)
	block.fw, block.bw = nil, nil
	q.front.pt, q.back.pt = block, block
	q.front.idx, q.back.idx = 0, -1
	return q
}

func (q *queue) push(node *Node) {
	if q.front.idx == piece_sz {
		q.front.idx = 0
		if q.front.pt.fw == nil {
			var block = new(piece)
			block.bw, block.fw = q.front.pt, nil
			q.front.pt.fw = block
		}
		q.front.pt = q.front.pt.fw
	}
	q.front.pt.space[q.front.idx] = node
	q.front.idx++
	q.cnt++
}

func (q *queue) pop() (node *Node, fail bool) {
	if q.cnt == 0 {
		return nil, true
	}
	q.cnt--
	q.back.idx++
	node = q.back.pt.space[q.back.idx]
	if q.back.idx == piece_sz-1 {
		q.back.idx = -1    //q.back.idx永远不指向(piece_sz-1)
		q.back.pt.bw = nil //只保留一块缓冲
		q.back.pt = q.back.pt.fw
	}
	return node, false
}
