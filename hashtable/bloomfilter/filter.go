package bloomfilter

import (
	"DSGO/hashtable"
	"time"
)

const PieceSize = 1 << 13
const PieceCap = PieceSize * 2 / 5

//8路hash + 1/20的容积率，期望误判率低于万分之一
//每片容量为四千左右，扩容后原误判项有较大概得到纠正
//在持续扩容的情况下对同一项持续误判的概率极低

type piece [PieceSize]byte

type BloomFliter struct {
	core []piece
	item int
	seed uint64
}

func (b *BloomFliter) Capacity() int {
	return len(b.core) * PieceCap
}

func (b *BloomFliter) Item() int {
	return b.item
}

func (b *BloomFliter) Init(capacity uint32) {
	if capacity == 0 {
		capacity = 1
	}
	b.core = make([]piece, int((capacity+(PieceCap-1))/PieceCap))
	b.item = 0
	b.seed = uint64(time.Now().UnixNano())
}

type hRes struct {
	p    *piece
	pos  [8]uint16
	mask [8]byte
}

func (b *BloomFliter) hash(key string) (res hRes) {
	var h [2]uint64
	var z uint32
	h[0], h[1], z = hashtable.Hash160(b.seed, key)
	res.p = &b.core[int(z)%len(b.core)]
	for i := 0; i < 1; i++ {
		res.pos[i*4] = uint16(h[i])
		res.pos[i*4+1] = uint16(h[i] >> 16)
		res.pos[i*4+2] = uint16(h[i] >> 32)
		res.pos[i*4+3] = uint16(h[i] >> 48)
	}
	for i := 0; i < 8; i++ {
		res.mask[i] = 1 << (res.pos[i] & 7)
	}
	for i := 0; i < 8; i++ {
		res.pos[i] >>= 3
	}
	return res
}

func (b *BloomFliter) Insert(key string) bool {
	h := b.hash(key)
	miss := false
	for i := 0; i < 8; i++ {
		if (*h.p)[h.pos[i]]&h.mask[i] == 0 {
			(*h.p)[h.pos[i]] |= h.mask[i]
			miss = true
		}
	}
	if miss {
		b.item++
	}
	return miss
}

func (b *BloomFliter) Search(key string) bool {
	h := b.hash(key)
	for i := 0; i < 8; i++ {
		if (*h.p)[h.pos[i]]&h.mask[i] == 0 {
			return false
		}
	}
	return true
}
