package sort

import (
	"constraints"
	"unsafe"
)

// 基数排序，不依赖比较操作，具有稳定性
// 复杂度为 O((w/m)N) & O(N+2^m)
func RadixSort[T constraints.Integer](list []T) {
	byteWidth := uint(unsafe.Sizeof(T(0)))
	bitWidth := byteWidth * 8
	base := T(T(1) << (bitWidth - 1))
	signed := (base >> bitWidth) != 0
	size := len(list)
	temp := make([]T, size)

	if signed {
		for i := 0; i < size; i++ {
			list[i] += base
		}
	}

	const bk = 1 << 8
	var memo [bk]int
	for step := uint(0); step < bitWidth; step += 8 {
		for i := 0; i < bk; i++ {
			memo[i] = 0
		}
		for i := 0; i < size; i++ {
			radix := uint8(list[i] >> step)
			memo[radix]++
		}
		off := 0
		for i := 0; i < bk; i++ {
			memo[i], off = off, off+memo[i]
		}
		for i := 0; i < size; i++ {
			radix := uint8(list[i] >> step)
			temp[memo[radix]] = list[i]
			memo[radix]++
		}
		list, temp = temp, list
	}

	if byteWidth%2 == 0 {
		if signed {
			for i := 0; i < size; i++ {
				list[i] -= base
			}
		}
	} else {
		if signed {
			for i := 0; i < size; i++ {
				list[i] = temp[i] - base
			}
		} else {
			copy(list, temp)
		}
	}
}
