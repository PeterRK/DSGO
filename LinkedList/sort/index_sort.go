package sort

import (
	asort "DSGO/Array/sort"
	"DSGO/LinkedList/list"
)

type Node = list.Node

const LOWER_BOUND_Y = asort.LOWER_BOUND_Y

type reference struct {
	ref *Node
	off int
}

func indexSimpleSort(index []reference) {
	if len(index) < 2 {
		return
	}
	for i := 1; i < len(index); i++ {
		key := index[i]
		if key.ref.Val < index[0].ref.Val ||
			(key.ref.Val == index[0].ref.Val && key.off < index[0].off) {
			for j := i; j > 0; j-- {
				index[j] = index[j-1]
			}
			index[0] = key
		} else {
			pos := i
			for index[pos-1].ref.Val > key.ref.Val ||
				(index[pos-1].ref.Val == key.ref.Val && index[pos-1].off > key.off) {
				index[pos] = index[pos-1]
				pos--
			}
			index[pos] = key
		}
	}
}

func indexHeapSort(index []reference) {
	for idx := len(index)/2 - 1; idx >= 0; idx-- {
		indexDown(index, idx)
	}
	for sz := len(index) - 1; sz > 0; sz-- {
		index[0], index[sz] = index[sz], index[0]
		indexDown(index[:sz], 0)
	}
}

func indexDown(index []reference, pos int) {
	key := index[pos]
	kid, last := pos*2+1, len(index)-1
	for kid < last {
		if index[kid+1].ref.Val > index[kid].ref.Val ||
			(index[kid+1].ref.Val == index[kid].ref.Val && index[kid+1].off > index[kid].off) {
			kid++
		}
		if key.ref.Val > index[kid].ref.Val ||
			(key.ref.Val == index[kid].ref.Val && key.off > index[kid].off) {
			break
		}
		index[pos] = index[kid]
		pos, kid = kid, kid*2+1
	}
	if kid == last && (key.ref.Val < index[kid].ref.Val ||
		(key.ref.Val == index[kid].ref.Val && key.off < index[kid].off)) {
		index[pos], pos = index[kid], kid
	}
	index[pos] = key
}

func indexTriPartition(index []reference) (fst, snd int) {
	sz := len(index)
	m, s := sz/2, sz/8
	a, b := m-s, m+s
	if index[a].ref.Val > index[b].ref.Val ||
		(index[a].ref.Val == index[b].ref.Val && index[a].off > index[b].off) {
		a, b = b, a
	}
	pivot1, pivot2 := index[a], index[b]
	index[a], index[b] = index[0], index[sz-1]

	a, b = 1, sz-2
	for a <= b && (index[a].ref.Val < pivot1.ref.Val ||
		(index[a].ref.Val == pivot1.ref.Val && index[a].off < pivot1.off)) {
		a++
	}
	for k := a; k <= b; k++ {
		if index[k].ref.Val > pivot2.ref.Val ||
			(index[k].ref.Val == pivot2.ref.Val && index[k].off > pivot2.off) {
			for k < b && (index[b].ref.Val > pivot2.ref.Val ||
				(index[b].ref.Val == pivot2.ref.Val && index[b].off > pivot2.off)) {
				b--
			}
			index[k], index[b] = index[b], index[k]
			b--
		}
		if index[k].ref.Val < pivot1.ref.Val ||
			(index[k].ref.Val == pivot1.ref.Val && index[k].off < pivot1.off) {
			index[k], index[a] = index[a], index[k]
			a++
		}
	}

	index[0], index[a-1] = index[a-1], pivot1
	index[sz-1], index[b+1] = index[b+1], pivot2
	return a - 1, b + 1
}

func IndexIntroSort(head *Node) *Node {
	index := make([]reference, 0, 16)
	for node := head; node != nil; node = node.Next {
		index = append(index, reference{ref: node, off: len(index)})
	}
	life := log2ceil(uint(len(index))) * 3 / 2
	indexIntroSort(index, life)

	//根据index对list进行最终处理
	var tail = list.FakeHead(&head)
	for i := 0; i < len(index); i++ {
		tail.Next = index[i].ref
		tail = tail.Next
	}
	tail.Next = nil
	return head
}
func indexIntroSort(index []reference, life uint) {
	for len(index) > LOWER_BOUND_Y {
		if life == 0 {
			indexHeapSort(index)
			return
		}
		life--
		fst, snd := indexTriPartition(index)
		indexIntroSort(index[:fst], life)
		indexIntroSort(index[snd+1:], life)
		index = index[fst+1 : snd]
	}
	indexSimpleSort(index)
}
