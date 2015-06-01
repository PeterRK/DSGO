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
	var u = new(leaf)
	u.inner = false
	return u
}

func (u *leaf) remove(place int) {
	u.cnt--
	for i := place; i < u.cnt; i++ {
		u.data[i] = u.data[i+1]
		//
	}
}

//peer为分裂项，peer为nil时表示不分裂
func (u *leaf) insert(place int, key int) (peer *leaf) {
	const full_sz, half_sz = LEAF_FULL, LEAF_HALF
	if u.cnt < full_sz {
		for i := u.cnt; i > place; i-- {
			u.data[i] = u.data[i-1]
			//
		}
		u.data[place] = key
		//
		u.cnt++
		return nil
	}

	peer = newLeaf()
	u.cnt, peer.cnt = half_sz, half_sz
	peer.next, u.next = u.next, peer
	if place < half_sz {
		for i := 0; i < half_sz; i++ {
			peer.data[i] = u.data[i+(half_sz-1)]
			//
		}
		for i := half_sz - 1; i > place; i-- {
			u.data[i] = u.data[i-1]
			//
		}
		u.data[place] = key
		//
	} else {
		for i := full_sz; i > place; i-- {
			peer.data[i-half_sz] = u.data[i-1]
			//
		}
		peer.data[place-half_sz] = key
		//
		for i := half_sz; i < place; i++ {
			peer.data[i-half_sz] = u.data[i]
			//
		}
	}
	return peer
}

//要求peer为unit后的节点，发生合并返回 true
func (u *leaf) combine(peer *leaf) bool {
	var total = u.cnt + peer.cnt
	if total <= LEAF_HALF+LEAF_QUARTER {
		for i := u.cnt; i < total; i++ {
			u.data[i] = peer.data[i-u.cnt]
			//
		}
		u.cnt, u.next = total, peer.next
		return true
	}
	//分流而不合并
	var u_sz = total / 2
	if u.cnt == u_sz {
		return false
	}
	var peer_sz = total - u_sz
	if peer.cnt > peer_sz { //后填前
		var diff = peer.cnt - peer_sz
		for i := u.cnt; i < u_sz; i++ {
			u.data[i] = peer.data[i-u.cnt]
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
		for i := u.cnt - 1; i >= u_sz; i-- {
			peer.data[i-u_sz] = u.data[i]
			//
		}
	}
	u.cnt, peer.cnt = u_sz, peer_sz
	return false
}
