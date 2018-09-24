package sort

import (
	"unsafe"
)

// 基数排序，不依赖比较操作，具有稳定性
// 复杂度为 O((w/m)N) & O(N+2^m)
func RadixSort(list []Unit) {
	const base = -int((^uint(0))>>1) - 1
	size := len(list)
	for i := 0; i < size; i++ {
		list[i].val += base
	}

	temp := make([]Unit, size)
	const MEMO_SIZE = 1 << 8
	memo := new([MEMO_SIZE]uint)

	const INT_BITS = uint(unsafe.Sizeof(temp[0].val)) * 8
	for step := uint(0); step < INT_BITS; step += 8 {
		for i := 0; i < MEMO_SIZE; i++ {
			memo[i] = 0
		}
		for i := 0; i < size; i++ {
			radix := uint8((list[i].val >> step) & 0xFF)
			memo[radix]++
		}
		line := uint(0)
		for i := 0; i < MEMO_SIZE; i++ {
			memo[i], line = line, line+memo[i]
		}
		for i := 0; i < size; i++ {
			radix := uint8((list[i].val >> step) & 0xFF)
			temp[memo[radix]] = list[i]
			memo[radix]++
		}
		list, temp = temp, list
	}

	//if INT_BITS%16 != 0 {
	//	copy(list, temp)
	//}
	for i := 0; i < size; i++ {
		list[i].val -= base
	}
}
