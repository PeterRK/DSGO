package trie

import (
	"DSGO/array"
	"DSGO/utils"
	"fmt"
)

const segCap = uint8(7)

type node struct {
	data [segCap]byte
	mark uint8 //高位一bit表示是否实体，低位表示段长
	kids []*node
}

type Trie struct {
	root node
	size int
}

func New() utils.StrSet {
	return new(Trie)
}

func (t *Trie) Clear() {
	t.size = 0
	t.root.mark = 0
	t.root.kids = t.root.kids[:0]
}

func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) IsEmpty() bool {
	return t.size == 0
}

// 在由小到大的序列中寻找第一个大于或等于ch的位置
func search(kids []*node, ch byte) int {
	a, b := 0, len(kids)
	for a < b {
		m := (a + b) / 2
		if ch > kids[m].data[0] {
			a = m + 1
		} else {
			b = m
		}
	}
	return a
}

//返回-1代表match
func (n *node) diff(key string, skip, size uint8) int {
	for i := skip; i < size; i++ {
		if key[i] != n.data[i] {
			return int(i)
		}
	}
	return -1
}

func (unit *node) compact() {
	if len(unit.kids) != 1 || unit.mark >= segCap {
		return
	}
	kid := unit.kids[0]
	if unit.mark+(kid.mark&0x7f) > segCap { //半缩
		i := uint8(0)
		for ; unit.mark < segCap; i++ {
			unit.data[unit.mark] = kid.data[i]
			unit.mark++
		}
		sz := uint8(0)
		for end := kid.mark & 0x7f; i < end; i++ {
			kid.data[sz] = kid.data[i]
			sz++
		}
		kid.mark = (kid.mark & 0x80) | sz
	} else { //全缩
		for i, end := uint8(0), kid.mark&0x7f; i < end; i++ {
			unit.data[unit.mark] = kid.data[i]
			unit.mark++
		}
		unit.mark |= kid.mark & 0x80
		unit.kids = kid.kids
	}
}

func (t *Trie) Search(key string) bool {
	unit, skip := &t.root, uint8(0)
	unit.compact()
	for sz := unit.mark & 0x7f; len(key) > int(sz); sz = unit.mark & 0x7f {
		if unit.diff(key, skip, sz) >= 0 {
			return false
		}
		skip = 1
		key = key[sz:]
		pos := search(unit.kids, key[0])
		if pos >= len(unit.kids) || unit.kids[pos].data[0] != key[0] {
			return false
		}
		unit = unit.kids[pos]
		unit.compact()
	}
	return (len(key)|0x80) == int(unit.mark) &&
		unit.diff(key, skip, unit.mark&0x7f) < 0

}

func createTail(key string) *node { //len(key) > 0
	head := new(node)
	tail := head
	for len(key) > int(segCap) {
		for i := uint8(0); i < segCap; i++ {
			tail.data[i] = key[i]
		}
		tail.mark = segCap
		tail.kids = make([]*node, 1)
		tail.kids[0] = new(node)
		tail = tail.kids[0]
		key = key[segCap:]
	}
	for i := 0; i < len(key); i++ {
		tail.data[i] = key[i]
	}
	tail.mark = uint8(len(key)) | 0x80
	return head
}

func (unit *node) split(pos int) *node {
	sz := int(unit.mark & 0x7f)
	kid := new(node)
	for i := pos; i < sz; i++ {
		kid.data[kid.mark] = unit.data[i]
		kid.mark++
	}
	kid.mark |= unit.mark & 0x80
	kid.kids = unit.kids
	unit.mark = uint8(pos)
	unit.kids = nil
	return kid
}

func (t *Trie) Insert(key string) bool {
	if t.insert(key) {
		t.size++
		return true
	}
	return false
}

func (t *Trie) insert(key string) bool {
	unit, skip := &t.root, uint8(0)
	unit.compact()
	brk := 0
	for sz := unit.mark & 0x7f; len(key) > int(sz); sz = unit.mark & 0x7f {
		brk = unit.diff(key, skip, sz)
		if brk >= 0 {
			goto Lsplit
		}
		skip = 1
		key = key[sz:]
		pos := search(unit.kids, key[0])
		if pos >= len(unit.kids) || unit.kids[pos].data[0] != key[0] {
			unit.kids = array.InsertTo(unit.kids, pos, createTail(key))
			return true
		}
		unit = unit.kids[pos]
		unit.compact()
	}
	brk = unit.diff(key, skip, uint8(len(key)))
	if brk < 0 {
		if len(key) == int(unit.mark&0x7f) {
			done := (unit.mark & 0x80) == 0
			unit.mark |= 0x80
			return done
		} else { //len(key) < sz
			unit.kids = []*node{unit.split(len(key))}
			unit.mark |= 0x80
			return true
		}
	}
Lsplit:
	kid1 := createTail(key[brk:])
	kid2 := unit.split(brk)
	if kid1.data[0] > kid2.data[0] {
		kid1, kid2 = kid2, kid1
	}
	unit.kids = []*node{kid1, kid2}
	return true
}

func (t *Trie) Remove(key string) bool {
	if t.remove(key) {
		t.size--
		return true
	}
	return false
}

func (t *Trie) remove(key string) bool {
	knot, branch := (*node)(nil), -1 //追踪删除关节
	unit, skip := &t.root, uint8(0)
	unit.compact()
	for sz := unit.mark & 0x7f; len(key) > int(sz); sz = unit.mark & 0x7f {
		if (unit.mark&0x80) != 0 || len(unit.kids) > 1 {
			knot = unit
		}
		if unit.diff(key, skip, sz) >= 0 {
			return false
		}
		skip = 1
		key = key[sz:]
		pos := search(unit.kids, key[0])
		if pos >= len(unit.kids) || unit.kids[pos].data[0] != key[0] {
			return false
		}
		if knot == unit {
			branch = pos
		}
		unit = unit.kids[pos]
		unit.compact()
	}
	if (len(key)|0x80) == int(unit.mark) &&
		unit.diff(key, skip, unit.mark&0x7f) < 0 {
		if len(unit.kids) == 0 && knot != nil {
			knot.kids = array.EraseFrom(knot.kids, branch, true)
		} else {
			unit.mark &= 0x7f
		}
		return true
	}
	return false
}

func (unit *node) debug(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	fmt.Print(string(unit.data[:unit.mark&0x7f]))
	if (unit.mark & 0x80) == 0 {
		fmt.Println()
	} else {
		fmt.Print("$\n")
	}
	for _, kid := range unit.kids {
		kid.debug(indent + 1)
	}
}

func (t *Trie) debug() {
	t.root.debug(0)
	fmt.Println("================")
}
