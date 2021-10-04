package binary

import (
	"constraints"
)

//二叉堆，底层采用数组。
//Build的复杂度为O(N)，Top的复杂度为O(1)，其余核心操作复杂度为O(logN)。
type Heap[T constraints.Ordered] struct {
	vec []T
}

func (hp *Heap[T]) Size() int {
	return len(hp.vec)
}

func (hp *Heap[T]) IsEmpty() bool {
	return len(hp.vec) == 0
}

func (hp *Heap[T]) Clear() {
	hp.vec = hp.vec[:0]
}

func (hp *Heap[T]) Top() T {
	if hp.IsEmpty() {
		panic("empty heap")
	}
	return hp.vec[0]
}

func (hp *Heap[T]) BuildInPlace(list []T) {
	size := len(list)
	hp.vec = list
	for idx := size/2 - 1; idx >= 0; idx-- {
		hp.down(idx)
	}
}

func (hp *Heap[T]) Build(list []T) {
	vec := make([]T, len(list))
	copy(vec, list)
	hp.BuildInPlace(vec)
}

func (hp *Heap[T]) Push(unit T) {
	place := len(hp.vec)
	hp.vec = append(hp.vec, unit)
	hp.up(place)
}

func (hp *Heap[T]) Pop() T {
	size := hp.Size()
	if size == 0 {
		panic("empty heap")
	}
	unit := hp.vec[0]
	if size == 1 {
		hp.vec = hp.vec[:0]
	} else {
		hp.vec[0] = hp.vec[size-1]
		hp.vec = hp.vec[:size-1]
		hp.down(0)
	}
	return unit
}

func (hp *Heap[T]) down(pos int) {
	tgt := hp.vec[pos]
	kid, last := pos*2+1, len(hp.vec)-1
	for kid < last {
		if hp.vec[kid+1] < hp.vec[kid] {
			kid++
		}
		if tgt <= hp.vec[kid] {
			break
		}
		hp.vec[pos] = hp.vec[kid]
		pos, kid = kid, kid*2+1
	}
	if kid == last && tgt > hp.vec[kid] {
		hp.vec[pos], pos = hp.vec[kid], kid
	}
	hp.vec[pos] = tgt
}

func (hp *Heap[T]) up(pos int) {
	tgt := hp.vec[pos]
	for pos > 0 {
		parent := (pos - 1) / 2
		if hp.vec[parent] <= tgt {
			break
		}
		hp.vec[pos], pos = hp.vec[parent], parent
	}
	hp.vec[pos] = tgt
}
