package bptree

const LEAF_HALF = BASE_SIZE*2 - 1
const LEAF_FULL = LEAF_HALF*2 - 1
const LEAF_QUARTER = LEAF_HALF / 2

type leaf struct {
	node //inner==false
	data [LEAF_FULL]int
	next *leaf
}

func newLeaf() *leaf {
	var unit = new(leaf)
	unit.inner = false
	return unit
}

func (unit *leaf) remove(place int) {
	unit.cnt--
	for i := place; i < unit.cnt; i++ {
		unit.data[i] = unit.data[i+1]
		//
	}
}

//peer为分裂项，peer为nil时表示不分裂
func (unit *leaf) insert(place int, key int) (peer *leaf) {
	const full_sz, half_sz = LEAF_FULL, LEAF_HALF
	if unit.cnt < full_sz {
		for i := unit.cnt; i > place; i-- {
			unit.data[i] = unit.data[i-1]
			//
		}
		unit.data[place] = key
		//
		unit.cnt++
		return nil
	}

	peer = newLeaf()
	unit.cnt, peer.cnt = half_sz, half_sz
	peer.next, unit.next = unit.next, peer
	if place < half_sz {
		for i := 0; i < half_sz; i++ {
			peer.data[i] = unit.data[i+(half_sz-1)]
			//
		}
		for i := half_sz - 1; i > place; i-- {
			unit.data[i] = unit.data[i-1]
			//
		}
		unit.data[place] = key
		//
	} else {
		for i := full_sz; i > place; i-- {
			peer.data[i-half_sz] = unit.data[i-1]
			//
		}
		peer.data[place-half_sz] = key
		//
		for i := half_sz; i < place; i++ {
			peer.data[i-half_sz] = unit.data[i]
			//
		}
	}
	return peer
}

//要求peer为unit后的节点，发生合并返回 true
func (unit *leaf) combine(peer *leaf) bool {
	var total = unit.cnt + peer.cnt
	if total <= LEAF_HALF+LEAF_QUARTER {
		for i := unit.cnt; i < total; i++ {
			unit.data[i] = peer.data[i-unit.cnt]
			//
		}
		unit.cnt, unit.next = total, peer.next
		return true
	}
	//分流而不合并
	var unit_sz = total / 2
	if unit.cnt == unit_sz {
		return false
	}
	var peer_sz = total - unit_sz
	if peer.cnt > peer_sz { //后填前
		var diff = peer.cnt - peer_sz
		for i := unit.cnt; i < unit_sz; i++ {
			unit.data[i] = peer.data[i-unit.cnt]
			//
		}
		for i := 0; i < peer_sz; i++ {
			peer.data[i] = peer.data[i+diff]
			//
		}
	} else { //前填后
		var diff = peer_sz - peer.cnt
		for i := peer.cnt - 1; i >= 0; i-- {
			peer.data[i+diff] = peer.data[i]
			//
		}
		for i := unit.cnt - 1; i >= unit_sz; i-- {
			peer.data[i-unit_sz] = unit.data[i]
			//
		}
	}
	unit.cnt, peer.cnt = unit_sz, peer_sz
	return false
}
