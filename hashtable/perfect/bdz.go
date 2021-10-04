package perfect

import (
	"DSGO/array"
	"DSGO/hashtable"
	"DSGO/utils"
	"fmt"
	"math"
	"time"
)

type bdz struct {
	seed   uint32
	width  uint32
	bitmap []uint32
}

type Hasher interface {
	Hash(string) uint32
}

func New(keys []string) Hasher {
	h := new(bdz)
	if h.init(keys) {
		return h
	}
	return nil
}

func bit2(bitmap []uint32, pos uint32) uint32 {
	blk, sft := pos>>4, (pos&15)<<1
	return (bitmap[blk] >> sft) & 3
}

func testBit(bitmap []uint32, pos uint32) bool {
	blk, sft := pos>>5, pos&31
	return (bitmap[blk] & (uint32(1) << sft)) != 0
}

func setBit(bitmap []uint32, pos uint32) {
	blk, sft := pos>>5, pos&31
	bitmap[blk] |= uint32(1) << sft
}

func setBitWithCheck(bitmap []uint32, pos uint32) bool {
	if testBit(bitmap, pos) {
		return false
	}
	setBit(bitmap, pos)
	return true
}

func (h *bdz) hash(key string) (slots [3]uint32) {
	a, b := hashtable.Hash128(uint64(h.seed), key)
	slots[0] = uint32(a) % h.width
	slots[1] = uint32(a>>32)%h.width + h.width
	slots[2] = uint32(b)%h.width + h.width*2
	return slots
}

func (h *bdz) Hash(key string) uint32 {
	slots := h.hash(key)
	mark := bit2(h.bitmap, slots[0]) + bit2(h.bitmap, slots[1]) + bit2(h.bitmap, slots[2])
	return slots[mark%3]
}

const endIdx = uint32(math.MaxUint32)

type vertex struct {
	slot uint32
	next uint32
}

type list struct {
	size uint8
	head uint32
}

type hypergraph struct {
	edges [][3]vertex //超图的边，每条边三个顶点
	nodes []list      //超图的点，表现为链表
}

func (h *bdz) init(keys []string) bool {
	if len(keys) > math.MaxInt32 { //注意，这里不用MaxUint32
		return false
	}
	h.width = uint32((uint64(len(keys))*41 + 99) / 100) //此比例来自BDZ论文
	h.width |= 1
	slotCnt := int(h.width * 3)
	h.bitmap = make([]uint32, (slotCnt+15)/16)
	ebmSize := (len(keys) + 31) / 32
	nbmSize := (slotCnt + 31) / 32
	bitmap := make([]uint32, utils.Max(ebmSize, nbmSize)) //nbmSize > ebmSize
	graph := hypergraph{
		edges: make([][3]vertex, len(keys)),
		nodes: make([]list, slotCnt)}

	free := make([]uint32, 0, len(keys))

	rand := utils.NewXorshift(uint32(time.Now().Unix()))
	for trys := 0; trys < 16; trys++ {
		h.seed = rand.Next()
		if trys != 0 {
			fmt.Printf("retry with seed %d\n", h.seed)
		}
		if !graph.init(keys, h.hash) { //根据随机种子生成超图
			return false
		}

		array.SetAll(bitmap[:ebmSize], 0)
		free = graph.tear(free, bitmap) //拆解超图获得边序列
		if len(free) != len(graph.edges) {
			continue
		}

		array.SetAll(bitmap[:nbmSize], 0)
		h.mapping(graph.edges, free, bitmap)
		return true
	}
	return false
}

func (g *hypergraph) init(keys []string, hash func(string) [3]uint32) bool {
	array.SetAll(g.nodes, list{0, endIdx})
	for i := 0; i < len(keys); i++ {
		slots := hash(keys[i])
		for j := 0; j < 3; j++ {
			v := &g.edges[i][j]
			v.slot = slots[j]
			n := &g.nodes[v.slot]
			v.next = n.head
			n.head = uint32(i)
			n.size++
			if n.size > 100 { //扎堆是不正常的
				return false
			}
		}
	}
	return true
}

func (g *hypergraph) tear(free []uint32, bitmap []uint32) []uint32 {
	free, head := free[:0], 0
	for i := 0; i < len(g.edges); i++ {
		for j := 0; j < 3; j++ {
			n := g.nodes[g.edges[i][j].slot]
			if n.size == 1 {
				if n.head != uint32(i) {
					panic("broken graph")
				}
				setBit(bitmap, n.head)
				free = append(free, n.head)
				break
			}
		}
	}
	for head < len(free) {
		curr := free[head]
		head++
		for j := 0; j < 3; j++ {
			v := &g.edges[curr][j]
			n := &g.nodes[v.slot]
			p := &n.head
			for *p != curr {
				p = &g.edges[*p][j].next
			}
			*p = v.next
			v.next = endIdx
			n.size--
			if n.size == 1 && setBitWithCheck(bitmap, n.head) {
				free = append(free, n.head)
			}
		}
	}
	return free
}

func (h *bdz) mapping(edges [][3]vertex, free []uint32, bitmap []uint32) {
	array.SetAll(h.bitmap, ^uint32(0))
	for i := len(free) - 1; i >= 0; i-- { //BDZ论文指出：逆序就可以保证不误占
		e := edges[free[i]]
		v0, v1, v2 := e[0].slot, e[1].slot, e[2].slot
		switch {
		case setBitWithCheck(bitmap, v0):
			setBitWithCheck(bitmap, v1)
			setBitWithCheck(bitmap, v2)
			setBit2on11(h.bitmap, v0, (6-(bit2(h.bitmap, v1)+bit2(h.bitmap, v2)))%3)
		case setBitWithCheck(bitmap, v1):
			setBitWithCheck(bitmap, v2)
			setBit2on11(h.bitmap, v1, (7-(bit2(h.bitmap, v0)+bit2(h.bitmap, v2)))%3)
		case setBitWithCheck(bitmap, v2):
			setBit2on11(h.bitmap, v2, (8-(bit2(h.bitmap, v0)+bit2(h.bitmap, v1)))%3)
		default:
			panic("all nodes are occupied")
		}
	}
}

func setBit2on11(bitmap []uint32, pos uint32, val uint32) {
	blk, sft := pos>>4, (pos&15)<<1
	bitmap[blk] ^= ((^val & 3) << sft)
}
