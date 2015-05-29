package bptree

const INDEX_HALF = BASE_SIZE
const INDEX_FULL = INDEX_HALF*2 - 1
const INDEX_QUARTER = INDEX_HALF / 2

type index struct {
	node //inner==true
	data [INDEX_FULL]int
	kids [INDEX_FULL]*index
}

func newIndex() *index {
	var unit = new(index)
	unit.inner = true
	return unit
}

func (unit *index) remove(place int) {
	unit.cnt--
	for i := place; i < unit.cnt; i++ {
		unit.data[i] = unit.data[i+1]
		unit.kids[i] = unit.kids[i+1]
	}
}

//peer为分裂项，peer为nil时表示不分裂
func (unit *index) insert(place int, kid *index) (peer *index) {
	const full_sz, half_sz = INDEX_FULL, INDEX_HALF
	if unit.cnt < full_sz {
		for i := unit.cnt; i > place; i-- {
			unit.data[i] = unit.data[i-1]
			unit.kids[i] = unit.kids[i-1]
		}
		unit.data[place] = kid.ceil()
		unit.kids[place] = kid
		unit.cnt++
		return nil
	}

	peer = newIndex()
	unit.cnt, peer.cnt = half_sz, half_sz

	if place < half_sz {
		for i := 0; i < half_sz; i++ {
			peer.data[i] = unit.data[i+(half_sz-1)]
			peer.kids[i] = unit.kids[i+(half_sz-1)]
		}
		for i := half_sz - 1; i > place; i-- {
			unit.data[i] = unit.data[i-1]
			unit.kids[i] = unit.kids[i-1]
		}
		unit.data[place] = kid.ceil()
		unit.kids[place] = kid
	} else {
		for i := full_sz; i > place; i-- {
			peer.data[i-half_sz] = unit.data[i-1]
			peer.kids[i-half_sz] = unit.kids[i-1]
		}
		peer.data[place-half_sz] = kid.ceil()
		peer.kids[place-half_sz] = kid
		for i := half_sz; i < place; i++ {
			peer.data[i-half_sz] = unit.data[i]
			peer.kids[i-half_sz] = unit.kids[i]
		}
	}
	return peer
}

//要求peer为unit后的节点，发生合并返回true
func (unit *index) combine(peer *index) bool {
	var total = unit.cnt + peer.cnt
	if total <= INDEX_HALF+INDEX_QUARTER {
		for i := unit.cnt; i < total; i++ {
			unit.data[i] = peer.data[i-unit.cnt]
			unit.kids[i] = peer.kids[i-unit.cnt]
		}
		unit.cnt = total
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
			unit.kids[i] = peer.kids[i-unit.cnt]
		}
		for i := 0; i < peer_sz; i++ {
			peer.data[i] = peer.data[i+diff]
			peer.kids[i] = peer.kids[i+diff]
		}
	} else { //前填后
		var diff = peer_sz - peer.cnt
		for i := peer.cnt - 1; i >= 0; i-- {
			peer.data[i+diff] = peer.data[i]
			peer.kids[i+diff] = peer.kids[i]
		}
		for i := unit.cnt - 1; i >= unit_sz; i-- {
			peer.data[i-unit_sz] = unit.data[i]
			peer.kids[i-unit_sz] = unit.kids[i]
		}
	}
	unit.cnt, peer.cnt = unit_sz, peer_sz
	return false
}
