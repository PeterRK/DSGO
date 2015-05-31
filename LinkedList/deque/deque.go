package deque

const piece_sz = 30

type piece struct {
	fw, bw *piece
	space  [piece_sz]int
}
type index struct {
	pt  *piece
	idx int
}
type deque struct {
	front, back index
	cnt         int
}

//dq.front.idx永远不指向0
//dq.back.idx永远不指向(piece_sz-1)

func (dq *deque) initialize() {
	dq.cnt = 0
	var block = new(piece)
	block.fw, block.bw = nil, nil
	dq.front.pt, dq.back.pt = block, block
	dq.front.idx, dq.back.idx = piece_sz/2, piece_sz/2-1
}
func (dq *deque) Clear() {
	dq.cnt = 0
	dq.back.pt = dq.front.pt
	dq.front.pt.fw, dq.back.pt.bw = nil, nil
	dq.front.idx, dq.back.idx = piece_sz/2, piece_sz/2-1
}

func (dq *deque) IsEmpty() bool {
	return dq.Size() == 0
}
func (dq *deque) Size() int {
	return dq.cnt
}

func (dq *deque) PushFront(key int) {
	if dq.front.idx == piece_sz {
		dq.front.idx = 0
		if dq.front.pt.fw == nil {
			var block = new(piece)
			block.bw, block.fw = dq.front.pt, nil
			dq.front.pt.fw = block
		}
		dq.front.pt = dq.front.pt.fw
	}
	dq.front.pt.space[dq.front.idx] = key
	dq.front.idx++
	dq.cnt++
}
func (dq *deque) PushBack(key int) {
	if dq.back.idx == -1 {
		dq.back.idx = piece_sz - 1
		if dq.back.pt.bw == nil {
			var block = new(piece)
			block.fw, block.bw = dq.back.pt, nil
			dq.back.pt.bw = block
		}
		dq.back.pt = dq.back.pt.bw
	}
	dq.back.pt.space[dq.back.idx] = key
	dq.back.idx--
	dq.cnt++
}

func (dq *deque) Front() (key int, fail bool) {
	if dq.IsEmpty() {
		return 0, true
	}
	return dq.front.pt.space[dq.front.idx-1], false
}
func (dq *deque) Back() (key int, fail bool) {
	if dq.IsEmpty() {
		return 0, true
	}
	return dq.back.pt.space[dq.back.idx+1], false
}

func (dq *deque) PopFront() (key int, fail bool) {
	if dq.IsEmpty() {
		return 0, true
	}
	dq.cnt--
	dq.front.idx--
	key = dq.front.pt.space[dq.front.idx]
	if dq.front.idx == 0 {
		dq.front.idx = piece_sz //dq.front.idx永远不指向0
		dq.front.pt.fw = nil    //只保留一块缓冲
		dq.front.pt = dq.front.pt.bw
	}
	return key, false
}
func (dq *deque) PopBack() (key int, fail bool) {
	if dq.IsEmpty() {
		return 0, true
	}
	dq.cnt--
	dq.back.idx++
	key = dq.back.pt.space[dq.back.idx]
	if dq.back.idx == piece_sz-1 {
		dq.back.idx = -1    //dq.back.idx永远不指向(piece_sz-1)
		dq.back.pt.bw = nil //只保留一块缓冲
		dq.back.pt = dq.back.pt.fw
	}
	return key, false
}
