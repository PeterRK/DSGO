package trie

const NODE_CAP = uint8(5)

type node struct {
	key  [NODE_CAP]byte
	cnt  uint8
	ref  uint16
	kids []*node
}
type Trie interface {
	Search(key string) uint16
	Insert(key string)
	Remove(key string, all bool)
}

func newNode() *node {
	var unit = new(node)
	//unit.cnt, unit.ref = 0, 0
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
	if root.cnt+kid.cnt > NODE_CAP { //半缩
		var j = uint8(0)
		for i := root.cnt; i < NODE_CAP; i++ {
			root.key[i] = kid.key[j]
			j++
		}
		root.cnt = NODE_CAP
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
func (root *node) Search(key string) uint16 {
	var mk = uint8(0)
	for idx := 0; idx < len(key); idx++ {
		if mk == root.cnt { //下探
			if len(root.kids) == 1 && root.ref == 0 && root.cnt < NODE_CAP {
				root.consume(root.kids[0])
			} else { //单纯下探
				var spot = root.searchKid(key[idx])
				if spot == len(root.kids) || root.kids[spot].key[0] != key[idx] {
					return 0
				}
				root = root.kids[spot]
				mk = 1
				continue
			}
		}
		if key[idx] != root.key[mk] {
			return 0
		}
		mk++
	}
	if mk != root.cnt {
		return 0
	}
	return root.ref
}

func createTail(str string) *node { //data非空
	var head = newNode()
	var last, idx = head, 0
	for ; idx+int(NODE_CAP) < len(str); idx += int(NODE_CAP) {
		for i := 0; i < int(NODE_CAP); i++ {
			last.key[i] = str[idx+i]
		}
		last.cnt = NODE_CAP
		last.kids = append(last.kids, newNode())
		last = last.kids[0]
	}
	for ; idx < len(str); idx++ {
		last.key[last.cnt] = str[idx]
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
func (root *node) Insert(key string) {
	if len(key) == 0 {
		return
	}
	var mk = uint8(0)
	for idx := 0; idx < len(key); idx++ {
		if mk == root.cnt { //下探
			var spot = root.searchKid(key[idx])
			if spot == len(root.kids) || root.kids[spot].key[0] != key[idx] {
				root.kids = append(root.kids, nil)
				for i := len(root.kids) - 1; i > spot; i-- {
					root.kids[i] = root.kids[i-1]
				}
				root.kids[spot] = createTail(key[idx:])
				return
			}
			root = root.kids[spot]
			mk = 1
		} else {
			if root.key[mk] != key[idx] {
				root.split(mk)
				root.kids = append(root.kids, createTail(key[idx:]))
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
func (root *node) Remove(key string, all bool) {
	if len(key) == 0 {
		return
	}
	var knot, branch = (*node)(nil), 0 //记录独苗分支节点
	var mk = uint8(0)
	for idx := 0; idx < len(key); idx++ {
		if mk == root.cnt { //下探
			var spot = root.searchKid(key[idx])
			if spot == len(root.kids) || root.kids[spot].key[0] != key[idx] {
				return
			}
			if root.ref != 0 || len(root.kids) > 1 {
				knot, branch = root, spot
			}
			root = root.kids[spot]
			mk = 1
		} else {
			if root.key[mk] != key[idx] {
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
