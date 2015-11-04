package hash

import (
	"unsafe"
)

//对于已知输入和基础hash函数f1(x),f2(x)，构造完全hash函数F(x)=(A*f1(x)+B*f2(x))/M。
//当返回M为0时表示构造失败。
func PerfectHash(words [][]byte, f1, f2 func([]byte) uint) (M uint, A, B uint) {
	var size = len(words)
	var codes = make([][2]uint, size)
	for i, word := range words {
		codes[i] = [2]uint{f1(word), f2(word)}
	}

	var shadow = make([][2]uint, size)
	if hashDuplicate(codes, shadow) {
		return 0, 0, 0
	}

	return 0, 0, 0
}

func hashDuplicate(list [][2]uint, shadow [][2]uint) bool {
	var size = len(list)
	sort(list, shadow, 0)
	sort(list, shadow, 1)
	for i := 1; i < size; i++ {
		if equal(list[i], list[i-1]) {
			return true
		}
	}
	return false
}
func equal(a [2]uint, b [2]uint) bool {
	return a[0] == b[0] && a[1] == b[1]
}
func sort(list [][2]uint, shadow [][2]uint, part int) {
	var size = len(list)
	var book [256]uint

	const UINT_LEN = uint(unsafe.Sizeof(uint(0))) * 8
	for step := uint(0); step < UINT_LEN; step += 8 {
		for i := 0; i < 256; i++ {
			book[i] = 0
		}
		for i := 0; i < size; i++ {
			var radix = uint8((list[i][part] >> step) & 0xFF)
			book[radix]++
		}
		var line = uint(0)
		for i := 0; i < 256; i++ {
			book[i], line = line, line+book[i]
		}
		for i := 0; i < size; i++ {
			var radix = uint8((list[i][part] >> step) & 0xFF)
			shadow[book[radix]] = list[i]
			book[radix]++
		}
		list, shadow = shadow, list
	}
}
