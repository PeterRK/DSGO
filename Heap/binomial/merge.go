package binomial

import (
	"unsafe"
)

func fakeHead(spt **node) *node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).peer)
	return (*node)(unsafe.Pointer(base - off))
}

//list是从少到多的
func (hp *Heap) merge(list *node) {
	var hp_list = hp.list //直接取对GC不友好，绕一下道
	var knot = fakeHead(&hp_list)
	for list != nil {
		var one, another = list, knot.peer
		if another == nil || one.level < another.level {
			list, one.peer = one.peer, another
			knot.peer = one
		} else if one.level > another.level {
			knot = knot.peer
		} else { //同级合并
			list, knot.peer = one.peer, another.peer

			if one.key > another.key {
				one, another = another, one
			}
			another.peer, one.child = one.child, another
			one.level++

			one.peer, list = list, one //可能会有一项逆序，不影响大局
		}
	}
	hp.list = hp_list
}
func (hp *Heap) Merge(victim *Heap) {
	if hp != victim && !victim.IsEmpty() {
		if hp.IsEmpty() {
			*hp = *victim
		} else {
			if hp.top.key > victim.top.key {
				hp.top = victim.top
			}
			hp.merge(victim.list)
		}
		victim.Clear()
	}
}
