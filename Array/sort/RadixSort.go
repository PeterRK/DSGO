package sort

import (
	"unsafe"
)

// 基数排序，不依赖比较操作，具有稳定性
// 复杂度为 O((w/m)N) & O(N+2^m)
func RadixSort(list []int) {
	const base = -int((^uint(0))>>1) - 1
	var size = len(list)
	for i := 0; i < size; i++ {
		list[i] += base
	}

	var shadow = make([]int, size)
	var book = new([256]uint)

	const UINT_LEN = uint(unsafe.Sizeof(uint(0))) * 8
	for step := uint(0); step < UINT_LEN; step += 8 {
		for i := 0; i < 256; i++ {
			book[i] = 0
		}
		for i := 0; i < size; i++ {
			var radix = uint8((list[i] >> step) & 0xFF)
			book[radix]++
		}
		var line = uint(0)
		for i := 0; i < 256; i++ {
			book[i], line = line, line+book[i]
		}
		for i := 0; i < size; i++ {
			var radix = uint8((list[i] >> step) & 0xFF)
			shadow[book[radix]] = list[i]
			book[radix]++
		}
		list, shadow = shadow, list
	}

	//if bytesOfUint%2 == 0 {
	for i := 0; i < size; i++ {
		list[i] -= base
	}
	//} else {
	//	for i := 0; i < size; i++ {
	//		shadow[i] = list[i] - base
	//	}
	//}
}
