package sort

import (
	"DSGO/utils"
)

type Ranker[T any] struct {
	Less func(a, b *T) bool
}

func (rk *Ranker[T]) simpleSort(list []T) {
	if len(list) < 2 {
		return
	}
	for i := 1; i < len(list); i++ {
		key := list[i]
		if rk.Less(&key, &list[0]) {
			for j := i; j > 0; j-- {
				list[j] = list[j-1]
			}
			list[0] = key
		} else {
			pos := i
			for rk.Less(&key, &list[pos-1]) {
				list[pos] = list[pos-1]
				pos--
			}
			list[pos] = key
		}
	}
}

func (rk *Ranker[T]) heapDown(list []T, pos int) {
	key := list[pos]
	kid, last := pos*2+1, len(list)-1
	for kid < last {
		if rk.Less(&list[kid], &list[kid+1]) {
			kid++
		}
		if !rk.Less(&key, &list[kid]) {
			break
		}
		list[pos] = list[kid]
		pos, kid = kid, kid*2+1
	}
	if kid == last && rk.Less(&key, &list[kid]) {
		list[pos], pos = list[kid], kid
	}
	list[pos] = key
}

func (rk *Ranker[T]) heapSort(list []T) {
	for idx := len(list)/2 - 1; idx >= 0; idx-- {
		rk.heapDown(list, idx)
	}
	for sz := len(list) - 1; sz > 0; sz-- {
		list[0], list[sz] = list[sz], list[0]
		rk.heapDown(list[:sz], 0)
	}
}

func (rk *Ranker[T]) partition(list []T) int {
	pivot := list[len(list)/2]
	a, b := 0, len(list)-1
	for {
		for rk.Less(&list[a], &pivot) {
			a++
		}
		for rk.Less(&pivot, &list[b]) {
			b--
		}
		if a >= b {
			break
		}
		list[a], list[b] = list[b], list[a]
		a++
		b--
	}
	return a
}

func (rk *Ranker[T]) introSort(list []T, life uint) {
	if len(list) < lowerBound {
		rk.simpleSort(list)
	} else if life == 0 {
		rk.heapSort(list)
	} else {
		m := rk.partition(list)
		rk.introSort(list[:m], life-1)
		rk.introSort(list[m:], life-1)
	}
}

func (rk *Ranker[T]) Sort(list []T) {
	life := utils.Log2Ceil(uint(len(list))) * 2
	rk.introSort(list, life)
}
