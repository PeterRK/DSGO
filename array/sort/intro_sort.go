package sort

import (
	"constraints"
	"math/bits"
)

// 内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
// 复杂度为O(NlogN) & O(logN)。
// 主要限制了QuickSort的最坏情况，适合递归实现(没有爆栈风险)。
func IntroSort[T constraints.Ordered](list []T) {
	life := bits.Len(uint(len(list))) * 2
	introSort(list, life)
}

func introSort[T constraints.Ordered](list []T, life int) {
	if len(list) < lowerBound {
		SimpleSort(list)
	} else if life--; life < 0 {
		HeapSort(list)
	} else {
		m := partition(list)
		introSort(list[:m], life)
		introSort(list[m:], life)
	}
}

func BlockIntroSort[T constraints.Ordered](list []T) {
	life := bits.Len(uint(len(list))) * 2
	blockIntroSort(list, life)
}

func blockIntroSort[E constraints.Ordered](list []E, life int) {
	for len(list) >= lowerBound {
		if life--; life < 0 {
			HeapSort(list)
			return
		}
		m := blockPartition(list)
		blockIntroSort(list[m:], life)
		list = list[:m]
	}
	SimpleSort(list)
}

// 三分内省排序
func IntroSortY[T constraints.Ordered](list []T) {
	life := bits.Len(uint(len(list))) * 3 / 2
	introSortY(list, life)
}

func introSortY[T constraints.Ordered](list []T, life int) {
	for len(list) > lowerBoundY {
		if life--; life < 0 {
			HeapSort(list)
			return
		}
		fst, snd := triPartition(list)
		introSortY(list[:fst], life)
		introSortY(list[snd+1:], life)
		if list[fst] == list[snd] {
			return
		}
		list = list[fst+1 : snd]
	}
	SimpleSort(list)
}
