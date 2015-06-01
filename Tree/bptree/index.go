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
	var u = new(index)
	u.inner = true
	return u
}

func (u *index) remove(place int) {
	u.cnt--
	for i := place; i < u.cnt; i++ {
		u.data[i] = u.data[i+1]
		u.kids[i] = u.kids[i+1]
	}
}

//peer为分裂项，peer为nil时表示不分裂
func (u *index) insert(place int, kid *index) (peer *index) {
	const full_sz, half_sz = INDEX_FULL, INDEX_HALF
	if u.cnt < full_sz {
		for i := u.cnt; i > place; i-- {
			u.data[i] = u.data[i-1]
			u.kids[i] = u.kids[i-1]
		}
		u.data[place] = kid.ceil()
		u.kids[place] = kid
		u.cnt++
		return nil
	}

	peer = newIndex()
	u.cnt, peer.cnt = half_sz, half_sz

	if place < half_sz {
		for i := 0; i < half_sz; i++ {
			peer.data[i] = u.data[i+(half_sz-1)]
			peer.kids[i] = u.kids[i+(half_sz-1)]
		}
		for i := half_sz - 1; i > place; i-- {
			u.data[i] = u.data[i-1]
			u.kids[i] = u.kids[i-1]
		}
		u.data[place] = kid.ceil()
		u.kids[place] = kid
	} else {
		for i := full_sz; i > place; i-- {
			peer.data[i-half_sz] = u.data[i-1]
			peer.kids[i-half_sz] = u.kids[i-1]
		}
		peer.data[place-half_sz] = kid.ceil()
		peer.kids[place-half_sz] = kid
		for i := half_sz; i < place; i++ {
			peer.data[i-half_sz] = u.data[i]
			peer.kids[i-half_sz] = u.kids[i]
		}
	}
	return peer
}

//要求peer为unit后的节点，发生合并返回true
func (u *index) combine(peer *index) bool {
	var total = u.cnt + peer.cnt
	if total <= INDEX_HALF+INDEX_QUARTER {
		for i := u.cnt; i < total; i++ {
			u.data[i] = peer.data[i-u.cnt]
			u.kids[i] = peer.kids[i-u.cnt]
		}
		u.cnt = total
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
			u.kids[i] = peer.kids[i-u.cnt]
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
		for i := u.cnt - 1; i >= u_sz; i-- {
			peer.data[i-u_sz] = u.data[i]
			peer.kids[i-u_sz] = u.kids[i]
		}
	}
	u.cnt, peer.cnt = u_sz, peer_sz
	return false
}
