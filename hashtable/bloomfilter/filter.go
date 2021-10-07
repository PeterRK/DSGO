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

type bloomFliter struct {
	core []piece
	size int
	seed uint64
}

type BloomFliter interface {
	Capacity() int
	Size() int
	Insert(string)
	Search(string) bool
}

func New(capacity uint32) BloomFliter {
	bf := new(bloomFliter)
	bf.init(capacity)
	return bf
}

func (bf *bloomFliter) Capacity() int {
	return len(bf.core) * PieceCap
}

func (bf *bloomFliter) Size() int {
	return bf.size
}

func (bf *bloomFliter) init(capacity uint32) {
	if capacity == 0 {
		capacity = 1
	}
	bf.core = make([]piece, int((capacity+(PieceCap-1))/PieceCap))
	//bf.size = 0
	bf.seed = uint64(time.Now().UnixNano())
}

type context struct {
	vec  []byte
	pos  [8]uint16
	mask [8]byte
}

func (ctx *context) insert() bool {
	miss := false
	for i := 0; i < 8; i++ {
		pos, mask := ctx.pos[i], ctx.mask[i]
		if (ctx.vec[pos] & mask) == 0 {
			ctx.vec[pos] |= mask
			miss = true
		}
	}
	return miss
}

func (ctx *context) search() bool {
	for i := 0; i < 8; i++ {
		pos, mask := ctx.pos[i], ctx.mask[i]
		if (ctx.vec[pos] & mask) == 0 {
			return false
		}
	}
	return true
}

func (bf *bloomFliter) hash(key string) (ctx context) {
	a, b, z := hashtable.Hash160(bf.seed, key)
	ctx.vec = bf.core[int(z)%len(bf.core)][:]
	ctx.pos[0] = uint16(a)
	ctx.pos[1] = uint16(a >> 16)
	ctx.pos[2] = uint16(a >> 32)
	ctx.pos[3] = uint16(a >> 48)
	ctx.pos[4] = uint16(b)
	ctx.pos[5] = uint16(b >> 16)
	ctx.pos[6] = uint16(b >> 32)
	ctx.pos[7] = uint16(b >> 48)
	for i := 0; i < 8; i++ {
		ctx.mask[i] = 1 << (ctx.pos[i] & 7)
	}
	for i := 0; i < 8; i++ {
		ctx.pos[i] >>= 3
	}
	return ctx
}

func (bf *bloomFliter) Insert(key string) {
	ctx := bf.hash(key)
	if ctx.insert() {
		bf.size++
	}
}

func (bf *bloomFliter) Search(key string) bool {
	ctx := bf.hash(key)
	return ctx.search()
}
