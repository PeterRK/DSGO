package trie

const capacity = uint8(5)

type node struct {
	key  [capacity]byte
	cnt  uint8
	ref  uint16
	kids []*node
}
type Trie interface {
	Search(data []byte) uint16
	Insert(data []byte)
	Remove(data []byte, all bool)
}

func newNode() *node {
	var unit = new(node)
	unit.cnt, unit.ref = 0, 0
	//unit.kids = make([]*node, 0, 4)
	return unit
}
func NewTrie() Trie {
	var root = newNode()
	root.ref = 1 //空串总是有效
	return root
}

func (unit *node) searchKid(ch byte) int {
	var start, end = 0, len(unit.kids)
	for start < end {
		var mid = (start + end) / 2
		if ch > unit.kids[mid].key[0] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func (root *node) consume(kid *node) {
	if root.cnt+kid.cnt > capacity { //半缩
		var j = uint8(0)
		for i := root.cnt; i < capacity; i++ {
			root.key[i] = kid.key[j]
			j++
		}
		root.cnt = capacity
		var i = uint8(0)
		for ; j < kid.cnt; j++ {
			kid.key[i] = kid.key[j]
			i++
		}
		kid.cnt = i
	} else { //全缩
		for i := uint8(0); i < kid.cnt; i++ {
			root.key[root.cnt] = kid.key[i]
			root.cnt++
		}
		root.ref, root.kids = kid.ref, kid.kids
	}
}

//除了查找，还要尝试节缩
func (root *node) Search(data []byte) uint16 {
	var mk = uint8(0)
	for idx := 0; idx < len(data); idx++ {
		if mk == root.cnt { //下探
			if len(root.kids) == 1 && root.ref == 0 && root.cnt < capacity {
				root.consume(root.kids[0])
			} else { //单纯下探
				var spot = root.searchKid(data[idx])
				if spot == len(root.kids) || root.kids[spot].key[0] != data[idx] {
					return 0
				}
				root = root.kids[spot]
				mk = 1
				continue
			}
		}
		if data[idx] != root.key[mk] {
			return 0
		}
		mk++
	}
	if mk != root.cnt {
		return 0
	}
	return root.ref
}

func createTail(data []byte) *node { //data非空
	var head = newNode()
	var last, idx = head, 0
	for ; idx+int(capacity) < len(data); idx += int(capacity) {
		for i := 0; i < int(capacity); i++ {
			last.key[i] = data[idx+i]
		}
		last.cnt = capacity
		last.kids = append(last.kids, newNode())
		last = last.kids[0]
	}
	for ; idx < len(data); idx++ {
		last.key[last.cnt] = data[idx]
		last.cnt++
	}
	last.ref = 1
	return head
}
func (unit *node) split(mk uint8) {
	var peer = newNode()
	peer.ref, unit.ref = unit.ref, 0
	peer.kids, unit.kids = unit.kids, append(peer.kids, peer)
	for i := mk; i < unit.cnt; i++ {
		peer.key[peer.cnt] = unit.key[i]
		peer.cnt++
	}
	unit.cnt = mk
}
func (root *node) Insert(data []byte) {
	if len(data) == 0 {
		return
	}
	var mk = uint8(0)
	for idx := 0; idx < len(data); idx++ {
		if mk == root.cnt { //下探
			var spot = root.searchKid(data[idx])
			if spot == len(root.kids) || root.kids[spot].key[0] != data[idx] {
				root.kids = append(root.kids, nil)
				for i := len(root.kids) - 1; i > spot; i-- {
					root.kids[i] = root.kids[i-1]
				}
				root.kids[spot] = createTail(data[idx:])
				return
			}
			root = root.kids[spot]
			mk = 1
		} else {
			if root.key[mk] != data[idx] {
				root.split(mk)
				root.kids = append(root.kids, createTail(data[idx:]))
				if root.kids[0].key[0] > root.kids[1].key[0] {
					root.kids[0], root.kids[1] = root.kids[1], root.kids[0]
				}
				return
			}
			mk++
		}
	}
	if mk != root.cnt {
		root.split(mk)
	}
	root.ref++
}

//清除标记，删除独苗分支，暂时不节缩
func (root *node) Remove(data []byte, all bool) {
	if len(data) == 0 {
		return
	}
	var knot, branch = (*node)(nil), 0 //记录独苗分支节点
	var mk = uint8(0)
	for idx := 0; idx < len(data); idx++ {
		if mk == root.cnt { //下探
			var spot = root.searchKid(data[idx])
			if spot == len(root.kids) || root.kids[spot].key[0] != data[idx] {
				return
			}
			if root.ref != 0 || len(root.kids) > 1 {
				knot, branch = root, spot
			}
			root = root.kids[spot]
			mk = 1
		} else {
			if root.key[mk] != data[idx] {
				return
			}
			mk++
		}
	}
	if mk != root.cnt || root.ref == 0 {
		return
	}
	if all {
		root.ref = 0
	} else {
		root.ref--
	}

	if root.ref == 0 && len(root.kids) == 0 {
		for i := branch + 1; i < len(knot.kids); i++ {
			knot.kids[i-1] = knot.kids[i]
		} //删除独苗
		knot.kids = knot.kids[:len(knot.kids)-1]
	}
}
