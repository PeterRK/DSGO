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

const (
	STACK = iota
	QUEUE
	DEQUE
)

func (this *deque) initialize(hint int) {
	this.cnt = 0
	var block = new(piece)
	block.fw, block.bw = nil, nil
	this.front.pt, this.back.pt = block, block
	switch hint {
	case STACK:
		this.front.idx, this.back.idx = 0, -1
	case QUEUE:
		this.front.idx, this.back.idx = piece_sz, piece_sz-1
	default: //DEQUE
		this.front.idx, this.back.idx = piece_sz/2, piece_sz/2-1
	}
}

func (this *deque) IsEmpty() bool {
	return this.front.pt == this.back.pt &&
		(this.front.idx-1) == this.back.idx
}
func (this *deque) Size() int {
	return this.cnt
}

func (this *deque) PushFront(key int) {
	if this.front.idx == piece_sz {
		this.front.idx = 0
		if this.front.pt.fw == nil {
			var block = new(piece)
			block.bw, block.fw = this.front.pt, nil
			this.front.pt.fw = block
			this.front.pt = block
		}
	}
	this.front.pt.space[this.front.idx] = key
	this.front.idx++
	this.cnt++
}

func (this *deque) PushBack(key int) {
	if this.back.idx == -1 {
		this.back.idx = piece_sz - 1
		if this.back.pt.bw == nil {
			var block = new(piece)
			block.fw, block.bw = this.back.pt, nil
			this.back.pt.bw = block
			this.back.pt = block
		}
	}
	this.back.pt.space[this.back.idx] = key
	this.back.idx--
	this.cnt++
}

func (this *deque) Front() (key int, fail bool) {
	if this.IsEmpty() {
		return 0, true
	}
	return this.front.pt.space[this.front.idx-1], false
}
func (this *deque) Back() (key int, fail bool) {
	if this.IsEmpty() {
		return 0, true
	}
	return this.back.pt.space[this.back.idx+1], false
}

func (this *deque) PopFront() (key int, fail bool) {
	if this.IsEmpty() {
		return 0, true
	}
	this.cnt--
	this.front.idx--
	key = this.front.pt.space[this.front.idx]
	if this.front.idx == 0 {
		this.front.idx = piece_sz //this.front.idx永远不指向0
		this.front.pt.fw = nil    //只保留一块缓冲
		this.front.pt = this.front.pt.bw
	}
	return key, false
}
func (this *deque) PopBack() (key int, fail bool) {
	if this.IsEmpty() {
		return 0, true
	}
	this.cnt--
	this.back.idx++
	key = this.back.pt.space[this.back.idx]
	if this.back.idx == piece_sz-1 {
		this.back.idx = -1 //this.back.idx永远不指向(piece_sz-1)
		this.back.pt.bw = nil
		this.back.pt = this.back.pt.fw
	}
	return key, false
}
