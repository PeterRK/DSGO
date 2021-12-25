package sort

import (
	"math/bits"
)

type Less[T any] func(a, b *T) bool

func (lt Less[T]) simpleSort(list []T) {
	if len(list) < 2 {
		return
	}
	for i := 1; i < len(list); i++ {
		key := list[i]
		if lt(&key, &list[0]) {
			for j := i; j > 0; j-- {
				list[j] = list[j-1]
			}
			list[0] = key
		} else {
			pos := i
			for ; lt(&key, &list[pos-1]); pos-- {
				list[pos] = list[pos-1]
			}
			list[pos] = key
		}
	}
}

func (lt Less[T]) heapDown(list []T, pos int) {
	key := list[pos]
	kid, last := pos*2+1, len(list)-1
	for kid < last {
		if lt(&list[kid], &list[kid+1]) {
			kid++
		}
		if !lt(&key, &list[kid]) {
			break
		}
		list[pos] = list[kid]
		pos, kid = kid, kid*2+1
	}
	if kid == last && lt(&key, &list[kid]) {
		list[pos], pos = list[kid], kid
	}
	list[pos] = key
}

func (lt Less[T]) heapSort(list []T) {
	for idx := len(list)/2 - 1; idx >= 0; idx-- {
		lt.heapDown(list, idx)
	}
	for sz := len(list) - 1; sz > 0; sz-- {
		list[0], list[sz] = list[sz], list[0]
		lt.heapDown(list[:sz], 0)
	}
}

func (lt Less[T]) partition(list []T) int {
	pivot := list[len(list)/2]
	a, b := 0, len(list)-1
	for {
		for lt(&list[a], &pivot) {
			a++
		}
		for lt(&pivot, &list[b]) {
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

func (lt Less[T]) introSort(list []T, life int) {
	if len(list) < lowerBound {
		lt.simpleSort(list)
	} else if life--; life < 0 {
		lt.heapSort(list)
	} else {
		m := lt.partition(list)
		lt.introSort(list[:m], life)
		lt.introSort(list[m:], life)
	}
}

func (lt Less[T]) Sort(list []T) {
	life := bits.Len(uint(len(list))) * 2
	lt.introSort(list, life)
}
