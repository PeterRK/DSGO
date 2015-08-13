package deque

import (
	"errors"
)

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

//dq.front.idx永远不为0，以piece_sz代之，此时要求dq.front.fw != nil
//dq.back.idx永远不为(piece_sz-1)，以-1代之，此时要求dq.front.bw != nil

func (dq *deque) initialize() {
	dq.reset(new(piece))
}
func (dq *deque) Clear() {
	dq.reset(dq.front.pt)
}
func (dq *deque) reset(block *piece) {
	dq.cnt = 0
	block.fw, block.bw = nil, nil
	dq.front.pt, dq.back.pt = block, block
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

func (dq *deque) Front() (int, error) {
	if dq.IsEmpty() {
		return 0, errors.New("empty")
	}
	return dq.front.pt.space[dq.front.idx-1], nil
}
func (dq *deque) Back() (int, error) {
	if dq.IsEmpty() {
		return 0, errors.New("empty")
	}
	return dq.back.pt.space[dq.back.idx+1], nil
}

func (dq *deque) PopFront() (int, error) {
	if dq.IsEmpty() {
		return 0, errors.New("empty")
	}
	dq.cnt--
	dq.front.idx--
	var key = dq.front.pt.space[dq.front.idx]
	if dq.front.idx == 0 {
		dq.front.idx = piece_sz //dq.front.idx永远不为0
		dq.front.pt.fw = nil    //只保留一块缓冲
		dq.front.pt = dq.front.pt.bw
	}
	return key, nil
}
func (dq *deque) PopBack() (int, error) {
	if dq.IsEmpty() {
		return 0, errors.New("empty")
	}
	dq.cnt--
	dq.back.idx++
	var key = dq.back.pt.space[dq.back.idx]
	if dq.back.idx == piece_sz-1 {
		dq.back.idx = -1    //dq.back.idx永远不为(piece_sz-1)
		dq.back.pt.bw = nil //只保留一块缓冲，因为
		dq.back.pt = dq.back.pt.fw
	}
	return key, nil
}
