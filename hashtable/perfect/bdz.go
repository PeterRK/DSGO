package perfect

import (
	"DSGO/array"
	"DSGO/hashtable"
	"DSGO/utils"
	"crypto/rand"
	"encoding/binary"
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

func testAndSetBit(bitmap []uint32, pos uint32) bool {
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
	prev uint32
	next uint32
}

type hypergraph struct {
	edges [][3]vertex //超图的边，每条边三个顶点
	nodes []uint32    //超图的点，表现为链表
}

func rand32() uint32 {
	var tmp [4]byte
	_, err := rand.Read(tmp[:])
	if err != nil {
		return uint32(time.Now().UnixNano())
	}
	return binary.LittleEndian.Uint32(tmp[:])
}

func (h *bdz) init(keys []string) bool {
	if len(keys) > math.MaxInt32 { //注意，这里不用MaxUint32
		return false
	}
	h.width = uint32((uint64(len(keys))*105 + 255) / 256) //此比例来自BDZ论文
	h.width |= 1
	slotCnt := int(h.width * 3)
	h.bitmap = make([]uint32, (slotCnt+15)/16)
	ebmSize := (len(keys) + 31) / 32
	nbmSize := (slotCnt + 31) / 32
	bitmap := make([]uint32, utils.Max(ebmSize, nbmSize)) //nbmSize > ebmSize
	graph := hypergraph{
		edges: make([][3]vertex, len(keys)),
		nodes: make([]uint32, slotCnt)}

	free := make([]uint32, 0, len(keys))

	for trys := 0; trys < 8; trys++ {
		h.seed = rand32()
		if trys != 0 {
			fmt.Printf("retry with seed %x\n", h.seed)
		}
		graph.init(keys, h.hash) //根据随机种子生成超图
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

func (g *hypergraph) init(keys []string, hash func(string) [3]uint32) {
	array.SetAll(g.nodes, endIdx)
	for i := 0; i < len(keys); i++ {
		slots := hash(keys[i])
		for j := 0; j < 3; j++ {
			v := &g.edges[i][j]
			v.slot = slots[j]
			v.prev = endIdx
			v.next = g.nodes[v.slot]
			g.nodes[v.slot] = uint32(i)
			if v.next != endIdx {
				g.edges[v.next][j].prev = uint32(i)
			}
		}
	}
}

func (g *hypergraph) tear(free []uint32, bitmap []uint32) []uint32 {
	free = free[:0]
	for i := len(g.edges) - 1; i >= 0; i-- {
		edge := g.edges[i]
		for j := 0; j < 3; j++ {
			v := edge[j]
			if v.prev == endIdx && v.next == endIdx &&
				testAndSetBit(bitmap, uint32(i)) {
				free = append(free, uint32(i))
			}
		}
	}
	for head := 0; head < len(free); head++ {
		curr := free[head]
		for j := 0; j < 3; j++ {
			v := &g.edges[curr][j]
			i := endIdx
			if v.prev != endIdx {
				i = v.prev
				g.edges[i][j].next = v.next
			}
			if v.next != endIdx {
				i = v.next
				g.edges[i][j].prev = v.prev
			}
			if i == endIdx {
				continue
			}
			u := &g.edges[i][j]
			if u.prev == endIdx && u.next == endIdx &&
				testAndSetBit(bitmap, uint32(i)) {
				free = append(free, i)
			}
		}
	}
	return free
}

func (h *bdz) mapping(edges [][3]vertex, free []uint32, bitmap []uint32) {
	array.SetAll(h.bitmap, ^uint32(0))
	//free从前到后的过程中伴随剩余点集的净收缩，逆之则可以保证每次处理的边不会无点可占
	for i := len(free) - 1; i >= 0; i-- {
		e := edges[free[i]]
		v0, v1, v2 := e[0].slot, e[1].slot, e[2].slot
		switch {
		case testAndSetBit(bitmap, v0):
			setBit(bitmap, v1)
			setBit(bitmap, v2)
			setBit2on11(h.bitmap, v0, (6-(bit2(h.bitmap, v1)+bit2(h.bitmap, v2)))%3)
		case testAndSetBit(bitmap, v1):
			setBit(bitmap, v2)
			setBit2on11(h.bitmap, v1, (7-(bit2(h.bitmap, v0)+bit2(h.bitmap, v2)))%3)
		case testAndSetBit(bitmap, v2):
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
