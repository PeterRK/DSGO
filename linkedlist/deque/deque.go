package deque

const PieceSize = 30

type piece[T any] struct {
	fw, bw *piece[T]
	data   [PieceSize]T
}

type deque[T any] struct {
	f, b struct {
		p *piece[T]
		i int
	}
	size int
}

type Deque[T any] interface {
	Size() int
	IsEmpty() bool
	Clear()
	PushFront(T)
	PushBack(T)
	PopFront() T
	PopBack() T
	Front() T
	Back() T
}

func New[T any]() Deque[T] {
	dq := new(deque[T])
	dq.init()
	return dq
}

//dq.f.i永远不为0，以PieceSize代之，此时要求dq.f.fw != nil
//dq.b.i永远不为PieceSize-1，以-1代之，此时要求dq.b.bw != nil

func (dq *deque[T]) reset(block *piece[T]) {
	dq.size = 0
	block.fw, block.bw = nil, nil
	dq.f.p, dq.b.p = block, block
	dq.f.i, dq.b.i = PieceSize/2, PieceSize/2-1
}

func (dq *deque[T]) init() {
	dq.reset(new(piece[T]))
}

func (dq *deque[T]) Clear() {
	dq.reset(dq.f.p)
}

func (dq *deque[T]) Size() int {
	return dq.size
}

func (dq *deque[T]) IsEmpty() bool {
	return dq.Size() == 0
}

func (dq *deque[T]) PushFront(unit T) {
	if dq.f.i == PieceSize {
		dq.f.i = 0
		if dq.f.p.fw == nil {
			block := new(piece[T])
			block.bw, block.fw = dq.f.p, nil
			dq.f.p.fw = block
		}
		dq.f.p = dq.f.p.fw
	}
	dq.f.p.data[dq.f.i] = unit
	dq.f.i++
	dq.size++
}

func (dq *deque[T]) PushBack(unit T) {
	if dq.b.i == -1 {
		dq.b.i = PieceSize - 1
		if dq.b.p.bw == nil {
			block := new(piece[T])
			block.fw, block.bw = dq.b.p, nil
			dq.b.p.bw = block
		}
		dq.b.p = dq.b.p.bw
	}
	dq.b.p.data[dq.b.i] = unit
	dq.b.i--
	dq.size++
}

func (dq *deque[T]) Front() T {
	if dq.IsEmpty() {
		panic("empty deque")
	}
	return dq.f.p.data[dq.f.i-1]
}

func (dq *deque[T]) Back() T {
	if dq.IsEmpty() {
		panic("empty deque")
	}
	return dq.b.p.data[dq.b.i+1]
}

func (dq *deque[T]) PopFront() T {
	if dq.IsEmpty() {
		panic("empty deque")
	}
	dq.size--
	dq.f.i--
	unit := dq.f.p.data[dq.f.i]
	if dq.f.i == 0 {
		dq.f.i = PieceSize //dq.f.idx永远不为0
		dq.f.p.fw = nil    //只保留一块缓冲
		dq.f.p = dq.f.p.bw
	}
	return unit
}

func (dq *deque[T]) PopBack() T {
	if dq.IsEmpty() {
		panic("empty deque")
	}
	dq.size--
	dq.b.i++
	unit := dq.b.p.data[dq.b.i]
	if dq.b.i == PieceSize-1 {
		dq.b.i = -1     //dq.b.idx永远不为(PieceSize-1)
		dq.b.p.bw = nil //只保留一块缓冲
		dq.b.p = dq.b.p.fw
	}
	return unit
}
