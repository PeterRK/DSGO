package sort

import (
	"DSGO/utils"
	"constraints"
)

// 内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
// 复杂度为O(NlogN) & O(logN)。
// 主要限制了QuickSort的最坏情况，适合递归实现(没有爆栈风险)。
func IntroSort[T constraints.Ordered](list []T) {
	life := utils.Log2Ceil(uint(len(list))) * 2
	introSort(list, life)
}

func introSort[T constraints.Ordered](list []T, life uint) {
	if len(list) < lowerBound {
		SimpleSort(list)
	} else if life == 0 {
		HeapSort(list)
	} else {
		m := partition(list)
		introSort(list[:m], life-1)
		introSort(list[m:], life-1)
	}
}

// 三分内省排序
func IntroSortY[T constraints.Ordered](list []T) {
	life := utils.Log2Ceil(uint(len(list))) * 3 / 2
	introSortY(list, life)
}

func introSortY[T constraints.Ordered](list []T, life uint) {
	for len(list) > lowerBoundY {
		if life == 0 {
			HeapSort(list)
			return
		}
		life--
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
